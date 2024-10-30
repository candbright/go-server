package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type ExistErr struct {
	Type string
	Id   string
}

func (error ExistErr) Error() string {
	return fmt.Sprintf("%s %s exist", error.Type, error.Id)
}

func ExistError(t, id string) error {
	return errors.WithStack(ExistErr{
		Type: t,
		Id:   id,
	})
}

type NotExistErr struct {
	Type string
	Id   string
}

func (error NotExistErr) Error() string {
	return fmt.Sprintf("%s %s not exist", error.Type, error.Id)
}

func NotExistError(t, id string) error {
	return errors.WithStack(NotExistErr{
		Type: t,
		Id:   id,
	})
}
