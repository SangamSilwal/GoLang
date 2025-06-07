package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormA struct {
	Foo string `json:"foo" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := FormA{}

	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `The body should be formA`)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Received",
		"data":    objA,
	})
}
