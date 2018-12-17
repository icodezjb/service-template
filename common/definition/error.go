package definition

import (
	"errors"

	"fmt"
)

// 业务层统一错误定义
var (
	ErrSuccess = CommonError{
		Code:    0,
		Message: "success",
	}

	ErrParams = CommonError{
		Code:    100,
		Message: "请求参数错误",
	}

	ErrModuleInvoke = CommonError{
		Code:    1000,
		Message: "系统错误",
	}

	ErrAccountNotExist = CommonError{
		Code:    1001,
		Message: "账号不存在",
	}

	ErrUserLogin = CommonError{
		Code:    1002,
		Message: "用户登录失败",
	}

	ErrUserAccountNotSafe = CommonError{
		Code:    1003,
		Message: "用户账号存在安全风险",
	}
)

// 内部统一错误标识
var (
	ErrUnknown = errors.New("miss a unknown error")
)

/********** 定义统一错误类型 **********/

type Error interface {
	error

	ErrCode() int
	ErrMessage() string
}

type CommonError struct {
	Code    int
	Message string

	source string
	err    error
}

func (e CommonError) WithError(err error) CommonError {
	e.err = err
	return e
}

func (e CommonError) WithSource(source string, params ...interface{}) CommonError {
	if len(params) == 0 {
		e.source = source
	} else {
		e.source = fmt.Sprintf(source, params...)
	}

	return e
}

/********** 实现Error接口 **********/

func (e CommonError) ErrCode() int {
	return e.Code
}

func (e CommonError) ErrMessage() string {
	return e.Message
}

func (e CommonError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return fmt.Sprintf("code: %d, message: %s, source:%s, raw error:%v", e.Code, e.Message, e.source, e.err)
}
