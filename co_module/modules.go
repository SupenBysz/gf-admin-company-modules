package co_module

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/internal/boot"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/company"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/text/gstr"
)

type Modules struct {
	conf     *co_model.Config
	company  co_interface.ICompany
	team     co_interface.ITeam
	employee co_interface.IEmployee
	my       co_interface.IMy
	i18n     *gi18n.Manager
	xDao     *co_dao.XDao
}

func (m *Modules) My() co_interface.IMy {
	return m.my
}

func (m *Modules) Company() co_interface.ICompany {
	return m.company
}

func (m *Modules) Team() co_interface.ITeam {
	return m.team
}

func (m *Modules) Employee() co_interface.IEmployee {
	return m.employee
}

func (m *Modules) GetConfig() *co_model.Config {
	return m.conf
}

func (m *Modules) T(ctx context.Context, content string) string {
	return m.i18n.Translate(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.i18n.TranslateFormat(ctx, format, values...)
}

func (m *Modules) SetI18n(i18n *gi18n.Manager) error {
	if i18n == nil {
		i18n = gi18n.New()
		i18n.SetLanguage("zh-CN")
		err := i18n.SetPath("i18n/" + gstr.ToLower(m.conf.KeyIndex))
		if err != nil {
			return err
		}
	}

	m.i18n = i18n
	return nil
}

func (m *Modules) Dao() *co_dao.XDao {
	return m.xDao
}

func NewModules(
	conf *co_model.Config,
	xDao *co_dao.XDao,
) *Modules {
	module := &Modules{
		conf: conf,
		xDao: xDao,
	}

	// 初始化默认多语言对象
	module.SetI18n(nil)

	module.company = company.NewCompany(module)
	module.employee = company.NewEmployee(module)
	module.team = company.NewTeam(module)
	module.my = company.NewMy(module)

	// 权限树追加权限
	co_consts.PermissionTree = append(co_consts.PermissionTree, boot.InitPermission(module)...)

	return module
}
