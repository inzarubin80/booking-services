package handler

import (
	"github.com/gin-gonic/gin"
	//"log"
	"fmt"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Println(message)	
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}