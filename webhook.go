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
	if hook.Event == "issue_comment" {
		processIssueComment(ctx.App, hook.Payload)
	}
	return nil
}

type PullRequest struct {
	Number int    `json:"number"`
	URL    string `json:"html_url"`
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
		app.updatePR(change.Pr.Number, change.Pr.URL)
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
	app.updatePR(change.Pr.Number, change.Pr.URL)
	return nil
}

type IssueComment struct {
	Action string      `json:"action"`
	Issue  PullRequest `json:"issue"`
	URL    string      `json:"html_url"`
}

func processIssueComment(app *App, payload []byte) error {
	change := IssueComment{}
	err := json.Unmarshal(payload, &change)
	if err != nil {
		return err
	}
	fmt.Println("What changed " + change.Action)
	if change.Action == "created" {
		app.updatePR(change.Issue.Number, change.Issue.URL)
	}
	return nil
}
