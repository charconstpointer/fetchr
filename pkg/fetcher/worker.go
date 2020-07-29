package fetcher

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type worker interface {
	AddJob(*job)
	Start(context.Context)
}
type Worker struct {
	j    []*job
	jobs chan (*job)
	rw   sync.RWMutex
}

func NewWorker() *Worker {
	return &Worker{
		j:    make([]*job, 0),
		jobs: make(chan *job),
	}
}

func (w *Worker) AddJob(j job) {
	w.rw.Lock()
	defer w.rw.Unlock()
	w.j = append(w.j, &j)
	w.jobs <- &j
}

func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case j := <-w.jobs:
			fmt.Printf("%v\n", *j)
		case _ = <-ctx.Done():
			log.Print("stopping worker")
			return
		}

	}
}

func (w *Worker) execute(j job) (interface{}, error) {
	log.Print("executing job")
	return nil, nil
}
