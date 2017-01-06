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

	reviewer, err := pr.reviewer()
	if err != nil {
		return err
	}
	return slack.direct(reviewer, pr.idleMessage())
}

func (pr *PR) idleMessage() string {
	minutes := int(pr.idle().Minutes())
	return fmt.Sprintf("PR <%s|%d> is idle for %d minutes. Please review. Sincery yours, bot", pr.url, pr.id, minutes)
}

func (pr *PR) process() error {
	var err error
	minutes := int(pr.idle().Minutes())
	fmt.Println(pr.idleMessage())
	if minutes > 240 {
		err = pr.notify()
	}
	return err
}
