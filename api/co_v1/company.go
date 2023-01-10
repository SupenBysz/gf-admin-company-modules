package co_v1

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetCompanyByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取组织单位|信息" tags:"组织单位"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"组织单位ID"`
}

type HasCompanyByNameReq struct {
	g.Meta `method:"post" summary:"判断名称是否存在" tags:"组织单位"`
	Name   string `json:"name" v:"required#名称不能为空" dc:"名称"`
}

type QueryCompanyListReq struct {
	g.Meta `method:"post" summary:"查询组织单位|列表" tags:"组织单位"`
	sys_model.SearchParams
}

type CreateCompanyReq struct {
	g.Meta `method:"post" summary:"创建组织单位|信息" tags:"组织单位"`
	co_model.Company
}

type UpdateCompanyReq struct {
	g.Meta `method:"post" summary:"更新组织单位|信息" tags:"组织单位"`
	co_model.Company
}
