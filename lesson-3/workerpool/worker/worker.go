package worker

import "log"

type Worker struct {
	id    int
	jobCh <-chan string
}

func (w *Worker) HandleJobs() {
	for job := range w.jobCh {
		log.Printf("Worker %d writes %s", w.id, job)
	}
}

func NewWorker(id int, jobCh <-chan string) *Worker {
	return &Worker{
		id:    id,
		jobCh: jobCh,
	}
}
