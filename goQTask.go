package goQTask

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
)

type Task interface {

	// task exec function
	Run() bool

	ExecTime() int64
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
func (q *QTask) AddTask(ts *Task) bool {
	if ts == nil {
		fmt.Println("task is empty,add valid")
		return false
	}
	q.qchan <- ts
	return true

}

func (q *QTask) Run() {
	for {

		select {
		case task := <-q.qchan:
			if task == nil {
				fmt.Println("recived a empty task")
				return
			}
			fmt.Println("task exec time=", task.ExecTime())
		default:
		}

	}
}
