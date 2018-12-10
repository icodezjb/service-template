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

	passportHandler business.PassportBusiness
}

func NewPassportController() *PassportController {
	return &PassportController{
		// 初始化对应的业务流程
		// 传空则表示默认为最新的处理接口
		// controller中统一使用business下的factory.go中的NewXXXXXXX方法创建业务对象
		// 避免使用实际的business/passport目录下的创建方法
		// 避免调用和具体实现耦合在一起
		// business目录下的interface.go factory.go 就相当于中间层，将调用方service层与实现方business层解耦
		passportHandler: business.NewPassportBusiness(definition.VersionLatest),
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

	/*************************************************************************************/

	// 加载对应的业务流程实现
	// 默认使用最新的版本
	token, bErr := passport.passportHandler.Login(businessCtx, passportBusiness.LoginParam{
		Account:  params.Account,
		Password: params.Password,
		ClientIP: ctx.ClientIP(),
	})
	if err != nil {
		passport.response(bErr, nil)
		return
	}

	// Example 创建对应版本的handler对应版本的business
	passportHandlerForVersion := business.NewPassportBusiness(definition.Version_1)
	token, bErr = passportHandlerForVersion.Login(businessCtx, passportBusiness.LoginParam{
		Account:  params.Account,
		Password: params.Password,
		ClientIP: ctx.ClientIP(),
	})
	if err != nil {
		passport.response(bErr, nil)
		return
	}

	/*************************************************************************************/

	ctx.SetCookie("plu", token, 3600, "/", "", true, true)
	passport.response(nil, nil)
}

func (passport *PassportController) Register(ctx *gin.Context) {

}

func (passport *PassportController) Logout(ctx *gin.Context) {

}
