package controller

import (
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, Response{
		Status:    SUCCESS,
		Message:   "Hello from GoBlog",
		Translate: "hello.from.GoBlog",
	})
}
