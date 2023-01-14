package co_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type cMy struct {
	modules co_interface.IModules
}

var My = func(modules co_interface.IModules) *cMy {
	return &cMy{
		modules: modules,
	}
}

// GetProfile 获取当前员工及用户信息
func (c *cMy) GetProfile(ctx context.Context, _ *co_v1.GetProfileReq) (*co_model.MyProfileRes, error) {
	result, err := c.modules.My().GetProfile(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCompany 获取当前公司信息
func (c *cMy) GetCompany(ctx context.Context, _ *co_v1.GetCompanyReq) (*co_model.MyCompanyRes, error) {
	result, err := c.modules.My().GetCompany(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// GetTeams 获取当前团队信息
func (c *cMy) GetTeams(ctx context.Context, _ *co_v1.GetTeamsReq) (co_model.MyTeamListRes, error) {

	result, err := c.modules.My().GetTeams(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
