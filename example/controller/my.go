package controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type MyController struct {
	i_controller.IMy
}

func My[
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
]) *MyController {
	return &MyController{
		IMy: co_controller.My(modules),
	}
}

// GetProfile 获取当前员工及用户信息
func (c *MyController) GetProfile(ctx context.Context, req *co_v1.GetProfileReq) (*co_model.MyProfileRes, error) {
	return c.IMy.GetProfile(ctx, &req.GetProfileReq)
}

// GetCompany 获取当前公司信息
func (c *MyController) GetCompany(ctx context.Context, req *co_v1.GetCompanyReq) (*co_model.MyCompanyRes, error) {
	return c.IMy.GetCompany(ctx, &req.GetCompanyReq)
}

// GetTeams 获取当前团队信息
func (c *MyController) GetTeams(ctx context.Context, req *co_v1.GetTeamsReq) (co_model.MyTeamListRes, error) {
	return c.IMy.GetTeams(ctx, &req.GetTeamsReq)
}

// SetAvatar 设置员工头像
func (c *MyController) SetAvatar(ctx context.Context, req *co_v1.SetAvatarReq) (api_v1.BoolRes, error) {
	return c.IMy.SetAvatar(ctx, &req.SetAvatarReq)
}

// SetMobile 设置手机号
func (c *MyController) SetMobile(ctx context.Context, req *co_v1.SetMobileReq) (api_v1.BoolRes, error) {
	return c.IMy.SetMobile(ctx, &req.SetMobileReq)
}

// SetMail 设置邮箱
func (c *MyController) SetMail(ctx context.Context, req *co_v1.SetMailReq) (api_v1.BoolRes, error) {
	return c.IMy.SetMail(ctx, &req.SetMailReq)
}

// GetAccountBills 我的账单|列表
func (c *MyController) GetAccountBills(ctx context.Context, req *co_v1.GetAccountBillsReq) (*co_model.MyAccountBillRes, error) {
	return c.IMy.GetAccountBills(ctx, &req.GetAccountBillsReq)
}

// GetAccounts 获取我的财务账号|列表
func (c *MyController) GetAccounts(ctx context.Context, req *co_v1.GetAccountsReq) (*co_model.FdAccountListRes, error) {
	return c.IMy.GetAccounts(ctx, &req.GetAccountsReq)
}

// GetBankCards 获取我的银行卡｜列表
func (c *MyController) GetBankCards(ctx context.Context, req *co_v1.GetBankCardsReq) (*co_model.FdBankCardListRes, error) {
	return c.IMy.GetBankCards(ctx, &req.GetBankCardsReq)
}

// GetInvoices 获取我的发票抬头｜列表
func (c *MyController) GetInvoices(ctx context.Context, req *co_v1.GetInvoicesReq) (*co_model.FdInvoiceListRes, error) {
	return c.IMy.GetInvoices(ctx, &req.GetInvoicesReq)
}

// UpdateAccount  修改我的财务账号
func (c *MyController) UpdateAccount(ctx context.Context, req *co_v1.UpdateAccountReq) (api_v1.BoolRes, error) {
	return c.IMy.UpdateAccount(ctx, &req.UpdateAccountReq)
}
