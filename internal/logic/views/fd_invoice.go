package views

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
)

type sFdInvoiceView struct {
}

// init函数用于初始化财务发票服务模块
// 这个函数在程序启动时自动调用
func init() {
	// 注册财务发票服务到服务发现系统中
	co_service.RegisterFdInvoiceView(NewFdInvoiceView())
}

// NewFdInvoiceView 创建并返回一个新的财务发票服务实例。
// 返回值是IFdInvoice接口类型，实际返回的是实现了该接口的sFdInvoiceView类型实例。
func NewFdInvoiceView() co_service.IFdInvoiceView {
	return &sFdInvoiceView{}
}

// GetFdInvoiceById 根据财务发票ID获取财务发票详细信息。
// 该方法首先尝试从数据库中获取财务发票信息。如果找到财务发票信息且makeResource参数为true，
// 则进一步处理数据以生成更多的资源信息。
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
//	id int64: 财务发票的唯一标识符。
//	makeResource bool: 指示是否需要生成额外的资源信息。
//
// 返回值:
//
//	*co_model.FdInvoiceViewRes: 财务发票详细信息的视图模型，如果找不到则返回nil。
//	error: 错误对象，如果操作成功则返回nil。
func (s *sFdInvoiceView) GetFdInvoiceById(ctx context.Context, id int64, makeResource bool) (*co_model.FdInvoiceViewRes, error) {
	// 从数据库中获取财务发票详细信息。
	data, err := daoctl.GetByIdWithError[co_model.FdInvoiceViewRes](co_dao.FdInvoiceView.Ctx(ctx), id)

	// 如果没有错误且makeResource为true，则进一步处理数据以生成更多的资源信息。
	if err == nil && makeResource {
		data = s.makeMore(ctx, data, makeResource)
	}

	// 返回获取的财务发票信息或错误。
	return data, err
}

// QueryFdInvoiceList 查询财务发票列表信息。
// 该方法根据提供的搜索参数查询财务发票信息，并可选地构建额外的资源信息。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息。
//	params - 搜索参数，用于指定查询的条件。
//	makeResource - 指示是否构建额外的资源信息。
//
// 返回值:
//
//	*co_model.FdInvoiceViewListRes - 包含财务发票列表的响应对象。
//	error - 错误信息，如果执行过程中发生错误。
func (s *sFdInvoiceView) QueryFdInvoiceList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.FdInvoiceViewListRes, error) {
	// 调用DAO层的方法来查询财务发票信息。
	data, err := daoctl.Query[co_model.FdInvoiceViewRes](co_dao.FdInvoiceView.Ctx(ctx), params, false)

	// 初始化结果对象，包含分页信息。
	result := &co_model.FdInvoiceViewListRes{
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

// makeMore 为财务发票视图数据添加更多关联信息
// 该函数主要用于为财务发票视图数据添加额外的关联信息，比如用户信息和所属单位信息
// 参数:
//
//	ctx - 上下文，用于传递请求范围的信息
//	data - 财务发票视图数据，将被添加更多关联信息
//	makeResource - 是否需要添加额外资源的标志
//
// 返回值:
//
//	返回添加了更多关联信息的财务发票视图数据
func (s *sFdInvoiceView) makeMore(ctx context.Context, data *co_model.FdInvoiceViewRes, makeResource bool) *co_model.FdInvoiceViewRes {
	// 如果data为nil或makeResource为false，则直接返回data，不做任何处理
	if data == nil || makeResource == false {
		return data
	}

	// 为data添加用户信息
	// 当data的Id大于0时，说明需要添加用户信息
	if data.UserId > 0 {
		base_funs.AttrMake[*co_model.FdInvoiceViewRes](ctx,
			co_dao.FdInvoiceView.Columns().Id,
			func() (res *sys_model.SysUser) {
				// 获取并设置财务发票的用户信息
				data.User, _ = sys_service.SysUser().GetSysUserById(ctx, data.Id)
				return data.User
			},
		)
	}

	// 为data添加所属单位信息
	// 当data的UnionMainId大于0时，说明需要添加所属单位信息
	if data.UnionMainId > 0 {
		base_funs.AttrMake[*co_model.FdInvoiceViewRes](ctx,
			co_dao.FdInvoiceView.Columns().Id,
			func() (res *co_entity.CompanyView) {
				// 获取并设置财务发票的所属单位信息
				data.UnionMain = daoctl.GetById[co_entity.CompanyView](co_dao.CompanyView.Ctx(ctx), data.UnionMainId)
				return data.UnionMain
			},
		)
	}

	// 返回添加了更多关联信息的data
	return data
}
