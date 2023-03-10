package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type ICompany interface {
	iModule
	// GetCompanyById 通过ID获取公司信息
	GetCompanyById(ctx context.Context, req *co_company_api.GetCompanyByIdReq) (*co_model.CompanyRes, error)
	// HasCompanyByName 公司名称是否存在
	HasCompanyByName(ctx context.Context, req *co_company_api.HasCompanyByNameReq) (api_v1.BoolRes, error)
	// QueryCompanyList 查询公司列表
	QueryCompanyList(ctx context.Context, req *co_company_api.QueryCompanyListReq) (*co_model.CompanyListRes, error)
	// CreateCompany 创建公司信息
	CreateCompany(ctx context.Context, req *co_company_api.CreateCompanyReq) (*co_model.CompanyRes, error)
	// UpdateCompany 更新公司信息
	UpdateCompany(ctx context.Context, req *co_company_api.UpdateCompanyReq) (*co_model.CompanyRes, error)
	// GetCompanyDetail 获取公司详情，包含完整商务联系人电话
	GetCompanyDetail(ctx context.Context, req *co_company_api.GetCompanyDetailReq) (*co_model.CompanyRes, error)
}
