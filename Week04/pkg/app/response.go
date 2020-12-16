package app

import (
	"service-notification/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}
var EmptRp = make(map[string]interface{})
func (g *Gin) Response(httpCode, errCode int, data interface{})  {
	g.C.JSON(httpCode, gin.H{
		"code" : errCode,
		"msg"  : e.GetMsg(errCode),
		"result" : data,
	})
	return
}