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

package options

import (
	"net"
	"os"
	"strconv"

	"github.com/dio/proxy/config"
)

var ForwardedArgs config.ForwardedArgs

type Flags struct {
	NodeID         string `help:"Specify the proxy node ID." default:"default"`
	XDSServerURL   string `help:"Specify the xDS server URL." default:"localhost:9901"`
	XDSClusterName string `help:"Specify the xDS cluster name." default:"xds_cluster"`
	XDSResources   string `help:"A directory containing resources." default:""`
	UseGoogleGRPC  bool   `help:"Specify to use Google gRPC client implementation."`
	AdminPort      int    `help:"Specify the admin interface port"`
	StatsPort      int    `help:"Specify the exposed stats port"`
	Output         string `help:"Specify the file to write the rendered config. Available values: stdout, stderr, or a valid file path"`
	Version        bool   `help:"Show application version."`
}

// NewFlags returns a parse-able instance to kong.Parse().
func NewFlags() *Flags {
	// We keep the forwarded args, and "modify" os.Args so it contains only maybe-valid-flags.
	os.Args = ForwardedArgs.Parse()
	// TODO(dio): Revert os.Args.
	return new(Flags)
}

// ToBootstrap returns boostrap object.
func (f *Flags) ToBootstrap() (*config.Bootstrap, error) {
	host, portstring, err := net.SplitHostPort(f.XDSServerURL)
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(portstring)
	if err != nil {
		return nil, err
	}

	return &config.Bootstrap{
		NodeID:           f.NodeID,
		XDSServerAddress: host,
		XDSServerPort:    port,
		XDSClusterName:   f.XDSClusterName,
		XDSResources:     f.XDSResources,
		StatsPort:        f.StatsPort,
		AdminPort:        f.AdminPort,
		Output:           f.Output,
		UseGoogleGRPC:    f.UseGoogleGRPC,
	}, nil
}
