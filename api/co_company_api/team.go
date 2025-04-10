package co_company_api

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type GetTeamByIdReq struct {
	Id      int64    `json:"id" v:"required#部门|团队|小组ID校验失败" dc:"部门|团队|小组ID"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type HasTeamByNameReq struct {
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionNameId int64  `json:"unionNameId" dc:"关联主体ID"`
	//ParentId    int64  `json:"parentId" dc:"关联的上级团队ID，没有的话或不限制就不填写"`
	ExcludeId int64 `json:"excludeId" dc:"要排除的部门/团队/小组ID"`
}

type QueryTeamListReq struct {
	base_model.SearchParams
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type CreateTeamReq struct {
	co_model.Team
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type UpdateTeamReq struct {
	co_model.Team
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type DeleteTeamReq struct {
	Id int64 `json:"id" v:"required#部门|团队|小组ID校验失败" dc:"部门|团队|小组ID"`
}
type QueryTeamListByEmployeeReq struct {
	EmployeeId  int64    `json:"employeeId" v:"required#员工ID校验失败" dc:"员工ID"`
	UnionMainId int64    `json:"unionMainId" dc:"关联主体，默认当前主体"`
	Include     []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type SetTeamMemberReq struct {
	Id          int64   `json:"id" v:"required#部门|团队|小组ID校验失败" dc:"部门|团队|小组ID"`
	EmployeeIds []int64 `json:"employeeIds" dc:"团队成员"`
}

type RemoveTeamMemberReq struct {
	Id          int64   `json:"id" v:"required#部门|团队|小组ID校验失败" dc:"部门|团队|小组ID"`
	EmployeeIds []int64 `json:"employeeIds" dc:"团队成员ids"`
}

type SetTeamOwnerReq struct {
	Id         int64 `json:"id" v:"required#部门|团队|小组ID校验失败" dc:"部门|团队|小组ID"`
	EmployeeId int64 `json:"employeeId" v:"required#团队管理者ID校验失败" dc:"团队管理者ID"`
}

type SetTeamCaptainReq struct {
	Id         int64 `json:"id" v:"required#部门|团队|小组ID校验失败" dc:"部门|团队|小组ID"`
	EmployeeId int64 `json:"employeeId" v:"required#团队队长ID校验失败" dc:"团队队长ID"`
}

type GetEmployeeListByTeamIdReq struct {
	TeamId  int64    `json:"teamId" v:"required#部门|团队|小组ID校验失败" dc:"部门|团队|小组ID"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type GetTeamInviteCodeReq struct {
	TeamId int64 `json:"teamId" v:"required#部门|团队|小组ID校验失败" dc:"部门|团队|小组ID"`
}

type JoinTeamByInviteCodeReq struct {
	InviteCode string `json:"inviteCode" dc:"邀请码" v:"required#邀请码不能为空"`
}
