package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_rest/queue"
	. "go_rest/util"
	"io"
)
type PayloadCollection struct {
	WindowsVersion  string    `json:"version"`
	Token           string    `json:"token"`
	Payloads        []Payload `json:"data"`
}

type Payload struct {
	// [redacted]
}

func PayloadHandler(c *gin.Context) {

	// Read the body into a string for json decoding
	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(c.Request.Body, 1024)).Decode(&content)
	if err != nil {
		//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//w.WriteHeader(http.StatusBadRequest)
		c.Header("Content-Type", "application/json; charset=UTF-8")

		return
	}

	// Go through each payload and queue items individually to be posted to S3
	/*for _, payload := range content.Payloads {

		// let's create a job with the payload
		work := queue.Job{}

		// Push the work onto the queue.
		queue.JobQueue <- work
	}*/
	work := queue.Job{}
	job := make(chan queue.Job,10)

	// Push the work onto the queue.
	job <- work

	SendResponse(c, nil, "成功")
}
