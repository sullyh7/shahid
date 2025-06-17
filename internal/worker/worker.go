package worker

import (
	"context"
	"fmt"
	"sync"
)

type Queue interface {
	Start(context.Context)
	Add(Job) error
}

type Job struct {
	VideoID   string
	ChannelID string
	Title     string
}

type WorkerPool struct {
	JobChan chan Job
	Size    int
	WG      sync.WaitGroup
	Process func(Job)
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < wp.Size; i++ {

		wp.WG.Add(1)
		go wp.worker(ctx, wp.JobChan)
	}
}

func (wp *WorkerPool) Add(job Job) error {
	select {
	case wp.JobChan <- job:
		return nil
	default:
		return fmt.Errorf("error adding job to the queue")
	}
}

func (wp *WorkerPool) Run(ctx context.Context) {
	wp.Start(ctx)
	wp.Wait()
}

func (wp *WorkerPool) Wait() {
	wp.WG.Wait()
}

func (wp *WorkerPool) worker(ctx context.Context, jobChan <-chan Job) {
	defer wp.WG.Done()
	for {
		select {
		case <-ctx.Done():
			return

		case job := <-jobChan:
			wp.Process(job)
			if ctx.Err() != nil {
				return
			}
		}
	}
}
