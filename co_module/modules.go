package co_module

import (
	"github.com/SupenBysz/gf-admin-community/utility/kmap"
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

var (
	modulesMap = kmap.New[string, *Modules]()
)

func NewModules(conf *co_model.Config) *Modules {
	result := &Modules{
		conf: conf,
	}
	result.company = company.NewCompany(result)
	result.employee = company.NewEmployee(result)
	result.team = company.NewTeam(result)
	modulesMap.Set(conf.KeyIndex, result)
	return result
}

func GetModule(moduleKey string) *Modules {
	return modulesMap.Get(moduleKey)
}
