package core

import (
	"fmt"
	"github.com/pkg/errors"
	"path"

	"github.com/candbright/go-ssh/ssh"
)

type Process struct {
	Version string
	RootDir string
	Screen  Screen
}

type ProcessConfig struct {
	Version string
	RootDir string
	Session ssh.Session
}

func NewProcess(cfg ProcessConfig) (*Process, error) {
	p := &Process{
		Version: cfg.Version,
		RootDir: cfg.RootDir,
		Screen: Screen{
			Session: cfg.Session,
			Name:    fmt.Sprintf("mc-%s", cfg.Version),
		},
	}
	return p, nil
}

func (p *Process) Active() bool {
	return p.Screen.Exists()
}

func (p *Process) ExecFile() string {
	return path.Join(p.RootDir, "bedrock_server")
}

func (p *Process) ScreenName() string {
	return fmt.Sprintf("mc-%s", p.Version)
}

func (p *Process) Start() error {
	if !p.Active() {
		err := p.Screen.Create()
		if err != nil {
			return err
		}
		err = p.Screen.ExecCmd("LD_LIBRARY_PATH=.", p.ExecFile())
		if err != nil {
			return err
		}
	}
	return errors.Errorf("screen %s already started!", p.ScreenName())
}

func (p *Process) Stop() error {
	if p.Active() {
		return p.Screen.Exit()
	}
	return errors.Errorf("screen %s not found!", p.ScreenName())
}

func (p *Process) Restart() error {
	err := p.Stop()
	if err != nil {
		return err
	}
	err = p.Start()
	if err != nil {
		return err
	}
	return nil
}

func (p *Process) ExecCmd(arg ...string) error {
	return p.Screen.ExecCmd(arg...)
}
