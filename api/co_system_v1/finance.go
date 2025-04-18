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
