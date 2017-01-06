package pool

import (
	"fmt"
	"runtime"
	"testing"
	// "time"
)

func TestGopool(t *testing.T) {
	work := NewWorker()
	for i := 0; i < 10000; i++ {
		work.Go(func() {
			fmt.Println(i)
		})
	}

	t.Log(runtime.NumGoroutine())
	// time.Sleep(time.Second * 1000)
}
