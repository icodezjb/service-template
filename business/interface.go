package business

import (
	"context"

	"github.com/buchenglei/service-template/business/passport"
	"github.com/buchenglei/service-template/common/definition"
)

type PassportBusiness interface {
	Login(context.Context, passport.LoginParam) (string, definition.Error)
}
