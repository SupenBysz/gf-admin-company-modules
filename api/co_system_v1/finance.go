package co_system_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetCurrencyByCodeReq struct {
	g.Meta       `path:"/getCurrencyByCode" method:"post" summary:"获取货币单位｜信息" tags:"系统/财务"`
	CurrencyCode string `json:"currencyCode" dc:"货币代码"`
}

type QueryCurrencyListReq struct {
	g.Meta `path:"/queryCurrencyList" method:"post" summary:"获取币种｜列表" tags:"系统/财务"`
	base_model.SearchParams
}

type QueryAccountRechargeViewReq struct {
	g.Meta `path:"/queryAccountRechargeView" method:"post" summary:"查询充值记录｜信息" tags:"系统/财务"`
	base_model.SearchParams
}

type GetAccountRechargeViewByIdReq struct {
	g.Meta `path:"/getAccountRechargeViewById" method:"post" summary:"根据充值ID获取充值记录｜信息" tags:"系统/财务"`
	Id     int64 `json:"id" v:"required#ID参数错误" dc:"充值记录ID"`
}

type GetRechargeByAccountIdReq struct {
	g.Meta `path:"/getRechargeByAccountId" method:"post" summary:"根据资金账户ID获取充值记录｜列表" tags:"系统/财务"`
	Id     int64 `json:"id" v:"required#ID参数错误" dc:"资金账户ID"`
}

type GetRechargeByUserIdReq struct {
	g.Meta `path:"/getRechargeByUserId" method:"post" summary:"根据用户ID获取充值记录｜列表" tags:"系统/财务"`
	Id     int64 `json:"id" v:"required#ID参数错误" dc:"用户ID"`
}

type GetRechargeByCompanyIdReq struct {
	g.Meta `path:"/getRechargeByCompanyId" method:"post" summary:"根据公司ID获取充值记录｜列表" tags:"系统/财务"`
	Id     int64 `json:"id" v:"required#ID参数错误" dc:"公司ID"`
}
