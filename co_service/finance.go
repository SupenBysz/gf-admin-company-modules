// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package co_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type (
	IFdCurrency interface {
		// GetCurrencyByCode 根据货币代码查找货币(主键)
		GetCurrencyByCode(ctx context.Context, currencyCode string) (response *co_model.FdCurrencyRes, err error)
		// QueryCurrencyList 获取币种列表
		QueryCurrencyList(ctx context.Context, search *base_model.SearchParams) (*co_model.FdCurrencyListRes, error)
		// GetCurrencyByCnName 根据国家查找货币信息
		GetCurrencyByCnName(ctx context.Context, cnName string) (response *co_model.FdCurrencyRes, err error)
	}
)

var (
	localFdCurrency IFdCurrency
)

func FdCurrency() IFdCurrency {
	if localFdCurrency == nil {
		panic("implement not found for interface IFdCurrency, forgot register?")
	}
	return localFdCurrency
}

func RegisterFdCurrency(i IFdCurrency) {
	localFdCurrency = i
}
