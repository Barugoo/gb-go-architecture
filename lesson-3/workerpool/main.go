package main

import (
	"fmt"
	"sync"
	"time"

	"workerpool/worker"
)

const numOfWorkers = 2

func main() {
	wg := &sync.WaitGroup{}

	jobCh := make(chan string)

	wg.Add(1)
	go func() {
		i := 0
		for {
			i++
			jobCh <- fmt.Sprintf("job %d", i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for i := 0; i < numOfWorkers; i++ {
		w := worker.NewWorker(i, jobCh)
		go w.HandleJobs()
	}

	wg.Wait()
}
