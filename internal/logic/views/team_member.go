package views

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
)

type sTeamMemberView struct {
}

func init() {
	co_service.RegisterTeamMemberView(NewTeamMemberView())
}

func NewTeamMemberView() co_service.ITeamMemberView {
	return &sTeamMemberView{}
}

// GetTeamMemberById 根据团队成员ID获取团队成员详细信息。
// 该方法首先尝试从数据库中获取团队成员信息。如果找到团队成员信息且makeResource参数为true，
// 则进一步处理数据以生成更多的资源信息。
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
//	id int64: 团队成员的唯一标识符。
//	makeResource bool: 指示是否需要生成额外的资源信息。
//
// 返回值:
//
//	*co_model.TeamMemberViewRes: 团队成员详细信息的视图模型，如果找不到则返回nil。
//	error: 错误对象，如果操作成功则返回nil。
func (s *sTeamMemberView) GetTeamMemberById(ctx context.Context, id int64, makeResource bool) (*co_model.TeamMemberViewRes, error) {
	// 从数据库中获取团队成员详细信息。
	data, err := daoctl.GetByIdWithError[co_model.TeamMemberViewRes](co_dao.CompanyTeamMemberView.Ctx(ctx), id)

	// 如果没有错误且makeResource为true，则进一步处理数据以生成更多的资源信息。
	if err == nil && makeResource {
		data = s.makeMore(ctx, data, makeResource)
	}

	// 返回获取的团队成员信息或错误。
	return data, err
}

// QueryTeamMemberList 查询团队成员列表信息。
// 该方法根据提供的搜索参数查询团队成员信息，并可选地构建额外的资源信息。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息。
//	params - 搜索参数，用于指定查询的条件。
//	makeResource - 指示是否构建额外的资源信息。
//
// 返回值:
//
//	*co_model.TeamMemberViewListRes - 包含团队成员列表的响应对象。
//	error - 错误信息，如果执行过程中发生错误。
func (s *sTeamMemberView) QueryTeamMemberList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.TeamViewMemberListRes, error) {
	// 调用DAO层的方法来查询团队成员信息。
	data, err := daoctl.Query[co_model.TeamMemberViewRes](co_dao.CompanyTeamMemberView.Ctx(ctx), params, false)

	// 初始化结果对象，包含分页信息。
	result := &co_model.TeamViewMemberListRes{
		PaginationRes: data.PaginationRes,
	}

	// 如果没有错误且查询到了记录且要求构建资源信息，则为每条记录构建更多的资源信息。
	if err == nil && len(data.Records) > 0 && makeResource {
		for i, record := range data.Records {
			// 为每个记录构建更多的资源信息，并将其添加到结果中。
			data := s.makeMore(ctx, &record, makeResource)
			result.Records[i] = *data
		}
	}

	// 返回结果和可能的错误。
	return result, err
}

// makeMore 为团队成员视图数据添加更多关联信息
// 该函数主要用于为团队成员视图数据添加额外的关联信息，比如团队信息、员工信息和邀人人信息
// 参数:
//
//	ctx - 上下文，用于传递请求范围的信息
//	data - 团队成员视图数据，将被添加更多关联信息
//	makeResource - 是否需要添加额外资源的标志
//
// 返回值:
//
//	返回添加了更多关联信息的团队成员视图数据
func (s *sTeamMemberView) makeMore(ctx context.Context, data *co_model.TeamMemberViewRes, makeResource bool) *co_model.TeamMemberViewRes {
	// 如果data为nil或makeResource为false，则直接返回data，不做任何处理
	if data == nil || makeResource == false {
		return data
	}

	// 为data添加团队信息
	// 当TeamId大于0时，说明需要添加团队信息
	if data.TeamId > 0 {
		base_funs.AttrMake[*co_model.TeamMemberViewRes](ctx,
			co_dao.CompanyTeamMemberView.Columns().TeamId,
			func() (res *co_entity.CompanyTeamView) {
				// 获取并设置团队成员的团队信息
				_ = g.Try(ctx, func(ctx context.Context) {
					team, _ := co_service.TeamView().GetTeamById(ctx, data.TeamId, false)
					if team != nil {
						data.Team = &team.CompanyTeamView
					}
				})
				return data.Team
			},
		)
	}

	// 为data添加员工信息
	// 当EmployeeId大于0时，说明需要添加员工信息
	if data.EmployeeId > 0 {
		base_funs.AttrMake[*co_model.TeamMemberViewRes](ctx,
			co_dao.CompanyTeamMemberView.Columns().EmployeeId,
			func() (res *co_entity.CompanyEmployeeView) {
				// 获取并设置团队成员的员工信息
				_ = g.Try(ctx, func(ctx context.Context) {
					employee, _ := co_service.EmployeeView().GetEmployeeById(ctx, data.EmployeeId, false)
					if employee != nil {
						data.Employee = &employee.CompanyEmployeeView
					}
				})
				return data.Employee
			},
		)
	}

	// 为data添加邀人人信息
	// 当InviteUserId大于0时，说明需要添加邀人人信息
	if data.InviteUserId > 0 {
		base_funs.AttrMake[*co_model.TeamMemberViewRes](ctx,
			co_dao.CompanyTeamMemberView.Columns().InviteUserId,
			func() (res *co_entity.CompanyEmployeeView) {
				// 获取并设置团队成员的邀人人信息
				_ = g.Try(ctx, func(ctx context.Context) {
					inviteUser, _ := co_service.EmployeeView().GetEmployeeById(ctx, data.InviteUserId, false)
					if inviteUser != nil {
						data.InviteUser = &inviteUser.CompanyEmployeeView
					}
				})
				return data.InviteUser
			},
		)
	}

	// 为data添加所属单位信息
	// 当UnionMainId大于0时，说明需要添加所属单位信息
	if data.UnionMainId > 0 {
		base_funs.AttrMake[*co_model.TeamMemberViewRes](ctx,
			co_dao.CompanyTeamMemberView.Columns().UnionMainId,
			func() (res *co_entity.CompanyView) {
				// 获取并设置团队成员的所属单位信息
				_ = g.Try(ctx, func(ctx context.Context) {
					unionMain, _ := co_service.CompanyView().GetCompanyById(ctx, data.UnionMainId, false)
					if unionMain != nil {
						data.UnionMain = &unionMain.CompanyView.CompanyView
					}
				})
				return data.UnionMain
			},
		)
	}

	// 返回添加了更多关联信息的data
	return data
}
