package queue

import "os"

var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

// Job represents the job to be run
type Job struct {
	Name string `json:"name"`
}

// A buffered channel that we can send work requests on.
var JobQueue = make(chan Job,10)
//var JobQueue chan Job
