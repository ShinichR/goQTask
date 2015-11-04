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
	Less(llrb.Item) bool
}

const (
	maxTime int64 = 0x0FFFFFFFFFFFFFFF
)

type QTask struct {
	tree     *llrb.LLRB
	waitTime int64
	qchan    chan Task
	mTask    map[string]Task
}

func NewQTask() *QTask {
	return &QTask{
		tree:     llrb.New(),
		qchan:    make(chan Task, 30),
		waitTime: maxTime,
		mTask:    make(map[string]Task),
	}
}
func (q *QTask) AddTask(ts Task) bool {
	if ts == nil {
		fmt.Println("task is empty,add valid")
		return false
	}
	//fmt.Println("AddTask ", ts.TaskName())
	q.qchan <- ts
	return true

}

func (q *QTask) Run() {
	timer := time.NewTimer(time.Duration(maxTime))
	for {

		select {
		case task := <-q.qchan:
			if task == nil {
				fmt.Println("recived a empty task")
				return
			}

			//fmt.Printf("I received a task. Current Time %d, it ask me to run it at %d\n",
			//	time.Now().Nanosecond(), task.ExecTime())
			if task.ExecTime() <= int64(time.Now().Nanosecond()) {
				fmt.Println("run task ", task.TaskName())
				go task.Run()
				continue
			}
			q.mTask[task.TaskName()] = task
			q.tree.InsertNoReplace(task)
			x := q.tree.Min()
			task = x.(Task)
			if task != nil {
				timer.Reset(time.Duration(task.ExecTime() - int64(time.Now().Nanosecond())))
				//fmt.Printf("wait %d\n", task.ExecTime()-int64(time.Now().Nanosecond()))
			}
			//debug
			/*
				q.tree.AscendGreaterOrEqual(q.tree.Min(), func(item llrb.Item) bool {
					task := item.(Task)
					if task == nil {
						return false
					}
					fmt.Printf("debug:----name:[%s].....time:[%d]\n", task.TaskName(), task.ExecTime())
					return true
				})
			*/
			//debug
		case <-timer.C:
			x := q.tree.Min()
			task := x.(Task)
			//fmt.Printf("run task:[%s],%d,%d\n", task.TaskName(), task.ExecTime(), int64(time.Now().Nanosecond()))
			for task.ExecTime() <= int64(time.Now().Nanosecond()) {
				go task.Run()
				q.tree.DeleteMin()
				x = q.tree.Min()

				if x == nil {
					timer.Reset(time.Duration(maxTime))
					task = nil
					break
				}
				task = x.(Task)
				//fmt.Printf("min task:[%s]\n", task.TaskName())
			}
			if task != nil {
				timer.Reset(time.Duration(task.ExecTime() - int64(time.Now().Nanosecond())))
				//fmt.Println("next task wait:", task.ExecTime()-int64(time.Now().Nanosecond()))
			}

		}

	}
}
