package controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/utility/kconv"
)

type CompanyController[TIRes co_model.ICompanyRes] struct {
	i_controller.ICompany[TIRes]
}

func Company[
	TIRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	TIRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) *CompanyController[TIRes] {
	return &CompanyController[TIRes]{
		ICompany: co_controller.Company(modules),
	}
}

// GetCompanyById 通过ID获取公司信息
func (c *CompanyController[TIRes]) GetCompanyById(ctx context.Context, req *co_v1.GetCompanyByIdReq) (*co_model.CompanyRes, error) {
	ret, err := c.ICompany.GetCompanyById(ctx, &req.GetCompanyByIdReq)
	return ret.Data(), err
}

// HasCompanyByName 公司名称是否存在
func (c *CompanyController[TIRes]) HasCompanyByName(ctx context.Context, req *co_v1.HasCompanyByNameReq) (api_v1.BoolRes, error) {
	return c.ICompany.HasCompanyByName(ctx, &req.HasCompanyByNameReq)
}

// QueryCompanyList 查询公司列表
func (c *CompanyController[TIRes]) QueryCompanyList(ctx context.Context, req *co_v1.QueryCompanyListReq) (*co_model.CompanyListRes, error) {
	ret, err := c.ICompany.QueryCompanyList(ctx, &req.QueryCompanyListReq)
	return kconv.Struct(ret, &co_model.CompanyListRes{}), err
}

// CreateCompany 创建公司信息
func (c *CompanyController[TIRes]) CreateCompany(ctx context.Context, req *co_v1.CreateCompanyReq) (*co_model.CompanyRes, error) {
	ret, err := c.ICompany.CreateCompany(ctx, &req.CreateCompanyReq)
	return ret.Data(), err
}

// UpdateCompany 更新公司信息
func (c *CompanyController[TIRes]) UpdateCompany(ctx context.Context, req *co_v1.UpdateCompanyReq) (*co_model.CompanyRes, error) {
	ret, err := c.ICompany.UpdateCompany(ctx, &req.UpdateCompanyReq)
	return ret.Data(), err
}

// GetCompanyDetail 查看更多信息含完整手机号
func (c *CompanyController[TIRes]) GetCompanyDetail(ctx context.Context, req *co_v1.GetCompanyDetailReq) (*co_model.CompanyRes, error) {
	ret, err := c.ICompany.GetCompanyDetail(ctx, &req.GetCompanyDetailReq)
	return ret.Data(), err
}
