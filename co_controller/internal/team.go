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
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/base-library/utility/kconv"
)

type TeamController[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
TIRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		TIRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
	team     co_interface.ITeam[TIRes]
	employee co_interface.IEmployee[ITEmployeeRes]

	dao co_dao.XDao
}

func Team[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
TIRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) i_controller.ITeam[TIRes] {
	return &TeamController[
		ITCompanyRes,
		ITEmployeeRes,
		TIRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		modules:  modules,
		dao:      *modules.Dao(),
		team:     modules.Team(),
		employee: modules.Employee(),
	}
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetTeamById(ctx context.Context, req *co_company_api.GetTeamByIdReq) (TIRes, error) {
	// 这种只满足两级edu_school_class::Create，
	// 还需要兼容这样子的Financial::BankCard::ViewBankCardDetail （先不考虑）

	permission := c.getPermission(ctx, co_permission.Employee.PermissionType(c.modules).ViewDetail)

	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.team.GetTeamById(c.makeMore(ctx), req.Id)
			return ret, err
		},
		permission,
	)

}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) HasTeamByName(ctx context.Context, req *co_company_api.HasTeamByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.HasTeamByName(ctx, req.Name, req.UnionNameId, 0, req.ExcludeId) == true, nil
		},
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryTeamList(ctx context.Context, req *co_company_api.QueryTeamListReq) (*base_model.CollectRes[TIRes], error) {

	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).List)
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {
			return c.team.QueryTeamList(c.makeMore(ctx), &req.SearchParams)
		},
		permission,
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) CreateTeam(ctx context.Context, req *co_company_api.CreateTeamReq) (TIRes, error) {

	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).Create)
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.team.CreateTeam(c.makeMore(ctx), &req.Team)
			return ret, err
		},
		permission,
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateTeam(ctx context.Context, req *co_company_api.UpdateTeamReq) (TIRes, error) {
	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).Update)
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.team.UpdateTeam(c.makeMore(ctx), &req.Team)
			return ret, err
		},
		permission,
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) DeleteTeam(ctx context.Context, req *co_company_api.DeleteTeamReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).Delete)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.DeleteTeam(ctx, req.Id)
		},
		permission,
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryTeamListByEmployee(ctx context.Context, req *co_company_api.QueryTeamListByEmployeeReq) (*base_model.CollectRes[TIRes], error) {
	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).List)
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {

			return c.team.QueryTeamListByEmployee(c.makeMore(ctx), req.EmployeeId, req.UnionMainId)
		},
		permission,
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetTeamMember(ctx context.Context, req *co_company_api.SetTeamMemberReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).SetMember)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.SetTeamMember(ctx, req.Id, req.EmployeeIds)
		},
		permission,
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) RemoveTeamMember(ctx context.Context, req *co_company_api.RemoveTeamMemberReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).SetMember)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.RemoveTeamMember(ctx, req.Id, req.EmployeeIds)
		},
		permission,
	)
}
func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetTeamOwner(ctx context.Context, req *co_company_api.SetTeamOwnerReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).SetCaptain)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.SetTeamOwner(ctx, req.Id, req.EmployeeId)
		},
		permission,
	)
}
func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetTeamCaptain(ctx context.Context, req *co_company_api.SetTeamCaptainReq) (api_v1.BoolRes, error) {
	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).SetCaptain)
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.SetTeamCaptain(ctx, req.Id, req.EmployeeId)
		},
		permission,
	)
}
func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeListByTeamId(ctx context.Context, req *co_company_api.GetEmployeeListByTeamIdReq) (*base_model.CollectRes[co_model.IEmployeeRes], error) {
	permission := c.getPermission(ctx, co_permission.Team.PermissionType(c.modules).MemberDetail)
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[co_model.IEmployeeRes], error) {

			ret, err := c.team.GetEmployeeListByTeamId(c.makeMore(ctx), req.TeamId)
			if err != nil {
				return nil, err
			}

			result := base_model.CollectRes[co_model.IEmployeeRes]{}
			for _, record := range ret.Records {
				i := new(ITEmployeeRes)
				res := kconv.Struct(record, i)
				result.Records = append(result.Records, *res)
			}

			return &result, err

			//return kconv.Struct(ret, &base_model.CollectRes[co_model.EmployeeRes]{}), err

		},
		permission,
	)
}
func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetTeamInviteCode(ctx context.Context, req *co_company_api.GetTeamInviteCodeReq) (*co_model.TeamInviteCodeRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamInviteCodeRes, error) {
			user := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := c.team.GetTeamInviteCode(ctx, req.TeamId, user.Id)

			return ret, err
		},
		//co_permission.Team.PermissionType(c.modules).ViewDetail, 无需校验
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) JoinTeamByInviteCode(ctx context.Context, req *co_company_api.JoinTeamByInviteCodeReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			user := sys_service.SysSession().Get(ctx).JwtClaimsUser

			ret, err := c.team.JoinTeamByInviteCode(ctx, req.InviteCode, user.Id)

			return ret == true, err
		},
		// co_permission.Team.PermissionType(c.modules).SetMember, 设置团队成员权限校验，由于扫码人员什么人都有，不进行权限判断
	)
}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
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
		ctx = base_funs.AttrBuilder[TIRes, ITCompanyRes](ctx, c.dao.Team.Columns().UnionMainId)
		ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Team.Columns().OwnerEmployeeId)
		ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Team.Columns().CaptainEmployeeId)
		ctx = base_funs.AttrBuilder[TIRes, TIRes](ctx, c.dao.Team.Columns().ParentId)
		ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	}

	if include.Contains("unionMain") {
		ctx = base_funs.AttrBuilder[TIRes, ITCompanyRes](ctx, c.dao.Team.Columns().UnionMainId)
	}

	if include.Contains("owner") {
		ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Team.Columns().OwnerEmployeeId)
	}

	if include.Contains("captain") {
		ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Team.Columns().CaptainEmployeeId)
	}

	if include.Contains("parent") {
		ctx = base_funs.AttrBuilder[TIRes, TIRes](ctx, c.dao.Team.Columns().ParentId)
	}

	// 因为需要附加公共模块user的数据，所以也要添加有关sys_user的附加数据订阅
	if include.Contains("user") {
		ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	}
	return ctx
}

