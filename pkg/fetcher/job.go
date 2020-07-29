package fetcher

import "log"

type job interface {
	Cancel()
}

type Job struct {
	D chan (struct{})
	p probe
}

func NewJob(p probe) *Job {
	return &Job{
		p: p,
		D: make(chan struct{}, 1),
	}
}

func (j *Job) Cancel() {
	select {
	case j.D <- struct{}{}:
		log.Printf("cancelling job %v", &j)
	default:
		log.Printf("can not cancell job %v", &j)
	}
}

func (j *Job) execute() (interface{}, error) {
	log.Print("executing job")
	return nil, nil
}
