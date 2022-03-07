package runner

import (
	"context"
	"os/exec"
	"runtime"
	"sync"
	"syscall"
)

type startWrapper struct {
	cmd   *exec.Cmd
	errCh chan error
}

type starter struct {
	once sync.Once
	cmd  *exec.Cmd
	ch   chan *startWrapper
}

func newStarter(cmd *exec.Cmd) *starter {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Pdeathsig: syscall.SIGKILL}
	return &starter{
		ch:  make(chan *startWrapper),
		cmd: cmd,
	}
}

func (s *starter) Start(ctx context.Context) error {
	s.once.Do(func() {
		go func() {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()

			for v := range s.ch {
				v.errCh <- v.cmd.Start()
			}
		}()
	})
	w := &startWrapper{
		cmd:   s.cmd,
		errCh: make(chan error, 1),
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.ch <- w:
	}
	return <-w.errCh
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
