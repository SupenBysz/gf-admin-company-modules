package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	base_funs "github.com/kysion/base-library/utility/base_funs"
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
		return nil, err
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
	// team相关附加信息
	ctx = base_funs.AttrBuilder[co_model.EmployeeRes, []co_model.Team](ctx, c.modules.Dao().Employee.Columns().UnionMainId)

	// 加上员工的附加信息订阅，
	ctx = base_funs.AttrBuilder[co_model.EmployeeRes, *co_model.EmployeeRes](ctx, c.modules.Dao().Employee.Columns().Id)

	// 因为需要附加公共模块user的数据，所以也要添加有关sys_user的附加数据订阅
	ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_entity.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	return ctx
}
