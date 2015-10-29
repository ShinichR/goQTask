package goQTask

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
	"time"
)

type Task interface {

	// task exec function
	Run() bool

	ExecTime() int64

	TaskName() string
}

const (
	maxTime int64 = 0x0FFFFFFFFFFFFFFF
)

type QTask struct {
	tree     *llrb.LLRB
	waitTime int64
	qchan    chan Task
}

func NewQTask() *QTask {
	return &QTask{
		tree:     llrb.New(),
		qchan:    make(chan Task),
		waitTime: maxTime,
	}
}
func (q *QTask) AddTask(ts Task) bool {
	if ts == nil {
		fmt.Println("task is empty,add valid")
		return false
	}
	q.qchan <- ts
	return true

}

func (q *QTask) Run() {
	timer := time.NewTimer(time.Duration(q.waitTime))
	for {

		select {
		case task := <-q.qchan:
			if task == nil {
				fmt.Println("recived a empty task")
				return
			}

			fmt.Printf("I received a task. Current Time %d, it ask me to run it at %d\n",
				time.Now().Nanosecond(), task.ExecTime())
			if int(task.ExecTime()) <= time.Now().Nanosecond() {
				fmt.Println("run task ", task.TaskName())
				go task.Run()
				continue
			}
			q.tree.InsertNoReplace(llrb.Int(int(task.ExecTime()) - time.Now().Nanosecond()))
			waitTime := q.tree.Min()
			if waitTime != nil {
				timer.Reset(time.Nanosecond * time.Duration(waitTime.(llrb.Int)))
				fmt.Printf("wait %d\n", waitTime.(llrb.Int))
			}
		case <-timer.C:
			fmt.Println("time to run task")
		default:
		}

	}
}
