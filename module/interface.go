package module

import "context"

// UserModule 定义用户模块需要实现的方法
type UserModule interface {
	AccountExists(ctx context.Context, account string) (bool, error)
	GetUserPassword(ctx context.Context, account string) (password, salt string, err error)
	CheckUserPassword(ctx context.Context, inputPwd, rawPwd, salt string) (isRight bool, token string, err error)
}

type SafeModule interface {
	CheckAccountState(ctx context.Context, account string) (isSafe bool, reason string, err error)
}

type RecordModule interface {
	RecordUserLogin(ctx context.Context, account, ip string) error
}

type MessageModule interface {
	AsyncSendUserLoginMessage(ctx context.Context, account string)
}
