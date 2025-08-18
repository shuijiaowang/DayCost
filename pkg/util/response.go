package util

import "github.com/gin-gonic/gin"

func SendSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func SendBadRequest(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"code": 400,
		"msg":  message,
	})
}

func SendUnauthorized(c *gin.Context, message string) {
	c.JSON(401, gin.H{
		"code": 401,
		"msg":  message,
	})
}