//
//func (c *TeamController[
//	ITCompanyRes,
//	ITEmployeeRes,
//	TIRes,
//	ITFdAccountRes,
//	ITFdAccountBillRes,
//	ITFdBankCardRes,
//	ITFdCurrencyRes,
//	ITFdInvoiceRes,
//	ITFdInvoiceDetailRes,
//]) getPermissionIdentifier(permission base_permission.IPermission) (identifierStr string) {
//	// 拼装标识符
//	return c.modules.GetConfig().Identifier.Team + "::" + permission.GetIdentifier()
//}

func (c *TeamController[
	ITCompanyRes,
	ITEmployeeRes,
	TIRes,
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
	identifierStr := c.modules.GetConfig().Identifier.Team + "::" + permission.GetIdentifier()
	// 注意：标识符匹配的话，需要找到数据库中的权限，然后传递进去
	sqlPermission, _ := sys_service.SysPermission().GetPermissionByIdentifier(ctx, identifierStr)
	if sqlPermission != nil {
		//permission = co_permission.Team.PermissionType(c.modules).ViewDetail.SetId(sqlPermission.Id).SetParentId(sqlPermission.ParentId).SetName(sqlPermission.Name).SetDescription(sqlPermission.Description).SetIdentifier(sqlPermission.Identifier).SetType(sqlPermission.Type).SetMatchMode(sqlPermission.MatchMode).SetIsShow(sqlPermission.IsShow).SetSort(sqlPermission.Sort)
		// CheckPermission 检验逻辑内部只用到了匹配模式 和 ID
		permission.SetId(sqlPermission.Id).SetParentId(sqlPermission.ParentId).SetIdentifier(sqlPermission.Identifier).SetMatchMode(sqlPermission.MatchMode)
	}

	return permission
}
