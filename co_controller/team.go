package co_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

type cTeam[T co_interface.IModules] struct {
	modules T
}

var Team = func(modules co_interface.IModules) *cTeam[co_interface.IModules] {
	return &cTeam[co_interface.IModules]{
		modules: modules,
	}
}

func (c *cTeam[T]) GetTeamById(ctx context.Context, req *co_v1.GetTeamByIdReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().GetTeamById(ctx, req.Id)
			return (*co_model.TeamRes)(ret), err
		},
		co_enum.Team.PermissionType(c.modules).ViewDetail,
	)
}
func (c *cTeam[T]) HasTeamByName(ctx context.Context, req *co_v1.HasTeamByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().HasTeamByName(ctx, req.Name, req.UnionMainId, req.ExcludeId) == true, nil
		},
	)
}

func (c *cTeam[T]) QueryTeamList(ctx context.Context, req *co_v1.QueryTeamListReq) (*co_model.TeamListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamListRes, error) {
			return c.modules.Team().QueryTeamList(ctx, &req.SearchParams)
		},
		co_enum.Team.PermissionType(c.modules).List,
	)
}

func (c *cTeam[T]) CreateTeam(ctx context.Context, req *co_v1.CreateTeamReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().CreateTeam(ctx, &req.Team)
			return (*co_model.TeamRes)(ret), err
		},
		co_enum.Team.PermissionType(c.modules).Create,
	)
}

func (c *cTeam[T]) UpdateTeam(ctx context.Context, req *co_v1.UpdateTeamReq) (*co_model.TeamRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamRes, error) {
			ret, err := c.modules.Team().UpdateTeam(ctx, req.Id, req.Name, req.Remark)
			return (*co_model.TeamRes)(ret), err
		},
		co_enum.Team.PermissionType(c.modules).Update,
	)
}

func (c *cTeam[T]) DeleteTeam(ctx context.Context, req *co_v1.DeleteTeamReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().DeleteTeam(ctx, req.Id)
		},
		co_enum.Team.PermissionType(c.modules).Delete,
	)
}

func (c *cTeam[T]) GetTeamMemberList(ctx context.Context, req *co_v1.GetTeamMemberListReq) (*co_model.EmployeeListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeListRes, error) {
			return c.modules.Team().GetTeamMemberList(ctx, req.Id)
		},
		co_enum.Team.PermissionType(c.modules).MemberDetail,
	)
}

func (c *cTeam[T]) QueryTeamListByEmployee(ctx context.Context, req *co_v1.QueryTeamListByEmployeeReq) (*co_model.TeamListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.TeamListRes, error) {
			return c.modules.Team().QueryTeamListByEmployee(ctx, req.EmployeeId, req.UnionMainId)
		},
		co_enum.Team.PermissionType(c.modules).List,
	)
}

func (c *cTeam[T]) SetTeamMember(ctx context.Context, req *co_v1.SetTeamMemberReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamMember(ctx, req.Id, req.EmployeeIds)
		},
		co_enum.Team.PermissionType(c.modules).SetMember,
	)
}
func (c *cTeam[T]) SetTeamOwner(ctx context.Context, req *co_v1.SetTeamOwnerReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamOwner(ctx, req.Id, req.EmployeeId)
		},
		co_enum.Team.PermissionType(c.modules).SetCaptain,
	)
}
func (c *cTeam[T]) SetTeamCaptain(ctx context.Context, req *co_v1.SetTeamCaptainReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Team().SetTeamCaptain(ctx, req.Id, req.EmployeeId)
		},
		co_enum.Team.PermissionType(c.modules).SetCaptain,
	)
}
