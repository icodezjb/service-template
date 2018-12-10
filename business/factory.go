package business

import (
	"github.com/buchenglei/service-template/business/passport"
	"github.com/buchenglei/service-template/common/definition"
)

func NewPassportBusiness(version definition.Version) (b PassportBusiness) {
	switch version {
	case definition.Version_1:
		b = passport.NewBusinessV1()
	default:
		// latest version
		b = passport.New()
	}

	return
}
