package controller

import (
	"github.com/buchenglei/service-template/common/util"

	"context"

	"log"

	"time"

	"github.com/buchenglei/service-template/common/definition"
	lib "github.com/buchenglei/service-template/common/lib"
	"github.com/buchenglei/service-template/config"
	"github.com/gin-gonic/gin"
)

type baseController struct {
	ctx *gin.Context
}

func (base *baseController) init(c *gin.Context) context.Context {
	base.ctx = c

	// 初始化新的context
	ctx := context.Background()

	// 携带请求的上下文信息
	ctx = context.WithValue(ctx, definition.MetadataRequestId, lib.GenRequestId())
	ctx = context.WithValue(ctx, definition.MetadataRequestSource, definition.RequestSourceThirdPart)
	ctx = context.WithValue(ctx, definition.MetadataRequestProtocol, definition.RequestProtocolHTTP)
	ctx = context.WithValue(ctx, definition.MetadataClientVersion, "v1")
	ctx = context.WithValue(ctx, definition.MetadataTimeReciveRequest, time.Now().UnixNano())

	log.Println("[request_id:%s] Reqeust start with params: %+v", "xxxxxx", c.Request)

	return ctx
}

func (base *baseController) response(ctx context.Context, err definition.Error, data interface{}) {
	if err == nil {
		err = definition.ErrSuccess
	}

	reqeustId := util.GetContextStringValue(ctx, definition.MetadataRequestId)
	start := util.GetContextInt64Value(ctx, definition.MetadataTimeReciveRequest)

	log.Println("[request_id:%s] Request end with response(reqeust_time:%dms): error:%s, data:%+v", reqeustId, util.CalRequestTime(start), err.Error(), data)

	response := make(map[string]interface{})
	response["code"] = err.ErrCode()
	response["message"] = err.ErrMessage()

	if data == nil {
		// 这么处理就是在空数据返回时，前端看到的是 "data": {}, 而不是 "data": null
		response["data"] = struct{}{}
	} else {
		response["data"] = data
	}

	// 在非正式环境需要额外输出一些调试信息
	if config.ServiceEnv != definition.EnvPro {
		response["request_id"] = reqeustId
	}

	base.ctx.JSON(200, response)
}
