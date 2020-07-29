package main

import (
	"context"
	"log"

	"github.com/charconstpointer/fetchr/pkg/fetcher"
)

func main() {
	w := fetcher.NewWorker()
	p := fetcher.NewProbe("https://google.com", 1)
	j := fetcher.NewJob(p)

	p2 := fetcher.NewProbe("https://polskieradio.pl", 3)
	j2 := fetcher.NewJob(p2)

	p3 := fetcher.NewProbe("https://httpbin.org/delay/10", 3)
	j3 := fetcher.NewJob(p3)

	ctx := context.Background()
	w.AddJob(j)
	w.AddJob(j2)
	w.AddJob(j3)

	go w.Start(ctx)
	for {
		select {
		case r := <-w.R:
			log.Printf("%v\n", r.Success)
		}
	}

}
