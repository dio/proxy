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

import "github.com/dio/proxy/internal/xds/config"

type Flags struct {
	ListenAddress string `help:"Specify the xDS server listen address." default:":9901"`
	Resources     string `help:"A directory containing resources." default:""`
	Version       bool   `help:"Show application version."`
}

func (f *Flags) ToBootstrap() *config.Bootstrap {
	return &config.Bootstrap{
		ListenAddress: f.ListenAddress,
		Resources:     f.Resources,
	}
}
