package account

import (
	"d2din.com/internal/v1/service/account"
	"d2din.com/internal/v1/service/company"
	model "d2din.com/internal/v1/structure/accounts"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	Account account.Service
	Company company.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		Account: account.New(db),
		Company: company.New(db),
	}
}
