package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

type TeamController[T co_interface.IModules] struct {
	i_controller.ITeam
	modules T
}

var Team = func(modules co_interface.IModules) i_controller.ITeam {
	return &TeamController[co_interface.IModules]{
		modules: modules,
	}
}

func (c *TeamController[T]) GetModules() co_interface.IModules {
	return c.modules
}

func (c *TeamController[T]) GetTeamById(ctx context.Context, req *co_company_api.GetTeamByIdReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().GetTeamById(ctx, req.Id)
			return (*co_model.TeamRes)(ret), err
		},
		co_enum.Team.PermissionType(c.modules).ViewDetail,
	)
}
func (c *TeamController[T]) HasTeamByName(ctx context.Context, req *co_company_api.HasTeamByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().HasTeamByName(ctx, req.Name, req.UnionMainId, req.ExcludeId) == true, nil
		},
	)
}

func (c *TeamController[T]) QueryTeamList(ctx context.Context, req *co_company_api.QueryTeamListReq) (*co_model.TeamListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamListRes, error) {
			return c.modules.Team().QueryTeamList(ctx, &req.SearchParams)
		},
		co_enum.Team.PermissionType(c.modules).List,
	)
}

func (c *TeamController[T]) CreateTeam(ctx context.Context, req *co_company_api.CreateTeamReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().CreateTeam(ctx, &req.Team)
			return (*co_model.TeamRes)(ret), err
		},
		co_enum.Team.PermissionType(c.modules).Create,
	)
}

func (c *TeamController[T]) UpdateTeam(ctx context.Context, req *co_company_api.UpdateTeamReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().UpdateTeam(ctx, req.Id, req.Name, req.Remark)
			return (*co_model.TeamRes)(ret), err
		},
		co_enum.Team.PermissionType(c.modules).Update,
	)
}

func (c *TeamController[T]) DeleteTeam(ctx context.Context, req *co_company_api.DeleteTeamReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().DeleteTeam(ctx, req.Id)
		},
		co_enum.Team.PermissionType(c.modules).Delete,
	)
}

func (c *TeamController[T]) GetTeamMemberList(ctx context.Context, req *co_company_api.GetTeamMemberListReq) (*co_model.EmployeeListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeListRes, error) {
			return c.modules.Team().GetTeamMemberList(ctx, req.Id)
		},
		co_enum.Team.PermissionType(c.modules).MemberDetail,
	)
}

func (c *TeamController[T]) QueryTeamListByEmployee(ctx context.Context, req *co_company_api.QueryTeamListByEmployeeReq) (*co_model.TeamListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamListRes, error) {
			return c.modules.Team().QueryTeamListByEmployee(ctx, req.EmployeeId, req.UnionMainId)
		},
		co_enum.Team.PermissionType(c.modules).List,
	)
}

func (c *TeamController[T]) SetTeamMember(ctx context.Context, req *co_company_api.SetTeamMemberReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamMember(ctx, req.Id, req.EmployeeIds)
		},
		co_enum.Team.PermissionType(c.modules).SetMember,
	)
}
func (c *TeamController[T]) SetTeamOwner(ctx context.Context, req *co_company_api.SetTeamOwnerReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamOwner(ctx, req.Id, req.EmployeeId)
		},
		co_enum.Team.PermissionType(c.modules).SetCaptain,
	)
}
func (c *TeamController[T]) SetTeamCaptain(ctx context.Context, req *co_company_api.SetTeamCaptainReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamCaptain(ctx, req.Id, req.EmployeeId)
		},
		co_enum.Team.PermissionType(c.modules).SetCaptain,
	)
}
