package main

import (
	"fmt"
	"github.com/ShinichR/goQTask"
	"time"
)

type task struct{}

func (t *task) Run() bool {

	fmt.Println("task run...\n")
	return true
}
func (t *task) ExecTime() int64 {
	return int64(time.Now().Nanosecond() + int(800))
}
func (t *task) TaskName() string {
	return "test"
}

func main() {

	qtask := goQTask.NewQTask()
	fmt.Println("qtask run...")
	/*t1 := time.NewTimer(time.Second * 1)
	t2 := time.NewTimer(time.Second * 2)
	for {
		select {
		case <-t1.C:
			fmt.Println("t1 out", time.Now().Nanosecond())
			t1.Reset(time.Second * 3)
		case <-t2.C:
			fmt.Println("t2 out", time.Now().Nanosecond())
		}
	}*/
	go qtask.Run()
	qtask.AddTask(&task{})
	for {

	}

}
