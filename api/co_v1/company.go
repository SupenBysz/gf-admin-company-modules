package co_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type GetCompanyByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取组织单位|信息" tags:"组织单位/公司"`
	co_company_api.GetCompanyByIdReq
}

type HasCompanyByNameReq struct {
	g.Meta `method:"post" summary:"判断名称是否存在" tags:"组织单位/公司"`
	co_company_api.HasCompanyByNameReq
}

type QueryCompanyListReq struct {
	g.Meta `method:"post" summary:"查询组织单位|列表" tags:"组织单位/公司"`
	co_company_api.QueryCompanyListReq
}

type CreateCompanyReq struct {
	g.Meta `method:"post" summary:"创建组织单位|信息" tags:"组织单位/公司"`
	co_company_api.CreateCompanyReq
}

type UpdateCompanyReq struct {
	g.Meta `method:"post" summary:"更新组织单位|信息" tags:"组织单位/公司"`
	co_company_api.UpdateCompanyReq
}

type GetCompanyDetailReq struct {
	g.Meta `method:"post" summary:"查看更多信息含完整手机号|信息" tags:"组织单位/公司"`
	co_company_api.GetCompanyDetailReq
}
