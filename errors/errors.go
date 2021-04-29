/**
 *@Description
 *@ClassName errors
 *@Date 2021/3/12 下午1:10
 *@Author ckhero
 */

package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/sd/lb"
	"google.golang.org/grpc/status"
)

type StatusError struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Reason  string `json:"reason"`
}

type StatusError2 struct {
	Code int32  `json:"code"`
}


func (e *StatusError) Is(target error) bool {
	err, ok := target.(*StatusError)
	if ok {
		return e.Code == err.Code
	}
	return false
}

func (e *StatusError) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(data)
	//return fmt.Sprintf("error: code = %d msg = %s reason + %s", e.Code, e.Msg, e.Reason)
}

func Error(code int32, reason, msg string) error {
	return &StatusError{
		Code: code,
		Msg:  msg,
		Reason: reason,
	}
}

func Errorf(code int32, reason, format string, a ...interface{}) error {
	return Error(code, reason, fmt.Sprintf(format, a...))
}

func Code(err error) int32 {
	if err == nil {
		return 0 // ok
	}
	if se := new(StatusError); errors.As(err, &se) {
		return se.Code
	}
	return 2 // unknown
}

func As(err error, target interface{}) bool {
	switch err.(type) {
	case lb.RetryError:
		err = err.(lb.RetryError).Final
		if ee, ok := status.FromError(err.(error)); ok {
			tmp :=ParseError(ee.Message())
			if tmp != nil {
				err = tmp
			}
		} else {
		}
	}
	return errors.As(err, target)
}

func ParseError(data string) *StatusError {
	var terr StatusError
	err := json.Unmarshal([]byte(data), &terr)
	if err != nil {
		return nil
	}
	return &terr
}
