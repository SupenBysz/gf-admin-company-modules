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

type CompanyController struct {
	i_controller.ICompany
}

var Company = func(modules co_interface.IModules) *CompanyController {
	return &CompanyController{
		co_controller.Company(modules),
	}
}

func (c *CompanyController) GetModules() co_interface.IModules {
	return c.ICompany.GetModules()
}

// GetCompanyById 通过ID获取公司信息
func (c *CompanyController) GetCompanyById(ctx context.Context, req *co_v1.GetCompanyByIdReq) (*co_model.CompanyRes, error) {
	return c.ICompany.GetCompanyById(ctx, &req.GetCompanyByIdReq)
}

// HasCompanyByName 公司名称是否存在
func (c *CompanyController) HasCompanyByName(ctx context.Context, req *co_v1.HasCompanyByNameReq) (api_v1.BoolRes, error) {
	return c.ICompany.HasCompanyByName(ctx, &req.HasCompanyByNameReq)
}

// QueryCompanyList 查询公司列表
func (c *CompanyController) QueryCompanyList(ctx context.Context, req *co_v1.QueryCompanyListReq) (*co_model.CompanyListRes, error) {
	return c.ICompany.QueryCompanyList(ctx, &req.QueryCompanyListReq)
}

// CreateCompany 创建公司信息
func (c *CompanyController) CreateCompany(ctx context.Context, req *co_v1.CreateCompanyReq) (*co_model.CompanyRes, error) {
	return c.ICompany.CreateCompany(ctx, &req.CreateCompanyReq)
}

// UpdateCompany 更新公司信息
func (c *CompanyController) UpdateCompany(ctx context.Context, req *co_v1.UpdateCompanyReq) (*co_model.CompanyRes, error) {
	return c.ICompany.UpdateCompany(ctx, &req.UpdateCompanyReq)
}
