package middleware

import (
	"github.com/gin-gonic/gin"
)

func MiddlewareAuth() gin.HandlerFunc  {
	return func(c *gin.Context) {
		/*appG := app.Gin{c}
		cid, _ := strconv.Atoi(c.PostForm("company_id"));
		if  cid == 0 {
			appG.Response(403, -1, "auth fail")
		}
		c.Abort()*/
		return
	}

}