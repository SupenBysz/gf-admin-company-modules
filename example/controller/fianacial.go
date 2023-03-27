package controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/utility/kconv"
)

type FinancialController[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	i_controller.IFinancial[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
}

func Financial[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
] {
	return &FinancialController[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		IFinancial: co_controller.Financial(modules),
	}
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccountBalance(ctx context.Context, req *co_v1.GetAccountBalanceReq) (api_v1.Int64Res, error) {
	return c.IFinancial.GetAccountBalance(ctx, &req.GetAccountBalanceReq)
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) InvoiceRegister(ctx context.Context, req *co_v1.CreateInvoiceReq) (*co_model.FdInvoiceRes, error) {
	ret, err := c.IFinancial.InvoiceRegister(ctx, &req.CreateInvoiceReq)
	return ret.Data(), err
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryInvoice(ctx context.Context, req *co_v1.QueryInvoiceReq) (*co_model.FdInvoiceListRes, error) {
	ret, err := c.IFinancial.QueryInvoice(ctx, &req.QueryInvoiceReq)
	return kconv.Struct(ret, &co_model.FdInvoiceListRes{}), err
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) DeletesFdInvoiceById(ctx context.Context, req *co_v1.DeleteInvoiceByIdReq) (api_v1.BoolRes, error) {
	return c.IFinancial.DeletesFdInvoiceById(ctx, &req.DeleteInvoiceByIdReq)
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) InvoiceDetailRegister(ctx context.Context, req *co_v1.CreateInvoiceDetailReq) (*co_model.FdInvoiceDetailRes, error) {
	ret, err := c.IFinancial.InvoiceDetailRegister(ctx, &req.CreateInvoiceDetailReq)
	return ret.Data(), err
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryInvoiceDetailList(ctx context.Context, req *co_v1.QueryInvoiceDetailListReq) (*co_model.FdInvoiceDetailListRes, error) {
	ret, err := c.IFinancial.QueryInvoiceDetailList(ctx, &req.QueryInvoiceDetailListReq)
	return kconv.Struct(ret, &co_model.FdInvoiceDetailListRes{}), err
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) MakeInvoiceDetailReq(ctx context.Context, req *co_v1.MakeInvoiceDetailReq) (api_v1.BoolRes, error) {
	return c.IFinancial.MakeInvoiceDetailReq(ctx, &req.MakeInvoiceDetailReq)
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) AuditInvoiceDetail(ctx context.Context, req *co_v1.AuditInvoiceDetailReq) (api_v1.BoolRes, error) {
	return c.IFinancial.AuditInvoiceDetail(ctx, &req.AuditInvoiceDetailReq)

}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) BankCardRegister(ctx context.Context, req *co_v1.BankCardRegisterReq) (*co_model.FdBankCardRes, error) {
	ret, err := c.IFinancial.BankCardRegister(ctx, &req.BankCardRegisterReq)
	return ret.Data(), err
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) DeleteBankCard(ctx context.Context, req *co_v1.DeleteBankCardReq) (api_v1.BoolRes, error) {
	return c.IFinancial.DeleteBankCard(ctx, &req.DeleteBankCardReq)
}

func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryBankCardList(ctx context.Context, req *co_v1.QueryBankCardListReq) (*co_model.FdBankCardListRes, error) {
	ret, err := c.IFinancial.QueryBankCardList(ctx, &req.QueryBankCardListReq)

	return kconv.Struct(ret, &co_model.FdBankCardListRes{}), err
}

// GetAccountDetail 查看财务账号明细
func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccountDetail(ctx context.Context, req *co_v1.GetAccountDetailReq) (ITFdAccountRes, error) {
	ret, err := c.IFinancial.GetAccountDetail(ctx, &req.GetAccountDetailReq)

	return ret, err
}

// UpdateAccountIsEnabled 修改财务账号启用状态
func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateAccountIsEnabled(ctx context.Context, req *co_v1.UpdateAccountIsEnabledReq) (api_v1.BoolRes, error) {
	ret, err := c.IFinancial.UpdateAccountIsEnabled(ctx, &req.UpdateAccountIsEnabledReq)

	return ret, err
}

// UpdateAccountLimitState 修改财务账号限制状态
func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateAccountLimitState(ctx context.Context, req *co_v1.UpdateAccountLimitStateReq) (api_v1.BoolRes, error) {
	ret, err := c.IFinancial.UpdateAccountLimitState(ctx, &req.UpdateAccountLimitStateReq)

	return ret, err
}

// GetAccountDetailById 根据财务账号id查询账单金额明细统计记录
func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccountDetailById(ctx context.Context, req *co_v1.GetAccountDetailByAccountIdReq) (*co_model.FdAccountDetailRes, error) {
	ret, err := c.IFinancial.GetAccountDetailById(ctx, &req.GetAccountDetailByAccountIdReq)

	return ret, err
}

// Increment 收入
func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) Increment(ctx context.Context, req *co_v1.IncrementReq) (api_v1.BoolRes, error) {
	ret, err := c.IFinancial.Increment(ctx, &req.IncrementReq)

	return ret, err
}

// Decrement 支出
func (c *FinancialController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) Decrement(ctx context.Context, req *co_v1.DecrementReq) (api_v1.BoolRes, error) {
	ret, err := c.IFinancial.Decrement(ctx, &req.DecrementReq)

	return ret, err
}
