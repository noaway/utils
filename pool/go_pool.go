package pool

import (
// "fmt"
// "time"
)

type Worker struct {
	jobs   chan func()
	notify chan int
}

func NewWorker() *Worker {
	w := new(Worker)
	w.jobs = make(chan func(), 100)
	w.notify = make(chan int, 1)
	go w.working()
	w.notify <- 10
	return w
}

//Create a new goroutine to run a task
func (w *Worker) working() {
	for {
		select {
		case f := <-w.jobs:
			f()
		case nums := <-w.notify:
			for i := 0; i < nums; i++ {
				go w.working()
			}
		}
	}
}

func (w *Worker) Go(f func()) {
	w.jobs <- f
}
