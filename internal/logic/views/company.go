package views

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
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

type sCompanyView struct {
}

func init() {
	co_service.RegisterCompanyView(NewCompanyView())
}

func NewCompanyView() co_service.ICompanyView {
	return &sCompanyView{}
}

// GetCompanyById 根据公司ID获取公司信息。
// 该方法首先尝试从数据库中获取公司信息。如果成功获取信息且makeResource参数为真，
// 则进一步处理公司信息以生成更多的资源数据。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的信息。
//	id - 公司的唯一标识符。
//	makeResource - 指示是否需要进一步处理公司信息以生成额外的资源数据。
//
// 返回值:
//
//	*model.CompanyViewRes - 公司信息的视图资源对象，如果找不到则返回nil。
//	error - 错误信息，如果执行过程中遇到任何问题则返回。
func (s *sCompanyView) GetCompanyById(ctx context.Context, id int64, makeResource bool) (*co_model.CompanyViewRes, error) {
	// 根据ID获取公司信息
	data, err := daoctl.GetByIdWithError[co_model.CompanyViewRes](co_dao.CompanyView.Ctx(ctx), id)

	// 如果获取公司信息成功且需要生成额外的资源数据，则进一步处理公司信息
	if err == nil && makeResource {
		// 如果需要，进一步处理公司信息
		data = s.makeMore(ctx, data, makeResource)
	}

	// 返回公司信息或错误
	return data, err
}

// QueryCompanyList 查询公司列表
// 该方法根据提供的搜索参数查询公司信息，并可选地处理额外的资源信息
// 主要用于获取分页的公司列表，每个公司信息的详细程度取决于makeResource参数
func (s *sCompanyView) QueryCompanyList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.CompanyViewListRes, error) {
	// 根据查询参数获取公司列表
	data, err := daoctl.Query[co_model.CompanyViewRes](co_dao.CompanyView.Ctx(ctx), params, false)

	if err != nil {
		// 如果查询出错，则返回错误
		return nil, err
	}

	// 初始化结果对象，包含分页信息
	result := &co_model.CompanyViewListRes{
		PaginationRes: data.PaginationRes,
	}

	// 如果查询无错误、结果非空且需要处理额外资源，则对每个公司信息进行进一步处理
	if len(data.Records) > 0 && makeResource {
		// 对每个公司信息，调用makeMore方法处理，并将结果更新到结果集中
		for i, record := range data.Records {
			data := s.makeMore(ctx, &record, makeResource)
			result.Records[i] = *data
		}
	}

	// 返回查询结果和可能的错误
	return result, err
}

// makeMore 为公司视图数据添加额外的相关信息。
// 该函数根据makeResource标志决定是否为data对象添加更多相关信息，包括管理员登录信息、员工信息和资质信息。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的信息。
//	data - 公司视图数据对象，将被添加更多信息。
//	makeResource - 布尔值，指示是否执行添加更多信息的操作。
//
// 返回值:
//
//	返回可能被更新的公司视图数据对象。
func (s *sCompanyView) makeMore(ctx context.Context, data *co_model.CompanyViewRes, makeResource bool) *co_model.CompanyViewRes {
	// 如果不需要进一步处理或数据为空，直接返回
	if makeResource != true || data == nil {
		return data
	}

	// 如果数据中包含有效的用户ID，注入管理员登录信息和员工信息。
	if data.UserId > 0 {
		// 按需注入管理员登录信息
		base_funs.AttrMake[co_model.CompanyViewRes](ctx, co_dao.CompanyView.Columns().UserId,
			func() (res *sys_model.SysUser) {
				_ = g.Try(ctx, func(ctx context.Context) {
					data.User, _ = sys_service.SysUser().GetSysUserById(ctx, data.UserId)
				})
				return data.User
			},
		)

		// 按需注入员工信息
		base_funs.AttrMake[co_model.CompanyViewRes](ctx, co_dao.CompanyView.Columns().UserId,
			func() (res *co_entity.CompanyEmployeeView) {
				_ = g.Try(ctx, func(ctx context.Context) {
					employee, _ := co_service.EmployeeView().GetEmployeeById(ctx, data.UserId, false)
					data.Employee = &employee.CompanyEmployeeView
				})
				return data.Employee
			},
		)
	}

	// 如果数据中包含有效的资质ID，注入资质信息。
	if data.LicenseId > 0 {
		base_funs.AttrMake[co_model.CompanyViewRes](ctx, co_dao.CompanyView.Columns().LicenseId,
			func() (res *co_entity.License) {
				// 注入资质信息
				_ = g.Try(ctx, func(ctx context.Context) {
					data.License, _ = co_service.License().GetLicenseById(ctx, data.LicenseId)
				})
				return data.License
			},
		)
	}

	// 返回可能被更新的数据对象
	return data
}
