package main

import "fmt"

func statusHandler(ctx *Context) error {
	fmt.Fprint(ctx.Res, "OK")
	return nil
}
