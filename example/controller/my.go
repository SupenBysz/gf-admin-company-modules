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

type MyController struct {
	i_controller.IMy
}

var My = func(modules co_interface.IModules) *MyController {
	return &MyController{
		co_controller.My(modules),
	}
}

func (c *MyController) GetModules() co_interface.IModules {
	return c.IMy.GetModules()
}

// GetProfile 获取当前员工及用户信息
func (c *MyController) GetProfile(ctx context.Context, req *co_v1.GetProfileReq) (*co_model.MyProfileRes, error) {
	return c.IMy.GetProfile(ctx, &req.GetProfileReq)
}

// GetCompany 获取当前公司信息
func (c *MyController) GetCompany(ctx context.Context, req *co_v1.GetCompanyReq) (*co_model.MyCompanyRes, error) {
	return c.IMy.GetCompany(ctx, &req.GetCompanyReq)
}

// GetTeams 获取当前团队信息
func (c *MyController) GetTeams(ctx context.Context, req *co_v1.GetTeamsReq) (co_model.MyTeamListRes, error) {
	return c.IMy.GetTeams(ctx, &req.GetTeamsReq)
}

// SetAvatar 设置员工头像
func (c *MyController) SetAvatar(ctx context.Context, req *co_v1.SetAvatarReq) (api_v1.BoolRes, error) {
	return c.IMy.SetAvatar(ctx, &req.SetAvatarReq)
}

// SetMobile 设置手机号
func (c *MyController) SetMobile(ctx context.Context, req *co_v1.SetMobileReq) (api_v1.BoolRes, error) {
	return c.IMy.SetMobile(ctx, &req.SetMobileReq)
}
