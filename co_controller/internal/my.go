package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/base-library/utility/kconv"
)

type MyController[
	TIRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
] struct {
	i_controller.IMy
	modules co_interface.IModules[
		TIRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]
}

func My[
	TIRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) i_controller.IMy {
	return &MyController[
		TIRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
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
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
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
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
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
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
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
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetAvatar(ctx context.Context, req *co_company_api.SetAvatarReq) (api_v1.BoolRes, error) {
	//permission := co_permission.Employee.PermissionType(c.modules).SetAvatar
	//identifierStr := c.modules.GetConfig().Identifier.Employee + "::" + permission.GetIdentifier()
	//// 注意：标识符匹配的话，需要找到数据库中的权限，然后传递进去
	//sqlPermission, _ := sys_service.SysPermission().GetPermissionByIdentifier(ctx, identifierStr)
	//if sqlPermission != nil {
	//	//permission = co_permission.Team.PermissionType(c.modules).ViewDetail.SetId(sqlPermission.Id).SetParentId(sqlPermission.ParentId).SetName(sqlPermission.Name).SetDescription(sqlPermission.Description).SetIdentifier(sqlPermission.Identifier).SetType(sqlPermission.Type).SetMatchMode(sqlPermission.MatchMode).SetIsShow(sqlPermission.IsShow).SetSort(sqlPermission.Sort)
	//	// CheckPermission 检验逻辑内部只用到了匹配模式 和 ID
	//	permission.SetId(sqlPermission.Id).SetParentId(sqlPermission.ParentId).SetIdentifier(sqlPermission.Identifier).SetMatchMode(sqlPermission.MatchMode)
	//}
	//
	////permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).SetAvatar)
	//return funs.CheckPermission(ctx,
	//	func() (api_v1.BoolRes, error) {
	ret, err := c.modules.My().SetMyAvatar(ctx, req.ImageId)
	return ret == true, err
	//	},
	//	permission,
	//)
}

// SetMobile 设置手机号
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetMobile(ctx context.Context, req *co_company_api.SetMobileReq) (api_v1.BoolRes, error) {
	ret, err := c.modules.My().SetMyMobile(ctx, req.Mobile, req.Captcha, req.Password)
	return ret == true, err
}

// SetMail 设置邮箱
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetMail(ctx context.Context, req *co_company_api.SetMailReq) (api_v1.BoolRes, error) {
	ret, err := c.modules.My().SetMyMail(ctx, req.OldMail, req.NewMail, req.Captcha, req.Password)
	return ret == true, err
}

// GetAccountBills 我的账单|列表
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountBills(ctx context.Context, req *co_company_api.GetAccountBillsReq) (*co_model.MyAccountBillRes, error) {
	ret, err := c.modules.My().GetAccountBills(ctx, &req.SearchParams)
	return ret, err
}

// GetAccounts 获取我的财务账号|列表
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccounts(ctx context.Context, _ *co_company_api.GetAccountsReq) (*co_model.FdAccountListRes, error) {
	ret, err := c.modules.My().GetAccounts(c.makeMore(ctx))
	return ret, err
}

// GetBankCards 获取我的银行卡｜列表
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetBankCards(ctx context.Context, _ *co_company_api.GetBankCardsReq) (*co_model.FdBankCardListRes, error) {
	ret, err := c.modules.My().GetBankCards(ctx)
	return ret, err
}

// GetInvoices 获取我的发票抬头｜列表
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetInvoices(ctx context.Context, _ *co_company_api.GetInvoicesReq) (*co_model.FdInvoiceListRes, error) {
	ret, err := c.modules.My().GetInvoices(ctx)
	return ret, err
}

// UpdateAccount  修改我的财务账号
func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateAccount(ctx context.Context, req *co_company_api.UpdateAccountReq) (api_v1.BoolRes, error) {
	ret, err := c.modules.My().UpdateAccount(ctx, req.AccountId, &req.UpdateAccount)
	return ret == true, err
}

func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) makeMore(ctx context.Context) context.Context {

	include := &garray.StrArray{}
	if ctx.Value("include") == nil {
		r := g.RequestFromCtx(ctx)
		array := r.GetForm("include").Array()
		arr := kconv.Struct(array, &[]string{})
		include = garray.NewStrArrayFrom(*arr)
	} else {
		array := ctx.Value("include")
		arr := kconv.Struct(array, &[]string{})
		include = garray.NewStrArrayFrom(*arr)
	}

	if include.Contains("*") {
		ctx = base_funs.AttrBuilder[ITTeamRes, ITEmployeeRes](ctx, c.modules.Dao().Team.Columns().OwnerEmployeeId)
		ctx = base_funs.AttrBuilder[ITTeamRes, ITEmployeeRes](ctx, c.modules.Dao().Team.Columns().CaptainEmployeeId)
		ctx = base_funs.AttrBuilder[ITTeamRes, TIRes](ctx, c.modules.Dao().Team.Columns().UnionMainId)
		ctx = base_funs.AttrBuilder[ITTeamRes, ITTeamRes](ctx, c.modules.Dao().Team.Columns().ParentId)
		ctx = base_funs.AttrBuilder[ITFdAccountRes, ITFdAccountRes](ctx, "id")
		ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	}
	// 附加数据1：团队负责人Owner
	if include.Contains("owner") {
		ctx = base_funs.AttrBuilder[ITTeamRes, ITEmployeeRes](ctx, c.modules.Dao().Team.Columns().OwnerEmployeeId)
	}

	// 附加数据2：团队队长Captain
	if include.Contains("captain") {
		ctx = base_funs.AttrBuilder[ITTeamRes, ITEmployeeRes](ctx, c.modules.Dao().Team.Columns().CaptainEmployeeId)
	}

	// 附加数据3：团队主体UnionMain
	if include.Contains("unionMain") {
		ctx = base_funs.AttrBuilder[ITTeamRes, TIRes](ctx, c.modules.Dao().Team.Columns().UnionMainId)
	}

	// 附加数据4：团队或小组父级
	if include.Contains("parent") {
		ctx = base_funs.AttrBuilder[ITTeamRes, ITTeamRes](ctx, c.modules.Dao().Team.Columns().ParentId)
	}

	// 附加数据5：用户信息附加数据
	if include.Contains("user") {
		ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	}

	// 财务账号明细Detail附加数据
	if include.Contains("detail") {
		ctx = base_funs.AttrBuilder[ITFdAccountRes, ITFdAccountRes](ctx, "id")
	}

	return ctx
}

func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) getPermissionIdentifier(permission base_permission.IPermission) (identifierStr string) {
	// 拼装标识符
	return c.modules.GetConfig().Identifier.Team + "::" + permission.GetIdentifier()
}

func (c *MyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetMyCompanyPermissionList(ctx context.Context, req *co_company_api.GetMyCompanyPermissionListReq) (*sys_model.MyPermissionListRes, error) {
	return c.modules.My().GetMyCompanyPermissionList(ctx, req.PermissionType)
}
