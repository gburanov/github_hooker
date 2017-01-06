package main

import (
	"fmt"
	"time"
)

type App struct {
	prs map[int]PR
}

type PR struct {
	id      int
	changed time.Time
}

func (a App) updatePR(number int) {
	fmt.Printf("Update PR %d\n", number)
	pr, ok := a.prs[number]
	if ok == false {
		pr = PR{id: number}
	}
	pr.changed = time.Now()
	a.prs[number] = pr
}

func (a App) Process() {
	for _, pr := range a.prs {
		pr.Process()
	}
}
