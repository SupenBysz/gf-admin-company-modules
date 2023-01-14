package co_v1

import "github.com/gogf/gf/v2/frame/g"

type GetProfileReq struct {
	g.Meta `path:"/getProfile" method:"post" summary:"我的基本信息|信息" tags:"我的"`
}

type GetCompanyReq struct {
	g.Meta `path:"/getCompany" method:"post" summary:"我的公司|信息" tags:"我的"`
}

type GetTeamsReq struct {
	g.Meta `path:"/getTeams" method:"post" summary:"我的团队或小组|信息" tags:"我的"`
}
