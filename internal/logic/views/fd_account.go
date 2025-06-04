package views

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
)

type sFdAccountView struct {
}

// init函数用于初始化财务账户服务模块
// 这个函数在程序启动时自动调用
func init() {
	// 注册财务账户服务到服务发现系统中
	co_service.RegisterFdAccountView(NewFdAccountView())
}

// NewFdAccountView 创建并返回一个新的财务账户服务实例。
// 返回值是IFdAccount接口类型，实际返回的是实现了该接口的sFdAccountView类型实例。
func NewFdAccountView() co_service.IFdAccountView {
	return &sFdAccountView{}
}

// GetFdAccountById 根据财务账户ID获取财务账户详细信息。
// 该方法首先尝试从数据库中获取财务账户信息。如果找到财务账户信息且makeResource参数为true，
// 则进一步处理数据以生成更多的资源信息。
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
//	id int64: 财务账户的唯一标识符。
//	makeResource bool: 指示是否需要生成额外的资源信息。
//
// 返回值:
//
//	*co_model.FdAccountViewRes: 财务账户详细信息的视图模型，如果找不到则返回nil。
//	error: 错误对象，如果操作成功则返回nil。
func (s *sFdAccountView) GetFdAccountById(ctx context.Context, id int64, makeResource bool) (*co_model.FdAccountViewRes, error) {
	// 从数据库中获取财务账户详细信息。
	data, err := daoctl.GetByIdWithError[co_model.FdAccountViewRes](co_dao.FdAccountView.Ctx(ctx), id)

	// 如果没有错误且makeResource为true，则进一步处理数据以生成更多的资源信息。
	if err == nil && makeResource {
		data = s.makeMore(ctx, data, makeResource)
	}

	// 返回获取的财务账户信息或错误。
	return data, err
}

// QueryFdAccountList 查询财务账户列表信息。
// 该方法根据提供的搜索参数查询财务账户信息，并可选地构建额外的资源信息。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息。
//	params - 搜索参数，用于指定查询的条件。
//	makeResource - 指示是否构建额外的资源信息。
//
// 返回值:
//
//	*co_model.FdAccountViewListRes - 包含财务账户列表的响应对象。
//	error - 错误信息，如果执行过程中发生错误。
func (s *sFdAccountView) QueryFdAccountList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.FdAccountViewListRes, error) {
	// 调用DAO层的方法来查询财务账户信息。
	data, err := daoctl.Query[co_model.FdAccountViewRes](co_dao.FdAccountView.Ctx(ctx), params, false)

	// 初始化结果对象，包含分页信息。
	result := &co_model.FdAccountViewListRes{
		PaginationRes: data.PaginationRes,
		Records:       data.Records,
	}

	// 如果没有错误且查询到了记录且要求构建资源信息，则为每条记录构建更多的资源信息。
	if err == nil && len(result.Records) > 0 && makeResource {
		for i, record := range result.Records {
			// 为每个记录构建更多的资源信息，并将其添加到结果中。
			item := s.makeMore(ctx, &record, makeResource)
			result.Records[i] = *item
		}
	}

	// 返回结果和可能的错误。
	return result, err
}

// GetUserDefaultFdAccountByUserId 根据用户ID和用户类型查询财务账户信息，并返回第一个匹配的财务账户
// 该方法根据提供的用户ID和用户类型查询财务账户信息，并返回第一个匹配的财务账户
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息
//	userId - 用户ID，用于指定查询的用户
//	userType - 用户类型，用于指定查询的用户类型
//
// 返回值:
//
//	*co_model.FdAccountViewRes - 财务账户视图数据，如果查询失败则返回nil
//	error - 错误信息，如果查询失败则返回错误
func (s *sFdAccountView) GetUserDefaultFdAccountByUserId(ctx context.Context, userId int64, userType sys_enum.UserType) (*co_model.FdAccountViewRes, error) {
	dataArr, err := co_service.FdAccountView().QueryFdAccountList(ctx, &base_model.SearchParams{
		Filter: []base_model.FilterInfo{
			{
				Field: co_dao.FdAccountView.Columns().UnionUserId,
				Where: "=",
				Value: userId,
			},
			{
				Field: co_dao.FdAccountView.Columns().CompanyType,
				Where: "=",
				Value: userType.Code(),
			},
		},
		OrderBy: []base_model.OrderBy{
			{
				Field: co_dao.FdAccountView.Columns().CreatedAt,
				Sort:  "ASC",
			},
		},
		Pagination: base_model.Pagination{
			PageNum:  1,
			PageSize: 1,
		},
	}, false)

	if err != nil || len(dataArr.Records) <= 0 {
		return nil, err
	}

	return &dataArr.Records[0], nil
}

