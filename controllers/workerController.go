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
	Job queue.Job `json:"job"`
}
// {"version":"1.0.0","token":"aaab","data":[{"job":{"name":"dawang"}}]}
func PayloadHandler(c *gin.Context) {

	// Read the body into a string for json decoding
	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(c.Request.Body, 1024)).Decode(&content)
	if err != nil {
		//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//w.WriteHeader(http.StatusBadRequest)
		SendResponse(c, err, "错误")
		return
	}

	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content.Payloads {

		// let's create a job with the payload
		work := payload.Job

		// Push the work onto the queue.
		queue.JobQueue <- work
	}

	SendResponse(c, nil, "成功")
}
