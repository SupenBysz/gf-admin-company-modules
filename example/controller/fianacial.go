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

type FinancialController struct {
	i_controller.IFinancial
}

var Financial = func(modules co_interface.IModules) *FinancialController {
	return &FinancialController{
		co_controller.Financial(modules),
	}
}

func (c *FinancialController) GetModules() co_interface.IModules {
	return c.IFinancial.GetModules()
}

func (c *FinancialController) GetAccountBalance(ctx context.Context, req *co_v1.GetAccountBalanceReq) (api_v1.Int64Res, error) {
	return c.IFinancial.GetAccountBalance(ctx, &req.GetAccountBalanceReq)
}

func (c *FinancialController) InvoiceRegister(ctx context.Context, req *co_v1.CreateInvoiceReq) (*co_model.FdInvoiceInfoRes, error) {
	return c.IFinancial.InvoiceRegister(ctx, &req.CreateInvoiceReq)
}

func (c *FinancialController) QueryInvoice(ctx context.Context, req *co_v1.QueryInvoiceReq) (*co_model.FdInvoiceListRes, error) {
	return c.IFinancial.QueryInvoice(ctx, &req.QueryInvoiceReq)
}

func (c *FinancialController) DeletesFdInvoiceById(ctx context.Context, req *co_v1.DeleteInvoiceByIdReq) (api_v1.BoolRes, error) {
	return c.IFinancial.DeletesFdInvoiceById(ctx, &req.DeleteInvoiceByIdReq)
}

func (c *FinancialController) InvoiceDetailRegister(ctx context.Context, req *co_v1.CreateInvoiceDetailReq) (*co_model.FdInvoiceDetailInfoRes, error) {
	return c.IFinancial.InvoiceDetailRegister(ctx, &req.CreateInvoiceDetailReq)
}

func (c *FinancialController) QueryInvoiceDetailList(ctx context.Context, req *co_v1.QueryInvoiceDetailListReq) (*co_model.FdInvoiceDetailListRes, error) {
	return c.IFinancial.QueryInvoiceDetailList(ctx, &req.QueryInvoiceDetailListReq)
}

func (c *FinancialController) MakeInvoiceDetailReq(ctx context.Context, req *co_v1.MakeInvoiceDetailReq) (api_v1.BoolRes, error) {
	return c.IFinancial.MakeInvoiceDetailReq(ctx, &req.MakeInvoiceDetailReq)
}

func (c *FinancialController) AuditInvoiceDetail(ctx context.Context, req *co_v1.AuditInvoiceDetailReq) (api_v1.BoolRes, error) {
	return c.IFinancial.AuditInvoiceDetail(ctx, &req.AuditInvoiceDetailReq)

}

func (c *FinancialController) BankCardRegister(ctx context.Context, req *co_v1.BankCardRegisterReq) (*co_model.BankCardInfoRes, error) {
	return c.IFinancial.BankCardRegister(ctx, &req.BankCardRegisterReq)
}

func (c *FinancialController) DeleteBankCard(ctx context.Context, req *co_v1.DeleteBankCardReq) (api_v1.BoolRes, error) {
	return c.IFinancial.DeleteBankCard(ctx, &req.DeleteBankCardReq)
}

func (c *FinancialController) QueryBankCardList(ctx context.Context, req *co_v1.QueryBankCardListReq) (*co_model.BankCardListRes, error) {
	return c.IFinancial.QueryBankCardList(ctx, &req.QueryBankCardListReq)
}
