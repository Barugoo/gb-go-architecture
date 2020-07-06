package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"ddos/worker"
)

func main() {
	address := flag.String("address", "", "")
	method := flag.String("method", "GET", "")
	uri := flag.String("uri", "", "")
	data := flag.String("data", "", "")
	threads := flag.Int("threads", 5, "")
	count := flag.Int("count", 100, "")

	flag.Parse()
	wg := &sync.WaitGroup{}

	jobCh := make(chan worker.Config)
	config := worker.Config{
		Address:  *address,
		Method:   *method,
		Uri:      *uri,
		Data:     *data,
		TimeChan: make(chan time.Duration),
	}

	wg.Add(1)
	go func() {
		for i := 0; i < *count; i++ {
			jobCh <- config
		}
	}()

	for i := 0; i < *threads; i++ {
		w := worker.NewWorker(jobCh)
		go w.HandleJobs()
	}

	go func() {
		attempts := 0
		var durationSum time.Duration

		for duration := range config.TimeChan {
			attempts += 1
			durationSum += duration

			avgDuration := time.Duration(int64(durationSum) / int64(attempts))
			fmt.Printf("%d request, average time is %d\n", attempts, avgDuration.Milliseconds())
			// fmt.Printf("%d\n", duration.Milliseconds())
		}

	}()

	wg.Wait()
}
