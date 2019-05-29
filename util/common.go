package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type jsonData struct {
	code int
	Data interface{}
	Msg string
}
func responseJson(c *gin.Context,data interface{})  {
	c.JSON(http.StatusOK,jsonData{
		code : 200,
		Msg : "kjdf",
		Data: data,
	})
}
