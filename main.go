package main

import (
	"context"
	"os/signal"
	"syscall"

	"esdumpweb/initial"
	"esdumpweb/router"
)

func main() {
	cm := initial.WireConfigManager()
	defer cm.Save()
	lm := initial.WireLoggerManager()
	defer lm.Close()

	ctx, cancel := signal.NotifyContext(context.TODO(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	defer cancel()

	go router.Wrapper(ctx, cancel)
	<-ctx.Done()
}
