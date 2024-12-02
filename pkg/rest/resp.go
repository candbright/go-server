package rest

import (
	"fmt"
	"net/http"
)

const UnknownErr = -1

type HTTPError struct {
	HttpStatus int         `json:"-"`
	Code       int         `json:"code,omitempty"`
	Err        error       `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (err HTTPError) Error() string {
	if err.Err == nil {
		return ""
	}
	return fmt.Sprintf("err code:%d\nstatus code:%d\nerr message:%s", err.Code, err.HttpStatus, err.Err)
}

func (err HTTPError) WithCode(code int) HTTPError {
	return HTTPError{
		HttpStatus: err.HttpStatus,
		Code:       code,
		Err:        err.Err,
		Data:       err.Data,
	}
}

func Json(data interface{}) HTTPError {
	return HTTPError{
		HttpStatus: http.StatusOK,
		Data:       data,
	}
}

func ErrorWithStatus(err error, status int) HTTPError {
	return HTTPError{
		HttpStatus: status,
		Code:       UnknownErr,
		Err:        err,
	}
}

type CodeError struct {
	Code int   `json:"code"`
	Err  error `json:"error,omitempty"`
}

func (err CodeError) Error() string {
	if err.Err == nil {
		return ""
	}
	return fmt.Sprintf("err code:%d\nerr message:%s", err.Code, err.Err)
}

func ErrorWithCode(err error, code int) CodeError {
	return CodeError{
		Code: code,
		Err:  err,
	}
}

func (err CodeError) WithStatus(status int) HTTPError {
	return HTTPError{
		HttpStatus: status,
		Code:       err.Code,
		Err:        err.Err,
	}
}
