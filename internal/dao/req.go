package dao

import (
	"context"
	"github.com/candbright/go-server/pkg/rest"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"sync"
)

var daoInstance Dao

func Init(driver string) error {
	dao, err := NewDao(driver)
	if err != nil {
		return err
	}
	daoInstance = dao
	return nil
}

type GetReq struct {
	ID int `json:"id" binding:"required"`
}

type Where struct {
	K string      `json:"k" binding:"required"`
	O string      `json:"o" binding:"required"`
	V interface{} `json:"v" binding:"required"`
}

type ScanReq struct {
	Page   int64    `json:"page" binding:"required,min=1"`
	Size   int64    `json:"size" binding:"required,min=1,max=100000"`
	Wheres []*Where `json:"wheres" binding:"omitempty"`
}

type BatchReq[T any] struct {
	Items []*T `json:"items" binding:"required"`
}

var modelCache sync.Map

type MappingHandler func(value reflect.Value) reflect.Value

type Mapping = map[string]MappingHandler

func TransferItems[O any, T any](items []*O) []*T {
	if len(items) == 0 {
		return []*T{}
	}

	out := make([]*T, len(items))
	for i, item := range items {
		out[i] = CopyContent[O, T](item)
	}

	return out
}

func CopyContent[O any, T any](src *O) *T {
	srcV := reflect.ValueOf(src)
	for srcV.Kind() == reflect.Ptr {
		srcV = srcV.Elem()
	}
	dst := new(T)
	dstV := reflect.ValueOf(dst)
	for dstV.Kind() == reflect.Ptr {
		dstV = dstV.Elem()
	}

	srcT := srcV.Type()
	dstT := dstV.Type()

	srcTMapping := getMapping(srcT)
	dstTMapping := getMapping(dstT)

	for srcName, srcHandler := range srcTMapping {
		dstHandler, ok := dstTMapping[srcName]
		if !ok {
			continue
		}
		dstHandler(dstV).Set(srcHandler(srcV))
	}
	return dst
}

func getMapping(t reflect.Type) Mapping {
	var mapping Mapping
	cache, loaded := modelCache.Load(t.String())
	if loaded {
		mapping = cache.(Mapping)
	} else {
		mapping = Mapping{}
		for i := 0; i < t.NumField(); i++ {
			fi := i
			field := t.Field(i)
			fType := field.Type
			if fType.Kind() == reflect.Pointer {
				panic("not support pointer")
			}
			if field.Type.Kind() == reflect.Struct {
				if field.Anonymous {
					fieldMapping := getMapping(fType)
					for k, handler := range fieldMapping {
						fHandler := handler
						mapping[k] = func(value reflect.Value) reflect.Value {
							return fHandler(value.Field(fi))
						}
					}
				}
			}

			name := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
			if name == "" {
				continue
			}
			mapping[name] = func(value reflect.Value) reflect.Value {
				return value.Field(fi)
			}
		}
		modelCache.Store(t.String(), mapping)
	}

	return mapping
}

func Get[M any](c *gin.Context) {
	rest.H(func(c *gin.Context, ctx context.Context) error {
		req := new(GetReq)
		err := c.ShouldBind(req)
		if err != nil {
			return err
		}
		//TODO
		return rest.Json(req)
	})(c)
}

func Scan[M any](c *gin.Context) {
	rest.H(func(c *gin.Context, ctx context.Context) error {
		req := new(ScanReq)
		err := c.ShouldBind(req)
		if err != nil {
			return err
		}
		resp, err := daoInstance.Scan(ctx, new(M), req)
		if err != nil {
			return err
		}
		return rest.Json(resp)
	})(c)
}

func Create[M, I any](c *gin.Context) {
	rest.H(func(c *gin.Context, ctx context.Context) error {
		req := new(BatchReq[I])
		err := c.ShouldBind(req)
		if err != nil {
			return err
		}
		items := TransferItems[I, M](req.Items)
		err = daoInstance.Create(ctx, items)
		if err != nil {
			return err
		}
		return nil
	})(c)
}
