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

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kong"

	"github.com/dio/proxy/internal/options"
	"github.com/dio/proxy/internal/proxy"
)

var (
	version = "dev"
	commit  = "dev"
)

func main() {
	flags := options.NewFlags()
	kong.Parse(flags)

	if flags.Version {
		fmt.Printf("proxy version: %s (commit: %s)\n", version, commit)
		return
	}

	b, err := flags.ToBootstrap()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := proxy.Run(context.Background(), b); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
