package controller

import (
	"github.com/buchenglei/service-template/common/util"

	"context"

	"github.com/buchenglei/service-template/common/definition"
	"github.com/gin-gonic/gin"
)

type responseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type baseController struct {
	ctx *gin.Context
}

func (base *baseController) init(ctx *gin.Context) {
	base.ctx = ctx
}

func (base *baseController) response(err util.Error, data interface{}) {
	if err == nil {
		err = definition.ErrSuccess
	}

	base.ctx.JSON(err.ErrHttpStatus(), responseData{
		Code: err.ErrCode(),
		Msg:  err.ErrMessage(),
		Data: data,
	})
}

func (base *baseController) newBaseContext() context.Context {
	ctx := context.Background()

	// 携带基础的上下文信息
	ctx = context.WithValue(ctx, definition.FieldRequestId, "xxxxxxxx")

	return ctx
}
