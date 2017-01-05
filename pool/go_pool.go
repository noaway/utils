package pool

import (
	"sync"
	"time"
)

var (
	jobs chan func()
	wg   sync.WaitGroup
)

func NewWork(g *GWorker) *GWorker {
	if g != nil {
		initjob(g.GoMaxNumber)
		jobs = make(chan func(), g.GoMaxNumber)
		return g
	}
	return nil
}

func worker(id int) {
	wg.Done()
	time.Sleep(time.Second * 2)
	for j := range jobs {
		j()
	}
}

func initjob(count int) {
	for w := 1; w <= count; w++ {
		if w%1000 == 0 {
			time.Sleep(time.Nanosecond * 200)
		}
		wg.Add(1)
		go worker(w)
	}
}

type GWorker struct {
	sync.RWMutex
	m           sync.Mutex
	GoMaxNumber int
}

func (self *GWorker) JoinWork(f func()) {
	jobs <- f
}

func (self *GWorker) Wait() {
	wg.Wait()
}
