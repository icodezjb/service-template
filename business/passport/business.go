package passport

import (
	"context"

	"errors"

	"log"

	"github.com/buchenglei/service-template/business/error"
	"github.com/buchenglei/service-template/common/definition"
	"github.com/buchenglei/service-template/common/util"
	"github.com/buchenglei/service-template/module"
	userModule "github.com/buchenglei/service-template/module/user"
)

type Business struct {
	userHandler    module.UserModule
	safeHandler    module.SafeModule
	recordHandler  module.RecordModule
	messageHandler module.MessageModule
}

func New() *Business {
	return &Business{
		userHandler: userModule.New(),
	}
}

func (b *Business) Login(ctx context.Context, param LoginParam) (string, util.Error) {
	requestId := util.GetContextStringValue(ctx, definition.FieldRequestId)
	// 检查账号是否存在
	exist, err := b.userHandler.AccountExists(ctx, param.Account)
	if err != nil {
		return "", error.ErrModuleInvoke.WithSource("(request_id: %s)Login->userHandler.AccountExists", requestId).WithError(err)
	}

	if !exist {
		return "", error.ErrAccountNotExist.WithSource(param.Account)
	}

	// 检查用户账户安全性，判断是否允许用户登录
	isSafe, reason, err := b.safeHandler.CheckAccountState(ctx, param.Account)
	if err != nil {
		return "", error.ErrModuleInvoke.WithSource("(request_id: %s)Login->userHandler.CheckAccountState", requestId).WithError(err)
	}

	if !isSafe {
		return "", error.ErrUserAccountNotSafe.WithSource(reason)
	}

	// 比较用户密码信息
	pwd, salt, err := b.userHandler.GetUserPassword(ctx, param.Account)
	if err != nil {
		return "", error.ErrModuleInvoke.WithSource("(request_id: %s)Login->userHandler.GetUserPassword", requestId).WithError(err)
	}

	succ, token, err := b.userHandler.CheckUserPassword(ctx, param.Password, pwd, salt)
	if err != nil {
		return "", error.ErrUserLogin.WithSource("(request_id: %s)Login->userHandler.CheckUserPassword", requestId).WithError(err)
	}
	if !succ {
		return "", error.ErrUserLogin.WithSource("(request_id: %s)Login->userHandler.CheckUserPassword", requestId).WithError(errors.New("用户密码不正确"))
	}

	// 记录用户登录行为
	if err = b.recordHandler.RecordUserLogin(ctx, param.Account, param.ClientIP); err != nil {
		log.Println("xxxxxxxxx")
	}

	// 发送用户登录消息
	if err = b.messageHandler.SendUserLoginMessage(ctx, param.Account); err != nil {
		log.Println("xxxxxxxxx")
	}

	return token, nil
}
