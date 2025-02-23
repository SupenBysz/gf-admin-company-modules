package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type FdCurrencyRes struct {
	co_entity.FdCurrency
}

type FdCurrencyListRes base_model.CollectRes[FdCurrencyRes]

func (m *FdCurrencyRes) Data() *FdCurrencyRes {
	return m
}

type IFdCurrencyRes interface {
	Data() *FdCurrencyRes
}
