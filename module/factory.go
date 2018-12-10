package module

import (
	"github.com/buchenglei/service-template/common/definition"
	"github.com/buchenglei/service-template/module/message"
	"github.com/buchenglei/service-template/module/user"
)

func NewUserModule(version definition.Version) (m UserModule) {
	switch version {
	default:
		m = user.New()
	}

	return
}

func NewMessageModule(version definition.Version) (m MessageModule) {
	switch version {
	default:
		m = message.New()
	}

	return
}
