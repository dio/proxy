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

package archives

import (
	"path/filepath"

	"github.com/codeclysm/extract"
)

var DefaultProxyVersion = "1.21.0" // TODO(dio): Define minimum version instead.

type Archive interface {
	Version() string
	BinaryName() string
	BinaryDir() string
	URLPattern() string
	Renamer() extract.Renamer
}

type Proxy struct {
	VersionUsed string
}

func (p *Proxy) Version() string {
	if p.VersionUsed != "" {
		return p.VersionUsed
	}
	return DefaultProxyVersion
}

func (p *Proxy) BinaryName() string {
	return "envoy"
}

func (p *Proxy) BinaryDir() string {
	return filepath.Join("versions", p.Version(), "bin")
}

func (p *Proxy) URLPattern() string {
	return "https://archive.tetratelabs.io/envoy/download/v%s/envoy-v%s-%s-amd64.tar.xz"
}

func (p *Proxy) Renamer() extract.Renamer {
	return func(name string) string {
		baseName := filepath.Base(name)
		if baseName == p.BinaryName() {
			return filepath.Join(p.BinaryDir(), baseName)
		}
		return name
	}
}
