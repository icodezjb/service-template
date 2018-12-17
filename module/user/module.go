package user

import (
	"context"

	"errors"

	"github.com/buchenglei/service-template/common/util"
	"github.com/buchenglei/service-template/data"
	"github.com/buchenglei/service-template/data/micro"
	"github.com/buchenglei/service-template/data/mysql"
)

type Module struct {
	redisData data.RedisData
}

func New() *Module {
	return &Module{
		redisData: data.NewRedisData(),
	}
}

func (m *Module) AccountExists(ctx context.Context, account string) (exist bool, err error) {
	// 如果在指定的超时时间内没有完成方法的执行，那么就会返回超时的报错
	// 这里只是做一个示例，如果需要更细力度的单独控制redis超时，或是mysql超时
	// 可以再将ctx向下传递，不少第三方库都支持ctx，这里仅仅只是demo
	err = util.DoWithTimeout(ctx, func() error {
		// 先检查redis中是否存在
		var innerErr error
		exist, innerErr = m.redisData.UserAccountExist(account)
		if exist {
			return innerErr
		}

		exist, innerErr = mysql.UserExist(0, account)
		return innerErr
	})

	return
}

func (*Module) GetUserPassword(ctx context.Context, account string) (password, salt string, err error) {
	panic("implement me")
}

func (*Module) CheckUserPassword(ctx context.Context, inputPwd, rawPwd, salt string) (isRight bool, token string, err error) {
	// 由于grpc本身会处理ctx，这里只需要传进去就可以了
	token, err = micro.CheckUserPassword(ctx, inputPwd, rawPwd, salt)
	if err != nil {
		return false, "", err
	}

	if token == "" {
		return false, "", errors.New("unknown")
	}

	return true, token, nil
}
