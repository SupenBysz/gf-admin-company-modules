package co_interface

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_hook"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
)

type (
	ICompany[TR co_model.ICompanyRes] interface {
		GetCompanyById(ctx context.Context, id int64) (response TR, err error)
		GetCompanyByName(ctx context.Context, name string) (response TR, err error)
		HasCompanyByName(ctx context.Context, name string, excludeIds ...int64) bool
		QueryCompanyList(ctx context.Context, filter *base_model.SearchParams) (*base_model.CollectRes[TR], error)
		CreateCompany(ctx context.Context, info *co_model.Company) (response TR, err error)
		UpdateCompany(ctx context.Context, info *co_model.Company) (response TR, err error)
		GetCompanyDetail(ctx context.Context, id int64) (response TR, err error)
		FilterUnionMainId(ctx context.Context, search *base_model.SearchParams) *base_model.SearchParams
	}
	IEmployee[TR co_model.IEmployeeRes] interface {
		GetModules() IModules[
			*co_model.CompanyRes,
			*co_model.EmployeeRes,
			*co_model.TeamRes,
			*co_model.FdAccountRes,
			*co_model.FdAccountBillRes,
			*co_model.FdBankCardRes,
			*co_model.FdCurrencyRes,
			*co_model.FdInvoiceRes,
			*co_model.FdInvoiceDetailRes,
		]
		SetXDao(dao co_dao.XDao)
		GetEmployeeById(ctx context.Context, id int64) (response TR, err error)
		GetEmployeeByName(ctx context.Context, name string) (response TR, err error)
		HasEmployeeByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool
		HasEmployeeByNo(ctx context.Context, no string, unionMainId int64, excludeIds ...int64) bool
		GetEmployeeBySession(ctx context.Context) (response TR, err error)
		QueryEmployeeList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error)
		CreateEmployee(ctx context.Context, info *co_model.Employee) (response TR, err error)
		UpdateEmployee(ctx context.Context, info *co_model.UpdateEmployee) (response TR, err error)
		UpdateEmployeeAvatar(ctx context.Context, id int64, avatar string) bool
		DeleteEmployee(ctx context.Context, id int64) (bool, error)
		GetEmployeeDetailById(ctx context.Context, id int64) (response TR, err error)
		GetEmployeeListByRoleId(ctx context.Context, roleId int64) (*base_model.CollectRes[TR], error)
		GetEmployeeListByTeamId(ctx context.Context, teamId int64) (*base_model.CollectRes[TR], error)
		SetEmployeeState(ctx context.Context, id int64, state int) (bool, error)
	}
	ITeam[TR co_model.ITeamRes] interface {
		SetXDao(dao co_dao.XDao)
		GetTeamById(ctx context.Context, id int64) (TR, error)
		GetTeamByName(ctx context.Context, name string) (TR, error)
		HasTeamByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool
		QueryTeamList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error)
		QueryTeamMemberList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[*co_model.TeamMemberRes], error)
		CreateTeam(ctx context.Context, info *co_model.Team) (TR, error)
		UpdateTeam(ctx context.Context, id int64, name string, remark string) (TR, error)
		QueryTeamListByEmployee(ctx context.Context, employeeId int64, unionMainId int64) (*base_model.CollectRes[TR], error)
		SetTeamMember(ctx context.Context, teamId int64, employeeIds []int64) (api_v1.BoolRes, error)
		SetTeamOwner(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error)
		SetTeamCaptain(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error)
		DeleteTeam(ctx context.Context, teamId int64) (api_v1.BoolRes, error)
		DeleteTeamMemberByEmployee(ctx context.Context, employeeId int64) (bool, error)
	}
	IMy interface {
		GetProfile(ctx context.Context) (*co_model.MyProfileRes, error)
		GetCompany(ctx context.Context) (*co_model.MyCompanyRes, error)
		GetTeams(ctx context.Context) (res co_model.MyTeamListRes, err error)
		SetMyMobile(ctx context.Context, newMobile string, captcha string, password string) (bool, error)
		SetMyAvatar(ctx context.Context, imageId int64) (bool, error)
		GetAccountBills(ctx context.Context, pagination *base_model.Pagination) (*co_model.MyAccountBillRes, error)
		GetAccounts(ctx context.Context) (*co_model.FdAccountListRes, error)
		GetBankCards(ctx context.Context) (*co_model.FdBankCardListRes, error)
		GetInvoices(ctx context.Context) (*co_model.FdInvoiceListRes, error)
		UpdateAccount(ctx context.Context, accountId int64, info *co_model.UpdateAccount) (api_v1.BoolRes, error)
	}

	IFdAccount[TR co_model.IFdAccountRes] interface {
		CreateAccount(ctx context.Context, info co_model.FdAccountRegister) (response TR, err error)
		GetAccountById(ctx context.Context, id int64) (response TR, err error)
		UpdateAccount(ctx context.Context, accountId int64, info *co_model.UpdateAccount) (bool, error)
		UpdateAccountIsEnable(ctx context.Context, id int64, isEnabled int) (bool, error)
		HasAccountByName(ctx context.Context, name string) (response TR, err error)
		UpdateAccountLimitState(ctx context.Context, id int64, limitState int) (bool, error)
		QueryAccountListByUserId(ctx context.Context, userId int64) (*base_model.CollectRes[TR], error)
		UpdateAccountBalance(ctx context.Context, accountId int64, amount int64, version int, inOutType int) (int64, error)
		GetAccountByUnionUserIdAndCurrencyCode(ctx context.Context, unionUserId int64, currencyCode string) (response TR, err error)
		GetAccountByUnionUserIdAndScene(ctx context.Context, unionUserId int64, accountType co_enum.AccountType, sceneType ...co_enum.SceneType) (response TR, err error)
		GetAccountDetailById(ctx context.Context, id int64) (res *co_model.FdAccountDetailRes, err error)
		Increment(ctx context.Context, id int64, amount int) (bool, error)
		Decrement(ctx context.Context, id int64, amount int) (bool, error)
		QueryDetailByUnionUserIdAndSceneType(ctx context.Context, unionUserId int64, sceneType co_enum.SceneType) (*base_model.CollectRes[co_model.FdAccountDetailRes], error)
	}
	IFdBankCard[TR co_model.IFdBankCardRes] interface {
		CreateBankCard(ctx context.Context, info co_model.BankCardRegister, user *sys_model.SysUser) (response TR, err error)
		GetBankCardById(ctx context.Context, id int64) (response TR, err error)
		GetBankCardByCardNumber(ctx context.Context, cardNumber string) (response TR, err error)
		UpdateBankCardState(ctx context.Context, bankCardId int64, state int) (bool, error)
		DeleteBankCardById(ctx context.Context, bankCardId int64) (bool, error)
		QueryBankCardListByUserId(ctx context.Context, userId int64) (*base_model.CollectRes[TR], error)
	}
	IFdCurrency[TR co_model.IFdCurrencyRes] interface {
		GetCurrencyByCurrencyCode(ctx context.Context, currencyCode string) (response TR, err error)
		GetCurrencyByCnName(ctx context.Context, cnName string) (response TR, err error)
	}
	IFdInvoice[TR co_model.IFdInvoiceRes] interface {
		CreateInvoice(ctx context.Context, info co_model.FdInvoiceRegister) (response TR, err error)
		GetInvoiceById(ctx context.Context, id int64) (response TR, err error)
		QueryInvoiceList(ctx context.Context, info *base_model.SearchParams, userId int64) (*base_model.CollectRes[TR], error)
		DeletesFdInvoiceById(ctx context.Context, invoiceId int64) (bool, error)
		GetFdInvoiceByTaxId(ctx context.Context, taxId string) (response TR, err error)
	}
	IFdInvoiceDetail[TR co_model.IFdInvoiceDetailRes] interface {
		CreateInvoiceDetail(ctx context.Context, info co_model.FdInvoiceDetailRegister) (response TR, err error)
		GetInvoiceDetailById(ctx context.Context, id int64) (response TR, err error)
		MakeInvoiceDetail(ctx context.Context, invoiceDetailId int64, makeInvoiceDetail co_model.FdMakeInvoiceDetail) (res bool, err error)
		AuditInvoiceDetail(ctx context.Context, invoiceDetailId int64, auditInfo co_model.FdInvoiceAuditInfo) (bool, error)
		QueryInvoiceDetailListByInvoiceId(ctx context.Context, invoiceId int64) (*base_model.CollectRes[TR], error)
		DeleteInvoiceDetail(ctx context.Context, id int64) (bool, error)
		QueryInvoiceDetail(ctx context.Context, info *base_model.SearchParams, userId int64, unionMainId int64) (*base_model.CollectRes[TR], error)
	}
	IFdAccountBill[TR co_model.IFdAccountBillRes] interface {
		InstallTradeHook(hookKey co_hook.AccountBillHookFilter, hookFunc co_hook.AccountBillHookFunc)
		GetTradeHook() base_hook.BaseHook[co_hook.AccountBillHookFilter, co_hook.AccountBillHookFunc]
		CreateAccountBill(ctx context.Context, info co_model.AccountBillRegister) (bool, error)
		GetAccountBillByAccountId(ctx context.Context, accountId int64, pagination *base_model.Pagination) (*base_model.CollectRes[TR], error)
	}
)

