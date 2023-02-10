package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

type CompanyController struct {
	i_controller.ICompany
	modules co_interface.IModules
	dao     *co_dao.XDao
}

var Company = func(modules co_interface.IModules) i_controller.ICompany {
	return &CompanyController{
		modules: modules,
		dao:     modules.Dao(),
	}
}

func (c *CompanyController) GetModules() co_interface.IModules {
	return c.modules
}

// GetCompanyById 通过ID获取公司信息
func (c *CompanyController) GetCompanyById(ctx context.Context, req *co_company_api.GetCompanyByIdReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().GetCompanyById(c.makeMore(ctx), req.Id)
			return ret, err
		},
		co_enum.Company.PermissionType(c.modules).ViewDetail,
	)
}

// HasCompanyByName 公司名称是否存在
func (c *CompanyController) HasCompanyByName(ctx context.Context, req *co_company_api.HasCompanyByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Company().HasCompanyByName(ctx, req.Name) == true, nil
		},
	)
}

// QueryCompanyList 查询公司列表
func (c *CompanyController) QueryCompanyList(ctx context.Context, req *co_company_api.QueryCompanyListReq) (*co_model.CompanyListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyListRes, error) {
			return c.modules.Company().QueryCompanyList(c.makeMore(ctx), &req.SearchParams)
		},
		co_enum.Company.PermissionType(c.modules).List,
	)
}

// CreateCompany 创建公司信息
func (c *CompanyController) CreateCompany(ctx context.Context, req *co_company_api.CreateCompanyReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().CreateCompany(c.makeMore(ctx), &req.Company)
			return ret, err
		},
		co_enum.Company.PermissionType(c.modules).Create,
	)
}

// UpdateCompany 更新公司信息
func (c *CompanyController) UpdateCompany(ctx context.Context, req *co_company_api.UpdateCompanyReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().UpdateCompany(c.makeMore(ctx), &req.Company)
			return ret, err
		},
		co_enum.Company.PermissionType(c.modules).Update,
	)
}

// GetCompanyDetail 获取公司详情，包含完整商务联系人电话
func (c *CompanyController) GetCompanyDetail(ctx context.Context, req *co_company_api.GetCompanyDetailReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().GetCompanyDetail(c.makeMore(ctx), req.Id)
			return ret, err
		},
		co_enum.Company.PermissionType(c.modules).ViewMobile,
	)
}

func (c *CompanyController) makeMore(ctx context.Context) context.Context {
	return funs.AttrBuilder[co_model.CompanyRes, co_model.EmployeeRes](ctx, c.dao.Company.Columns().UserId)
}
