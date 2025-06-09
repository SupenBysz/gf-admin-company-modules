package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type CompanyView struct {
	co_entity.CompanyView
}

type CompanyViewRes struct {
	CompanyView `json:"employeeView"`
	License     *co_entity.License             `json:"license" dc:"资质信息"`
	Employee    *co_entity.CompanyEmployeeView `json:"employee" dc:"员工信息"`
	User        *sys_model.SysUser             `json:"user" dc:"用户信息"`
}

type CompanyViewListRes base_model.CollectRes[CompanyViewRes]
