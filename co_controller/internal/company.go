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

type CompanyController[T co_interface.IModules] struct {
	i_controller.ICompany
	modules T
}

var Company = func(modules co_interface.IModules) i_controller.ICompany {
	return &CompanyController[co_interface.IModules]{
		modules: modules,
	}
}

func (c *CompanyController[T]) GetModules() co_interface.IModules {
	return c.modules
}

// GetCompanyById 通过ID获取公司信息
func (c *CompanyController[T]) GetCompanyById(ctx context.Context, req *co_company_api.GetCompanyByIdReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().GetCompanyById(ctx, req.Id)
			return (*co_model.CompanyRes)(ret), err
		},
		co_enum.Company.PermissionType(c.modules).ViewDetail,
	)
}

// HasCompanyByName 公司名称是否存在
func (c *CompanyController[T]) HasCompanyByName(ctx context.Context, req *co_company_api.HasCompanyByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Company().HasCompanyByName(ctx, req.Name) == true, nil
		},
	)
}

// QueryCompanyList 查询公司列表
func (c *CompanyController[T]) QueryCompanyList(ctx context.Context, req *co_company_api.QueryCompanyListReq) (*co_model.CompanyListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyListRes, error) {
			return c.modules.Company().QueryCompanyList(ctx, &req.SearchParams)
		},
		co_enum.Company.PermissionType(c.modules).List,
	)
}

// CreateCompany 创建公司信息
func (c *CompanyController[T]) CreateCompany(ctx context.Context, req *co_company_api.CreateCompanyReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().CreateCompany(ctx, &req.Company)
			return (*co_model.CompanyRes)(ret), err
		},
		co_enum.Company.PermissionType(c.modules).Create,
	)
}

// UpdateCompany 更新公司信息
func (c *CompanyController[T]) UpdateCompany(ctx context.Context, req *co_company_api.UpdateCompanyReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().UpdateCompany(ctx, &req.Company)
			return (*co_model.CompanyRes)(ret), err
		},
		co_enum.Company.PermissionType(c.modules).Update,
	)
}
