// package irc provides an interface to IRC.
package irc

import (
	"fmt"
	"log"

	"github.com/juju/errgo"
	"github.com/thoj/go-ircevent"
)

var (
	connectionNotMadeError = errgo.New("IRC connection not made")
)

// Client provides a client to IRC.
type Client struct {
	// Config is the configuration of the Client.
	Config Config

	// Conn is the IRC connection itself.
	Conn *irc.Connection
}

// Config provides configuration for the IRC client.
type Config struct {
	// Address is the address of the IRC channel to connect to - e.g: "irc.quakenet.org:6667".
	Address string
	// Channel is the channel you want to connect to.
	Channel string

	// Nickname is the nickname to use to connect.
	Nickname string
	// Username is the username to use to connect.
	Username string
}

// NewClient returns a new IRC Client, given a configuration.
func NewClient(config Config) (*Client, error) {
	conn := irc.IRC(config.Nickname, config.Username)
	if conn == nil {
		return nil, connectionNotMadeError
	}

	client := &Client{
		Config: config,
		Conn:   conn,
	}

	log.Print("Connecting to IRC")
	if err := client.Conn.Connect(client.Config.Address); err != nil {
		return nil, errgo.Mask(err)
	}

	return client, nil
}

// GetNames connects to the channel, issues a NAMES command,
// and returns the raw message on namesResponseChan.
func (c *Client) GetNames(namesResponseChan chan string) {
	c.Conn.AddCallback("001", func(e *irc.Event) {
		log.Println("Connected")

		log.Println("Joining channel:", c.Config.Channel)
		c.Conn.Join(c.Config.Channel)

		log.Println("Sending NAMES command")
		c.Conn.SendRaw(fmt.Sprintf("NAMES %v", c.Config.Channel))
	})

	c.Conn.AddCallback("353", func(e *irc.Event) {
		log.Println("Got NAMES response")
		namesResponseChan <- e.Message()

		log.Println("Closing IRC connection")
		c.Conn.Quit()
	})

	c.Conn.Loop()
}
