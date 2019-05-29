package router

import (
	"github.com/spf13/viper"
	"go_rest/queue"
)

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan queue.Job
	maxWorkers int
	pool       int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan queue.Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		maxWorkers: viper.GetInt("max_workers"),
	}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := queue.NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-queue.JobQueue:
			// a job request has been received
			go func(job queue.Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
