package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendErrJSON 有错误发生时，发送错误JSON
func SendErrJSON(msg string, args ...interface{}) {
	fmt.Println(msg)
	if len(args) == 0 {
		panic("缺少 *gin.Context")
	}
	var c *gin.Context
	if len(args) == 1 {
		theCtx, ok := args[0].(*gin.Context)
		if !ok {
			panic("缺少 *gin.Context")
		}
		c = theCtx
	} else if len(args) == 2 {
		status, ok := args[0].(int)
		if !ok {
			panic("errNo不正确")
		}
		theCtx, ok := args[1].(*gin.Context)
		if !ok {
			panic("缺少 *gin.Context")
		}
		c = theCtx
		c.JSON(status, gin.H{
			"msg": msg,
		})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": msg,
	})
	// 终止请求链
	c.Abort()
}
