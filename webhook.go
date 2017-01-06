package main

import (
	"encoding/json"
	"fmt"

	githubhook "gopkg.in/rjz/githubhook.v0"
)

const secret string = "12345"

func webhookHandler(ctx *Context) error {
	hook, err := githubhook.Parse([]byte(secret), ctx.Req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("Processing hook " + hook.Event)
	if hook.Event == "pull_request" {
		processPullRequest(ctx.App, hook.Payload)
	}
	if hook.Event == "pull_request_review_comment" {
		processPullRequestCommit(ctx.App, hook.Payload)
	}
	return nil
}

type PullRequest struct {
	Number int `json:"number"`
}

type PullRequestChange struct {
	Action string      `json:"action"`
	Pr     PullRequest `json:"pull_request"`
}

func processPullRequest(app *App, payload []byte) error {
	change := PullRequestChange{}
	err := json.Unmarshal(payload, &change)
	if err != nil {
		return err
	}
	fmt.Println("What changed " + change.Action)
	if change.Action == "closed" {
		app.closePR(change.Pr.Number)
	} else {
		app.updatePR(change.Pr.Number)
	}
	return nil
}

func processPullRequestCommit(app *App, payload []byte) error {
	change := PullRequestChange{}
	err := json.Unmarshal(payload, &change)
	if err != nil {
		return err
	}
	fmt.Println("What changed " + change.Action)
	app.updatePR(change.Pr.Number)
	return nil
}
