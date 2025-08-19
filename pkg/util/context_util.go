// pkg/util/context_util.go
package util

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetUserID 安全获取上下文中的用户ID
func GetUserID(c *gin.Context) (int, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("未获取到用户信息")
	}

	id, ok := userID.(int)
	if !ok {
		return 0, errors.New("无效的用户ID类型")
	}

	return id, nil
}
