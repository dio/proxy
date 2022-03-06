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
	"syscall"
)

// New returns a new runner.
func New(binaryPath string) *Runner {
	return &Runner{
		binaryPath: binaryPath,
	}
}

// Runner runs proxy at binary path.
type Runner struct {
	binaryPath string
}

// Run runs proxy with the specified arguments.
func (r *Runner) Run(ctx context.Context, args []string) error {
	// We don't use CommandContext here so we can send exactly SIGTERM instead of kill -9 or SIGINT
	// when killing the process.
	cmd := exec.Command(filepath.Clean(r.binaryPath), args...) //nolint:gosec
	// TODO(dio): Setpdeathsig to true and execute cmd.Start in a locked thread through channel on Linux.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// TODO(dio): Add log streamer when we have decided the log library that we want to use.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	// TODO(dio): Do checking on admin address path availability, so we know that the process is
	// ready.

	<-ctx.Done()

	if cmd.Process != nil {
		// TODO(dio): Make the following cross-platform.
		_ = syscall.Kill(cmd.Process.Pid, syscall.SIGTERM) // Kill the proxy process using SIGTERM.
	}
	return cmd.Wait()
}
