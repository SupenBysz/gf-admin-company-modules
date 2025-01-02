package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/base-library/utility/kconv"
)

type CompanyController[
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
	i_controller.ICompany[TIRes]
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
	dao co_dao.XDao
}

func Company[
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
]) i_controller.ICompany[TIRes] {
	return &CompanyController[
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
		dao:     *modules.Dao(),
	}
}

// GetCompanyById 通过ID获取公司信息
func (c *CompanyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetCompanyById(ctx context.Context, req *co_company_api.GetCompanyByIdReq) (TIRes, error) {
	permission := c.getPermission(ctx, co_permission.Company.PermissionType(c.modules).ViewDetail)
	return funs.CheckPermission(ctx,

		func() (TIRes, error) {
			ret, err := c.modules.Company().GetCompanyById(c.makeMore(ctx), req.Id)
			return ret, err
		},
		permission,
	)
}

// HasCompanyByName 公司名称是否存在
func (c *CompanyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) HasCompanyByName(ctx context.Context, req *co_company_api.HasCompanyByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Company().HasCompanyByName(ctx, req.Name) == true, nil
		},
	)
}

// QueryCompanyList 查询公司列表
func (c *CompanyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryCompanyList(ctx context.Context, req *co_company_api.QueryCompanyListReq) (*base_model.CollectRes[TIRes], error) {
	permission := c.getPermission(ctx, co_permission.Company.PermissionType(c.modules).List)
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {
			return c.modules.Company().QueryCompanyList(c.makeMore(ctx), &req.SearchParams, req.IsExport)
		},
		permission,
	)
}

// CreateCompany 创建公司信息
func (c *CompanyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) CreateCompany(ctx context.Context, req *co_company_api.CreateCompanyReq) (TIRes, error) {
	permission := c.getPermission(ctx, co_permission.Company.PermissionType(c.modules).Create)

	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			return c.modules.Company().CreateCompany(c.makeMore(ctx), &req.Company, nil)
		},
		permission,
	)
}

// UpdateCompany 更新公司信息
func (c *CompanyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateCompany(ctx context.Context, req *co_company_api.UpdateCompanyReq) (TIRes, error) {
	permission := c.getPermission(ctx, co_permission.Company.PermissionType(c.modules).Update)

	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			return c.modules.Company().UpdateCompany(c.makeMore(ctx), &req.Company)
		},
		permission,
	)
}

// GetCompanyDetail 获取公司详情，包含完整商务联系人电话
func (c *CompanyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetCompanyDetail(ctx context.Context, req *co_company_api.GetCompanyDetailReq) (TIRes, error) {
	permission := c.getPermission(ctx, co_permission.Company.PermissionType(c.modules).ViewMobile)

	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			return c.modules.Company().GetCompanyDetail(c.makeMore(ctx), req.Id)
		},
		permission,
	)
}

func (c *CompanyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetCompanyState(ctx context.Context, req *co_company_api.SetCompanyStateReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Company.PermissionType(c.modules).SetState)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Company().SetCompanyState(ctx, req.Id, co_enum.Company.State.New(req.State, ""))
			return ret == true, err
		},
		permission,
	)
}

func (c *CompanyController[
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
		ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Company.Columns().UserId)
		ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	}

	if include.Contains("adminUser") {
		ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Company.Columns().UserId)
	}

	if include.Contains("user") {
		ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	}

	return ctx
}

func (c *CompanyController[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) getPermission(ctx context.Context, permission base_permission.IPermission) base_permission.IPermission {
	// 这种只满足两级edu_school_class::Create，
	// 还需要兼容这样子的Financial::BankCard::ViewBankCardDetail （先不考虑）

	//identifierStr := c.getPermissionIdentifier(permission)
	identifierStr := c.modules.GetConfig().Identifier.Company + "::" + permission.GetIdentifier()
	// 注意：标识符匹配的话，需要找到数据库中的权限，然后传递进去
	sqlPermission, _ := sys_service.SysPermission().GetPermissionByIdentifier(ctx, identifierStr)
	if sqlPermission != nil {
		//permission = co_permission.Team.PermissionType(c.modules).ViewDetail.SetId(sqlPermission.Id).SetParentId(sqlPermission.ParentId).SetName(sqlPermission.Name).SetDescription(sqlPermission.Description).SetIdentifier(sqlPermission.Identifier).SetType(sqlPermission.Type).SetMatchMode(sqlPermission.MatchMode).SetIsShow(sqlPermission.IsShow).SetSort(sqlPermission.Sort)
		// CheckPermission 检验逻辑内部只用到了匹配模式 和 ID
		permission.SetId(sqlPermission.Id).SetParentId(sqlPermission.ParentId).SetIdentifier(sqlPermission.Identifier).SetMatchMode(sqlPermission.MatchMode)
	}

	return permission
}
