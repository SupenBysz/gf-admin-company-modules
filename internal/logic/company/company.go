package company

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/masker"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type sCompany struct {
	modules co_interface.IModules
}

func NewCompany(modules co_interface.IModules) co_interface.ICompany {
	return &sCompany{
		modules: modules,
	}
}

// InjectHook 注入Audit的Hook
func (s *sCompany) InjectHook() {
	sys_service.Jwt().InstallHook(sys_enum.User.Type.Operator, s.JwtHookFunc)
}

// JwtHookFunc Jwt钩子函数
func (s *sCompany) JwtHookFunc(ctx context.Context, claims *sys_model.JwtCustomClaims) (*sys_model.JwtCustomClaims, error) {
	// 获取到当前user的主体id
	employee, err := s.modules.Employee().GetEmployeeById(ctx, claims.Id)
	if employee == nil {
		return claims, err
	}

	company, err := s.modules.Company().GetCompanyById(ctx, employee.UnionMainId)
	if company == nil || err != nil {
		return claims, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("主体id获取失败"), "主体id获取失败", co_dao.Company(s.modules).Table())
	}

	claims.IsAdmin = claims.Type == -1 || claims.Id == company.UserId
	claims.UnionMainId = company.Id

	return claims, nil
}

// GetCompanyById 根据ID获取获取公司信息
func (s *sCompany) GetCompanyById(ctx context.Context, id int64) (*co_entity.Company, error) {
	data, err := daoctl.GetByIdWithError[co_entity.Company](co_dao.Company(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler), id)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.GetConfig().CompanyName+s.modules.T(ctx, "error_Data_NotFound"), co_dao.Company(s.modules).Table())
	}

	return s.masker(data), nil
}

// HasCompanyByName 判断名称是否存在
func (s *sCompany) HasCompanyByName(ctx context.Context, name string, excludeId ...int64) bool {
	model := co_dao.Company(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler)

	if len(excludeId) > 0 {
		model = model.WhereNotIn(co_dao.Company(s.modules).Columns().Id, excludeId)
	}

	count, _ := model.Where(co_do.Company{Name: name}).Count()
	return count > 0
}

// QueryCompanyList 查询公司列表
func (s *sCompany) QueryCompanyList(ctx context.Context, filter *sys_model.SearchParams) (*co_model.CompanyListRes, error) {
	data, err := daoctl.Query[*co_entity.Company](co_dao.Company(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler), filter, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.GetConfig().CompanyName+s.modules.T(ctx, "error_Data_NotFound"), co_dao.Company(s.modules).Table())
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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.GetConfig().CompanyName+s.modules.T(ctx, "error_Data_NotFound"), co_dao.Company(s.modules).Table())
	}
	return s.saveCompany(ctx, info)
}

// SaveCompany 保存公司信息
func (s *sCompany) saveCompany(ctx context.Context, info *co_model.Company) (*co_entity.Company, error) {
	// 名称重名检测
	if s.HasCompanyByName(ctx, info.Name, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.GetConfig().CompanyName+"名称已被占用，请修改后再试", co_dao.Company(s.modules).Table())
	}

	// 构建公司ID
	UnionMainId := idgen.NextId()

	data := co_do.Company{
		Id:            info.Id,
		Name:          info.Name,
		ContactName:   info.ContactName,
		ContactMobile: info.ContactMobile,
		Remark:        info.Remark,
	}

	// 获取登录用户
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 启用事务
	err := co_dao.Company(s.modules).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var employee *co_entity.CompanyEmployee
		// 是否创建默认员工和角色
		if s.modules.GetConfig().IsCreateDefaultEmployeeAndRole && info.Id == 0 {
			// 构建员工信息
			employee, err = s.modules.Employee().CreateEmployee(ctx, &co_model.Employee{
				No:          "001",
				Name:        info.ContactName,
				Mobile:      info.ContactMobile,
				UnionMainId: UnionMainId,
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
			data.UserId = 0
			data.CreatedBy = sessionUser.Id
			data.CreatedAt = gtime.Now()
			if employee != nil {
				data.UserId = employee.Id
			}

			_, err = co_dao.Company(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(data)

		} else {
			data.UpdatedBy = sessionUser.Id
			data.UpdatedAt = gtime.Now()
			_, err = co_dao.Company(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.Company{Id: info.Id}).Update(data)
		}
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.GetConfig().CompanyName+s.modules.T(ctx, "error_Data_Save_Failed"), co_dao.Company(s.modules).Table())
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetCompanyById(ctx, data.Id.(int64))
}

// Masker 信息脱敏
func (s *sCompany) masker(company *co_entity.Company) *co_entity.Company {
	if company == nil {
		return nil
	}

	company.ContactMobile = masker.MaskString(company.ContactMobile, masker.MaskPhone)
	return company
}
