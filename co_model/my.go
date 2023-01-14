package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type MyCompanyRes co_entity.Company

type MyTeamRes struct {
	co_entity.CompanyTeam
	EmployeeListRes `json:"memberItems" description:"团队或小组成员"`
}

type MyTeamListRes []MyTeamRes

type MyProfileRes struct {
	sys_entity.SysUser
	*co_entity.CompanyEmployee `json:"employeeInfo" description:"员工信息"`
}
