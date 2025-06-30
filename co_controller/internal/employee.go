package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/base-library/utility/kconv"
)

type EmployeeController[
ITCompanyRes co_model.ICompanyRes,
TIRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
ITFdRechargeRes co_model.IFdRechargeRes,
] struct {
	modules co_interface.IModules[
		ITCompanyRes,
		TIRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]
	employee co_interface.IEmployee[TIRes]
	team     co_interface.ITeam[ITTeamRes]

	dao co_dao.XDao
}

func Employee[
ITCompanyRes co_model.ICompanyRes,
TIRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) i_controller.IEmployee[TIRes] {
	return &EmployeeController[
		ITCompanyRes,
		TIRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]{
		modules:  modules,
		dao:      *modules.Dao(),
		employee: modules.Employee(),
		team:     modules.Team(),
	}
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetEmployeeById(ctx context.Context, req *co_company_api.GetEmployeeByIdReq) (TIRes, error) {
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).ViewDetail)
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			return c.employee.GetEmployeeById(c.makeMore(ctx), req.Id)
		},
		permission,
	)
}

// GetEmployeeDetailById 获取员工详情信息
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetEmployeeDetailById(ctx context.Context, req *co_company_api.GetEmployeeDetailByIdReq) (res TIRes, err error) {
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).MoreDetail)
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			return c.employee.GetEmployeeDetailById(c.makeMore(ctx), req.Id)
		},
		permission,
	)
}

// HasEmployeeByName 员工名称是否存在
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) HasEmployeeByName(ctx context.Context, req *co_company_api.HasEmployeeByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.employee.HasEmployeeByName(ctx, req.Name, req.UnionNameId, req.ExcludeId) == true, nil
		},
	)
}

// HasEmployeeByNo 员工工号是否存在
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) HasEmployeeByNo(ctx context.Context, req *co_company_api.HasEmployeeByNoReq) (api_v1.BoolRes, error) {
	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.employee.HasEmployeeByNo(ctx, req.No, unionMainId, req.ExcludeId) == true, nil
		},
	)
}

// QueryEmployeeList 查询员工列表
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryEmployeeList(ctx context.Context, req *co_company_api.QueryEmployeeListReq) (*base_model.CollectRes[TIRes], error) {
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).List)

	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {
			return c.employee.QueryEmployeeList(c.makeMore(ctx), &req.SearchParams)
		},
		permission,
	)
}

// CreateEmployee 创建员工信息
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) CreateEmployee(ctx context.Context, req *co_company_api.CreateEmployeeReq) (TIRes, error) {
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser
	req.UnionMainId = session.UnionMainId
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).Create)
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			user, err := sys_service.SysUser().CreateUser(ctx, req.UserInnerRegister, sys_enum.User.State.Normal, sys_enum.User.Type.New(session.Type, ""))

			if err != nil {
				var ret TIRes
				return ret, err
			}

			ret, err := c.employee.CreateEmployee(c.makeMore(ctx), &req.Employee, user)
			return ret, err
		},
		permission,
	)
}

// UpdateEmployee 更新员工信息
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateEmployee(ctx context.Context, req *co_company_api.UpdateEmployeeReq) (TIRes, error) {
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).Update)
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.employee.UpdateEmployee(c.makeMore(ctx), &req.UpdateEmployee)
			return ret, err
		},
		permission,
	)
}

// DeleteEmployee 删除员工信息
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) DeleteEmployee(ctx context.Context, req *co_company_api.DeleteEmployeeReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).Delete)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.employee.DeleteEmployee(ctx, req.Id)
			return ret == true, err
		},
		permission,
	)
}

// GetEmployeeListByRoleId 根据角色ID获取所有所属员工
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetEmployeeListByRoleId(ctx context.Context, req *co_company_api.GetEmployeeListByRoleIdReq) (*base_model.CollectRes[TIRes], error) {
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).ViewDetail)
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {
			return c.employee.GetEmployeeListByRoleId(c.makeMore(ctx), req.Id)
		},
		permission,
	)
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetEmployeeRoles(ctx context.Context, req *co_company_api.SetEmployeeRolesReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).SetRoles)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := sys_service.SysUser().SetUserRoles(
				ctx,
				req.UserId,
				req.RoleIds,
				sessionUser.UnionMainId,
			)
			return ret == true, err
		},
		permission,
	)
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetCommissionRate(ctx context.Context, req *co_company_api.SetEmployeeCommissionRateReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).SetCommission)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := c.modules.Employee().SetCommissionRate(ctx, req.UserId, req.CommissionRate, sessionUser.Id)
			return ret == true, err
		},
		permission,
	)
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetEmployeeState(ctx context.Context, req *co_company_api.SetEmployeeStateReq) (api_v1.BoolRes, error) {
	// 注意：标识符匹配的话，需要找到数据库中的权限，然后传递进去
	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).SetState)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.employee.SetEmployeeState(ctx, req.Id, req.State)
			return ret == true, err
		},
		permission,
	)
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
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
		ctx = base_funs.AttrBuilder[TIRes, []ITTeamRes](ctx, c.dao.Employee.Columns().UnionMainId)
		ctx = base_funs.AttrBuilder[TIRes, TIRes](ctx, c.dao.Employee.Columns().Id)
		ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	}

	// 最新附加数据规范：前端有需求，通过请求参数传递，后端在控制层才进行订阅数据，然后在service逻辑层进行数据附加
	if include.Contains("teamList") {
		ctx = base_funs.AttrBuilder[TIRes, []ITTeamRes](ctx, c.dao.Employee.Columns().UnionMainId)
	}

	// 因为需要附加公共模块user的数据，所以也要添加有关sys_user的附加数据订阅
	if include.Contains("user") {
		ctx = base_funs.AttrBuilder[TIRes, TIRes](ctx, c.dao.Employee.Columns().Id)
		ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	}
	return ctx
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) getPermission(ctx context.Context, permission base_permission.IPermission) base_permission.IPermission {

	//identifierStr := c.getPermissionIdentifier(permission)
	identifierStr := c.modules.GetConfig().Identifier.Employee + "::" + permission.GetIdentifier()
	// 注意：标识符匹配的话，需要找到数据库中的权限，然后传递进去
	sqlPermission, _ := sys_service.SysPermission().GetPermissionByIdentifier(ctx, identifierStr)
	if sqlPermission != nil {
		//permission = co_permission.Team.PermissionType(c.modules).ViewDetail.SetId(sqlPermission.Id).SetParentId(sqlPermission.ParentId).SetName(sqlPermission.Name).SetDescription(sqlPermission.Description).SetIdentifier(sqlPermission.Identifier).SetType(sqlPermission.Type).SetMatchMode(sqlPermission.MatchMode).SetIsShow(sqlPermission.IsShow).SetSort(sqlPermission.Sort)
		// CheckPermission 检验逻辑内部只用到了匹配模式 和 ID
		permission.SetId(sqlPermission.Id).SetParentId(sqlPermission.ParentId).SetIdentifier(sqlPermission.Identifier).SetMatchMode(sqlPermission.MatchMode)
	}

	return permission
}
