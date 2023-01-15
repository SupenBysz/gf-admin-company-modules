package co_module

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/company"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type Modules struct {
	conf     *co_model.Config
	company  co_interface.ICompany
	team     co_interface.ITeam
	employee co_interface.IEmployee
	my       co_interface.IMy
	i18n     *gi18n.Manager
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
		err := i18n.SetPath("i18n/" + m.conf.KeyIndex)
		if err != nil {
			return err
		}
	}

	m.i18n = i18n
	return nil
}

func NewModules[C any, E any, T any, TM any](
	conf *co_model.Config,
	Company func(module ...co_interface.IModules) C,
	Employee func(module ...co_interface.IModules) E,
	Team func(module ...co_interface.IModules) T,
	TeamMember func(module ...co_interface.IModules) TM,
) *Modules {
	result := &Modules{
		conf: conf,
	}
	result.company = company.NewCompany(result)
	result.employee = company.NewEmployee(result)
	result.team = company.NewTeam(result)
	result.my = company.NewMy(result)

	if result != nil {
		Company(result)
	}
	if result != nil {
		Employee(result)
	}
	if result != nil {
		Team(result)
	}
	if result != nil {
		TeamMember(result)
	}

	return result
}
