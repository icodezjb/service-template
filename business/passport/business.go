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
	userModule    module.UserModule
	safeModule    module.SafeModule
	recordModule  module.RecordModule
	messageModule module.MessageModule
}

func New() *Business {
	return &Business{
		userModule:    module.NewUserModule(definition.VersionLatest),
		messageModule: module.NewMessageModule(definition.VersionLatest),
	}
}

func (b *Business) Login(ctx context.Context, param LoginParam) (string, definition.Error) {
	requestId := util.GetContextStringValue(ctx, definition.MetadataRequestId)
	// 检查账号是否存在
	exist, err := b.userModule.AccountExists(ctx, param.Account)
	if err != nil {
		return "", definition.ErrModuleInvoke.WithSource("Login->userModule.AccountExists").WithError(err)
	}

	log.Println("[reqeust_id:%s] xxxxxxxx", requestId)

	if !exist {
		return "", definition.ErrAccountNotExist.WithSource(param.Account)
	}

	// 检查用户账户安全性，判断是否允许用户登录
	isSafe, reason, err := b.safeModule.CheckAccountState(ctx, param.Account)
	if err != nil {
		return "", definition.ErrModuleInvoke.WithSource("Login->userModule.CheckAccountState").WithError(err)
	}

	if !isSafe {
		return "", definition.ErrUserAccountNotSafe.WithSource(reason)
	}

	// 比较用户密码信息
	pwd, salt, err := b.userModule.GetUserPassword(ctx, param.Account)
	if err != nil {
		return "", definition.ErrModuleInvoke.WithSource("Login->userModule.GetUserPassword").WithError(err)
	}

	succ, token, err := b.userModule.CheckUserPassword(ctx, param.Password, pwd, salt)
	if err != nil {
		return "", definition.ErrUserLogin.WithSource("Login->userModule.CheckUserPassword").WithError(err)
	}
	if !succ {
		return "", definition.ErrUserLogin.WithSource("Login->userModule.CheckUserPassword").WithError(errors.New("用户密码不正确"))
	}

	// 记录用户登录行为
	if err = b.recordModule.RecordUserLogin(ctx, param.Account, param.ClientIP); err != nil {
		log.Println("xxxxxxxxx")
	}

	// 发送用户登录消息
	b.messageModule.AsyncSendUserLoginMessage(ctx, param.Account)

	return token, nil
}
