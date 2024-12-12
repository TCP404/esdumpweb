package router

import (
	"net/http"

	"esdumpweb/initial"

	"github.com/gin-gonic/gin"
)

func indexHandle(c *gin.Context) {
	conf := initial.WireConfig()
	ret := gin.H{
		"host":         conf.AddrName,
		"index":        conf.Index,
		"timeField":    conf.TimeField,
		"product":      conf.Product,
		"saveLocation": conf.SaveDir,
		"condition":    conf.Condition,
	}
	c.HTML(http.StatusOK, "index.tpl", ret)
}

func shutdownHandle(c *gin.Context) {
	defer Shutdown()
	c.JSON(http.StatusOK, gin.H{"msg": "shutdown"})
}

// func updateHandler(c *gin.Context) {}
