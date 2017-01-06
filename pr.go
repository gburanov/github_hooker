package main

import (
	"fmt"
	"time"
)

type PR struct {
	id      int
	changed time.Time
}

func (pr PR) Process() {
	fmt.Printf("PR %d time %s\n", pr.id, pr.changed)
}
