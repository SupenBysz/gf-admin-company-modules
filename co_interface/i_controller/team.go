package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type ITeam[ITTeamRes co_model.ITeamRes] interface {
	SetTeam(team co_interface.ITeam[*co_model.TeamRes], employee co_interface.IEmployee[*co_model.EmployeeRes])

	// GetTeamById 根据id获取团队信息
	GetTeamById(ctx context.Context, req *co_company_api.GetTeamByIdReq) (ITTeamRes, error)

	// HasTeamByName 判断团队名称是否存在
	HasTeamByName(ctx context.Context, req *co_company_api.HasTeamByNameReq) (api_v1.BoolRes, error)

	// QueryTeamList 查询团队列表
	QueryTeamList(ctx context.Context, req *co_company_api.QueryTeamListReq) (*base_model.CollectRes[ITTeamRes], error)

	// CreateTeam 创建团队
	CreateTeam(ctx context.Context, req *co_company_api.CreateTeamReq) (ITTeamRes, error)

	// UpdateTeam 更新团队信息
	UpdateTeam(ctx context.Context, req *co_company_api.UpdateTeamReq) (ITTeamRes, error)

	// DeleteTeam 删除团队信息
	DeleteTeam(ctx context.Context, req *co_company_api.DeleteTeamReq) (api_v1.BoolRes, error)

	// QueryTeamListByEmployee 根据员工获取团队列表
	QueryTeamListByEmployee(ctx context.Context, req *co_company_api.QueryTeamListByEmployeeReq) (*base_model.CollectRes[ITTeamRes], error)

	// SetTeamMember 设置团队成员
	SetTeamMember(ctx context.Context, req *co_company_api.SetTeamMemberReq) (api_v1.BoolRes, error)

	// SetTeamOwner 设置团队负责人
	SetTeamOwner(ctx context.Context, req *co_company_api.SetTeamOwnerReq) (api_v1.BoolRes, error)

	// SetTeamCaptain 设置团队队长或组长
	SetTeamCaptain(ctx context.Context, req *co_company_api.SetTeamCaptainReq) (api_v1.BoolRes, error)
}
