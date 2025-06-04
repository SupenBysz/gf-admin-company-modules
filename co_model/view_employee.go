package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type EmployeeView struct {
	co_entity.CompanyEmployeeView
}

type EmployeeViewRes struct {
	EmployeeView `json:"employeeView"`
	TeamList     *[]TeamViewRes         `json:"teamList" dc:"团队列表"`
	User         *sys_model.SysUser     `json:"user" dc:"用户"`
	UnionMain    *co_entity.CompanyView `json:"unionMain" dc:"所属单位"`
}

type EmployeeViewListRes base_model.CollectRes[EmployeeViewRes]
