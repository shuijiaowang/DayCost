// pkg/middleware/error.go
package middleware

import (
	"DayCost/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

// ErrorHandler 全局异常处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印错误堆栈信息，方便调试
				debug.PrintStack()
				util.Result(c, http.StatusInternalServerError, "服务器内部错误", nil)
			}
		}()

		c.Next()
	}
}
