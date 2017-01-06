package main

import (
	"fmt"
	"time"
)

type PR struct {
	id      int
	changed time.Time
}

func (pr PR) idle() time.Duration {
	return time.Since(pr.changed)
}

func (pr PR) notify() {
	sendToSlack(pr.idleMessage())
}

func (pr PR) idleMessage() string {
	minutes := int(pr.idle().Minutes())
	return fmt.Sprintf("PR %d is idle for %d minutes", pr.id, minutes)
}

func (pr PR) Process() {
	minutes := int(pr.idle().Minutes())
	fmt.Println(pr.idleMessage())
	if minutes > 2 {
		pr.notify()
	}
}
