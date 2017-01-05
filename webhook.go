package main

import "fmt"

func webhookHandler(ctx *Context) error {
	fmt.Printf("EVENT")
	return nil
}
