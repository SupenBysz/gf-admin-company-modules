package co_module

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/company"
	"github.com/SupenBysz/gf-admin-company-modules/utility/dao_helper"
	"github.com/gogf/gf/v2/container/gmap"
)

type ModulesDao[TDaoCompay any, TDaoEmployee any, TDaoTeam any, TDaoTeamMember any] struct {
	Company    dao_helper.IDao[TDaoCompay]
	Employee   dao_helper.IDao[TDaoEmployee]
	Team       dao_helper.IDao[TDaoTeam]
	TeamMember dao_helper.IDao[TDaoTeamMember]
}

type Modules[C any, E any, T any, TM any] struct {
	conf     *co_model.Config
	company  co_interface.ICompany
	team     co_interface.ITeam
	employee co_interface.IEmployee
	dao      ModulesDao[C, E, T, TM]
}

func (m *Modules[C, E, T, TM]) Company() co_interface.ICompany {
	return m.company
}

func (m *Modules[C, E, T, TM]) Team() co_interface.ITeam {
	return m.team
}

func (m *Modules[C, E, T, TM]) Employee() co_interface.IEmployee {
	return m.employee
}

func (m *Modules[C, E, T, TM]) GetConfig() *co_model.Config {
	return m.conf
}

func (m *Modules[C, E, T, TM]) T(ctx context.Context, content string) string {
	return m.GetConfig().I18n.Translate(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules[C, E, T, TM]) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.GetConfig().I18n.TranslateFormat(ctx, format, values...)
}

var (
	modulesMap = gmap.NewStrAnyMap()
)

func NewModules[TDaoCompay any, TDaoEmployee any, TDaoTeam any, TDaoTeamMember any](conf *co_model.Config) *Modules[TDaoCompay, TDaoEmployee, TDaoTeam, TDaoTeamMember] {
	result := &Modules[TDaoCompay, TDaoEmployee, TDaoTeam, TDaoTeamMember]{
		conf: conf,
	}
	result.company = company.NewCompany(result)
	result.employee = company.NewEmployee(result)
	result.team = company.NewTeam(result)
	modulesMap.Set(conf.KeyIndex, result)
	return result
}

func GetModule[TDaoCompay any, TDaoEmployee any, TDaoTeam any, TDaoTeamMember any](moduleKey string) *Modules[TDaoCompay, TDaoEmployee, TDaoTeam, TDaoTeamMember] {
	v, ok := modulesMap.Get(moduleKey).(*Modules[TDaoCompay, TDaoEmployee, TDaoTeam, TDaoTeamMember])
	if ok {
		return v
	}
	return nil
}
