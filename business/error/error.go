package error

import (
	"github.com/buchenglei/service-template/common/util"
)

// 业务层统一错误定义
var (
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
