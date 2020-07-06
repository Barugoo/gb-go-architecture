package worker

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Address  string
	Method   string
	Uri      string
	Data     string
	TimeChan chan time.Duration
}

type Worker struct {
	jobCh <-chan Config
}

func (w *Worker) sendElapsedTime(start time.Time, ch chan<- time.Duration) {
	elapsed := time.Since(start)
	ch <- elapsed
}

func (w *Worker) HandleJobs() {
	for config := range w.jobCh {
		start := time.Now()

		client := &http.Client{}
		req, err := http.NewRequest(
			config.Method,
			fmt.Sprintf("%s/%s", config.Address, config.Uri),
			nil)

		if err != nil {
			log.Println(err)
		}

		_, err = client.Do(req)

		if err != nil {
			log.Println(err)
		}

		w.sendElapsedTime(start, config.TimeChan)
	}
}

func NewWorker(jobCh <-chan Config) *Worker {
	return &Worker{
		jobCh: jobCh,
	}
}
