package co_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
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
	return funs.ProxyFunc1[*co_model.TeamRes](
		ctx, req.Id,
		c.modules.Team().GetTeamById, nil,
		co_enum.Team.PermissionType(c.modules).ViewDetail,
	)
}
func (c *cTeam[T]) HasTeamByName(ctx context.Context, req *co_v1.HasTeamByNameReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc[api_v1.BoolRes](
		ctx,
		func(ctx context.Context) (api_v1.BoolRes, error) {
			return c.modules.Team().HasTeamByName(ctx, req.Name, req.UnionMainId, req.ExcludeId) == true, nil
		}, false,
	)
}

func (c *cTeam[T]) QueryTeamList(ctx context.Context, req *co_v1.QueryTeamListReq) (*co_model.TeamListRes, error) {
	return funs.ProxyFunc1[*co_model.TeamListRes](
		ctx, &req.SearchParams,
		c.modules.Team().QueryTeamList,
		&co_model.TeamListRes{
			PaginationRes: sys_model.PaginationRes{
				Pagination: req.Pagination,
				PageTotal:  0,
				Total:      0,
			},
		},
		co_enum.Team.PermissionType(c.modules).List,
	)
}

func (c *cTeam[T]) CreateTeam(ctx context.Context, req *co_v1.CreateTeamReq) (*co_model.TeamRes, error) {
	return funs.ProxyFunc1[*co_model.TeamRes](
		ctx, &req.Team,
		c.modules.Team().CreateTeam,
		nil,
		co_enum.Team.PermissionType(c.modules).Create,
	)
}

func (c *cTeam[T]) UpdateTeam(ctx context.Context, req *co_v1.UpdateTeamReq) (*co_model.TeamRes, error) {
	return funs.ProxyFunc3[*co_model.TeamRes](ctx,
		req.Id, req.Name, req.Remark,
		c.modules.Team().UpdateTeam, nil,
		co_enum.Team.PermissionType(c.modules).Update,
	)
}

func (c *cTeam[T]) DeleteTeam(ctx context.Context, req *co_v1.DeleteTeamReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc1[api_v1.BoolRes](
		ctx, req.Id,
		c.modules.Team().DeleteTeam, false,
		co_enum.Team.PermissionType(c.modules).Delete,
	)
}

func (c *cTeam[T]) GetTeamMemberList(ctx context.Context, req *co_v1.GetTeamMemberListReq) (*co_model.EmployeeListRes, error) {
	return funs.ProxyFunc1[*co_model.EmployeeListRes](
		ctx, req.Id,
		c.modules.Team().GetTeamMemberList,
		&co_model.EmployeeListRes{
			PaginationRes: sys_model.PaginationRes{
				Pagination: sys_model.Pagination{
					Page:     1,
					PageSize: 100,
				},
				PageTotal: 0,
				Total:     0,
			},
		},
		co_enum.Team.PermissionType(c.modules).MemberDetail,
	)
}

func (c *cTeam[T]) QueryTeamByEmployeeList(ctx context.Context, req *co_v1.QueryTeamByEmployeeListReq) (*co_model.TeamListRes, error) {
	return funs.ProxyFunc2[*co_model.TeamListRes](
		ctx, req.EmployeeId, req.UnionMainId,
		c.modules.Team().QueryTeamByEmployeeList,
		&co_model.TeamListRes{
			PaginationRes: sys_model.PaginationRes{
				Pagination: sys_model.Pagination{
					Page:     1,
					PageSize: 100,
				},
				PageTotal: 0,
				Total:     0,
			},
		},
		co_enum.Team.PermissionType(c.modules).List,
	)
}

func (c *cTeam[T]) SetTeamMember(ctx context.Context, req *co_v1.SetTeamMemberReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc2[api_v1.BoolRes](
		ctx, req.Id, req.EmployeeIds,
		c.modules.Team().SetTeamMember, false,
		co_enum.Team.PermissionType(c.modules).SetMember,
	)
}
func (c *cTeam[T]) SetTeamOwner(ctx context.Context, req *co_v1.SetTeamOwnerReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc2[api_v1.BoolRes](
		ctx, req.Id, req.EmployeeId,
		c.modules.Team().SetTeamOwner, false,
		co_enum.Team.PermissionType(c.modules).SetOwner,
	)
}
func (c *cTeam[T]) SetTeamCaptain(ctx context.Context, req *co_v1.SetTeamCaptainReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc2[api_v1.BoolRes](
		ctx, req.Id, req.EmployeeId,
		c.modules.Team().SetTeamCaptain, false,
		co_enum.Team.PermissionType(c.modules).SetCaptain,
	)
}
