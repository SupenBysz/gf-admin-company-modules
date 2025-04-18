package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/kconv"
)

type FinanceController[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
] struct {
	i_controller.IFinance[
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]
}

func Finance[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
] {
	return &FinanceController[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]{
		IFinance: co_controller.Finance(modules),
	}
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountBalance(ctx context.Context, req *co_v1.GetAccountBalanceReq) (api_v1.Int64Res, error) {
	return c.IFinance.GetAccountBalance(ctx, &req.GetAccountBalanceReq)
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) InvoiceRegister(ctx context.Context, req *co_v1.CreateInvoiceReq) (*co_model.FdInvoiceRes, error) {
	ret, err := c.IFinance.InvoiceRegister(ctx, &req.CreateInvoiceReq)
	return ret.Data(), err
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryInvoice(ctx context.Context, req *co_v1.QueryInvoiceReq) (*co_model.FdInvoiceListRes, error) {
	ret, err := c.IFinance.QueryInvoice(ctx, &req.QueryInvoiceReq)
	return kconv.Struct(ret, &co_model.FdInvoiceListRes{}), err
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) DeletesFdInvoiceById(ctx context.Context, req *co_v1.DeleteInvoiceByIdReq) (api_v1.BoolRes, error) {
	return c.IFinance.DeletesFdInvoiceById(ctx, &req.DeleteInvoiceByIdReq)
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) InvoiceDetailRegister(ctx context.Context, req *co_v1.CreateInvoiceDetailReq) (*co_model.FdInvoiceDetailRes, error) {
	ret, err := c.IFinance.InvoiceDetailRegister(ctx, &req.CreateInvoiceDetailReq)
	return ret.Data(), err
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryInvoiceDetailList(ctx context.Context, req *co_v1.QueryInvoiceDetailListReq) (*co_model.FdInvoiceDetailListRes, error) {
	ret, err := c.IFinance.QueryInvoiceDetailList(ctx, &req.QueryInvoiceDetailListReq)
	return kconv.Struct(ret, &co_model.FdInvoiceDetailListRes{}), err
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) MakeInvoiceDetailReq(ctx context.Context, req *co_v1.MakeInvoiceDetailReq) (api_v1.BoolRes, error) {
	return c.IFinance.MakeInvoiceDetailReq(ctx, &req.MakeInvoiceDetailReq)
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) AuditInvoiceDetail(ctx context.Context, req *co_v1.AuditInvoiceDetailReq) (api_v1.BoolRes, error) {
	return c.IFinance.AuditInvoiceDetail(ctx, &req.AuditInvoiceDetailReq)

}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) BankCardRegister(ctx context.Context, req *co_v1.BankCardRegisterReq) (*co_model.FdBankCardRes, error) {
	ret, err := c.IFinance.BankCardRegister(ctx, &req.BankCardRegisterReq)
	return ret.Data(), err
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) DeleteBankCard(ctx context.Context, req *co_v1.DeleteBankCardReq) (api_v1.BoolRes, error) {
	return c.IFinance.DeleteBankCard(ctx, &req.DeleteBankCardReq)
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryBankCardList(ctx context.Context, req *co_v1.QueryBankCardListReq) (*co_model.FdBankCardListRes, error) {
	ret, err := c.IFinance.QueryBankCardList(ctx, &req.QueryBankCardListReq)

	return kconv.Struct(ret, &co_model.FdBankCardListRes{}), err
}

// GetAccountDetail 查看财务账号明细
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountDetail(ctx context.Context, req *co_v1.GetAccountDetailReq) (ITFdAccountRes, error) {
	ret, err := c.IFinance.GetAccountDetail(ctx, &req.GetAccountDetailReq)

	return ret, err
}

// UpdateAccountIsEnabled 修改财务账号启用状态
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateAccountIsEnabled(ctx context.Context, req *co_v1.UpdateAccountIsEnabledReq) (api_v1.BoolRes, error) {
	ret, err := c.IFinance.UpdateAccountIsEnabled(ctx, &req.UpdateAccountIsEnabledReq)

	return ret, err
}

// UpdateAccountLimitState 修改财务账号限制状态
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateAccountLimitState(ctx context.Context, req *co_v1.UpdateAccountLimitStateReq) (api_v1.BoolRes, error) {
	ret, err := c.IFinance.UpdateAccountLimitState(ctx, &req.UpdateAccountLimitStateReq)

	return ret, err
}

// GetAccountDetailById 根据财务账号id查询账单金额明细统计记录
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountDetailById(ctx context.Context, req *co_v1.GetAccountDetailByAccountIdReq) (*co_model.FdAccountDetailRes, error) {
	ret, err := c.IFinance.GetAccountDetailById(ctx, &req.GetAccountDetailByAccountIdReq)

	return ret, err
}

// UpdateAccountBalance 财务账号金额冲正
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateAccountBalance(ctx context.Context, req *co_v1.UpdateAccountBalanceReq) (api_v1.Int64Res, error) {
	ret, err := c.IFinance.UpdateAccountBalance(ctx, &req.UpdateAccountBalanceReq)

	return ret, err
}

// GetCurrencyByCode 财务账号金额冲正
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetCurrencyByCode(ctx context.Context, req *co_v1.GetCurrencyByCodeReq) (*co_model.FdCurrencyRes, error) {
	return c.IFinance.GetCurrencyByCode(ctx, &req.GetCurrencyByCodeReq)
}

// QueryCurrencyList 获取币种列表
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryCurrencyList(ctx context.Context, req *co_v1.QueryCurrencyListReq) (*co_model.FdCurrencyListRes, error) {
	ret, err := c.IFinance.QueryCurrencyList(ctx, &req.QueryCurrencyListReq)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	resp := co_model.FdCurrencyListRes{}

	if ret == nil {
		return &resp, nil
	}

	_ = gconv.Struct(ret, &resp)
	return &resp, nil
}

// QueryAccountBills  根据财务账号id查询账单
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryAccountBills(ctx context.Context, req *co_v1.QueryAccountBillsReq) (*base_model.CollectRes[ITFdAccountBillRes], error) {
	return c.IFinance.QueryAccountBills(ctx, &req.QueryAccountBillsReq)
}

//
//// Increment 收入
//func (c *FinanceController[
//	ITCompanyRes,
//	ITEmployeeRes,
//	ITTeamRes,
//	ITFdAccountRes,
//	ITFdAccountBillRes,
//	ITFdBankCardRes,
//	ITFdCurrencyRes,
//	ITFdInvoiceRes,
//	ITFdInvoiceDetailRes,
//]) Increment(ctx context.Context, req *co_v1.IncrementReq) (api_v1.BoolRes, error) {
//	ret, err := c.IFinance.Increment(ctx, &req.IncrementReq)
//
//	return ret, err
//}
//
//// Decrement 支出
//func (c *FinanceController[
//	ITCompanyRes,
//	ITEmployeeRes,
//	ITTeamRes,
//	ITFdAccountRes,
//	ITFdAccountBillRes,
//	ITFdBankCardRes,
//	ITFdCurrencyRes,
//	ITFdInvoiceRes,
//	ITFdInvoiceDetailRes,
//]) Decrement(ctx context.Context, req *co_v1.DecrementReq) (api_v1.BoolRes, error) {
//	ret, err := c.IFinance.Decrement(ctx, &req.DecrementReq)
//
//	return ret, err
//}

// SetAccountAllowExceed 设置财务账号是否允许存在负余额
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetAccountAllowExceed(ctx context.Context, req *co_v1.SetAccountAllowExceedReq) (api_v1.BoolRes, error) {
	ret, err := c.IFinance.SetAccountAllowExceed(ctx, &req.SetAccountAllowExceedReq)

	return ret, err
}
