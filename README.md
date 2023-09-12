# forkJoinPool

`forkJoinPool` is a simple implementation of a thread pool in the Go programming language. This thread pool allows you to execute asynchronous tasks using multiple worker threads.

## Features

- Create a pool with a specified number of worker threads and a queue size.
- Support for submitting tasks as functions with a variable number of arguments and executing them in worker threads.
- Ability to gracefully shut down the pool and wait for the completion of all tasks.

# Installation
You can install this package using go get:

```go
go get github.com/StormCrownSC/forkJoinPool
```

## Usage

Here's an example of using `forkJoinPool`:

```go
package main

import (
	"fmt"
	"github.com/StormCrownSC/forkJoinPool"
)

func main() {
	// Create a pool with 4 worker threads and a queue size of 10.
	pool, err := forkJoinPool.NewPool(4, 10)

    if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Submit a task to the pool.
	pool.Submit(func(args ...interface{}) {
		fmt.Println("Task executed with args:", args)
	}, i, "Hello")

	// Waiting for tasks to complete and shutting down the pool
    pool.Wait()
    pool.Shutdown()

}
```

When creating a pool using NewPool, you can specify two parameters:
- numWorkers: This parameter determines the number of worker threads in the pool. It must be greater than 0.
- QueueSize: This parameter sets the size of the task queue. It should not be less than 0. If set to 0, it means an unlimited queue size.

In the example above, we created a pool with 4 worker threads (numWorkers) and a queue size of 10 (QueueSize). Please ensure that numWorkers is greater than 0 and QueueSize is not less than 0 when configuring your pool.