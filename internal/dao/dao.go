package dao

import (
	"context"
	"github.com/pkg/errors"
)

type Dao interface {
	Scan(ctx context.Context, model interface{}, req *ScanReq) (*ScanRsp, error)
	Create(ctx context.Context, value interface{}) error
	Delete(ctx context.Context, value interface{}) (*ExecRsp, error)
}

func NewDao(driver string) (Dao, error) {
	switch driver {
	case "native":
		return &NativeDao{}, nil
	default:
		return nil, errors.New("not support driver")
	}
}
