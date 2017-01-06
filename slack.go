package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func sendToSlack(message string) error {
	api := slack.New(slack_token)

	channelID, timestamp, err := api.PostMessage("C3NS9BXCN", message, slack.PostMessageParameters{})
	if err != nil {
		return err
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
