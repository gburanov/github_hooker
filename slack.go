package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

const broadcastChannel string = "C3NS9BXCN"

type Slack struct {
	api *slack.Client
}

func (s *Slack) init() {
	s.api = slack.New(slack_token)
}

func (s *Slack) broadcast(message string) error {
	channelID, timestamp, err := s.api.PostMessage(broadcastChannel, message, slack.PostMessageParameters{})
	if err != nil {
		return err
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
	return nil
}

func (slack *Slack) direct(user string, message string) error {
	fmt.Printf("Message successfully sent to user %s\n", user)
	return nil
}
