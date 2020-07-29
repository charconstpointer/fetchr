package main

import (
	"context"
	"time"

	"github.com/charconstpointer/fetchr/pkg/fetcher"
)

func main() {
	w := fetcher.NewWorker()
	p := fetcher.NewProbe()
	j := fetcher.NewJob(p)

	ctx := context.Background()

	go w.Start(ctx)

	w.AddJob(j)
	time.Sleep(1000 * time.Second)
}
