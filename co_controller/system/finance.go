package system

import (
	"context"

	"github.com/SupenBysz/gf-admin-company-modules/api/co_system_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
)

type cSystemFinance struct{}

var SystemFinance i_controller.ISystemFinance = &cSystemFinance{}

// GetCurrencyByCode 根据货币代码查找货币(主键)
func (c *cSystemFinance) GetCurrencyByCode(ctx context.Context, req *co_system_v1.GetCurrencyByCodeReq) (*co_model.FdCurrencyRes, error) {
	return co_service.FdCurrency().GetCurrencyByCode(ctx, req.CurrencyCode)
}

// QueryCurrencyList 获取币种列表
func (c *cSystemFinance) QueryCurrencyList(ctx context.Context, req *co_system_v1.QueryCurrencyListReq) (*co_model.FdCurrencyListRes, error) {
	return co_service.FdCurrency().QueryCurrencyList(ctx, &req.SearchParams)
}
