package controller

import (
	"net/http"

	"github.com/buchenglei/service-template/common/util"
)

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
)
