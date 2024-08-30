package co_company_api

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type GetCompanyByIdReq struct {
	Id      int64    `json:"id" v:"required#ID校验失败" dc:"组织单位ID"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type HasCompanyByNameReq struct {
	Name string `json:"name" v:"required#名称不能为空" dc:"名称"`
}

type QueryCompanyListReq struct {
	base_model.SearchParams
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
	IsExport bool `json:"isExport" dc:"isExport为true的话，会输出列表所有数据"`
}

type CreateCompanyReq struct {
	co_model.Company
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type UpdateCompanyReq struct {
	co_model.Company
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type GetCompanyDetailReq struct {
	Id      int64    `json:"id" v:"required#ID校验失败" dc:"组织单位ID"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type SetCompanyStateReq struct {
	Id    int64 `json:"id" v:"required#ID校验失败" dc:"组织单位ID"`
	State int   `json:"state" v:"in:0,1,2" dc:"状态：0未激活，1正常，-1停用"`
}
