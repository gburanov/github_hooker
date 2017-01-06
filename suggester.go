package main

import "github.com/gburanov/pr_helper/sqs_lib"

func (pr *PR) reviewer() (string, error) {
	queue, err := sqs_lib.InputQueue()
	if err != nil {
		return "", err
	}
	err = queue.Send(pr.url)
	return "gburanov", err
}
