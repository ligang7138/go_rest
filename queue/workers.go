package queue

import (
	"fmt"
	"github.com/spf13/viper"
)

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool  chan chan Job
	JobChannel  chan Job
	quit        chan bool
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job ,viper.GetInt("max_queue")),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {

	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
				case job := <-w.JobChannel:
					// we have received a work request. 处理业务逻辑
					/*if err := job.Payload.UploadToS3(); err != nil {
						log.Errorf("Error uploading to S3: %s", err.Error())
						w.JobChannel <- job
					}*/

					fmt.Println("w-",job)
				case <-w.quit:
					fmt.Println("w-")
					// we have received a signal to stop
					return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
