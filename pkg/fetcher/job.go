package fetcher

import (
	"log"
	"time"
)

type job interface {
	Cancel()
}

type Job struct {
	D chan (struct{})
	T *time.Ticker
	p Probe
}
type Result struct {
	Success bool
	Res     string
	Dur     int
	Date    time.Time
}

func NewJob(p *Probe) *Job {
	return &Job{
		p: *p,
		D: make(chan struct{}, 1),
		T: time.NewTicker(time.Duration(p.interval) * time.Second),
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

func (j *Job) Probe() Probe {
	return j.p
}
