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

package config

import (
	"os"
)

// Forwarded arguments. This is used by handler when building handler#Args for spawning the proxy
// process.
type ForwardedArgs []string

// Parse splits os.Args into two by looking for "--". The first part will be passed for further
// flag parsing. The second one will be passed directly to the spawned proxy process.
//
// A sample usage for Parse(): internal/flags/flags.go#NewFlags for preparing kong parseable object.
// One may "populate" this confing.ForwardedArgs in any preferrable ways.
func (a *ForwardedArgs) Parse() []string {
	maybeArgs := os.Args
	for index, arg := range os.Args {
		if arg == "--" {
			maybeArgs = os.Args[:index]
			*a = os.Args[index+1:]
			break
		}
	}
	return maybeArgs
}

// Bootstrap holds the required parameters for bootstrapping proxy configuration.
type Bootstrap struct {
	NodeID           string
	XDSServerAddress string
	XDSServerPort    int
	XDSClusterName   string
	XDSResources     string
	AdminPort        int
	StatsPort        int
	Output           string
	UseGoogleGRPC    bool
}
