package fetcher

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type worker interface {
	AddJob(*job)
	Start(context.Context)
}
type Worker struct {
	j    []*Job
	jobs chan (*Job)
	R    chan (*Result)
	rw   sync.RWMutex
}

func NewWorker() *Worker {
	return &Worker{
		j:    make([]*Job, 0),
		jobs: make(chan *Job),
		R:    make(chan *Result),
	}
}

func (w *Worker) AddJob(j *Job) {
	w.rw.Lock()
	defer w.rw.Unlock()
	w.j = append(w.j, j)

}

func (w *Worker) Start(ctx context.Context) {
	for i := 0; i < len(w.j); i++ {
		go func(i int) {
			job := (w.j[i])
			for {
				select {
				case _ = <-(*job).T.C:
					res, err := w.execute(job)
					if err != nil {
						log.Print(err)
						break
					}
					select {
					case w.R <- res:
					default:
						log.Println("can't persist job result")
					}
				case _ = <-(*job).D:
					log.Print("stopping worker")
					return
				case _ = <-ctx.Done():
					log.Print("stopping worker")
					return
				}
			}
		}(i)
	}
}

func (w *Worker) execute(j *Job) (*Result, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	start := time.Now()
	r, err := client.Get(j.p.url)
	stop := time.Since(start)
	if err != nil {
		return &Result{
			res:     err.Error(),
			dur:     int(stop.Nanoseconds()),
			success: false,
			date:    time.Now(),
		}, nil
	}

	res, err := w.parseResp(r)
	if err != nil {
		return &Result{
			Res:     res,
			Dur:     int(stop.Nanoseconds()),
			Success: false,
			Date:    time.Now(),
		}, nil
	}
	return &Result{
		Res:     res,
		Dur:     int(stop.Nanoseconds()),
		Success: true,
		Date:    time.Now(),
	}, nil
}
func (w *Worker) parseResp(r *http.Response) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
