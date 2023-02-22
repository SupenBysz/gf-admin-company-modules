package co_company_api

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type GetEmployeeByIdReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type HasEmployeeByNameReq struct {
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionNameId int64  `json:"unionNameId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的员工ID"`
}

type HasEmployeeByNoReq struct {
	No        string `json:"no" dc:"工号"`
	ExcludeId int64  `json:"excludeId" dc:"要排除的员工ID"`
}

type QueryEmployeeListReq struct {
	base_model.SearchParams
}

type CreateEmployeeReq struct {
	co_model.Employee
}

type UpdateEmployeeReq struct {
	co_model.Employee
}

type DeleteEmployeeReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type GetEmployeeDetailByIdReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type GetEmployeeListByRoleIdReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"角色ID"`
}

type GetEmployeeListByTeamId struct {
	TeamId int64 `json:"teamId" v:"required#团队ID校验失败" dc:"团队或小组ID"`
}
