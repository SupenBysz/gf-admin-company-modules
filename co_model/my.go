package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type MyCompanyRes co_entity.Company

type MyTeamRes struct {
	TeamRes
	EmployeeListRes `json:"memberItems" description:"团队或小组成员"`
}

type MyTeamListRes []MyTeamRes

type MyProfileRes struct {
	User     *sys_model.SysUser `json:"user" description:"员工信息"`
	Employee *EmployeeRes       `json:"employee" description:"员工信息"`
}
