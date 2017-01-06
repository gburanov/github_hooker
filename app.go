package main

import (
	"fmt"
	"time"
)

type App struct {
	prs map[int]PR
}

func (a *App) Init() {
	a.prs = make(map[int]PR)
}

func (a *App) updatePR(number int, url string) {
	fmt.Printf("Update PR %d\n", number)
	pr, ok := a.prs[number]
	if ok == false {
		pr = PR{id: number, url: url, notified: false, app: a}
	}
	pr.changed = time.Now()
	pr.notified = false
	a.prs[number] = pr
}

func (a *App) closePR(number int) {
	fmt.Printf("Close PR %d\n", number)
	delete(a.prs, number)
}

func (a *App) Process() {
	for _, pr := range a.prs {
		if !pr.notified {
			pr.Process()
		}
	}
}
