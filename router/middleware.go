package router

import (
	"net/http"

	"esdumpweb/initial"
	"esdumpweb/kit/verify"
	"esdumpweb/schema"

	"github.com/gin-gonic/gin"
)

func dumpAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := verify.GetReqParams[schema.DumpReq](c)
		conf := initial.WireConfig()
		host, ok := conf.Addrs[req.AddrName]
		if !ok || len(host.Username) == 0 || len(host.Password) == 0 {
			c.JSON(http.StatusForbidden, gin.H{"message": "请配置用户名和密码"})
			c.Abort()
			return
		}
		c.Set("username", host.Username)
		c.Set("password", host.Password)
		c.Next()
	}
}
