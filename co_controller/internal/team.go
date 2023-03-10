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
)

type TeamController struct {
	i_controller.ITeam
	modules co_interface.IModules
	dao     *co_dao.XDao
}

var Team = func(modules co_interface.IModules) i_controller.ITeam {
	return &TeamController{
		modules: modules,
		dao:     modules.Dao(),
	}
}

func (c *TeamController) GetModules() co_interface.IModules {
	return c.modules
}

func (c *TeamController) GetTeamById(ctx context.Context, req *co_company_api.GetTeamByIdReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().GetTeamById(c.makeMore(ctx), req.Id)
			return ret, err
		},
		co_permission.Team.PermissionType(c.modules).ViewDetail,
	)
}

func (c *TeamController) HasTeamByName(ctx context.Context, req *co_company_api.HasTeamByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().HasTeamByName(ctx, req.Name, req.UnionNameId, req.ExcludeId) == true, nil
		},
	)
}

func (c *TeamController) QueryTeamList(ctx context.Context, req *co_company_api.QueryTeamListReq) (*co_model.TeamListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamListRes, error) {
			return c.modules.Team().QueryTeamList(c.makeMore(ctx), &req.SearchParams)
		},
		co_permission.Team.PermissionType(c.modules).List,
	)
}

func (c *TeamController) CreateTeam(ctx context.Context, req *co_company_api.CreateTeamReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().CreateTeam(c.makeMore(ctx), &req.Team)
			return ret, err
		},
		co_permission.Team.PermissionType(c.modules).Create,
	)
}

func (c *TeamController) UpdateTeam(ctx context.Context, req *co_company_api.UpdateTeamReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().UpdateTeam(c.makeMore(ctx), req.Id, req.Name, req.Remark)
			return ret, err
		},
		co_permission.Team.PermissionType(c.modules).Update,
	)
}

func (c *TeamController) DeleteTeam(ctx context.Context, req *co_company_api.DeleteTeamReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().DeleteTeam(ctx, req.Id)
		},
		co_permission.Team.PermissionType(c.modules).Delete,
	)
}

func (c *TeamController) GetTeamMemberList(ctx context.Context, req *co_company_api.GetTeamMemberListReq) (*co_model.EmployeeListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeListRes, error) {
			return c.modules.Team().GetTeamMemberList(c.makeMore(ctx), req.Id)
		},
		co_permission.Team.PermissionType(c.modules).MemberDetail,
	)
}

func (c *TeamController) QueryTeamListByEmployee(ctx context.Context, req *co_company_api.QueryTeamListByEmployeeReq) (*co_model.TeamListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamListRes, error) {

			return c.modules.Team().QueryTeamListByEmployee(c.makeMore(ctx), req.EmployeeId, req.UnionMainId)
		},
		co_permission.Team.PermissionType(c.modules).List,
	)
}

func (c *TeamController) SetTeamMember(ctx context.Context, req *co_company_api.SetTeamMemberReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamMember(ctx, req.Id, req.EmployeeIds)
		},
		co_permission.Team.PermissionType(c.modules).SetMember,
	)
}
func (c *TeamController) SetTeamOwner(ctx context.Context, req *co_company_api.SetTeamOwnerReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamOwner(ctx, req.Id, req.EmployeeId)
		},
		co_permission.Team.PermissionType(c.modules).SetCaptain,
	)
}
func (c *TeamController) SetTeamCaptain(ctx context.Context, req *co_company_api.SetTeamCaptainReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamCaptain(ctx, req.Id, req.EmployeeId)
		},
		co_permission.Team.PermissionType(c.modules).SetCaptain,
	)
}

func (c *TeamController) makeMore(ctx context.Context) context.Context {
	ctx = funs.AttrBuilder[co_model.TeamRes, *co_model.CompanyRes](ctx, c.dao.Team.Columns().UnionMainId)
	ctx = funs.AttrBuilder[co_model.TeamRes, *co_model.EmployeeRes](ctx, c.dao.Team.Columns().OwnerEmployeeId)
	ctx = funs.AttrBuilder[co_model.TeamRes, *co_model.EmployeeRes](ctx, c.dao.Team.Columns().CaptainEmployeeId)

	ctx = funs.AttrBuilder[sys_model.SysUser, *sys_entity.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	return ctx
}
