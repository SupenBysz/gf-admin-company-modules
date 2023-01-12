package co_v1

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetTeamByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取团队或小组|信息" tags:"团队|小组"`
	Id     int64 `json:"id" v:"required#团队ID校验失败" dc:"团队或小组ID"`
}

type HasTeamByNameReq struct {
	g.Meta      `method:"post" summary:"判断名称是否存在" tags:"团队|小组"`
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionMainId int64  `json:"unionMainId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的团队ID"`
}

type QueryTeamListReq struct {
	g.Meta `method:"post" summary:"查询团队或小组|列表" tags:"团队|小组"`
	sys_model.SearchParams
}

type CreateTeamReq struct {
	g.Meta `method:"post" summary:"创建团队或小组|信息" tags:"团队|小组"`
	co_model.Team
}

type UpdateTeamReq struct {
	g.Meta `method:"post" summary:"更新团队或小组|信息" tags:"团队|小组"`
	Id     int64  `json:"id" v:"required#团队ID校验失败" dc:"团队或小组ID"`
	Name   string `json:"name" v:"required#名称不能为空" dc:"名称"`
	Remark string `json:"remark" dc:"备注"`
}

type DeleteTeamReq struct {
	g.Meta `method:"post" summary:"删除团队或小组|信息" tags:"团队|小组"`
	Id     int64 `json:"id" v:"required#团队ID校验失败" dc:"团队或小组ID"`
}

type GetTeamMemberListReq struct {
	g.Meta `method:"post" summary:"获取团队成员|列表" tags:"团队|小组"`
	Id     int64 `json:"id" v:"required#团队ID校验失败" dc:"团队或小组ID"`
}

type QueryTeamListByEmployeeReq struct {
	g.Meta      `method:"post" summary:"根据员工查询团队|列表" tags:"团队|小组"`
	EmployeeId  int64 `json:"employeeId" v:"required#员工ID校验失败" dc:"员工ID"`
	UnionMainId int64 `json:"unionMainId" dc:"关联主体，默认当前主体"`
}

type SetTeamMemberReq struct {
	g.Meta      `method:"post" summary:"设置团队成员" tags:"团队|小组"`
	Id          int64   `json:"id" v:"required#团队ID校验失败" dc:"团队或小组ID"`
	EmployeeIds []int64 `json:"employeeIds" dc:"团队成员"`
}

type SetTeamOwnerReq struct {
	g.Meta     `method:"post" summary:"设置团队管理者" tags:"团队|小组"`
	Id         int64 `json:"id" v:"required#团队ID校验失败" dc:"团队或小组ID"`
	EmployeeId int64 `json:"employeeId" v:"required#团队管理者ID校验失败" dc:"团队管理者ID"`
}

type SetTeamCaptainReq struct {
	g.Meta     `method:"post" summary:"设置团队队长" tags:"团队|小组"`
	Id         int64 `json:"id" v:"required#团队ID校验失败" dc:"团队或小组ID"`
	EmployeeId int64 `json:"employeeId" v:"required#团队队长ID校验失败" dc:"团队队长ID"`
}
