package co_router

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func ModulesGroup(module co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	CompanyGroup(module, group)
	EmployeeGroup(module, group)
	TeamGroup(module, group)
	return group
}

func CompanyGroup(module co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := module.GetConfig().RoutePrefix + "/" + gstr.LcFirst(module.GetConfig().Identifier.Company)
	group.POST(routePrefix+"/createCompany", co_controller.Company(module).CreateCompany)
	group.POST(routePrefix+"/updateCompany", co_controller.Company(module).UpdateCompany)
	group.POST(routePrefix+"/hasCompanyByName", co_controller.Company(module).HasCompanyByName)
	group.POST(routePrefix+"/getCompanyById", co_controller.Company(module).GetCompanyById)
	group.POST(routePrefix+"/queryCompanyList", co_controller.Company(module).QueryCompanyList)
	return group
}

func EmployeeGroup(module co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := module.GetConfig().RoutePrefix + "/" + gstr.LcFirst(module.GetConfig().Identifier.Employee)
	group.POST(routePrefix+"/getEmployeeById", co_controller.Employee(module).GetEmployeeById)
	group.POST(routePrefix+"/getEmployeeDetailById", co_controller.Employee(module).GetEmployeeDetailById)
	group.POST(routePrefix+"/hasEmployeeByName", co_controller.Employee(module).HasEmployeeByName)
	group.POST(routePrefix+"/hasEmployeeByNo", co_controller.Employee(module).HasEmployeeByNo)
	group.POST(routePrefix+"/queryEmployeeList", co_controller.Employee(module).QueryEmployeeList)
	group.POST(routePrefix+"/createEmployee", co_controller.Employee(module).CreateEmployee)
	group.POST(routePrefix+"/updateEmployee", co_controller.Employee(module).UpdateEmployee)
	group.POST(routePrefix+"/deleteEmployee", co_controller.Employee(module).DeleteEmployee)
	group.POST(routePrefix+"/setEmployeeMobile", co_controller.Employee(module).SetEmployeeMobile)
	group.POST(routePrefix+"/setEmployeeAvatar", co_controller.Employee(module).SetEmployeeAvatar)
	return group
}

func TeamGroup(module co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := module.GetConfig().RoutePrefix + "/" + gstr.LcFirst(module.GetConfig().Identifier.Team)
	group.POST(routePrefix+"/getTeamById", co_controller.Team(module).GetTeamById)
	group.POST(routePrefix+"/hasTeamByName", co_controller.Team(module).HasTeamByName)
	group.POST(routePrefix+"/queryTeamList", co_controller.Team(module).QueryTeamList)
	group.POST(routePrefix+"/createTeam", co_controller.Team(module).CreateTeam)
	group.POST(routePrefix+"/updateTeam", co_controller.Team(module).UpdateTeam)
	group.POST(routePrefix+"/getTeamMemberList", co_controller.Team(module).GetTeamMemberList)
	group.POST(routePrefix+"/queryTeamListByEmployee", co_controller.Team(module).QueryTeamListByEmployee)
	group.POST(routePrefix+"/setTeamMember", co_controller.Team(module).SetTeamMember)
	group.POST(routePrefix+"/setTeamOwner", co_controller.Team(module).SetTeamOwner)
	group.POST(routePrefix+"/setTeamCaptain", co_controller.Team(module).SetTeamCaptain)
	group.POST(routePrefix+"/deleteTeam", co_controller.Team(module).DeleteTeam)
	return group
}