// makeMore 为财务账户视图数据添加更多关联信息
// 该函数主要用于为财务账户视图数据添加额外的关联信息，比如用户信息和所属单位信息
// 参数:
//
//	ctx - 上下文，用于传递请求范围的信息
//	data - 财务账户视图数据，将被添加更多关联信息
//	makeResource - 是否需要添加额外资源的标志
//
// 返回值:
//
//	返回添加了更多关联信息的财务账户视图数据
func (s *sFdAccountView) makeMore(ctx context.Context, data *co_model.FdAccountViewRes, makeResource bool) *co_model.FdAccountViewRes {
	// 如果data为nil或makeResource为false，则直接返回data，不做任何处理
	if data == nil || makeResource == false {
		return data
	}

	// 为data添加财务账户的详细信息
	// 添加财务账户的详细信息
	base_funs.AttrMake[*co_model.FdAccountViewRes](ctx,
		co_dao.FdAccountView.Columns().Id,
		func() (res *co_model.FdAccountDetailView) {
			// 获取并设置财务账户的详细信息
			_ = g.Try(ctx, func(ctx context.Context) {
				FdAccountDetail, _ := daoctl.GetByIdWithError[co_model.FdAccountDetailView](co_dao.FdAccountDetailView.Ctx(ctx), data.FdAccountView.Id)
				if FdAccountDetail != nil {
					data.FdAccountDetailView = *FdAccountDetail
				}
			})
			return &data.FdAccountDetailView
		},
	)

	// 为data添加用户信息
	// 当data的Id大于0时，说明需要添加用户信息
	if data.UnionUserId > 0 {
		base_funs.AttrMake[*co_model.FdAccountViewRes](ctx,
			co_dao.FdAccountView.Columns().UnionUserId,
			func() (res *sys_model.SysUser) {
				// 获取并设置财务账户的用户信息
				_ = g.Try(ctx, func(ctx context.Context) {
					User, err := sys_service.SysUser().GetSysUserById(ctx, data.UnionUserId)
					if User != nil && err == nil {
						data.User = User
					}
				})

				return data.User
			},
		)

		base_funs.AttrMake[*co_model.FdAccountViewRes](ctx,
			co_dao.FdAccountView.Columns().UnionUserId,
			func() (res *co_entity.CompanyEmployeeView) {
				// 获取并设置财务账户的用户信息
				_ = g.Try(ctx, func(ctx context.Context) {
					employee, err := daoctl.GetByIdWithError[co_entity.CompanyEmployeeView](co_dao.CompanyEmployeeView.Ctx(ctx), data.UnionUserId)
					if employee != nil && err == nil {
						data.Employee = employee
					}
				})

				return data.Employee
			},
		)
	}

	// 为data添加所属单位信息
	// 当data的UnionMainId大于0时，说明需要添加所属单位信息
	if data.FdAccountView.UnionMainId > 0 {
		base_funs.AttrMake[*co_model.FdAccountViewRes](ctx,
			co_dao.FdAccountView.Columns().UnionMainId,
			func() (res *co_entity.CompanyView) {
				// 获取并设置财务账户的所属单位信息
				data.UnionMain = daoctl.GetById[co_entity.CompanyView](co_dao.CompanyView.Ctx(ctx), data.FdAccountView.UnionMainId)
				return data.UnionMain
			},
		)
	}

	// 返回添加了更多关联信息的data
	return data
}
