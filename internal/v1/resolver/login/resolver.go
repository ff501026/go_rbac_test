package login

import (
	"d2din.com/internal/v1/service/account"
	"d2din.com/internal/v1/service/jwe"
	jweModel "d2din.com/internal/v1/structure/jwe"
	loginsModel "d2din.com/internal/v1/structure/logins"
	"gorm.io/gorm"
)

type Resolver interface {
	Web(input *loginsModel.Login) interface{}
	Refresh(input *jweModel.Refresh) interface{}
}

type resolver struct {
	Account account.Service
	JWE     jwe.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		Account: account.New(db),
		JWE:     jwe.New(),
	}
}
