package co_module

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/company"
	"github.com/gogf/gf/v2/container/gmap"
)

type ModulesDao[TDaoCompay any, TDaoEmployee any, TDaoTeam any, TDaoTeamMember any] struct {
	Company    co_dao.IDao[TDaoCompay]
	Employee   co_dao.IDao[TDaoEmployee]
	Team       co_dao.IDao[TDaoTeam]
	TeamMember co_dao.IDao[TDaoTeamMember]
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
