package controller

import (
	"github.com/buchenglei/service-template/business"
	passportBusinessPkg "github.com/buchenglei/service-template/business/passport"

	"context"
	"time"

	"github.com/buchenglei/service-template/common/definition"
	"github.com/buchenglei/service-template/common/util"
	"github.com/gin-gonic/gin"
)

type PassportController struct {
	baseController

	passportBusiness business.PassportBusiness
}

func NewPassportController() *PassportController {
	return &PassportController{
		// 初始化对应的业务流程
		// 传空则表示默认为最新的处理接口
		// controller中统一使用business下的factory.go中的NewXXXXXXX方法创建业务对象
		// 避免使用实际的business/passport目录下的创建方法
		// 避免调用和具体实现耦合在一起
		// business目录下的interface.go factory.go 就相当于中间层，将调用方service层与实现方business层解耦
		passportBusiness: business.NewPassportBusiness(definition.VersionLatest),
	}
}

func (passport *PassportController) Login(c *gin.Context) {
	// 定义该业务的请求超时时间
	businessCtx, cancel := context.WithTimeout(passport.init(c), 3*time.Second)
	defer cancel()

	// 获取请求参数
	params := struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{}

	err := c.BindJSON(&params)
	if err != nil {
		passport.response(businessCtx, definition.ErrParams.WithError(err), nil)
		return
	}

	/*************************************************************************************/

	// 加载对应的业务流程实现
	// 默认使用最新的版本
	// Example 创建对应版本的handler对应版本的business
	version := definition.Version(util.GetContextStringValue(businessCtx, definition.MetadataClientVersion))

	passportHandler := passport.passportBusiness
	if version != definition.VersionLatest {
		passportHandler = business.NewPassportBusiness(version)
	}

	token, bErr := passportHandler.Login(businessCtx, passportBusinessPkg.LoginParam{
		Account:  params.Account,
		Password: params.Password,
		ClientIP: c.ClientIP(),
	})
	if err != nil {
		passport.response(businessCtx, bErr, nil)
		return
	}

	/*************************************************************************************/

	c.SetCookie("plu", token, 3600, "/", "", true, true)
	passport.response(businessCtx, nil, nil)
}

func (passport *PassportController) Register(ctx *gin.Context) {

}

func (passport *PassportController) Logout(ctx *gin.Context) {

}
