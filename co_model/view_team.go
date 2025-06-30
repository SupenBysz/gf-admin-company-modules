package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type TeamView struct {
	co_entity.CompanyTeamView
}

type TeamViewRes struct {
	TeamView
	Owner     *co_entity.CompanyEmployeeView `json:"owner" dc:"部门｜团队所有者/业务总监/业务经理/团队队长"`
	Captain   *co_entity.CompanyEmployeeView `json:"captain" dc:"部门｜团队队长编号/小组组长"`
	UnionMain *co_entity.CompanyView         `json:"unionMain" dc:"所属单位"`
	Parent    *TeamViewRes                   `json:"parent" dc:"部门/团队/小组父级ID"`
}

type TeamViewListRes base_model.CollectRes[TeamViewRes]
