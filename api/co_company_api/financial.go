package co_company_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type BankCardRegisterReq struct {
	co_model.BankCardRegister
}

type DeleteBankCardReq struct {
	BankCardId int64
}

type QueryBankCardListReq struct {
	UserId int64
	sys_model.SearchParams
}

type GetAccountBalanceReq struct {
	AccountId int64
}

type CreateInvoiceReq struct {
	co_model.FdInvoiceRegister
}

type QueryInvoiceReq struct {
	UserId int64
	sys_model.SearchParams
}

type DeleteInvoiceByIdReq struct {
	InvoiceId int64
}

type CreateInvoiceDetailReq struct {
	co_model.FdInvoiceDetailRegister
}

type QueryInvoiceDetailListReq struct {
	UnionMainId int64
	UserId      int64
	sys_model.SearchParams
}

type MakeInvoiceDetailReq struct {
	InvoiceDetailId int64 `json:"invoiceDetailId"  v:"required|max-length:64#请输入发票详情id|id最大支持64个字符" dc:"发票详情id"`
	co_model.FdMakeInvoiceDetail
}

type AuditInvoiceDetailReq struct {
	InvoiceDetailId int64 `json:"invoiceDetailId"  v:"required|max-length:64#请输入发票详情id|id最大支持64个字符" dc:"发票详情id" `
	AuditInfo       co_model.FdInvoiceAuditInfo
}

// type InvoiceDetailListRes []co_model.FdInvoiceDetailListRes