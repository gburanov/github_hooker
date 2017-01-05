package main

import "time"

type App struct {
	prs []PR
}

type PR struct {
	id      int
	changed time.Time
}
