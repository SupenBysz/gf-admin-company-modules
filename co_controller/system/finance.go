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

// QueryAccountRechargeView 查询财务账户充值记录
func (c *cSystemFinance) QueryAccountRechargeView(ctx context.Context, req *co_system_v1.QueryAccountRechargeViewReq) (*co_model.FdRechargeViewListRes, error) {
	return co_service.FdRechargeView().QueryAccountRecharge(ctx, &req.SearchParams)
}

// GetAccountRechargeViewById 根据ID查询财务账户充值记录
func (c *cSystemFinance) GetAccountRechargeViewById(ctx context.Context, req *co_system_v1.GetAccountRechargeViewByIdReq) (*co_model.FdRechargeViewRes, error) {
	return co_service.FdRechargeView().GetAccountRechargeById(ctx, req.Id)
}

// GetRechargeByCompanyId 根据公司ID获取充值记录
func (c *cSystemFinance) GetRechargeByCompanyId(ctx context.Context, req *co_system_v1.GetRechargeByCompanyIdReq) (*co_model.FdRechargeViewListRes, error) {
	return co_service.FdRechargeView().GetRechargeByCompanyId(ctx, req.Id)
}

// GetRechargeByAccountId 根据资金账户ID获取充值记录
func (c *cSystemFinance) GetRechargeByAccountId(ctx context.Context, req *co_system_v1.GetRechargeByAccountIdReq) (*co_model.FdRechargeViewListRes, error) {
	return co_service.FdRechargeView().GetRechargeByAccountId(ctx, req.Id)
}

// GetRechargeByUserId 根据用户ID获取充值记录
func (c *cSystemFinance) GetRechargeByUserId(ctx context.Context, req *co_system_v1.GetRechargeByUserIdReq) (*co_model.FdRechargeViewListRes, error) {
	return co_service.FdRechargeView().GetRechargeByUserId(ctx, req.Id)
}
