package co_company_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type GetCompanyByIdReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"服务商ID"`
}

type HasCompanyByNameReq struct {
	Name string `json:"name" v:"required#名称不能为空" dc:"名称"`
}

type QueryCompanyListReq struct {
	sys_model.SearchParams
}

type CreateCompanyReq struct {
	co_model.Company
}

type UpdateCompanyReq struct {
	co_model.Company
}

type GetCompanyDetailReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"服务商ID"`
}
