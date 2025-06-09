package views

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
)

type sEmployeeView struct {
}

// init函数用于初始化员工服务模块
// 这个函数在程序启动时自动调用
func init() {
	// 注册员工服务到服务发现系统中
	co_service.RegisterEmployeeView(NewEmployeeView())
}

// NewEmployeeView 创建并返回一个新的员工服务实例。
// 返回值是IEmployee接口类型，实际返回的是实现了该接口的sEmployeeView类型实例。
func NewEmployeeView() co_service.IEmployeeView {
	return &sEmployeeView{}
}

// GetEmployeeById 根据员工ID获取员工详细信息。
// 该方法首先尝试从数据库中获取员工信息。如果找到员工信息且makeResource参数为true，
// 则进一步处理数据以生成更多的资源信息。
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
//	id int64: 员工的唯一标识符。
//	makeResource bool: 指示是否需要生成额外的资源信息。
//
// 返回值:
//
//	*co_model.EmployeeViewRes: 员工详细信息的视图模型，如果找不到则返回nil。
//	error: 错误对象，如果操作成功则返回nil。
func (s *sEmployeeView) GetEmployeeById(ctx context.Context, id int64, makeResource bool) (*co_model.EmployeeViewRes, error) {
	// 从数据库中获取员工详细信息。
	data, err := daoctl.GetByIdWithError[co_model.EmployeeViewRes](co_dao.CompanyEmployeeView.Ctx(ctx), id)

	// 如果没有错误且makeResource为true，则进一步处理数据以生成更多的资源信息。
	if err == nil && makeResource {
		data = s.makeMore(ctx, data, makeResource)
	}

	// 返回获取的员工信息或错误。
	return data, err
}

// QueryEmployeeList 查询员工列表信息。
// 该方法根据提供的搜索参数查询员工信息，并可选地构建额外的资源信息。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息。
//	params - 搜索参数，用于指定查询的条件。
//	makeResource - 指示是否构建额外的资源信息。
//
// 返回值:
//
//	*co_model.EmployeeViewListRes - 包含员工列表的响应对象。
//	error - 错误信息，如果执行过程中发生错误。
func (s *sEmployeeView) QueryEmployeeList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.EmployeeViewListRes, error) {
	// 调用DAO层的方法来查询员工信息。
	data, err := daoctl.Query[co_model.EmployeeViewRes](co_dao.CompanyEmployeeView.Ctx(ctx), params, false)

	// 初始化结果对象，包含分页信息。
	result := &co_model.EmployeeViewListRes{
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

func (s *sEmployeeView) QueryMyInviteEmployee(ctx context.Context, inviteUserId int64) (*[]co_model.EmployeeViewRes, error) {
	inviteUserList, _ := sys_service.SysInvite().QueryInvitePersonList(ctx, inviteUserId)

	result := make([]co_model.EmployeeViewRes, 0)

	if len(inviteUserList.Records) == 0 {
		return &result, nil
	}

	userIds := make([]int64, 0)

	for _, record := range inviteUserList.Records {
		userIds = append(userIds, record.ByUserId)
	}

	err := co_dao.CompanyEmployeeView.Ctx(ctx).WhereIn(co_dao.CompanyEmployee.Columns().Id, userIds).WhereOr(co_dao.CompanyEmployeeView.Columns().CreatedBy, inviteUserId).OrderDesc(co_dao.CompanyEmployeeView.Columns().CreatedAt).Scan(&result)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询员工列表失败", co_dao.CompanyEmployeeView.Table())
	}

	return &result, nil
}

// makeMore 按需加载附加数据

// makeMore 为员工视图数据添加更多关联信息
// 该函数主要用于为员工视图数据添加额外的关联信息，比如团队信息和用户信息
// 参数:
//
//	ctx - 上下文，用于传递请求范围的信息
//	data - 员工视图数据，将被添加更多关联信息
//	makeResource - 是否需要添加额外资源的标志
//
// 返回值:
//
//	返回添加了更多关联信息的员工视图数据
func (s *sEmployeeView) makeMore(ctx context.Context, data *co_model.EmployeeViewRes, makeResource bool) *co_model.EmployeeViewRes {
	// 如果data为nil或makeResource为false，则直接返回data，不做任何处理
	if data == nil || makeResource == false {
		return data
	}

	// 为data添加用户信息
	// 当data的Id大于0时，说明需要添加用户信息
	base_funs.AttrMake[*co_model.EmployeeViewRes](ctx,
		co_dao.CompanyEmployeeView.Columns().Id,
		func() (res *sys_model.SysUser) {
			// 获取并设置员工的用户信息
			data.User, _ = sys_service.SysUser().GetSysUserById(ctx, data.Id)
			return data.User
		},
	)

	// 为data添加团队信息
	// 当data的UnionMainId大于0时，说明需要添加团队信息
	if data.UnionMainId > 0 {
		// 根据外部订阅，自动匹配到对应类型后附加数据信息
		base_funs.AttrMake[*co_model.EmployeeViewRes](ctx,
			co_dao.CompanyEmployeeView.Columns().Id,
			func() *[]co_model.TeamViewRes {
				// 尝试获取并设置员工的团队信息
				_ = g.Try(ctx, func(ctx context.Context) {
					// 获取到该员工的所有团队成员信息记录ids
					ids, err := co_dao.CompanyTeamMemberView.Ctx(ctx).
						Where(co_do.CompanyTeamMember{EmployeeId: data.Id}).Fields([]string{co_dao.CompanyTeamMember.Columns().TeamId}).All()

					temIds := ids.Array()

					// 如果发生错误或没有找到任何记录，则返回
					if err != nil || len(ids) == 0 {
						data.TeamList = nil
						return
					}

					// 记录该员工所在所有团队,并将结果赋值给data.TeamList
					if len(temIds) > 0 {
						_ = co_dao.CompanyTeamView.Ctx(ctx).
							WhereIn(co_dao.CompanyTeamView.Columns().Id, temIds).Scan(&data.TeamList)
					}
				})

				// 尝试获取并设置员工的UnionMain信息
				_ = g.Try(ctx, func(ctx context.Context) {
					data.UnionMain = daoctl.GetById[co_entity.CompanyView](co_dao.CompanyView.Ctx(ctx), data.UnionMainId)
				})
				return data.TeamList
			},
		)
	}

	// 返回添加了更多关联信息的data
	return data
}
