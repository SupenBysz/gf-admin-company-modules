package co_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type GetTeamByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取团队或小组｜信息" tags:"组织单位/团队｜小组"`
	co_company_api.GetTeamByIdReq
}

type HasTeamByNameReq struct {
	g.Meta `method:"post" summary:"判断名称是否存在" tags:"组织单位/团队｜小组"`
	co_company_api.HasTeamByNameReq
}

type QueryTeamListReq struct {
	g.Meta `method:"post" summary:"查询团队或小组｜列表" tags:"组织单位/团队｜小组"`
	co_company_api.QueryTeamListReq
}

type CreateTeamReq struct {
	g.Meta `method:"post" summary:"创建团队或小组｜信息" tags:"组织单位/团队｜小组"`
	co_company_api.CreateTeamReq
}

type UpdateTeamReq struct {
	g.Meta `method:"post" summary:"更新团队或小组｜信息" tags:"组织单位/团队｜小组"`
	co_company_api.UpdateTeamReq
}

type DeleteTeamReq struct {
	g.Meta `method:"post" summary:"删除团队或小组｜信息" tags:"组织单位/团队｜小组"`
	co_company_api.DeleteTeamReq
}

type QueryTeamListByEmployeeReq struct {
	g.Meta `method:"post" summary:"根据员工查询团队｜列表" tags:"组织单位/团队｜小组"`
	co_company_api.QueryTeamListByEmployeeReq
}

type SetTeamMemberReq struct {
	g.Meta `method:"post" summary:"设置团队成员" tags:"组织单位/团队｜小组"`
	co_company_api.SetTeamMemberReq
}

type SetTeamOwnerReq struct {
	g.Meta `method:"post" summary:"设置团队管理者" tags:"组织单位/团队｜小组"`
	co_company_api.SetTeamOwnerReq
}

type SetTeamCaptainReq struct {
	g.Meta `method:"post" summary:"设置团队队长" tags:"组织单位/团队｜小组"`
	co_company_api.SetTeamCaptainReq
}

type GetEmployeeListByTeamIdReq struct {
	g.Meta `method:"post" summary:"根据团队ID获取成员｜列表" tags:"组织单位/团队｜小组"`
	co_company_api.GetEmployeeListByTeamIdReq
}

type GetTeamInviteCodeReq struct {
	g.Meta `method:"post" summary:"查看团队邀约码" tags:"组织单位/团队｜小组"`
	co_company_api.GetTeamInviteCodeReq
}

type JoinTeamByInviteCodeReq struct {
	g.Meta `method:"post" summary:"通过邀请码加入团队" tags:"组织单位/团队｜小组"`
	co_company_api.JoinTeamByInviteCodeReq
}
