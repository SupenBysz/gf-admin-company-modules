package company

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/masker"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type sCompany struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewCompany(modules co_interface.IModules, xDao *co_dao.XDao) co_interface.ICompany {
	result := &sCompany{
		modules: modules,
		dao:     xDao,
	}
	result.injectHook()
	return result
}

// InjectHook 注入Audit的Hook
func (s *sCompany) injectHook() {
	sys_service.Jwt().InstallHook(s.modules.GetConfig().UserType, s.jwtHookFunc)
}

// JwtHookFunc Jwt钩子函数
func (s *sCompany) jwtHookFunc(ctx context.Context, claims *sys_model.JwtCustomClaims) (*sys_model.JwtCustomClaims, error) {
	// 获取到当前user的主体id
	employee, err := s.modules.Employee().GetEmployeeById(ctx, claims.Id)
	if employee == nil {
		return claims, err
	}

	// 这里还没登录成功不能使用 GetCompanyById，因为里面包含获取当前登录用户的 session 存在矛盾
	company, err := daoctl.GetByIdWithError[co_entity.Company](
		s.dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler),
		employee.UnionMainId,
	)
	if company == nil || err != nil {
		return claims, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName}{#error_Data_Get_Failed}"), s.dao.Company.Table())
	}

	claims.IsAdmin = claims.Type == -1 || claims.Id == company.UserId
	claims.UnionMainId = company.Id
	claims.ParentId = company.ParentId

	return claims, nil
}

// GetCompanyById 根据ID获取获取公司信息
func (s *sCompany) GetCompanyById(ctx context.Context, id int64) (*co_entity.Company, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
	data, err := daoctl.GetByIdWithError[co_entity.Company](
		s.dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler).
			Where(co_do.Company{ParentId: sessionUser.ParentId}),
		id,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.masker(data), nil
}

// GetCompanyByName 根据Name获取获取公司信息
func (s *sCompany) GetCompanyByName(ctx context.Context, name string) (*co_entity.Company, error) {
	data, err := daoctl.ScanWithError[co_entity.Company](
		s.dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler).
			Where(co_do.Company{Name: name}),
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.masker(data), nil
}

// HasCompanyByName 判断名称是否存在
func (s *sCompany) HasCompanyByName(ctx context.Context, name string, excludeIds ...int64) bool {
	model := s.dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler)

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
	data, err := daoctl.Query[*co_entity.Company](
		s.dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler).
			Where(co_do.Company{ParentId: sessionUser.ParentId}),
		filter,
		false,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	if data.Total > 0 {
		items := make([]*co_entity.Company, 0)
		// 脱敏处理
		for _, item := range data.Records {
			items = append(items, s.masker(item))
		}
		data.Records = items
	}

	return (*co_model.CompanyListRes)(data), nil
}

// CreateCompany 创建公司信息
func (s *sCompany) CreateCompany(ctx context.Context, info *co_model.Company) (*co_entity.Company, error) {
	info.Id = 0
	return s.saveCompany(ctx, info)
}

// UpdateCompany 更新公司信息
func (s *sCompany) UpdateCompany(ctx context.Context, info *co_model.Company) (*co_entity.Company, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}
	return s.saveCompany(ctx, info)
}

// SaveCompany 保存公司信息
func (s *sCompany) saveCompany(ctx context.Context, info *co_model.Company) (*co_entity.Company, error) {
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
		var employee *co_entity.CompanyEmployee
		// 是否创建默认员工和角色
		if s.modules.GetConfig().IsCreateDefaultEmployeeAndRole && info.Id == 0 {
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

		if info.Id == 0 {
			data.Id = UnionMainId
			data.UserId = employee.Id
			data.ParentId = sessionUser.UnionMainId
			data.CreatedBy = sessionUser.Id
			data.CreatedAt = gtime.Now()

			affected, err := daoctl.InsertWithError(
				s.dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler),
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
				s.dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler).
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
func (s *sCompany) GetCompanyDetail(ctx context.Context, id int64) (*co_entity.Company, error) {
	// 获取登录用户
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	data, err := daoctl.GetByIdWithError[co_entity.Company](
		s.dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler).
			Where(co_do.Company{ParentId: sessionUser.ParentId}), id,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}
	return data, nil
}

// Masker 信息脱敏
func (s *sCompany) masker(company *co_entity.Company) *co_entity.Company {
	if company == nil {
		return nil
	}

	company.ContactMobile = masker.MaskString(company.ContactMobile, masker.MaskPhone)
	return company
}
