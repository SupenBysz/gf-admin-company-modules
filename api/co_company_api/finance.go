package co_company_api

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type BankCardRegisterReq struct {
	co_model.BankCardRegister
}

type DeleteBankCardReq struct {
	BankCardId int64 `json:"bankCardId" dc:"银行卡ID"`
}

type QueryBankCardListReq struct {
	UserId int64 `json:"userId" dc:"用户ID"`
	base_model.SearchParams
}

type GetAccountBalanceReq struct {
	AccountId int64 `json:"accountId" dc:"财务账号id"`
}

type CreateInvoiceReq struct {
	co_model.FdInvoiceRegister
}

type QueryInvoiceReq struct {
	UserId int64 `json:"userId"  dc:"用户ID" v:"required#请输入用户ID"'`
	base_model.SearchParams
}

type DeleteInvoiceByIdReq struct {
	InvoiceId int64 `json:"invoiceId" dc:"发票抬头ID"`
}

type CreateInvoiceDetailReq struct {
	co_model.FdInvoiceDetailRegister
}

type QueryInvoiceDetailListReq struct {
	UnionMainId int64 `json:"unionMainId" dc:"主体ID"`
	UserId      int64 `json:"userId" dc:"用户ID"`
	base_model.SearchParams
}

type MakeInvoiceDetailReq struct {
	InvoiceDetailId int64 `json:"invoiceDetailId"  v:"required|max-length:64#请输入发票详情id|id最大支持64个字符" dc:"发票详情id"`
	co_model.FdMakeInvoiceDetail
}

type AuditInvoiceDetailReq struct {
	InvoiceDetailId int64 `json:"invoiceDetailId"  v:"required|max-length:64#请输入发票详情id|id最大支持64个字符" dc:"发票详情id" `
	AuditInfo       co_model.FdInvoiceAuditInfo
}

type GetAccountDetailReq struct {
	AccountId int64    `json:"accountId" dc:"财务账号id"`
	Include   []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type UpdateAccountIsEnabledReq struct {
	AccountId int64 `json:"accountId" dc:"财务账号id"`
	IsEnabled int   `json:"isEnabled" dc:"是否启用: 1启用、0禁用"`
}

type UpdateAccountLimitStateReq struct {
	AccountId  int64 `json:"accountId" dc:"财务账号id"`
	LimitState int   `json:"limitState" dc:"限制状态：0不限制，1限制支出、2限制收入"`
}

type SetAccountCurrencyCodeReq struct {
	AccountId    int64  `json:"accountId" dc:"财务账号id"`
	CurrencyCode string `json:"currencyCode" dc:"货币代码"`
}

type GetAccountDetailByAccountIdReq struct {
	AccountId int64 `json:"accountId" dc:"财务账号id"`
}

// IncrementReq TODO 这两个接口只是用于测试，后续需要去除
type IncrementReq struct {
	AccountId int64 `json:"accountId" dc:"财务账号id"`
	Amount    int   `json:"amount" dc:"收入金额"`
}

// DecrementReq TODO 这两个接口只是用于测试，后续需要去除
type DecrementReq struct {
	AccountId int64 `json:"accountId" dc:"财务账号id"`
	Amount    int   `json:"amount" dc:"支出金额"`
}

type SetAccountAllowExceedReq struct {
	AccountId   int64 `json:"accountId" dc:"财务账号ID"  v:"required#财务账号不允许为空"`
	AllowExceed int   `json:"allowExceed" dc:"是否允许存在负余额: 0禁止、1允许，主体默认财务账号是允许超出" v:"required|in:0,1#请正确填写是否允许"`
}

type UpdateAccountBalanceReq struct {
	AccountId int64 `json:"accountId" dc:"财务账号id" v:"required#请输入财务账号id"`
	Amount    int64 `json:"amount" dc:"冲正金额"`
}

type QueryAccountBillsReq struct {
	AccountId int64 `json:"accountId" dc:"财务账号id" v:"required#请输入财务账号id"`
	base_model.SearchParams
}
