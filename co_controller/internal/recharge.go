package internal

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/kysion/base-library/base_model"
)

type cRechargeController[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	TTFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	TR co_model.IFdRechargeRes,
] struct {
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		TTFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		TR,
	]
	recharge co_interface.IFdRecharge[TR]
	dao      co_dao.XDao
}

func Recharge[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	TTFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	TR co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TTFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) i_controller.IFdRecharge[TR] {
	return &cRechargeController[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		TTFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		TR,
	]{
		modules:  modules,
		dao:      *modules.Dao(),
		recharge: modules.Recharge(),
	}
}

func (c *cRechargeController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) GetAccountRechargeById(ctx context.Context, req *co_company_api.GetAccountRechargeByIdReq) (TR, error) {
	return c.recharge.GetAccountRechargeById(ctx, req.Id)
}

func (c *cRechargeController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) SetAccountRechargeAudit(ctx context.Context, req *co_company_api.SetAccountRechargeAuditReq) (api_v1.BoolRes, error) {
	audit, err := c.recharge.SetAccountRechargeAudit(ctx, req.Id, req.State, req.Reply)

	return audit == true, err
}

func (c *cRechargeController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) QueryAccountRecharge(ctx context.Context, req *co_company_api.QueryAccountRechargeReq) (*base_model.CollectRes[TR], error) {
	return c.recharge.QueryAccountRecharge(ctx, &req.SearchParams)
}

func (c *cRechargeController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) AccountRecharge(ctx context.Context, req *co_company_api.AccountRecharge) (TR, error) {
	session := sys_service.SysSession().Get(ctx)
	return c.recharge.AccountRecharge(ctx, req.Info, &session.JwtClaimsUser.SysUser)
}
