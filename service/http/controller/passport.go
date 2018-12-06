package controller

import (
	"github.com/buchenglei/service-template/business"
	passportBusiness "github.com/buchenglei/service-template/business/passport"

	"context"
	"time"

	"github.com/buchenglei/service-template/common/definition"
	"github.com/gin-gonic/gin"
)

type PassportController struct {
	baseController

	passporthandler business.PassportBusiness
}

func NewPassportController() *PassportController {
	return &PassportController{
		// 初始化对应的业务流程
	}
}

func (passport *PassportController) Login(ctx *gin.Context) {
	// 获取请求参数
	params := struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{}

	err := ctx.BindJSON(&params)
	if err != nil {
		passport.response(definition.ErrParams.WithError(err), nil)
		return
	}

	// 定义该业务的请求超时时间
	businessCtx, cancel := context.WithTimeout(passport.newBaseContext(), 3*time.Second)
	defer cancel()

	// 加载对应的业务流程实现
	token, bErr := passport.passporthandler.Login(businessCtx, passportBusiness.LoginParam{
		Account:  params.Account,
		Password: params.Password,
		ClientIP: ctx.ClientIP(),
	})
	if err != nil {
		passport.response(bErr, nil)
		return
	}

	ctx.SetCookie("plu", token, 3600, "/", "", true, true)
	passport.response(nil, nil)
}

func (passport *PassportController) Register(ctx *gin.Context) {

}

func (passport *PassportController) Logout(ctx *gin.Context) {

}
