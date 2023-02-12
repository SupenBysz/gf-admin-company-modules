package company

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/masker"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
)

type sCompany struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewCompany(modules co_interface.IModules) co_interface.ICompany {
	return &sCompany{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetCompanyById 根据ID获取获取公司信息
func (s *sCompany) GetCompanyById(ctx context.Context, id int64) (*co_model.CompanyRes, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	data, err := daoctl.GetByIdWithError[co_model.CompanyRes](
		s.dao.Company.Ctx(ctx),
		id,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Get_Failed}"), s.dao.Company.Table())
		}
	}
	if err == sql.ErrNoRows || data != nil && data.Id != sessionUser.UnionMainId && data.ParentId != sessionUser.UnionMainId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.masker(s.makeMore(ctx, data)), nil
}

// GetCompanyByName 根据Name获取获取公司信息
func (s *sCompany) GetCompanyByName(ctx context.Context, name string) (*co_model.CompanyRes, error) {
	data, err := daoctl.ScanWithError[co_model.CompanyRes](
		s.dao.Company.Ctx(ctx).
			Where(co_do.Company{Name: name}),
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.masker(s.makeMore(ctx, data)), nil
}

// HasCompanyByName 判断名称是否存在
func (s *sCompany) HasCompanyByName(ctx context.Context, name string, excludeIds ...int64) bool {
	model := s.dao.Company.Ctx(ctx)

	if len(excludeIds) > 0 {
		var ids []int64
		for _, id := range excludeIds {
			if id > 0 {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			model = model.WhereNotIn(s.dao.Company.Columns().Id, ids)
		}
	}

	count, _ := model.Where(co_do.Company{Name: name}).Count()
	return count > 0
}

// QueryCompanyList 查询公司列表
func (s *sCompany) QueryCompanyList(ctx context.Context, filter *sys_model.SearchParams) (*co_model.CompanyListRes, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
	data, err := daoctl.Query[*co_model.CompanyRes](
		s.dao.Company.Ctx(ctx).
			Where(co_do.Company{ParentId: sessionUser.UnionMainId}),
		filter,
		false,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	if data.Total > 0 {
		items := make([]*co_model.CompanyRes, 0)
		// 脱敏处理
		for _, item := range data.Records {
			items = append(items, s.masker(s.makeMore(ctx, item)))
		}
		data.Records = items
	}

	return (*co_model.CompanyListRes)(data), nil
}

// CreateCompany 创建公司信息
func (s *sCompany) CreateCompany(ctx context.Context, info *co_model.Company) (*co_model.CompanyRes, error) {
	info.Id = 0
	return s.saveCompany(ctx, info)
}

// UpdateCompany 更新公司信息
func (s *sCompany) UpdateCompany(ctx context.Context, info *co_model.Company) (*co_model.CompanyRes, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}
	return s.saveCompany(ctx, info)
}

// SaveCompany 保存公司信息
func (s *sCompany) saveCompany(ctx context.Context, info *co_model.Company) (*co_model.CompanyRes, error) {
	// 名称重名检测
	if s.HasCompanyByName(ctx, info.Name, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#CompanyName} {#error_NameAlreadyExists}"), s.dao.Company.Table())
	}

	// 获取登录用户
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 构建公司ID
	UnionMainId := idgen.NextId()

	data := co_do.Company{
		Id:            info.Id,
		Name:          info.Name,
		ContactName:   info.ContactName,
		ContactMobile: info.ContactMobile,
		Remark:        info.Remark,
	}

	// 启用事务
	err := s.dao.Company.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var employee *co_model.EmployeeRes
		if info.Id == 0 {
			// 是否创建默认员工和角色
			if s.modules.GetConfig().IsCreateDefaultEmployeeAndRole {
				// 构建员工信息
				employee, err = s.modules.Employee().CreateEmployee(ctx, &co_model.Employee{
					No:          "001",
					Name:        info.ContactName,
					Mobile:      info.ContactMobile,
					UnionMainId: UnionMainId,
					State:       co_enum.Employee.State.Normal.Code(),
					HiredAt:     gtime.Now(),
				})
				if err != nil {
					return err
				}

				// 构建角色信息
				roleData := sys_model.SysRole{
					Name:        "管理员",
					UnionMainId: UnionMainId,
					IsSystem:    true,
				}
				roleInfo, err := sys_service.SysRole().Create(ctx, roleData)
				if err != nil {
					return err
				}
				_, err = sys_service.SysUser().SetUserRoleIds(ctx, []int64{roleInfo.Id}, employee.Id)
				if err != nil {
					return err
				}
			}

			if employee != nil {
				// 如果需要创建默认的用户和角色的时候才会有employee，所以进行非空判断，不然会有空指针错误
				data.UserId = employee.Id
			} else {
				data.UserId = 0
			}

			data.Id = UnionMainId
			data.ParentId = sessionUser.UnionMainId
			data.CreatedBy = sessionUser.Id
			data.CreatedAt = gtime.Now()

			affected, err := daoctl.InsertWithError(
				s.dao.Company.Ctx(ctx),
				data,
			)
			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
			}
		} else {
			if gstr.Contains(info.ContactMobile, "***") || info.ContactMobile == "" {
				data.ContactMobile = nil
			}

			data.UpdatedBy = sessionUser.Id
			data.UpdatedAt = gtime.Now()
			_, err = daoctl.UpdateWithError(
				s.dao.Company.Ctx(ctx).
					Where(co_do.Company{Id: info.Id}),
				data,
			)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
			}
		}
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetCompanyById(ctx, data.Id.(int64))
}

