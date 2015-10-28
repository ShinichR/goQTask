package main

import (
	"fmt"
	"github.com/ShinichR/goQTask"
)

type task struct{}

func (t *task) Run() bool {

	fmt.Println("task run...")
	return true
}
func (t *task) ExecTime() int64 {
	return 2
}

func main() {

	qtask := goQTask.NewQTask()
	fmt.Println("qtask run...")
	go qtask.Run()
	qtask.AddTask(&task{})
	for {

	}

}
