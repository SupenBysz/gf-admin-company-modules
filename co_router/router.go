package co_router

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller/system"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func SystemGroup(group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Group("/finance", func(group *ghttp.RouterGroup) {
			group.Bind(system.SystemFinance)
		})
	})
}

func ModulesGroup(modules co_interface.IModules[
	co_model.ICompanyRes,
	co_model.IEmployeeRes,
	co_model.ITeamRes,
	co_model.IFdAccountRes,
	co_model.IFdAccountBillsRes,
	co_model.IFdBankCardRes,
	co_model.IFdInvoiceRes,
	co_model.IFdInvoiceDetailRes,
	co_model.IFdRechargeRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	CompanyGroup(modules, group)
	EmployeeGroup(modules, group)
	TeamGroup(modules, group)
	MyGroup(modules, group)
	FinanceGroup(modules, group)

	return group
}

func CompanyGroup(modules co_interface.IModules[
	co_model.ICompanyRes,
	co_model.IEmployeeRes,
	co_model.ITeamRes,
	co_model.IFdAccountRes,
	co_model.IFdAccountBillsRes,
	co_model.IFdBankCardRes,
	co_model.IFdInvoiceRes,
	co_model.IFdInvoiceDetailRes,
	co_model.IFdRechargeRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Company)
	controller := co_controller.Company(modules)

	group.POST(routePrefix+"/createCompany", controller.CreateCompany)
	group.POST(routePrefix+"/updateCompany", controller.UpdateCompany)
	group.POST(routePrefix+"/hasCompanyByName", controller.HasCompanyByName)
	group.POST(routePrefix+"/getCompanyById", controller.GetCompanyById)
	group.POST(routePrefix+"/queryCompanyList", controller.QueryCompanyList)
	group.POST(routePrefix+"/getCompanyDetail", controller.GetCompanyDetail)
	group.POST(routePrefix+"/setCompanyState", controller.SetCompanyState)
	return group
}

func EmployeeGroup(modules co_interface.IModules[
	co_model.ICompanyRes,
	co_model.IEmployeeRes,
	co_model.ITeamRes,
	co_model.IFdAccountRes,
	co_model.IFdAccountBillsRes,
	co_model.IFdBankCardRes,
	co_model.IFdInvoiceRes,
	co_model.IFdInvoiceDetailRes,
	co_model.IFdRechargeRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Employee)
	controller := co_controller.Employee(modules)
	group.POST(routePrefix+"/getEmployeeById", controller.GetEmployeeById)
	group.POST(routePrefix+"/getEmployeeDetailById", controller.GetEmployeeDetailById)
	group.POST(routePrefix+"/hasEmployeeByName", controller.HasEmployeeByName)
	group.POST(routePrefix+"/hasEmployeeByNo", controller.HasEmployeeByNo)
	group.POST(routePrefix+"/queryEmployeeList", controller.QueryEmployeeList)
	group.POST(routePrefix+"/createEmployee", controller.CreateEmployee)
	group.POST(routePrefix+"/updateEmployee", controller.UpdateEmployee)
	group.POST(routePrefix+"/deleteEmployee", controller.DeleteEmployee)
	group.POST(routePrefix+"/getEmployeeListByRoleId", controller.GetEmployeeListByRoleId)
	group.POST(routePrefix+"/setEmployeeRoles", controller.SetEmployeeRoles)
	group.POST(routePrefix+"/setEmployeeState", controller.SetEmployeeState)

	return group
}

func TeamGroup(modules co_interface.IModules[
	co_model.ICompanyRes,
	co_model.IEmployeeRes,
	co_model.ITeamRes,
	co_model.IFdAccountRes,
	co_model.IFdAccountBillsRes,
	co_model.IFdBankCardRes,
	co_model.IFdInvoiceRes,
	co_model.IFdInvoiceDetailRes,
	co_model.IFdRechargeRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Team)
	controller := co_controller.Team(modules)
	group.POST(routePrefix+"/getTeamById", controller.GetTeamById)
	group.POST(routePrefix+"/hasTeamByName", controller.HasTeamByName)
	group.POST(routePrefix+"/queryTeamList", controller.QueryTeamList)
	group.POST(routePrefix+"/createTeam", controller.CreateTeam)
	group.POST(routePrefix+"/updateTeam", controller.UpdateTeam)
	group.POST(routePrefix+"/queryTeamListByEmployee", controller.QueryTeamListByEmployee)
	group.POST(routePrefix+"/setTeamMember", controller.SetTeamMember)
	group.POST(routePrefix+"/removeTeamMember", controller.RemoveTeamMember)
	group.POST(routePrefix+"/setTeamOwner", controller.SetTeamOwner)
	group.POST(routePrefix+"/setTeamCaptain", controller.SetTeamCaptain)
	group.POST(routePrefix+"/deleteTeam", controller.DeleteTeam)
	group.POST(routePrefix+"/getEmployeeListByTeamId", controller.GetEmployeeListByTeamId)
	group.POST(routePrefix+"/getTeamInviteCode", controller.GetTeamInviteCode)
	group.POST(routePrefix+"/joinTeamByInviteCode", controller.JoinTeamByInviteCode)

	return group
}

func MyGroup(modules co_interface.IModules[
	co_model.ICompanyRes,
	co_model.IEmployeeRes,
	co_model.ITeamRes,
	co_model.IFdAccountRes,
	co_model.IFdAccountBillsRes,
	co_model.IFdBankCardRes,
	co_model.IFdInvoiceRes,
	co_model.IFdInvoiceDetailRes,
	co_model.IFdRechargeRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	controller := co_controller.My(modules)
	routePrefix := modules.GetConfig().RoutePrefix + "/my"
	group.POST(routePrefix+"/getProfile", controller.GetProfile)
	group.POST(routePrefix+"/getCompany", controller.GetCompany)
	group.POST(routePrefix+"/getTeams", controller.GetTeams)
	group.POST(routePrefix+"/setAvatar", controller.SetAvatar)
	group.POST(routePrefix+"/setEmployeeMobile", controller.SetMobile)
	group.POST(routePrefix+"/setEmployeeMail", controller.SetMail)

	// 我的财务相关
	group.POST(routePrefix+"/getAccountBills", controller.GetAccountBills)
	group.POST(routePrefix+"/getAccounts", controller.GetAccounts)
	group.POST(routePrefix+"/getBankCards", controller.GetBankCards)
	group.POST(routePrefix+"/getInvoices", controller.GetInvoices)
	group.POST(routePrefix+"/updateAccount", controller.UpdateAccount)
	return group
}

func FinanceGroup(modules co_interface.IModules[
	co_model.ICompanyRes,
	co_model.IEmployeeRes,
	co_model.ITeamRes,
	co_model.IFdAccountRes,
	co_model.IFdAccountBillsRes,
	co_model.IFdBankCardRes,
	co_model.IFdInvoiceRes,
	co_model.IFdInvoiceDetailRes,
	co_model.IFdRechargeRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	// routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Company)

	controller := co_controller.Finance(modules)
	routePrefix := modules.GetConfig().RoutePrefix + "/finance"
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

	group.POST(routePrefix+"/getAccountDetail", controller.GetAccountDetail)
	group.POST(routePrefix+"/updateAccountIsEnabled", controller.UpdateAccountIsEnabled)
	group.POST(routePrefix+"/updateAccountLimitState", controller.UpdateAccountLimitState)
	group.POST(routePrefix+"/reversedAmount", controller.ReversedAmount)

	group.POST(routePrefix+"/getAccountDetailByAccountId", controller.GetAccountDetailById)
	//group.POST(routePrefix+"/increment", controller.Increment)
	//group.POST(routePrefix+"/decrement", controller.Decrement)
	group.POST(routePrefix+"/setAccountAllowExceed", controller.SetAccountAllowExceed)

	return group
}