type IConfig interface {
	GetConfig() *co_model.Config
}

type ModuleFactory[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	NewEmployee func(modules IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]) IEmployee[ITEmployeeRes]

	NewTeam func(modules IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]) ITeam[ITTeamRes]
}

type IBaseFactory interface {
	NewEmployee(info co_dao.XDao) IEmployee[*co_model.EmployeeRes]
	//NewEmployee(info IEmployee[co_model.IEmployeeRes]) IEmployee[co_model.IEmployeeRes]

	//NewTeam(info ITeam[co_model.ITeamRes]) ITeam[co_model.ITeamRes]
	NewTeam(info co_dao.XDao) ITeam[*co_model.TeamRes]
}

type IModules[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] interface {
	IConfig
	Company() ICompany[ITCompanyRes]
	Team() ITeam[ITTeamRes]
	Employee() IEmployee[ITEmployeeRes]
	My() IMy
	Account() IFdAccount[ITFdAccountRes]
	AccountBill() IFdAccountBill[ITFdAccountBillRes]
	BankCard() IFdBankCard[ITFdBankCardRes]
	Currency() IFdCurrency[ITFdCurrencyRes]
	Invoice() IFdInvoice[ITFdInvoiceRes]
	InvoiceDetail() IFdInvoiceDetail[ITFdInvoiceDetailRes]

	SetI18n(i18n *gi18n.Manager) error
	T(ctx context.Context, content string) string
	Tf(ctx context.Context, format string, values ...interface{}) string
	Dao() *co_dao.XDao
}
