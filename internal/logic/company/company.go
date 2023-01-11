package company

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
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
		return claims, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("主体id获取失败"), "主体id获取失败", co_dao.Company.Table())
	}

	claims.IsAdmin = claims.Type == -1 || claims.Id == company.UserId
	claims.UnionMainId = company.Id

	return claims, nil
}

// GetCompanyById 根据ID获取获取公司信息
func (s *sCompany) GetCompanyById(ctx context.Context, id int64) (*co_entity.Company, error) {
	data, err := daoctl.GetByIdWithError[co_entity.Company](co_dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler), id)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.GetConfig().CompanyName+"信息不存在", co_dao.Company.Table())
	}

	return s.masker(data), nil
}

// HasCompanyByName 判断名称是否存在
func (s *sCompany) HasCompanyByName(ctx context.Context, name string, excludeId ...int64) bool {
	model := co_dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler)

	if len(excludeId) > 0 {
		model = model.WhereNotIn(co_dao.Company.Columns().Id, excludeId)
	}

	count, _ := model.Where(co_do.Company{Name: name}).Count()
	return count > 0
}

// QueryCompanyList 查询公司列表
func (s *sCompany) QueryCompanyList(ctx context.Context, filter *sys_model.SearchParams) (*co_model.CompanyListRes, error) {
	data, err := daoctl.Query[*co_entity.Company](co_dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler), filter, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.GetConfig().CompanyName+"信息不存在", co_dao.Company.Table())
	}

	if data.Total > 0 {
		items := make([]*co_entity.Company, 0)
		// 脱敏处理
		for _, item := range *data.List {
			items = append(items, s.masker(item))
		}
		data.List = &items
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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.GetConfig().CompanyName+"信息不存在", co_dao.Company.Table())
	}
	return s.saveCompany(ctx, info)
}

// SaveCompany 保存公司信息
func (s *sCompany) saveCompany(ctx context.Context, info *co_model.Company) (*co_entity.Company, error) {
	// 名称重名检测
	if s.HasCompanyByName(ctx, info.Name, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.GetConfig().CompanyName+"名称已被占用，请修改后再试", co_dao.Company.Table())
	}

	data := co_do.Company{
		Name:          info.Name,
		ContactName:   info.ContactName,
		ContactMobile: info.ContactMobile,
		Remark:        info.Remark,
	}

	// 获取登录用户
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 启用事务
	err := co_dao.Company.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		if info.Id == 0 {
			data.Id = idgen.NextId()
			data.CreatedBy = user.Id
			data.CreatedAt = gtime.Now()

			_, err = co_dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(data)

		} else {
			data.UpdatedBy = user.Id
			data.UpdatedAt = gtime.Now()
			_, err = co_dao.Company.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.Company{Id: info.Id}).Update(data)
		}
		return err
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
