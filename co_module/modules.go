package co_module

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/company"
)

type Modules struct {
	conf     *co_model.Config
	company  co_interface.ICompany
	team     co_interface.ITeam
	employee co_interface.IEmployee
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
	return m.GetConfig().I18n.Translate(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.GetConfig().I18n.TranslateFormat(ctx, format, values...)
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
