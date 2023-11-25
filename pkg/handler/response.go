package handler

import "github.com/gin-gonic/gin"

type messageResponse struct {
	Message string
}

func responseErrorMessage(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, messageResponse{message})
}
