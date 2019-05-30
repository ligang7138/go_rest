package router

import (
	"fmt"
	"go_rest/queue"
)

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	// 向分发器注册任务管道的工作通道
	WorkerPool chan chan queue.Job
	maxWorkers int
	pool       int
}


func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan queue.Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		maxWorkers: maxWorkers,
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
				fmt.Println(job.Name+"邮件已发送")
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job

			}(job)
		}
	}
}