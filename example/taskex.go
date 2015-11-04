package main

import (
	"fmt"
	"github.com/ShinichR/goQTask"
	"github.com/petar/GoLLRB/llrb"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

type task struct {
	execTime int64
	taskname string
}

func (t *task) Run() bool {

	fmt.Println(t.taskname, "is running...", time.Now().Nanosecond())
	return true
}
func (t *task) ExecTime() int64 {
	//fmt.Println(t.taskname, " execTime :", t.execTime)
	return t.execTime
}
func (t *task) TaskName() string {
	return t.taskname
}

func (t *task) Less(x llrb.Item) bool {
	tt := x.(*task)
	return t.execTime < tt.execTime
}

func taskTest() {
	var i = 0
	qtask := goQTask.NewQTask()
	end := 100000000
	start := 10000000
	go qtask.Run()
	for i < 20 {
		randtask := rand.Intn(end-start) + start
		qtask.AddTask(&task{int64(time.Now().Nanosecond()) + int64(randtask), strconv.Itoa(randtask)})
		fmt.Println("add task:[", strconv.Itoa(randtask), "]")
		time.Sleep(time.Nanosecond * 1000000)
		i++
	}

}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("qtask run...")

	go taskTest()

	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)

}
