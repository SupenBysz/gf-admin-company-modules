package co_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type GetAccountRechargeByIdReq struct {
	g.Meta `path:"/getAccountRechargeById" method:"post" summary:"根据ID获取充值记录" tags:"组织单位/财务服务"`
	co_company_api.GetAccountRechargeByIdReq
}

type SetAccountRechargeAuditReq struct {
	g.Meta `path:"/setAccountRechargeAudit" method:"post" summary:"审核充值记录" tags:"组织单位/财务服务"`
	co_company_api.SetAccountRechargeAuditReq
}

type QueryAccountRechargeReq struct {
	g.Meta `path:"/queryAccountRecharge" method:"post" summary:"查询充值记录" tags:"组织单位/财务服务"`
	co_company_api.QueryAccountRechargeReq
}

type AccountRecharge struct {
	g.Meta `path:"/accountRecharge" method:"post" summary:"充值记录" tags:"组织单位/财务服务"`
	co_company_api.AccountRecharge
}
