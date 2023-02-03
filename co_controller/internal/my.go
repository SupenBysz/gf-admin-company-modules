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

type MyController struct {
	i_controller.IMy
	modules co_interface.IModules
}

var My = func(modules co_interface.IModules) i_controller.IMy {
	return &MyController{
		modules: modules,
	}
}

func (c *MyController) GetModules() co_interface.IModules {
	return c.modules
}

// GetProfile 获取当前员工及用户信息
func (c *MyController) GetProfile(ctx context.Context, _ *co_company_api.GetProfileReq) (*co_model.MyProfileRes, error) {
	result, err := c.modules.My().GetProfile(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCompany 获取当前公司信息
func (c *MyController) GetCompany(ctx context.Context, _ *co_company_api.GetCompanyReq) (*co_model.MyCompanyRes, error) {
	result, err := c.modules.My().GetCompany(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// GetTeams 获取当前团队信息
func (c *MyController) GetTeams(ctx context.Context, _ *co_company_api.GetTeamsReq) (co_model.MyTeamListRes, error) {

	result, err := c.modules.My().GetTeams(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// SetAvatar 设置员工头像
func (c *MyController) SetAvatar(ctx context.Context, req *co_company_api.SetAvatarReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Employee().SetEmployeeAvatar(ctx, req.ImageId)
			return ret == true, err
		},
		co_enum.Employee.PermissionType(c.modules).SetAvatar,
	)
}
