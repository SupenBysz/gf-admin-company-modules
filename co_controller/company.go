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

type cCompany[T co_interface.IModules] struct {
	modules T
}

var Company = func(modules co_interface.IModules) *cCompany[co_interface.IModules] {
	return &cCompany[co_interface.IModules]{
		modules: modules,
	}
}

// GetCompanyById 通过ID获取公司信息
func (c *cCompany[T]) GetCompanyById(ctx context.Context, req *co_v1.GetCompanyByIdReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().GetCompanyById(ctx, req.Id)
			return (*co_model.CompanyRes)(ret), err
		},
		co_enum.Company.PermissionType(c.modules).ViewDetail,
	)
}

// HasCompanyByName 公司名称是否存在
func (c *cCompany[T]) HasCompanyByName(ctx context.Context, req *co_v1.HasCompanyByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Company().HasCompanyByName(ctx, req.Name) == true, nil
		},
	)
}

// QueryCompanyList 查询公司列表
func (c *cCompany[T]) QueryCompanyList(ctx context.Context, req *co_v1.QueryCompanyListReq) (*co_model.CompanyListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyListRes, error) {
			return c.modules.Company().QueryCompanyList(ctx, &req.SearchParams)
		},
		co_enum.Company.PermissionType(c.modules).List,
	)
}

// CreateCompany 创建公司信息
func (c *cCompany[T]) CreateCompany(ctx context.Context, req *co_v1.CreateCompanyReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().CreateCompany(ctx, &req.Company)
			return (*co_model.CompanyRes)(ret), err
		},
		co_enum.Company.PermissionType(c.modules).Create,
	)
}

// UpdateCompany 更新公司信息
func (c *cCompany[T]) UpdateCompany(ctx context.Context, req *co_v1.UpdateCompanyReq) (*co_model.CompanyRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.CompanyRes, error) {
			ret, err := c.modules.Company().UpdateCompany(ctx, &req.Company)
			return (*co_model.CompanyRes)(ret), err
		},
		co_enum.Company.PermissionType(c.modules).Update,
	)
}
