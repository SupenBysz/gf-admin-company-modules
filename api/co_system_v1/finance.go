package co_system_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetCurrencyByCodeReq struct {
	g.Meta       `path:"/getCurrencyByCode" method:"post" summary:"获取货币单位信息" tags:"系统/财务"`
	CurrencyCode string `json:"currencyCode" dc:"货币代码"`
}

type QueryCurrencyListReq struct {
	g.Meta `path:"/queryCurrencyList" method:"post" summary:"获取币种列表" tags:"系统/财务"`
	base_model.SearchParams
}

type QueryAccountRechargeViewReq struct {
	g.Meta `path:"/queryAccountRechargeView" method:"post" summary:"查询充值记录" tags:"系统/财务"`
	base_model.SearchParams
}

type GetAccountRechargeViewByIdReq struct {
	g.Meta `path:"/getAccountRechargeViewById" method:"post" summary:"根据ID获取充值记录" tags:"系统/财务"`
	Id     int64 `json:"id" dc:"充值记录ID"`
}
