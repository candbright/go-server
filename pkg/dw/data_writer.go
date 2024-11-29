package dw

import (
	"github.com/candbright/go-ssh/ssh"
	"github.com/pkg/errors"
)

type Config struct {
	Session   ssh.Session
	Path      string
	Marshal   func(v any) ([]byte, error)
	Unmarshal func(data []byte, v any) error
}

type DataWriter[T any] struct {
	cfg  Config
	Data T
}

func New[T any](cfg Config) *DataWriter[T] {
	if cfg.Session == nil {
		session, err := ssh.NewSession()
		if err != nil {
			panic(err)
		}
		cfg.Session = session
	}
	manager := &DataWriter[T]{
		cfg: cfg,
	}
	err := manager.Read()
	if err != nil {
		panic(err)
	}
	return manager
}

func (manager *DataWriter[T]) Read() error {
	fileBytes, err := manager.cfg.Session.ReadFile(manager.cfg.Path)
	if err != nil {
		return errors.WithStack(err)
	}
	if len(fileBytes) == 0 {
		return nil
	}
	err = manager.cfg.Unmarshal(fileBytes, &manager.Data)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (manager *DataWriter[T]) Write() error {
	marshalBytes, err := manager.cfg.Marshal(manager.Data)
	if err != nil {
		return errors.WithStack(err)
	}
	err = manager.cfg.Session.WriteString(manager.cfg.Path, string(marshalBytes))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