// GetCompanyDetail 获取公司详情，包含完整商务联系人电话
func (s *sCompany) GetCompanyDetail(ctx context.Context, id int64) (*co_model.CompanyRes, error) {
	data, err := daoctl.GetByIdWithError[co_model.CompanyRes](
		s.dao.Company.Ctx(ctx),
		id,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.makeMore(ctx, data), nil
}

func (s *sCompany) FilterUnionMainId(ctx context.Context, search *sys_model.SearchParams) *sys_model.SearchParams {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	filter := make([]sys_model.FilterInfo, 0)
	// 遍历所有过滤条件：
	for _, field := range search.Filter {
		// 过滤所有自定义主体ID条件
		if gstr.ToLower(field.Field) == gstr.ToLower("unionMainId") {
			unionMainId := gconv.Int64(field.Value)
			if unionMainId == sessionUser.UnionMainId || unionMainId <= 0 {
				filter = append(filter, field)
				continue
			}
			company, err := s.modules.Company().GetCompanyById(ctx, unionMainId)
			if err != nil || (company != nil && company.ParentId != unionMainId) {
				field.Value = sessionUser.UnionMainId
				filter = append(filter, field)
			}
		} else {
			filter = append(filter, field)
		}
	}
	search.Filter = filter

	return search
}

// makeMore 按需加载附加数据
func (s *sCompany) makeMore(ctx context.Context, data *co_model.CompanyRes) *co_model.CompanyRes {
	if data == nil {
		return nil
	}

	if data.UserId > 0 {
		// 附加数据 employee
		funs.AttrMake[co_model.CompanyRes](ctx, s.dao.Company.Columns().UserId,
			func() *co_model.EmployeeRes {
				data.AdminUser, _ = s.modules.Employee().GetEmployeeById(ctx, data.UserId)
				if data.AdminUser == nil {
					return nil
				}

				user, _ := sys_service.SysUser().GetSysUserById(ctx, data.UserId)
				if user != nil {
					gconv.Struct(user.SysUser, &data.AdminUser.User)
					gconv.Struct(user.Detail, &data.AdminUser.Detail)
				}

				return data.AdminUser
			},
		)
	}

	return data
}

// Masker 信息脱敏
func (s *sCompany) masker(company *co_model.CompanyRes) *co_model.CompanyRes {
	if company == nil {
		return nil
	}

	company.ContactMobile = masker.MaskString(company.ContactMobile, masker.MaskPhone)

	return company
}
