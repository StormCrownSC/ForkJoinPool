# forkJoinPool

`forkJoinPool` is a simple implementation of a thread pool in the Go programming language. This thread pool allows you to execute asynchronous tasks using multiple worker threads.

## Features

- Create a pool with a specified number of worker threads and a queue size.
- Support for submitting tasks as functions with a variable number of arguments and executing them in worker threads.
- Ability to gracefully shut down the pool and wait for the completion of all tasks.

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
	pool := forkJoinPool.NewPool(4, 10)

	// Submit a task to the pool.
	pool.Submit(func(args ...interface{}) {
		fmt.Println("Task executed with args:", args)
	}, 42, "Hello")

	// Shutdown the pool gracefully and wait for task completion.
	pool.Shutdown()
	pool.Wait()
}
```

# Installation
You can install this package using go get:

```go get github.com/StormCrownSC/forkJoinPool```