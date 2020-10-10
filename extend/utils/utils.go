package utils

import (
	"TinyURL/extend/code"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ResponseFormat(c *gin.Context, respStatus *code.Code, data interface{}) {
	if respStatus == nil {
		fmt.Println("没有发现请求状态参数")
		respStatus = code.RequestParamError
	}
	c.JSON(respStatus.Status, gin.H{
		"code": respStatus.Code,
		"msg":  respStatus.Message,
		"data": data,
	})
}

func MakeSha1(source string) string {
	sha1Hash :=sha1.New()
	sha1Hash.Write([]byte(source))
	return hex.EncodeToString(sha1Hash.Sum(nil))
}
