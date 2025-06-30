package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type TeamMemberView struct {
	co_entity.CompanyTeamMemberView
}

type TeamMemberViewRes struct {
	TeamMemberView
	Team       *co_entity.CompanyTeamView     `json:"team" dc:"团队信息"`
	Employee   *co_entity.CompanyEmployeeView `json:"employee" dc:"员工信息"`
	InviteUser *co_entity.CompanyEmployeeView `json:"inviteUser" dc:"邀人人"`
	UnionMain  *co_entity.CompanyView         `json:"unionMain" dc:"所属单位"`
}

type TeamViewMemberListRes base_model.CollectRes[TeamMemberViewRes]
