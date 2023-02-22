package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type IFinancial[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] interface {
	GetAccountBalance(ctx context.Context, req *co_company_api.GetAccountBalanceReq) (api_v1.Int64Res, error)

	InvoiceRegister(ctx context.Context, req *co_company_api.CreateInvoiceReq) (ITFdInvoiceRes, error)

	QueryInvoice(ctx context.Context, req *co_company_api.QueryInvoiceReq) (*base_model.CollectRes[ITFdInvoiceRes], error)

	DeletesFdInvoiceById(ctx context.Context, req *co_company_api.DeleteInvoiceByIdReq) (api_v1.BoolRes, error)

	InvoiceDetailRegister(ctx context.Context, req *co_company_api.CreateInvoiceDetailReq) (ITFdInvoiceDetailRes, error)

	QueryInvoiceDetailList(ctx context.Context, req *co_company_api.QueryInvoiceDetailListReq) (*base_model.CollectRes[ITFdInvoiceDetailRes], error)

	MakeInvoiceDetailReq(ctx context.Context, req *co_company_api.MakeInvoiceDetailReq) (api_v1.BoolRes, error)

	AuditInvoiceDetail(ctx context.Context, req *co_company_api.AuditInvoiceDetailReq) (api_v1.BoolRes, error)

	BankCardRegister(ctx context.Context, req *co_company_api.BankCardRegisterReq) (ITFdBankCardRes, error)

	DeleteBankCard(ctx context.Context, req *co_company_api.DeleteBankCardReq) (api_v1.BoolRes, error)

	QueryBankCardList(ctx context.Context, req *co_company_api.QueryBankCardListReq) (*base_model.CollectRes[ITFdBankCardRes], error)
}
