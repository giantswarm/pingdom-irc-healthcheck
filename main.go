package main

import (
	"log"
	"os"
	"strings"

	"github.com/giantswarm/pingdom-irc-healthcheck/irc"
)

func main() {
	nameToFind := "slackbridge"

	config := irc.Config{
		Address:  "irc.freenode.org:6667",
		Channel:  "#giantswarm",
		Nickname: "bridgesupport",
		Username: "bridgesupport",
	}

	client, err := irc.NewClient(config)
	if err != nil {
		log.Println(err)
		os.Exit(1)

	}

	namesResponseChan := make(chan string)
	go client.GetNames(namesResponseChan)
	namesResponse := <-namesResponseChan

	names := strings.Split(namesResponse, " ")
	found := false
	for _, name := range names {
		if name == nameToFind {
			found = true
		}
	}

	if found {
		log.Println("Found user:", nameToFind)
	} else {
		log.Println("Did not find user:", nameToFind)
	}
}
