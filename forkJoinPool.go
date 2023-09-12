package forkJoinPool

import (
	"fmt"
	"time"
)

type Worker struct {
	ID     int
	taskCh chan func()
}

type Pool struct {
	NumWorkers int
	QueueSize  int
	workers    []*Worker
	taskQueue  chan func()
}

func NewPool(numWorkers, QueueSize int) (*Pool, error) {
	if numWorkers <= 0 {
		return nil, fmt.Errorf("numWorkers must be greater than 0")
	}

	if QueueSize < 0 {
		return nil, fmt.Errorf("QueueSize cannot be negative")
	}

	pool := &Pool{
		NumWorkers: numWorkers,
		QueueSize:  QueueSize,
		taskQueue:  make(chan func(), QueueSize),
	}

	for i := 0; i < int(numWorkers); i++ {
		worker := &Worker{
			ID:     i,
			taskCh: pool.taskQueue,
		}
		pool.workers = append(pool.workers, worker)
		go worker.Start()
	}

	return pool, nil
}

func (worker *Worker) Start() {
	for act := range worker.taskCh {
		if act != nil {
			act()
		}
	}
}

func (pool *Pool) Shutdown() {
	close(pool.taskQueue)
}

func (pool *Pool) Wait() {
	time.Sleep(10 * time.Millisecond)
	for {
		if len(pool.taskQueue) == 0 {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (pool *Pool) Submit(act func(...interface{}), args ...interface{}) {
	go func() {
		taskWithArgs := func() {
			act(args...)
		}

		if pool.QueueSize > len(pool.taskQueue) || pool.QueueSize == 0 {
			pool.taskQueue <- taskWithArgs
		}
	}()
}
