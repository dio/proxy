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

package xdsserver

import (
	"context"
	"net"
	"sync"
	"time"

	clusterservice "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	endpointservice "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	listenerservice "github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	routeservice "github.com/envoyproxy/go-control-plane/envoy/service/route/v3"
	runtimeservice "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	secretservice "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/dio/proxy/internal/xds/config"
)

const (
	grpcKeepaliveTime        = 30 * time.Second
	grpcKeepaliveTimeout     = 5 * time.Second
	grpcKeepaliveMinTime     = 30 * time.Second
	grpcMaxConcurrentStreams = 1000000
)

func New(c *config.Bootstrap) *Server {
	return &Server{
		c:     c,
		cache: cache.NewSnapshotCache(true, cache.IDHash{}, Logger{}),
	}
}

type SnaphotUpdater interface {
	UpdateSnaphot(ctx context.Context, nodeID string, snapshot *cache.Snapshot) error
}

type Server struct {
	grpcServer *grpc.Server
	cache      cache.SnapshotCache
	mu         sync.RWMutex
	c          *config.Bootstrap
	// This is set when grpcServer.Serve() is called regardless the call returns error or not.
	served bool
}

func (s *Server) Run(ctx context.Context) error {
	// See: https://github.com/envoyproxy/go-control-plane/blob/52a61f1448c5ba16cc3e7d6dece1ee12050d22f7/internal/example/server.go#L56-L59.
	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions,
		grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    grpcKeepaliveTime,
			Timeout: grpcKeepaliveTimeout,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             grpcKeepaliveMinTime,
			PermitWithoutStream: true,
		}),
	)
	s.grpcServer = grpc.NewServer(grpcOptions...)

	// Attach all handlers.
	registerServer(s.grpcServer, server.NewServer(ctx, s.cache, &callbacks{}))

	l, err := net.Listen("tcp", s.c.ListenAddress)
	if err != nil {
		return err
	}

	errors := make(chan error, 1)
	go func(s *Server) {
		errors <- s.grpcServer.Serve(l)
	}(s)

	s.mu.Lock()
	s.served = true
	s.mu.Unlock()

	select {
	case err = <-errors:
	case <-ctx.Done():
	}

	return err
}

func (s *Server) Interrupt(error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.grpcServer == nil || !s.served {
		return
	}
	s.grpcServer.GracefulStop()
}

func (s *Server) UpdateSnaphot(ctx context.Context, nodeID string, snapshot *cache.Snapshot) error {
	if snapshot == nil {
		return nil
	}
	if err := snapshot.Consistent(); err != nil {
		return err
	}
	return s.cache.SetSnapshot(ctx, nodeID, snapshot)
}

func registerServer(grpcServer *grpc.Server, srv server.Server) {
	discoverygrpc.RegisterAggregatedDiscoveryServiceServer(grpcServer, srv)
	endpointservice.RegisterEndpointDiscoveryServiceServer(grpcServer, srv)
	clusterservice.RegisterClusterDiscoveryServiceServer(grpcServer, srv)
	routeservice.RegisterRouteDiscoveryServiceServer(grpcServer, srv)
	listenerservice.RegisterListenerDiscoveryServiceServer(grpcServer, srv)
	secretservice.RegisterSecretDiscoveryServiceServer(grpcServer, srv)
	runtimeservice.RegisterRuntimeDiscoveryServiceServer(grpcServer, srv)
}
