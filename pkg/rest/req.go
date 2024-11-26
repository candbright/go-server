package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func H(f func(c *gin.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := f(c)
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