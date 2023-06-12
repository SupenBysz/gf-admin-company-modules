package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/kysion/base-library/utility/base_funs"
)

type MyController[
TIRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	i_controller.IMy
	modules co_interface.IModules[
		TIRes,
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

func My[
TIRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) i_controller.IMy {
	return &MyController[
		TIRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		modules: modules,
	}
}

// GetProfile 获取当前员工及用户信息 (附加数据：user、user_detail、employee、teamList)
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetProfile(ctx context.Context, _ *co_company_api.GetProfileReq) (*co_model.MyProfileRes, error) {
	result, err := c.modules.My().GetProfile(c.makeMore(ctx))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCompany 获取当前公司信息
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetCompany(ctx context.Context, _ *co_company_api.GetCompanyReq) (*co_model.MyCompanyRes, error) {
	result, err := c.modules.My().GetCompany(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// GetTeams 获取当前团队信息  (附加数据：user、user_detail、employee、teamList)
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetTeams(ctx context.Context, _ *co_company_api.GetTeamsReq) (co_model.MyTeamListRes, error) {

	result, err := c.modules.My().GetTeams(c.makeMore(ctx))
	if err != nil {
		return co_model.MyTeamListRes{}, err
	}

	return result, nil
}

// SetAvatar 设置员工头像
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetAvatar(ctx context.Context, req *co_company_api.SetAvatarReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.My().SetMyAvatar(ctx, req.ImageId)
			return ret == true, err
		},
		co_permission.Employee.PermissionType(c.modules).SetAvatar,
	)
}

// SetMobile 设置手机号
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetMobile(ctx context.Context, req *co_company_api.SetMobileReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.My().SetMyMobile(ctx, req.Mobile, req.Captcha, req.Password)
			return ret == true, err
		},
		co_permission.Employee.PermissionType(c.modules).SetMobile,
	)
}

// GetAccountBills 我的账单|列表
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccountBills(ctx context.Context, req *co_company_api.GetAccountBillsReq) (*co_model.MyAccountBillRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.MyAccountBillRes, error) {
			ret, err := c.modules.My().GetAccountBills(ctx, &req.Pagination)
			return ret, err
		},
		co_permission.Financial.PermissionType(c.modules).GetAccountDetail,
	)
}

// GetAccounts 获取我的财务账号|列表
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccounts(ctx context.Context, _ *co_company_api.GetAccountsReq) (*co_model.FdAccountListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.FdAccountListRes, error) {
			ret, err := c.modules.My().GetAccounts(ctx)
			return ret, err
		},
		co_permission.Financial.PermissionType(c.modules).GetAccountDetail,
	)
}

// GetBankCards 获取我的银行卡｜列表
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetBankCards(ctx context.Context, _ *co_company_api.GetBankCardsReq) (*co_model.FdBankCardListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.FdBankCardListRes, error) {
			ret, err := c.modules.My().GetBankCards(ctx)
			return ret, err
		},
		co_permission.Financial.PermissionType(c.modules).BankCardList,
	)
}

// GetInvoices 获取我的发票抬头｜列表
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetInvoices(ctx context.Context, _ *co_company_api.GetInvoicesReq) (*co_model.FdInvoiceListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.FdInvoiceListRes, error) {
			ret, err := c.modules.My().GetInvoices(ctx)
			return ret, err
		},
		co_permission.Financial.PermissionType(c.modules).InvoiceList,
	)
}

// UpdateAccount  修改我的财务账号
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateAccount(ctx context.Context, req *co_company_api.UpdateAccountReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.My().UpdateAccount(ctx, req.AccountId, &req.UpdateAccount)
			return ret == true, err
		},
		co_permission.Financial.PermissionType(c.modules).UpdateAccountDetail,
	)
}

func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) makeMore(ctx context.Context) context.Context {
	// 附加数据1：团队负责人Owner
	ctx = base_funs.AttrBuilder[co_model.TeamRes, *co_model.EmployeeRes](ctx, c.modules.Dao().Team.Columns().OwnerEmployeeId)

	// 附加数据2：团队队长Captain
	ctx = base_funs.AttrBuilder[co_model.TeamRes, *co_model.EmployeeRes](ctx, c.modules.Dao().Team.Columns().CaptainEmployeeId)

	// 附加数据3：团队主体UnionMain
	ctx = base_funs.AttrBuilder[co_model.TeamRes, *co_model.CompanyRes](ctx, c.modules.Dao().Team.Columns().UnionMainId)

	// 附加数据4：团队或小组父级
	ctx = base_funs.AttrBuilder[co_model.TeamRes, *co_model.TeamRes](ctx, c.modules.Dao().Team.Columns().ParentId)

	return ctx
}
