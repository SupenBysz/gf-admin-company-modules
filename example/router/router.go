package router

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/example/controller"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func ModulesGroup[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillsRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	CompanyGroup(modules, group)
	EmployeeGroup(modules, group)
	TeamGroup(modules, group)
	MyGroup(modules, group)
	FinanceGroup(modules, group)

	return group
}

func CompanyGroup[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillsRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Company)
	ctrl := controller.Company(modules)

	group.POST(routePrefix+"/createCompany", ctrl.CreateCompany)
	group.POST(routePrefix+"/updateCompany", ctrl.UpdateCompany)
	group.POST(routePrefix+"/hasCompanyByName", ctrl.HasCompanyByName)
	group.POST(routePrefix+"/getCompanyById", ctrl.GetCompanyById)
	group.POST(routePrefix+"/queryCompanyList", ctrl.QueryCompanyList)
	group.POST(routePrefix+"/getCompanyDetail", ctrl.GetCompanyDetail)
	group.POST(routePrefix+"/setCompanyState", ctrl.SetCompanyState)
	return group
}

func EmployeeGroup[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillsRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Employee)
	ctrl := controller.Employee(modules)

	group.POST(routePrefix+"/getEmployeeById", ctrl.GetEmployeeById)
	group.POST(routePrefix+"/getEmployeeDetailById", ctrl.GetEmployeeDetailById)
	group.POST(routePrefix+"/hasEmployeeByName", ctrl.HasEmployeeByName)
	group.POST(routePrefix+"/hasEmployeeByNo", ctrl.HasEmployeeByNo)
	group.POST(routePrefix+"/queryEmployeeList", ctrl.QueryEmployeeList)
	group.POST(routePrefix+"/createEmployee", ctrl.CreateEmployee)
	group.POST(routePrefix+"/updateEmployee", ctrl.UpdateEmployee)
	group.POST(routePrefix+"/deleteEmployee", ctrl.DeleteEmployee)
	group.POST(routePrefix+"/getEmployeeListByRoleId", ctrl.GetEmployeeListByRoleId)
	group.POST(routePrefix+"/setEmployeeRoles", ctrl.SetEmployeeRoles)
	group.POST(routePrefix+"/setEmployeeState", ctrl.SetEmployeeState)

	return group
}

func TeamGroup[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillsRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := modules.GetConfig().RoutePrefix + "/" + gstr.LcFirst(modules.GetConfig().Identifier.Team)
	ctrl := controller.Team(modules)

	group.POST(routePrefix+"/getTeamById", ctrl.GetTeamById)
	group.POST(routePrefix+"/hasTeamByName", ctrl.HasTeamByName)
	group.POST(routePrefix+"/queryTeamList", ctrl.QueryTeamList)
	group.POST(routePrefix+"/createTeam", ctrl.CreateTeam)
	group.POST(routePrefix+"/updateTeam", ctrl.UpdateTeam)
	// group.POST(routePrefix+"/queryTeamListByEmployee", ctrl.QueryTeamListByEmployee)
	group.POST(routePrefix+"/setTeamMember", ctrl.SetTeamMember)
	group.POST(routePrefix+"/removeTeamMember", ctrl.RemoveTeamMember)
	group.POST(routePrefix+"/setTeamOwner", ctrl.SetTeamOwner)
	group.POST(routePrefix+"/setTeamCaptain", ctrl.SetTeamCaptain)
	group.POST(routePrefix+"/deleteTeam", ctrl.DeleteTeam)
	group.POST(routePrefix+"/getEmployeeListByTeamId", ctrl.GetEmployeeListByTeamId)
	group.POST(routePrefix+"/getTeamInviteCode", ctrl.GetTeamInviteCode)
	group.POST(routePrefix+"/joinTeamByInviteCode", ctrl.JoinTeamByInviteCode)

	return group
}

func MyGroup[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillsRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	ctrl := controller.My(modules)
	routePrefix := modules.GetConfig().RoutePrefix + "/my"
	group.POST(routePrefix+"/getProfile", ctrl.GetProfile)
	group.POST(routePrefix+"/getCompany", ctrl.GetCompany)
	group.POST(routePrefix+"/getTeams", ctrl.GetTeams)
	group.POST(routePrefix+"/setAvatar", ctrl.SetAvatar)
	group.POST(routePrefix+"/setEmployeeMobile", ctrl.SetMobile)
	group.POST(routePrefix+"/setEmployeeMail", ctrl.SetMail)

	// 我的财务相关
	group.POST(routePrefix+"/getAccountBills", ctrl.GetAccountBills)
	group.POST(routePrefix+"/getAccounts", ctrl.GetAccounts)
	group.POST(routePrefix+"/getBankCards", ctrl.GetBankCards)
	group.POST(routePrefix+"/getInvoices", ctrl.GetInvoices)
	group.POST(routePrefix+"/updateAccount", ctrl.UpdateAccount)

	return group
}

func FinanceGroup[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillsRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	ctrl := controller.Finance(modules)
	routePrefix := modules.GetConfig().RoutePrefix + "/finance"

	group.POST(routePrefix+"/registerBankCard", ctrl.BankCardRegister)
	group.POST(routePrefix+"/deleteBankCard", ctrl.DeleteBankCard)
	group.POST(routePrefix+"/queryBankCardList", ctrl.QueryBankCardList)
	group.POST(routePrefix+"/getAccountBalance", ctrl.GetAccountBalance)
	group.POST(routePrefix+"/invoiceRegister", ctrl.InvoiceRegister)
	group.POST(routePrefix+"/queryInvoice", ctrl.QueryInvoice)
	group.POST(routePrefix+"/deleteInvoiceById", ctrl.DeletesFdInvoiceById)
	group.POST(routePrefix+"/invoiceDetailRegister", ctrl.InvoiceDetailRegister)
	group.POST(routePrefix+"/queryInvoiceDetailList", ctrl.QueryInvoiceDetailList)
	group.POST(routePrefix+"/makeInvoiceDetail", ctrl.MakeInvoiceDetailReq)
	group.POST(routePrefix+"/auditInvoiceDetail", ctrl.AuditInvoiceDetail)

	group.POST(routePrefix+"/getAccountDetail", ctrl.GetAccountDetail)
	group.POST(routePrefix+"/updateAccountIsEnabled", ctrl.UpdateAccountIsEnabled)
	group.POST(routePrefix+"/updateAccountLimitState", ctrl.UpdateAccountLimitState)
	group.POST(routePrefix+"/setAccountCurrencyCode", ctrl.SetAccountCurrencyCode)
	group.POST(routePrefix+"/getCurrencyByCode", ctrl.GetCurrencyByCode)

	group.POST(routePrefix+"/getAccountDetailByAccountId", ctrl.GetAccountDetailById)
	//group.POST(routePrefix+"/increment", ctrl.Increment)
	//group.POST(routePrefix+"/decrement", ctrl.Decrement)
	group.POST(routePrefix+"/setAccountAllowExceed", ctrl.SetAccountAllowExceed)

	return group
}
