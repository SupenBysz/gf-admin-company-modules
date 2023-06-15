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
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/kysion/base-library/base_model"
	base_funs "github.com/kysion/base-library/utility/base_funs"
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
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.team.GetTeamById(c.makeMore(ctx), req.Id)
			return ret, err
		},
		co_permission.Team.PermissionType(c.modules).ViewDetail,
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
			return c.team.HasTeamByName(ctx, req.Name, req.UnionNameId, req.ExcludeId) == true, nil
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
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {
			return c.team.QueryTeamList(c.makeMore(ctx), &req.SearchParams)
		},
		co_permission.Team.PermissionType(c.modules).List,
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
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.team.CreateTeam(c.makeMore(ctx), &req.Team)
			return ret, err
		},
		co_permission.Team.PermissionType(c.modules).Create,
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
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.team.UpdateTeam(c.makeMore(ctx), req.Id, req.Name, req.Remark)
			return ret, err
		},
		co_permission.Team.PermissionType(c.modules).Update,
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
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.DeleteTeam(ctx, req.Id)
		},
		co_permission.Team.PermissionType(c.modules).Delete,
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
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {

			return c.team.QueryTeamListByEmployee(c.makeMore(ctx), req.EmployeeId, req.UnionMainId)
		},
		co_permission.Team.PermissionType(c.modules).List,
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
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.SetTeamMember(ctx, req.Id, req.EmployeeIds)
		},
		co_permission.Team.PermissionType(c.modules).SetMember,
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
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.SetTeamOwner(ctx, req.Id, req.EmployeeId)
		},
		co_permission.Team.PermissionType(c.modules).SetCaptain,
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
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.team.SetTeamCaptain(ctx, req.Id, req.EmployeeId)
		},
		co_permission.Team.PermissionType(c.modules).SetCaptain,
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
	ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Team.Columns().UnionMainId)
	ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Team.Columns().OwnerEmployeeId)
	ctx = base_funs.AttrBuilder[TIRes, ITEmployeeRes](ctx, c.dao.Team.Columns().CaptainEmployeeId)

	ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_entity.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	return ctx
}
