package views

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
)

type sTeamView struct {
}

func init() {
	co_service.RegisterTeamView(NewTeamView())
}

func NewTeamView() co_service.ITeamView {
	return &sTeamView{}
}

// GetTeamById 根据团队ID获取团队详细信息。
// 该方法首先尝试从数据库中获取团队信息。如果找到团队信息且makeResource参数为true，
// 则进一步处理数据以生成更多的资源信息。
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
//	id int64: 团队的唯一标识符。
//	makeResource bool: 指示是否需要生成额外的资源信息。
//
// 返回值:
//
//	*co_model.TeamViewRes: 团队详细信息的视图模型，如果找不到则返回nil。
//	error: 错误对象，如果操作成功则返回nil。
func (s *sTeamView) GetTeamById(ctx context.Context, id int64, makeResource bool) (*co_model.TeamViewRes, error) {
	// 从数据库中获取团队详细信息。
	data, err := daoctl.GetByIdWithError[co_model.TeamViewRes](co_dao.CompanyTeamView.Ctx(ctx), id)

	// 如果没有错误且makeResource为true，则进一步处理数据以生成更多的资源信息。
	if err == nil && makeResource {
		data = s.makeMore(ctx, data, makeResource)
	}

	// 返回获取的团队信息或错误。
	return data, err
}

// QueryTeamList 查询团队列表信息。
// 该方法根据提供的搜索参数查询团队信息，并可选地构建额外的资源信息。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息。
//	params - 搜索参数，用于指定查询的条件。
//	makeResource - 指示是否构建额外的资源信息。
//
// 返回值:
//
//	*co_model.TeamViewListRes - 包含团队列表的响应对象。
//	error - 错误信息，如果执行过程中发生错误。
func (s *sTeamView) QueryTeamList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.TeamViewListRes, error) {
	// 调用DAO层的方法来查询团队信息。
	data, err := daoctl.Query[co_model.TeamViewRes](co_dao.CompanyTeamView.Ctx(ctx), params, false)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	// 初始化结果对象，包含分页信息。
	result := &co_model.TeamViewListRes{
		PaginationRes: base_funs.If(data != nil, data.PaginationRes, base_model.PaginationRes{}),
	}

	if errors.Is(err, sql.ErrNoRows) {
		return result, nil
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

// makeMore 为团队视图数据添加更多关联信息
// 该函数主要用于为团队视图数据添加额外的关联信息，比如用户信息
// 参数:
//
//	ctx - 上下文，用于传递请求范围的信息
//	data - 团队视图数据，将被添加更多关联信息
//	makeResource - 是否需要添加额外资源的标志
//
// 返回值:
//
//	返回添加了更多关联信息的团队视图数据
func (s *sTeamView) makeMore(ctx context.Context, data *co_model.TeamViewRes, makeResource bool) *co_model.TeamViewRes {
	// 如果data为nil或makeResource为false，则直接返回data，不做任何处理
	if data == nil || makeResource == false {
		return data
	}

	// 为data添加用户信息
	// 当OwnerEmployeeId大于0时，说明需要添加用户信息
	if data.OwnerEmployeeId > 0 {
		base_funs.AttrMake[*co_model.TeamViewRes](ctx,
			co_dao.CompanyTeamView.Columns().Id,
			func() (res *co_entity.CompanyEmployeeView) {
				// 获取并设置团队的成员信息
				employee, _ := co_service.EmployeeView().GetEmployeeById(ctx, data.OwnerEmployeeId, false)
				if employee != nil {
					data.Owner = &employee.CompanyEmployeeView
				}
				return data.Owner
			},
		)
	}

	if data.CaptainEmployeeId > 0 {
		base_funs.AttrMake[*co_model.TeamViewRes](ctx,
			co_dao.CompanyTeamView.Columns().Id,
			func() (res *co_entity.CompanyEmployeeView) {
				// 获取并设置团队的成员信息
				employee, _ := co_service.EmployeeView().GetEmployeeById(ctx, data.OwnerEmployeeId, false)
				if employee != nil {
					data.Captain = &employee.CompanyEmployeeView
				}
				return data.Captain
			},
		)
	}

	// 附加数据4：团队或小组父级
	if data.ParentId > 0 {
		base_funs.AttrMake[*co_model.TeamViewRes](ctx,
			co_dao.CompanyTeamView.Columns().ParentId,
			func() *co_model.TeamViewRes {
				if data.ParentId == 0 {
					return nil
				}

				team, _ := co_service.TeamView().GetTeamById(ctx, data.ParentId, true)
				if team != nil {
					data.Parent = team
				}

				return team
			},
		)
	}

	// 返回添加了更多关联信息的data
	return data
}
