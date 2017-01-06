package main

import "github.com/gburanov/pr_helper/sqs_lib"

func (pr *PR) reviewer() {
	queue := sqs_lib.InputQueue()
	queue.SendMessage(pr.url)
}
