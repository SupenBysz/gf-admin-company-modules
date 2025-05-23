package co_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type GetProfileReq struct {
	g.Meta `method:"post" summary:"我的基本信息|信息" tags:"组织单位/我的"`
	co_company_api.GetProfileReq
}

type GetCompanyReq struct {
	g.Meta `method:"post" summary:"我的公司|信息" tags:"组织单位/我的"`
	co_company_api.GetCompanyReq
}

type GetTeamsReq struct {
	g.Meta `method:"post" summary:"我的团队或小组|信息" tags:"组织单位/我的"`
	co_company_api.GetTeamsReq
}

type SetAvatarReq struct {
	g.Meta `method:"post" summary:"设置头像|信息" tags:"组织单位/我的"`
	co_company_api.SetAvatarReq
}

type SetMobileReq struct {
	g.Meta `method:"post" summary:"设置业务手机号|信息" tags:"组织单位/我的" dc:"注意：此手机号不是用于登陆的手机号，通常属于工作的手机号联系方式"`
	co_company_api.SetMobileReq
}

type SetMailReq struct {
	g.Meta `method:"post" summary:"设置业务邮箱|信息" tags:"组织单位/我的" dc:"注意：此邮箱不是用于登陆的邮箱，通常属于工作的邮箱联系方式"`
	co_company_api.SetMailReq
}

type GetAccountBillsReq struct {
	g.Meta `method:"post" summary:"我的账单|列表" tags:"组织单位/我的财务"`
	co_company_api.GetAccountBillsReq
}

type GetAccountsReq struct {
	g.Meta `method:"post" summary:"我的财务账号|列表" tags:"组织单位/我的财务"`
	co_company_api.GetAccountsReq
}

type GetBankCardsReq struct {
	g.Meta `method:"post" summary:"我的银行卡｜列表" tags:"组织单位/我的财务"`
	co_company_api.GetBankCardsReq
}

type GetInvoicesReq struct {
	g.Meta `method:"post" summary:"我的发票抬头｜列表" tags:"组织单位/我的财务"`
	co_company_api.GetInvoicesReq
}

type UpdateAccountReq struct {
	g.Meta `method:"post" summary:"修改我的财务账号｜信息" tags:"组织单位/我的财务"`
	co_company_api.UpdateAccountReq
}
