package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func H(f func(c *gin.Context, ctx context.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		timeout := 10 * time.Second
		t, err := strconv.ParseInt(c.Query("timeout"), 10, 64)
		if err == nil && t > 0 {
			timeout = time.Duration(t) * time.Second
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err = f(c, ctx)
		if err == nil {
			c.JSON(http.StatusNoContent, nil)
			return
		}
		var resp HTTPError
		switch e := err.(type) {
		case HTTPError:
			ce, ok := e.Err.(CodeError)
			if ok {
				resp = HTTPError{
					HttpStatus: e.HttpStatus,
					Code:       ce.Code,
					Err:        ce.Err,
				}
			} else {
				resp = e
			}
		case CodeError:
			resp = HTTPError{
				HttpStatus: http.StatusBadRequest,
				Code:       e.Code,
				Err:        e.Err,
			}
		default:
			resp = HTTPError{
				HttpStatus: http.StatusBadRequest,
				Code:       UnknownErr,
				Err:        err,
			}
		}
		if resp.Err != nil {
			resp.Err = errors.Cause(resp.Err)
			c.JSON(resp.HttpStatus, gin.H{
				"code":    resp.Code,
				"message": resp.Err.Error(),
			})
		} else {
			c.JSON(resp.HttpStatus, resp.Data)
		}
	}
}
