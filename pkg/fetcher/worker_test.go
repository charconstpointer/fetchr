package fetcher

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type WorkerTestSuite struct {
	suite.Suite
	w *Worker
}

func TestWorkerTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerTestSuite))
}

func (suite *WorkerTestSuite) SetupTest() {
	suite.w = NewWorker()
}

func (suite *WorkerTestSuite) TestAddJob() {
	p := NewProbe(3, "https://google.com", 10)
	j := NewJob(p)
	ctx := context.Background()

	go suite.w.Start(ctx)
	suite.w.AddJob(j)
}

func (suite *WorkerTestSuite) TestStart() {
	ctx := context.Background()

	go suite.w.Start(ctx)

	p := NewProbe(1, "https://google.com", 10)
	j := NewJob(p)

	suite.w.AddJob(j)
}

func (suite *WorkerTestSuite) Testexecute() {

	w := NewWorker()

	p := NewProbe(2, "https://google.com", 10)
	j := NewJob(p)

	r, _ := w.execute(j)
	suite.True(r.Success)

}
