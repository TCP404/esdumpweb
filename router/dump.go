package router

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"esdumpweb/constant"
	"esdumpweb/initial"
	"esdumpweb/kit/verify"
	"esdumpweb/schema"

	"github.com/TCP404/esdumpcore/core"
	"github.com/TCP404/esdumpcore/schedule"
	"github.com/TCP404/eutil"
	"github.com/TCP404/eutil/etl"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type (
	E = core.Hit
	L = core.Hit
)

var (
	etlEngineLock    eutil.Mutex
	etlEngine        *etl.ETL[E, L]
	etlEngineRunning bool
	dataTotal        int64
)

func dumpHandle(c *gin.Context) {
	if checkETLEngine() {
		c.JSON(http.StatusConflict, gin.H{"message": "Task is running, please wait"})
		return
	}

	req := verify.GetReqParams[schema.DumpReq](c)
	slog.Info("received request", slog.String("req", req.String()))

	scheduler, err := buildScheduler(c, req)
	if err != nil {
		msg := "failed to build scheduler"
		slog.Error(msg, "err", errors.Cause(err).Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "total": 0})
		return
	}

	query, err := scheduler.BuildQuery()
	if err != nil {
		msg := "failed to build query"
		slog.Error(msg, "err", errors.Cause(err).Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "total": 0})
		return
	}

	total, err := scheduler.QueryTotal(query)
	if err != nil {
		msg := "failed to query count"
		slog.Error(msg, "err", errors.Cause(err).Error(), "body", string(query.BodyBytes))
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "total": 0})
		return
	}

	slog.Info("Query count", slog.Int64("total:", total))
	if total == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No data to export", "total": 0})
		return
	}

	transformFunc := func(data []E) ([]E, error) {
		return data, nil
	}
	query = query.With(core.WithSize(constant.MaxSize))
	ins, err := scheduler.BuildWithETL(query, transformFunc, uint64(total))
	if err != nil {
		msg := "failed to build ETL engine"
		slog.Error(msg, "err", errors.Cause(err).Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "total": 0})
		return
	}

	go saveConfig(req)
	go func() {
		etlEngineLock.WithLock(func() {
			etlEngineRunning = true
			etlEngine = ins
			dataTotal = total
		})
		defer etlEngineLock.WithLock(func() {
			etlEngineRunning = false
			etlEngine = nil
			dataTotal = 0
		})

		runETL(scheduler)
	}()
	c.JSON(http.StatusOK, gin.H{"message": "Task submitted successfully", "total": total, "output": req.SaveLocation})
}

func checkETLEngine() (ok bool) {
	etlEngineLock.WithRLock(func() {
		ok = etlEngineRunning && etlEngine != nil
	})
	return
}

func buildScheduler(c *gin.Context, req schema.DumpReq) (*schedule.Scheduler, error) {
	username, password := c.GetString("username"), c.GetString("password")
	return schedule.New(
		req.AddrHost, username, password, req.Index,
		req.TimeField,
		req.StartTime,
		req.EndTime,
		req.SaveLocation, req.OutputHandler,
		req.BodyBool,
	)
}

func runETL(scheduler *schedule.Scheduler) {
	if err := scheduler.Init(); err != nil {
		slog.Error("scheduler initial", slog.String("err", err.Error()))
		return
	}
	defer scheduler.Close()

	ctx, cancel := signal.NotifyContext(context.TODO(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	defer cancel()

	slog.Info("dump start")
	if err := etlEngine.Run(ctx); err != nil {
		slog.Error("dump failed", slog.String("err", err.Error()))
	} else {
		slog.Info("dump success")
	}
	slog.Info("dump end")
}

func saveConfig(req schema.DumpReq) {
	conf := initial.WireConfigManager().Config
	conf.AddrName = req.AddrName
	conf.Index = req.Index
	conf.TimeField = req.TimeField
	conf.Product = req.Product
	conf.Condition = req.Condition
	if err := initial.WireConfigManager().Flush(); err != nil {
		slog.Error("save config failed", "err", err)
	}
}

func processHandle(c *gin.Context) {
	etlEngineLock.WithRLock(func() {
		if etlEngine == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message":  "Task not found",
				"progress": 0,
				"total":    0,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "Task progress retrieved successfully",
			"progress": etlEngine.Progress(),
			"total":    dataTotal,
		})
	})
}
