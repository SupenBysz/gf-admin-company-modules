package controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type TeamController struct {
	i_controller.ITeam
}

var Team = func(modules co_interface.IModules) *TeamController {
	return &TeamController{
		co_controller.Team(modules),
	}
}

func (c *TeamController) GetModules() co_interface.IModules {
	return c.ITeam.GetModules()
}

func (c *TeamController) GetTeamById(ctx context.Context, req *co_v1.GetTeamByIdReq) (*co_model.TeamRes, error) {
	return c.ITeam.GetTeamById(ctx, &req.GetTeamByIdReq)
}

func (c *TeamController) HasTeamByName(ctx context.Context, req *co_v1.HasTeamByNameReq) (api_v1.BoolRes, error) {
	return c.ITeam.HasTeamByName(ctx, &req.HasTeamByNameReq)
}

func (c *TeamController) QueryTeamList(ctx context.Context, req *co_v1.QueryTeamListReq) (*co_model.TeamListRes, error) {
	return c.ITeam.QueryTeamList(ctx, &req.QueryTeamListReq)
}

func (c *TeamController) CreateTeam(ctx context.Context, req *co_v1.CreateTeamReq) (*co_model.TeamRes, error) {
	return c.ITeam.CreateTeam(ctx, &req.CreateTeamReq)
}

func (c *TeamController) UpdateTeam(ctx context.Context, req *co_v1.UpdateTeamReq) (*co_model.TeamRes, error) {
	return c.ITeam.UpdateTeam(ctx, &req.UpdateTeamReq)
}

func (c *TeamController) DeleteTeam(ctx context.Context, req *co_v1.DeleteTeamReq) (api_v1.BoolRes, error) {
	return c.ITeam.DeleteTeam(ctx, &req.DeleteTeamReq)
}

func (c *TeamController) GetTeamMemberList(ctx context.Context, req *co_v1.GetTeamMemberListReq) (*co_model.EmployeeListRes, error) {
	return c.ITeam.GetTeamMemberList(ctx, &req.GetTeamMemberListReq)
}

func (c *TeamController) QueryTeamListByEmployee(ctx context.Context, req *co_v1.QueryTeamListByEmployeeReq) (*co_model.TeamListRes, error) {
	return c.ITeam.QueryTeamListByEmployee(ctx, &req.QueryTeamListByEmployeeReq)
}

func (c *TeamController) SetTeamMember(ctx context.Context, req *co_v1.SetTeamMemberReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamMember(ctx, &req.SetTeamMemberReq)
}

func (c *TeamController) SetTeamOwner(ctx context.Context, req *co_v1.SetTeamOwnerReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamOwner(ctx, &req.SetTeamOwnerReq)
}

func (c *TeamController) SetTeamCaptain(ctx context.Context, req *co_v1.SetTeamCaptainReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamCaptain(ctx, &req.SetTeamCaptainReq)
}
