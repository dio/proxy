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
	"os/exec"
)

type starter struct {
	cmd *exec.Cmd
}

func newStarter(cmd *exec.Cmd) *starter {
	return &starter{cmd: cmd}
}

func (s *starter) Start(ctx context.Context) error {
	return s.cmd.Start()
}

func (s *starter) Wait() error {
	return s.cmd.Wait()
}

func (s *starter) Kill() error {
	if s.cmd.Process != nil {
		return s.cmd.Process.Kill() // TODO(dio): Checkout gopsutil.
	}
	return nil
}
