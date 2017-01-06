package main

import (
	"fmt"
	"time"
)

type PR struct {
	app      *App
	id       int
	changed  time.Time
	url      string
	notified bool
}

func (pr *PR) idle() time.Duration {
	return time.Since(pr.changed)
}

func (pr *PR) notify() error {
	slack := Slack{}
	slack.init()
	err := slack.broadcast(pr.idleMessage())
	if err != nil {
		return err
	}
	pr.notified = true
	pr.app.prs[pr.id] = *pr

	reviewer := pr.reviewer()
	return slack.notify(reviewer, pr.idleMessage())
}

func (pr *PR) idleMessage() string {
	minutes := int(pr.idle().Minutes())
	return fmt.Sprintf("PR <%s|%d> is idle for %d minutes. Please review. Sincery yours, bot", pr.url, pr.id, minutes)
}

func (pr *PR) Process() {
	minutes := int(pr.idle().Minutes())
	fmt.Println(pr.idleMessage())
	if minutes > 240 {
		pr.notify()
	}
}
