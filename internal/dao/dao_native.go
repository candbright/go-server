package dao

import (
	"context"
	"github.com/candbright/go-server/pkg/config"
	"github.com/candbright/go-server/pkg/dw"
	"github.com/candbright/go-ssh/ssh"
	"github.com/pkg/errors"
	"path/filepath"
	"reflect"
)

type NativeDao struct {
	cache map[string]*dw.DataWriter[interface{}]
}

func (d *NativeDao) GetDw(model interface{}) (*dw.DataWriter[interface{}], error) {
	if d.cache == nil {
		d.cache = make(map[string]*dw.DataWriter[interface{}])
	}
	key := reflect.TypeOf(model).Elem().Name()
	if v, ok := d.cache[key]; ok {
		return v, nil
	}
	session, err := ssh.NewSession()
	if err != nil {
		return nil, err
	}
	savePath := filepath.Join(config.Global.Get("dao.path"), key)
	newDw := dw.Default[interface{}](session, savePath)
	d.cache[key] = newDw
	return newDw, nil
}

func (d *NativeDao) Scan(ctx context.Context, model interface{}, req *ScanReq) (*ScanRsp, error) {
	dataWriter, err := d.GetDw(model)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(dataWriter.Data)
	if v.Kind() == reflect.Slice {
		data := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			data[i] = v.Index(i).Interface()
		}
		return &ScanRsp{
			Items: data,
			Page:  req.Page,
			Size:  req.Size,
			Total: int64(len(data)),
		}, nil
	} else {
		return &ScanRsp{
			Items: make([]interface{}, 0),
			Page:  req.Page,
			Size:  req.Size,
			Total: 0,
		}, nil
	}
}

func (d *NativeDao) Create(ctx context.Context, value interface{}) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice {
		return errors.New("value is not a slice")
	}
	if v.Len() == 0 {
		return nil
	}
	dataWriter, err := d.GetDw(v.Index(0).Interface())
	if err != nil {
		return err
	}
	data := dataWriter.Data
	if slice, ok := data.([]interface{}); ok {
		for i := 0; i < v.Len(); i++ {
			dataWriter.Data = append(slice, v.Index(i).Interface())
		}
		err = dataWriter.Write()
		if err != nil {
			return err
		}
		return nil
	} else {
		dataWriter.Data = value
		err = dataWriter.Write()
		if err != nil {
			return err
		}
		return nil
	}
}

func (d *NativeDao) Delete(ctx context.Context, value interface{}) (*ExecRsp, error) {
	dataWriter, err := d.GetDw(value)
	if err != nil {
		return nil, err
	}
	data := dataWriter.Data
	if slice, ok := data.([]interface{}); ok {
		for i := 0; i < len(slice); i++ {
			equal := reflect.DeepEqual(value, slice[i])
			if equal {
				dataWriter.Data = append(slice[:i], slice[i+1:]...)
				err = dataWriter.Write()
				if err != nil {
					return nil, err
				}
				return &ExecRsp{
					RowsAffected: 1,
				}, nil
			}
		}

	}
	return &ExecRsp{
		RowsAffected: 0,
	}, nil
}
