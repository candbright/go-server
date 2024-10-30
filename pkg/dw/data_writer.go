package dw

import (
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	Path      string
	Marshal   func(v any) ([]byte, error)
	Unmarshal func(data []byte, v any) error
}

type DataWriter[T any] struct {
	cfg  Config
	Data T
}

func New[T any](cfg Config) *DataWriter[T] {
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
	fileBytes, err := os.ReadFile(manager.cfg.Path)
	if err != nil {
		return errors.WithStack(err)
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
	err = os.WriteFile(manager.cfg.Path, marshalBytes, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
