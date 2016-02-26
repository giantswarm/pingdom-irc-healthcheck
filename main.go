package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/giantswarm/pingdom-irc-healthcheck/irc"
)

const (
	PingdomTemplate = `<pingdom_http_custom_check>
	<status>%v</status>
	<response_time>%v</response_time>
</pingdom_http_custom_check>`
)

var (
	IRCAddress = "irc.freenode.org:6667"
	Channel    = "#giantswarm"
	Nickname   = "bridgesupport"
	Username   = "bridgesupport"
	Port       = "8000"
	NameToFind = "slackbridge"
)

func main() {
	log.Println("Starting server on port:", Port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+Port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	log.Println("Handling request")

	config := irc.Config{
		Address:  IRCAddress,
		Channel:  Channel,
		Nickname: Nickname,
		Username: Username,
	}

	client, err := irc.NewClient(config)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}

	namesResponseChan := make(chan string)
	go client.GetNames(namesResponseChan)
	namesResponse := <-namesResponseChan

	names := strings.Split(namesResponse, " ")
	found := false
	for _, name := range names {
		if name == NameToFind {
			found = true
		}
	}

	status := "NOT OK"
	if found {
		status = "OK"
	}

	elapsedSeconds := time.Since(start).Seconds()

	fmt.Fprintf(w, PingdomTemplate, status, elapsedSeconds)
}
