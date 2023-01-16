package co_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type GetProfileReq struct {
	g.Meta `method:"post" summary:"我的基本信息|信息" tags:"我的"`
	*co_company_api.GetProfileReq
}

type GetCompanyReq struct {
	g.Meta `method:"post" summary:"我的公司|信息" tags:"我的"`
	*co_company_api.GetCompanyReq
}

type GetTeamsReq struct {
	g.Meta `method:"post" summary:"我的团队或小组|信息" tags:"我的"`
	*co_company_api.GetTeamsReq
}
