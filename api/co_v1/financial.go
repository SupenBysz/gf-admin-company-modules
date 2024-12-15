package co_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type BankCardRegisterReq struct {
	g.Meta ` method:"post" summary:"添加银行卡" tags:"组织单位/财务服务"`
	co_company_api.BankCardRegisterReq
}

type DeleteBankCardReq struct {
	g.Meta ` method:"post" summary:"删除银行卡" tags:"组织单位/财务服务"`
	co_company_api.DeleteBankCardReq
}

type QueryBankCardListReq struct {
	g.Meta ` method:"post" summary:"获取银行卡|列表" tags:"组织单位/财务服务"`
	co_company_api.QueryBankCardListReq
}

type GetAccountBalanceReq struct {
	g.Meta ` method:"post" summary:"查看账户余额" tags:"组织单位/财务服务"`
	co_company_api.GetAccountBalanceReq
}

type CreateInvoiceReq struct {
	g.Meta ` method:"post" summary:"添加发票抬头" tags:"组织单位/财务服务"`
	co_company_api.CreateInvoiceReq
}

type QueryInvoiceReq struct {
	g.Meta ` method:"post" summary:"获取发票抬头|列表" tags:"组织单位/财务服务"`
	co_company_api.QueryInvoiceReq
}

type DeleteInvoiceByIdReq struct {
	g.Meta ` method:"post" summary:"删除发票抬头" tags:"组织单位/财务服务"`
	co_company_api.DeleteInvoiceByIdReq
}

type CreateInvoiceDetailReq struct {
	g.Meta ` method:"post" summary:"申请开发票" tags:"组织单位/财务服务"`
	co_company_api.CreateInvoiceDetailReq
}

type QueryInvoiceDetailListReq struct {
	g.Meta ` method:"post" summary:"获取发票详情|列表" tags:"组织单位/财务服务"`
	co_company_api.QueryInvoiceDetailListReq
}

type MakeInvoiceDetailReq struct {
	g.Meta ` method:"post" summary:"开发票" tags:"组织单位/财务服务"`
	co_company_api.MakeInvoiceDetailReq
}

type AuditInvoiceDetailReq struct {
	g.Meta ` method:"post" summary:"审核发票" tags:"组织单位/财务服务"`
	co_company_api.AuditInvoiceDetailReq
}

type GetAccountDetailReq struct {
	g.Meta ` method:"post" summary:"获取财务账号详细数据" tags:"组织单位/财务服务"`
	co_company_api.GetAccountDetailReq
}

type UpdateAccountIsEnabledReq struct {
	g.Meta ` method:"post" summary:"修改财务账号启用状态" tags:"组织单位/财务服务"`
	co_company_api.UpdateAccountIsEnabledReq
}

type UpdateAccountLimitStateReq struct {
	g.Meta ` method:"post" summary:"修改财务账号限制状态" tags:"组织单位/财务服务"`
	co_company_api.UpdateAccountLimitStateReq
}

type GetAccountDetailByAccountIdReq struct {
	g.Meta ` method:"post" summary:"获取财务账号金额明细" tags:"组织单位/财务服务"`
	co_company_api.GetAccountDetailByAccountIdReq
}

type SetAccountAllowExceedReq struct {
	g.Meta ` method:"post" summary:"设置财务账号是否允许存在负余额" tags:"组织单位/财务服务"`
	co_company_api.SetAccountAllowExceedReq
}

type UpdateAccountBalanceReq struct {
	g.Meta ` method:"post" summary:"财务账号金额冲正" tags:"组织单位/财务服务"`
	co_company_api.UpdateAccountBalanceReq
}

type GetCurrencyByCodeReq struct {
	g.Meta ` method:"post" summary:"根据货币代码获取货币详情" tags:"组织单位/财务服务"`
	co_company_api.GetCurrencyByCodeReq
}

type QueryCurrencyListReq struct {
	g.Meta ` method:"post" summary:"获取货币列表" tags:"组织单位/财务服务"`
	co_company_api.QueryCurrencyListReq
}

type QueryAccountBillsReq struct {
	g.Meta ` method:"post" summary:"查询财务账号账单明细" tags:"组织单位/财务服务"`
	co_company_api.QueryAccountBillsReq
}
