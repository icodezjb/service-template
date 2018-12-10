package definition

import (
	"net/http"

	"errors"

	"github.com/buchenglei/service-template/common/util"
)

// 业务层统一错误定义
var (
	ErrSuccess = util.CommonError{
		Code:       0,
		HttpStatus: http.StatusOK,
		Message:    "success",
	}

	ErrParams = util.CommonError{
		Code:       100,
		HttpStatus: http.StatusBadRequest,
		Message:    "请求参数错误",
	}

	ErrModuleInvoke = util.CommonError{
		Code:    1000,
		Message: "系统错误",
	}

	ErrAccountNotExist = util.CommonError{
		Code:    1001,
		Message: "账号不存在",
	}

	ErrUserLogin = util.CommonError{
		Code:    1002,
		Message: "用户登录失败",
	}

	ErrUserAccountNotSafe = util.CommonError{
		Code:    1003,
		Message: "用户账号存在安全风险",
	}
)

// 内部统一错误标识
var (
	ErrUnknown = errors.New("miss a unknown error")
)
