package core

import "github.com/candbright/go-ssh/ssh"

type Screen struct {
	Session ssh.Session
	Name    string
}

func (s Screen) Create() error {
	return s.Session.Run("screen", "-dmS", s.Name)
}

func (s Screen) Exists() bool {
	err := s.Session.Run("screen", "-ls", s.Name)
	return err == nil
}

func (s Screen) Exit() error {
	return s.Session.Run("screen", "-X", "-S", s.Name, "quit")
}

func (s Screen) ExecCmd(arg ...string) error {
	arg = append([]string{"-X", s.Name}, arg...)
	return s.Session.Run("screen", arg...)
}
