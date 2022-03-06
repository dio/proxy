package xds

import (
	"context"
	"os"
	"syscall"

	"github.com/dio/proxy/internal/xds/config"
	xdsserver "github.com/dio/proxy/internal/xds/server"
	"github.com/dio/proxy/internal/xds/watcher"
	"github.com/oklog/run"
)

// Run runs the main handler.
func Run(ctx context.Context, c *config.Bootstrap) error {
	var g run.Group
	g.Add(run.SignalHandler(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM))

	s := xdsserver.New(c)
	{
		ctx, cancel := context.WithCancel(ctx)
		g.Add(func() error {
			return s.Run(ctx)
		}, func(err error) {
			s.Interrupt(err)
			cancel()
		})
	}

	w := watcher.New(c, s)
	{
		ctx, cancel := context.WithCancel(ctx)
		g.Add(func() error {
			return w.Run(ctx)
		}, func(err error) {
			cancel()
		})
	}

	if err := g.Run(); err != nil {
		if _, ok := err.(run.SignalError); ok {
			return nil
		}
		return err
	}
	return nil
}
