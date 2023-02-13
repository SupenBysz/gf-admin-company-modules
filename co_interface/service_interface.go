package co_interface

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type (
	ICompany interface {
		GetCompanyById(ctx context.Context, id int64) (*co_model.CompanyRes, error)
		GetCompanyByName(ctx context.Context, name string) (*co_model.CompanyRes, error)
		HasCompanyByName(ctx context.Context, name string, excludeIds ...int64) bool
		QueryCompanyList(ctx context.Context, filter *sys_model.SearchParams) (*co_model.CompanyListRes, error)
		CreateCompany(ctx context.Context, info *co_model.Company) (*co_model.CompanyRes, error)
		UpdateCompany(ctx context.Context, info *co_model.Company) (*co_model.CompanyRes, error)
		GetCompanyDetail(ctx context.Context, id int64) (*co_model.CompanyRes, error)
		FilterUnionMainId(ctx context.Context, search *sys_model.SearchParams) *sys_model.SearchParams
	}
	IEmployee interface {
		GetEmployeeById(ctx context.Context, id int64) (*co_model.EmployeeRes, error)
		GetEmployeeByName(ctx context.Context, name string) (*co_model.EmployeeRes, error)
		HasEmployeeByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool
		HasEmployeeByNo(ctx context.Context, no string, unionMainId int64, excludeIds ...int64) bool
		GetEmployeeBySession(ctx context.Context) (*co_model.EmployeeRes, error)
		QueryEmployeeList(ctx context.Context, search *sys_model.SearchParams) (*co_model.EmployeeListRes, error)
		CreateEmployee(ctx context.Context, info *co_model.Employee) (*co_model.EmployeeRes, error)
		UpdateEmployee(ctx context.Context, info *co_model.Employee) (*co_model.EmployeeRes, error)
		DeleteEmployee(ctx context.Context, id int64) (bool, error)
		GetEmployeeDetailById(ctx context.Context, id int64) (*co_model.EmployeeRes, error)
		GetEmployeeListByRoleId(ctx context.Context, roleId int64) (*co_model.EmployeeListRes, error)
	}
	ITeam interface {
		GetTeamById(ctx context.Context, id int64) (*co_model.TeamRes, error)
		GetTeamByName(ctx context.Context, name string) (*co_model.TeamRes, error)
		HasTeamByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool
		QueryTeamList(ctx context.Context, search *sys_model.SearchParams) (*co_model.TeamListRes, error)
		QueryTeamMemberList(ctx context.Context, search *sys_model.SearchParams) (*co_model.TeamMemberListRes, error)
		CreateTeam(ctx context.Context, info *co_model.Team) (*co_model.TeamRes, error)
		UpdateTeam(ctx context.Context, id int64, name string, remark string) (*co_model.TeamRes, error)
		GetTeamMemberList(ctx context.Context, id int64) (*co_model.EmployeeListRes, error)
		QueryTeamListByEmployee(ctx context.Context, employeeId int64, unionMainId int64) (*co_model.TeamListRes, error)
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
		SetMyMobile(ctx context.Context, newMobile int64, captcha string, password string) (bool, error)
		SetMyAvatar(ctx context.Context, imageId int64) (bool, error)
	}

	IFdBankCard interface {
		CreateBankCard(ctx context.Context, info co_model.BankCardRegister, user *sys_model.SysUser) (*co_entity.FdBankCard, error)
		GetBankCardById(ctx context.Context, id int64) (*co_entity.FdBankCard, error)
		GetBankCardByCardNumber(ctx context.Context, cardNumber string) (*co_entity.FdBankCard, error)
		UpdateBankCardState(ctx context.Context, bankCardId int64, state int) (bool, error)
		DeleteBankCardById(ctx context.Context, bankCardId int64) (bool, error)
		QueryBankCardListByUserId(ctx context.Context, userId int64) (*co_model.BankCardListRes, error)
	}
	IFdCurrency interface {
		GetCurrencyByCurrencyCode(ctx context.Context, currencyCode string) (*co_entity.FdCurrency, error)
		GetCurrencyByCnName(ctx context.Context, cnName string) (*co_entity.FdCurrency, error)
	}
	IFdInvoice interface {
		CreateInvoice(ctx context.Context, info co_model.FdInvoiceRegister) (*co_entity.FdInvoice, error)
		GetInvoiceById(ctx context.Context, id int64) (*co_entity.FdInvoice, error)
		QueryInvoiceList(ctx context.Context, info *sys_model.SearchParams, userId int64) (*co_model.FdInvoiceListRes, error)
		DeletesFdInvoiceById(ctx context.Context, invoiceId int64) (bool, error)
		GetFdInvoiceByTaxId(ctx context.Context, taxId string) (*co_entity.FdInvoice, error)
	}
	IFdInvoiceDetail interface {
		CreateInvoiceDetail(ctx context.Context, info co_model.FdInvoiceDetailRegister) (*co_entity.FdInvoiceDetail, error)
		GetInvoiceDetailById(ctx context.Context, id int64) (*co_entity.FdInvoiceDetail, error)
		MakeInvoiceDetail(ctx context.Context, invoiceDetailId int64, makeInvoiceDetail co_model.FdMakeInvoiceDetail) (res bool, err error)
		AuditInvoiceDetail(ctx context.Context, invoiceDetailId int64, auditInfo co_model.FdInvoiceAuditInfo) (bool, error)
		QueryInvoiceDetailListByInvoiceId(ctx context.Context, invoiceId int64) (*co_model.FdInvoiceDetailListRes, error)
		DeleteInvoiceDetail(ctx context.Context, id int64) (bool, error)
		QueryInvoiceDetail(ctx context.Context, info *sys_model.SearchParams, userId int64, unionMainId int64) (*co_model.FdInvoiceDetailListRes, error)
	}
	IFdAccount interface {
		CreateAccount(ctx context.Context, info co_model.FdAccountRegister) (*co_entity.FdAccount, error)
		GetAccountById(ctx context.Context, id int64) (*co_entity.FdAccount, error)
		UpdateAccountIsEnable(ctx context.Context, id int64, isEnabled int64) (bool, error)
		HasAccountByName(ctx context.Context, name string) (*co_entity.FdAccount, error)
		UpdateAccountLimitState(ctx context.Context, id int64, limitState int64) (bool, error)
		QueryAccountListByUserId(ctx context.Context, userId int64) (*co_model.AccountList, error)
		UpdateAccountBalance(ctx context.Context, accountId int64, amount int64, version int, inOutType int) (int64, error)
		GetAccountByUnionUserIdAndCurrencyCode(ctx context.Context, unionUserId int64, currencyCode string) (*co_entity.FdAccount, error)
	}
	IFdAccountBill interface {
		InstallHook(filter co_model.AccountBillHookFilter, hookFunc co_model.AccountBillHookFunc)
		UnInstallHook(filter co_model.AccountBillHookFilter)
		ClearAllHook()
		CreateAccountBill(ctx context.Context, info co_model.AccountBillRegister) (bool, error)
		GetAccountBillByAccountId(ctx context.Context, accountId int64, pagination *sys_model.Pagination) (*co_model.AccountBillListRes, error)
	}
)

type IModules interface {
	Company() ICompany
	Team() ITeam
	Employee() IEmployee
	GetConfig() *co_model.Config
	My() IMy
	BankCard() IFdBankCard
	Currency() IFdCurrency
	Invoice() IFdInvoice
	InvoiceDetail() IFdInvoiceDetail
	Account() IFdAccount
	AccountBill() IFdAccountBill

	SetI18n(i18n *gi18n.Manager) error
	T(ctx context.Context, content string) string
	Tf(ctx context.Context, format string, values ...interface{}) string
	Dao() *co_dao.XDao
}
