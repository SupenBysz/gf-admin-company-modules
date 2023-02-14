package router

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/example/controller"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func ModulesGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	CompanyGroup(modules, group)
	EmployeeGroup(modules, group)
	TeamGroup(modules, group)
	MyGroup(modules, group)
	// FinancialGroup(modules, group)

	return group
}

func CompanyGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Company)
	controller := controller.Company(modules)

	group.POST(routePrefix+"/createCompany", controller.CreateCompany)
	group.POST(routePrefix+"/updateCompany", controller.UpdateCompany)
	group.POST(routePrefix+"/hasCompanyByName", controller.HasCompanyByName)
	group.POST(routePrefix+"/getCompanyById", controller.GetCompanyById)
	group.POST(routePrefix+"/queryCompanyList", controller.QueryCompanyList)
	group.POST(routePrefix+"/getCompanyDetail", controller.GetCompanyDetail)
	return group
}

func EmployeeGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Employee)
	controller := controller.Employee(modules)
	group.POST(routePrefix+"/getEmployeeById", controller.GetEmployeeById)
	group.POST(routePrefix+"/getEmployeeDetailById", controller.GetEmployeeDetailById)
	group.POST(routePrefix+"/hasEmployeeByName", controller.HasEmployeeByName)
	group.POST(routePrefix+"/hasEmployeeByNo", controller.HasEmployeeByNo)
	group.POST(routePrefix+"/queryEmployeeList", controller.QueryEmployeeList)
	group.POST(routePrefix+"/createEmployee", controller.CreateEmployee)
	group.POST(routePrefix+"/updateEmployee", controller.UpdateEmployee)
	group.POST(routePrefix+"/deleteEmployee", controller.DeleteEmployee)
	group.POST(routePrefix+"/getEmployeeListByRoleId", controller.GetEmployeeListByRoleId)
	return group
}

func TeamGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Team)
	controller := controller.Team(modules)
	group.POST(routePrefix+"/getTeamById", controller.GetTeamById)
	group.POST(routePrefix+"/hasTeamByName", controller.HasTeamByName)
	group.POST(routePrefix+"/queryTeamList", controller.QueryTeamList)
	group.POST(routePrefix+"/createTeam", controller.CreateTeam)
	group.POST(routePrefix+"/updateTeam", controller.UpdateTeam)
	group.POST(routePrefix+"/getTeamMemberList", controller.GetTeamMemberList)
	group.POST(routePrefix+"/queryTeamListByEmployee", controller.QueryTeamListByEmployee)
	group.POST(routePrefix+"/setTeamMember", controller.SetTeamMember)
	group.POST(routePrefix+"/setTeamOwner", controller.SetTeamOwner)
	group.POST(routePrefix+"/setTeamCaptain", controller.SetTeamCaptain)
	group.POST(routePrefix+"/deleteTeam", controller.DeleteTeam)
	return group
}

func MyGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	controller := controller.My(modules)
	routePrefix := modules.GetConfig().RoutePrefix + "/my"
	group.POST(routePrefix+"/getProfile", controller.GetProfile)
	group.POST(routePrefix+"/getCompany", controller.GetCompany)
	group.POST(routePrefix+"/getTeams", controller.GetTeams)
	group.POST(routePrefix+"/setAvatar", controller.SetAvatar)
	group.POST(routePrefix+"/setMobile", controller.SetMobile)
	return group
}

func FinancialGroup(modules co_interface.IModules, group *ghttp.RouterGroup) *ghttp.RouterGroup {
	controller := controller.Financial(modules)
	routePrefix := modules.GetConfig().RoutePrefix + "/financial"
	group.POST(routePrefix+"/registerBankCard", controller.BankCardRegister)
	group.POST(routePrefix+"/deleteBankCard", controller.DeleteBankCard)
	group.POST(routePrefix+"/queryBankCardList", controller.QueryBankCardList)
	group.POST(routePrefix+"/getAccountBalance", controller.GetAccountBalance)
	group.POST(routePrefix+"/invoiceRegister", controller.InvoiceRegister)
	group.POST(routePrefix+"/queryInvoice", controller.QueryInvoice)
	group.POST(routePrefix+"/deleteInvoiceById", controller.DeletesFdInvoiceById)
	group.POST(routePrefix+"/invoiceDetailRegister", controller.InvoiceDetailRegister)
	group.POST(routePrefix+"/queryInvoiceDetailList", controller.QueryInvoiceDetailList)
	group.POST(routePrefix+"/makeInvoiceDetail", controller.MakeInvoiceDetailReq)
	group.POST(routePrefix+"/auditInvoiceDetail", controller.AuditInvoiceDetail)

	return group
}
