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
