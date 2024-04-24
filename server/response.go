package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回成功
func ResponseSuccess(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, gin.H{
		"data": obj,
	})
}

func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"message": msg,
	})
}
