package router

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/TCP404/eutil/fetch"
	"github.com/gin-gonic/gin"

	"esdumpweb/initial"
	"esdumpweb/kit/verify"
	"esdumpweb/schema"
)

func loginHandle(c *gin.Context) {
	req := verify.GetReqParams[schema.LoginReq](c)

	url := fmt.Sprintf("http://%s:%s@%s/", req.Username, req.Password, req.AddrHost)
	p := fetch.Fetch(http.MethodGet, url).
		Then(func(res *http.Response) {
			// save to config file
			initial.WireConfig().Addrs[req.AddrName] = initial.HostConfig{
				Host:     req.AddrHost,
				Username: req.Username,
				Password: req.Password,
			}
		}).
		Catch(func(res *http.Response, err error) {
			slog.Error(fmt.Sprintf("login failed: %v", err))
		})

	if !p.Await() {
		c.JSON(http.StatusForbidden, gin.H{"message": "登录失败"})
		return
	}

	if err := initial.WireConfigManager().Flush(); err != nil {
		slog.Error("save config failed", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "保存配置失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}

func logoutHandle(c *gin.Context) {
	initial.WireConfig().Addrs = make(map[string]initial.HostConfig)
	if err := initial.WireConfigManager().Flush(); err != nil {
		slog.Error("save config failed", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "保存配置失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "登出成功"})
}
