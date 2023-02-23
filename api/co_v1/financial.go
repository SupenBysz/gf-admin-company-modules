package co_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type BankCardRegisterReq struct {
	g.Meta ` method:"post" summary:"添加银行卡" tags:"财务服务"`
	co_company_api.BankCardRegisterReq
}

type DeleteBankCardReq struct {
	g.Meta ` method:"post" summary:"删除银行卡" tags:"财务服务"`
	co_company_api.DeleteBankCardReq
}

type QueryBankCardListReq struct {
	g.Meta ` method:"post" summary:"获取银行卡|列表" tags:"财务服务"`
	co_company_api.QueryBankCardListReq
}

type GetAccountBalanceReq struct {
	g.Meta ` method:"post" summary:"查看账户余额" tags:"财务服务"`
	co_company_api.GetAccountBalanceReq
}

type CreateInvoiceReq struct {
	g.Meta ` method:"post" summary:"添加发票抬头" tags:"财务服务"`
	co_company_api.CreateInvoiceReq
}

type QueryInvoiceReq struct {
	g.Meta ` method:"post" summary:"获取发票抬头|列表" tags:"财务服务"`
	co_company_api.QueryInvoiceReq
}

type DeleteInvoiceByIdReq struct {
	g.Meta ` method:"post" summary:"删除发票抬头" tags:"财务服务"`
	co_company_api.DeleteInvoiceByIdReq
}

type CreateInvoiceDetailReq struct {
	g.Meta ` method:"post" summary:"申请开发票" tags:"财务服务"`
	co_company_api.CreateInvoiceDetailReq
}

type QueryInvoiceDetailListReq struct {
	g.Meta ` method:"post" summary:"获取发票详情|列表" tags:"财务服务"`
	co_company_api.QueryInvoiceDetailListReq
}

type MakeInvoiceDetailReq struct {
	g.Meta ` method:"post" summary:"开发票" tags:"财务服务"`
	co_company_api.MakeInvoiceDetailReq
}

type AuditInvoiceDetailReq struct {
	g.Meta ` method:"post" summary:"审核发票" tags:"财务服务"`
	co_company_api.AuditInvoiceDetailReq
}
