package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type IFinancial interface {
	iModule

	GetAccountBalance(ctx context.Context, req *co_company_api.GetAccountBalanceReq) (api_v1.Int64Res, error)

	InvoiceRegister(ctx context.Context, req *co_company_api.CreateInvoiceReq) (*co_model.FdInvoiceInfoRes, error)

	QueryInvoice(ctx context.Context, req *co_company_api.QueryInvoiceReq) (*co_model.FdInvoiceListRes, error)

	DeletesFdInvoiceById(ctx context.Context, req *co_company_api.DeleteInvoiceByIdReq) (api_v1.BoolRes, error)

	InvoiceDetailRegister(ctx context.Context, req *co_company_api.CreateInvoiceDetailReq) (*co_model.FdInvoiceDetailInfoRes, error)

	QueryInvoiceDetailList(ctx context.Context, req *co_company_api.QueryInvoiceDetailListReq) (*co_model.FdInvoiceDetailListRes, error)

	MakeInvoiceDetailReq(ctx context.Context, req *co_company_api.MakeInvoiceDetailReq) (api_v1.BoolRes, error)

	AuditInvoiceDetail(ctx context.Context, req *co_company_api.AuditInvoiceDetailReq) (api_v1.BoolRes, error)

	BankCardRegister(ctx context.Context, req *co_company_api.BankCardRegisterReq) (*co_model.BankCardInfoRes, error)

	DeleteBankCard(ctx context.Context, req *co_company_api.DeleteBankCardReq) (api_v1.BoolRes, error)

	QueryBankCardList(ctx context.Context, req *co_company_api.QueryBankCardListReq) (*co_model.BankCardListRes, error)
}
