package passport

import (
	"context"

	"errors"

	"log"

	"github.com/buchenglei/service-template/common/definition"
	"github.com/buchenglei/service-template/common/util"
	"github.com/buchenglei/service-template/module"
)

type Business struct {
	userHandler    module.UserModule
	safeHandler    module.SafeModule
	recordHandler  module.RecordModule
	messageHandler module.MessageModule
}

func New() *Business {
	return &Business{
		userHandler:    module.NewUserModule(definition.VersionLatest),
		messageHandler: module.NewMessageModule(definition.VersionLatest),
	}
}

func (b *Business) Login(ctx context.Context, param LoginParam) (string, definition.Error) {
	requestId := util.GetContextStringValue(ctx, definition.MetadataRequestId)
	// 检查账号是否存在
	exist, err := b.userHandler.AccountExists(ctx, param.Account)
	if err != nil {
		return "", definition.ErrModuleInvoke.WithSource("Login->userHandler.AccountExists").WithError(err)
	}

	log.Println("[reqeust_id:%s] xxxxxxxx", requestId)

	if !exist {
		return "", definition.ErrAccountNotExist.WithSource(param.Account)
	}

	// 检查用户账户安全性，判断是否允许用户登录
	isSafe, reason, err := b.safeHandler.CheckAccountState(ctx, param.Account)
	if err != nil {
		return "", definition.ErrModuleInvoke.WithSource("Login->userHandler.CheckAccountState").WithError(err)
	}

	if !isSafe {
		return "", definition.ErrUserAccountNotSafe.WithSource(reason)
	}

	// 比较用户密码信息
	pwd, salt, err := b.userHandler.GetUserPassword(ctx, param.Account)
	if err != nil {
		return "", definition.ErrModuleInvoke.WithSource("Login->userHandler.GetUserPassword").WithError(err)
	}

	succ, token, err := b.userHandler.CheckUserPassword(ctx, param.Password, pwd, salt)
	if err != nil {
		return "", definition.ErrUserLogin.WithSource("Login->userHandler.CheckUserPassword").WithError(err)
	}
	if !succ {
		return "", definition.ErrUserLogin.WithSource("Login->userHandler.CheckUserPassword").WithError(errors.New("用户密码不正确"))
	}

	// 记录用户登录行为
	if err = b.recordHandler.RecordUserLogin(ctx, param.Account, param.ClientIP); err != nil {
		log.Println("xxxxxxxxx")
	}

	// 发送用户登录消息
	b.messageHandler.AsyncSendUserLoginMessage(ctx, param.Account)

	return token, nil
}
