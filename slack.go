package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func sendToSlack(message string) {
	api := slack.New("YOUR_TOKEN_HERE")
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "Your PR bot",
		Text:    message,
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage("C3NS9BXCN", message, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
