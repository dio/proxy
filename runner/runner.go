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

package runner

import (
	"context"
	_ "embed" // to allow embedding files
	"os"
	"os/exec"
	"path/filepath"
)

// New returns a new runner.
func New(binaryPath string, silence bool) *Runner {
	return &Runner{
		binaryPath: binaryPath,
		silence:    silence,
	}
}

// Runner runs proxy at binary path.
type Runner struct {
	binaryPath string
	silence    bool
}

// Run runs proxy with the specified arguments.
func (r *Runner) Run(ctx context.Context, args []string) error {
	// We don't use CommandContext here so we can send exactly SIGTERM instead of kill -9 or SIGINT
	// when killing the process.
	cmd := exec.Command(filepath.Clean(r.binaryPath), args...) //nolint:gosec
	// TODO(dio): Add log streamer when we have decided the log library that we want to use.
	if !r.silence {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	starter := newStarter(cmd)

	err := starter.Start(ctx)
	if err != nil {
		return err
	}

	// TODO(dio): Do checking on admin address path availability, so we know that the process is
	// ready.

	go func() {
		<-ctx.Done()
		_ = starter.Kill()
	}()

	err = starter.Wait()
	if err != nil {
		return err
	}

	if cmd.Process != nil {
		return cmd.Process.Kill()
	}
	return nil
}
