// Copyright 2022 Dhi Aurrahman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/oklog/run"

	"github.com/dio/proxy/config"
	"github.com/dio/proxy/handler"
	"github.com/dio/proxy/internal/downloader"
	xdsconfig "github.com/dio/proxy/internal/xds/config"
	xdsserver "github.com/dio/proxy/internal/xds/server"
	"github.com/dio/proxy/internal/xds/watcher"
	"github.com/dio/proxy/runner"
)

// Run runs the main handler.
func Run(ctx context.Context, c *config.Bootstrap) error {
	var g run.Group
	g.Add(run.SignalHandler(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM))

	binaryPath, err := downloader.Download(ctx)
	if err != nil {
		return err
	}

	// Handle config preparation, config watching, TLS establishment.
	h := handler.New(c)
	args, err := h.Args()
	if err != nil {
		return err
	}
	defer func() {
		_ = args.Cleanup()
	}()

	if c.Output != "" {
		return nil
	}

	if c.Resources != "" { // When user asks to watch a directory, activate the embedded xDS server.
		xdsBootstrap := &xdsconfig.Bootstrap{
			Resources:     c.Resources,
			ListenAddress: fmt.Sprintf(":%d", c.XDSServerPort),
		}

		xdsServer := xdsserver.New(xdsBootstrap)
		{

			runCtx, cancel := context.WithCancel(ctx)
			g.Add(func() error {
				return xdsServer.Run(runCtx)
			}, func(err error) {
				cancel()
			})
		}

		w := watcher.New(xdsBootstrap, xdsServer)
		{
			runCtx, cancel := context.WithCancel(ctx)
			g.Add(func() error {
				return w.Run(runCtx)
			}, func(err error) {
				cancel()
			})
		}
	}

	{
		runCtx, cancel := context.WithCancel(ctx)
		g.Add(
			func() error {
				return h.Run(runCtx)
			},
			func(err error) {
				cancel()
			})
	}

	{
		r := runner.New(binaryPath)
		runCtx, cancel := context.WithCancel(ctx)
		g.Add(func() error {
			return r.Run(runCtx, args.Values)
		},
			func(err error) {
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
