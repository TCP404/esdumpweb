package router

import (
	"context"
	"fmt"
	"html/template"
	"log/slog"

	"esdumpweb/kit/verify"
	"esdumpweb/resources"
	"esdumpweb/schema"

	"github.com/TCP404/eutil"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	appLock eutil.Mutex
	cancel  context.CancelFunc
)

func Shutdown() {
	appLock.WithLock(func() {
		if cancel != nil {
			cancel()
		}
	})
}

func Wrapper(ctx context.Context, cancelFunc context.CancelFunc) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error(fmt.Sprintf("panic: %v", r))
		}
	}()

	appLock.WithLock(func() { cancel = cancelFunc })

	if err := start(ctx); err != nil {
		slog.Error(err.Error())
	}
}

func App() *gin.Engine {
	app := gin.New()

	if err := setup(app); err != nil {
		panic(errors.Wrap(err, "failed to setup router"))
	}
	if err := printIP(); err != nil {
		panic(errors.Wrap(err, "failed to print ip"))
	}
	return app
}

func start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			slog.Info("server shutdown")
			return nil
		default:
			gin.SetMode(gin.ReleaseMode)
			if err := App().Run(":80"); err != nil {
				return errors.Wrap(err, "server runtime error")
			}
		}
	}
}

func printIP() error {
	ip, err := eutil.GetLocalIP()
	if err != nil {
		return errors.Wrap(err, "failed to get local ip")
	}
	slog.Info(fmt.Sprintf("Server started at http://%s:80 or http://localhost:80", ip))
	return nil
}

func setup(app *gin.Engine) error {
	app.Use(gin.Logger(), gin.Recovery())

	if err := loadResources(app); err != nil {
		return err
	}

	indexG := app.Group("/")
	{
		indexG.GET("/", indexHandle)
		indexG.GET("/shutdown", shutdownHandle)
		indexG.GET("/feedback", feedbackHandle)
		// indexG.GET("/update", updateHandle)
	}
	dumpG := app.Group("/")
	{
		dumpG.GET("/process", processHandle)
		dumpG.POST("/dump", verify.ValidateX(schema.DumpFixer), dumpAuth(), dumpHandle)
	}
	loginG := app.Group("/")
	{
		loginG.POST("/login", verify.ValidateX(schema.LoginFixer), loginHandle)
	}

	return nil
}

func loadResources(app *gin.Engine) error {
	t, err := template.ParseFS(resources.FS, "template/*.tpl")
	if err != nil {
		return err
	}
	app.SetHTMLTemplate(t)
	return nil
}
