package runner

import (
	"context"
	"os/exec"
	"syscall"
)

type starter struct {
	cmd *exec.Cmd
}

func newStarter(cmd *exec.Cmd) *starter {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
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
		// TODO(dio): Make the following cross-platform.
		return syscall.Kill(s.cmd.Process.Pid, syscall.SIGTERM) // Kill the proxy process using SIGTERM.
	}
	return nil
}
