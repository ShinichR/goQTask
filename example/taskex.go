package main

import (
	"fmt"
	"github.com/ShinichR/goQTask"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type task struct {
	execTime int64
	taskname string
}

func New(e int64, name string) *task {
	return &task{
		execTime: e,
		taskname: name,
	}
}
func (t *task) Run() bool {

	fmt.Println(t.taskname, "task run...\n")
	return true
}
func (t *task) ExecTime() int64 {
	//fmt.Println(t.taskname, " execTime :", int64(time.Now().Nanosecond())+t.execTime)
	return int64(time.Now().Nanosecond()) + t.execTime
}
func (t *task) TaskName() string {
	return t.taskname
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
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

	qtask.AddTask(&task{1000, "100"})
	qtask.AddTask(&task{7000, "200"})
	qtask.AddTask(&task{4000, "300"})
	qtask.AddTask(&task{2000, "400"})
	qtask.AddTask(&task{2500, "500"})

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
	qtask.AddTask(nil)
}
