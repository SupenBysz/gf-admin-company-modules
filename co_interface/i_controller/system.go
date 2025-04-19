package i_controller

import (
	"context"

	"github.com/SupenBysz/gf-admin-company-modules/api/co_system_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type ISystemFinance interface {
	GetCurrencyByCode(ctx context.Context, req *co_system_v1.GetCurrencyByCodeReq) (*co_model.FdCurrencyRes, error)
	QueryCurrencyList(ctx context.Context, req *co_system_v1.QueryCurrencyListReq) (*co_model.FdCurrencyListRes, error)
}
