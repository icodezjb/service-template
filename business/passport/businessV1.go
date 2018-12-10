package passport

import (
	"context"

	"github.com/buchenglei/service-template/common/util"
)

type BusinessV1 struct{}

func NewBusinessV1() *BusinessV1 {
	return &BusinessV1{}
}

func (*BusinessV1) Login(context.Context, LoginParam) (string, util.Error) {
	panic("implement me")
}
