package co_interface

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
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
		// GetCompanyById 根据ID获取获取公司信息
		GetCompanyById(ctx context.Context, id int64) (response TR, err error)
		// GetCompanyByName 根据Name获取获取公司信息
		GetCompanyByName(ctx context.Context, name string) (response TR, err error)
		// HasCompanyByName 判断名称是否存在
		HasCompanyByName(ctx context.Context, name string, excludeIds ...int64) bool
		// QueryCompanyList 查询公司列表
		QueryCompanyList(ctx context.Context, filter *base_model.SearchParams, isExport ...bool) (*base_model.CollectRes[TR], error)
		// CreateCompany 创建公司信息
		CreateCompany(ctx context.Context, info *co_model.Company, bindUser *sys_model.SysUser) (response TR, err error)
		// UpdateCompany 更新公司信息
		UpdateCompany(ctx context.Context, info *co_model.Company) (response TR, err error)
		// GetCompanyDetail 获取公司详情，包含完整商务联系人电话
		GetCompanyDetail(ctx context.Context, id int64) (response TR, err error)
		// SetCompanyState 设置主体状态
		SetCompanyState(ctx context.Context, companyId int64, companyState co_enum.CompanyState) (bool, error)
		// SetCompanyAdminUser 设置主体的管理员用户
		SetCompanyAdminUser(ctx context.Context, sysUserId, unionMainId int64) (bool, error)
		// FilterUnionMainId 跨主体查询条件过滤
		FilterUnionMainId(ctx context.Context, search *base_model.SearchParams) *base_model.SearchParams
		// SetCommissionRate 设置公司佣金比例
		SetCommissionRate(ctx context.Context, companyId int64, commissionRate int, actionUserId int64) (bool, error)
	}
	IEmployee[TR co_model.IEmployeeRes] interface {
		SetXDao(dao co_dao.XDao)
		// GetEmployeeById 根据ID获取员工信息
		GetEmployeeById(ctx context.Context, id int64) (response TR, err error)
		// GetEmployeeByName 根据Name获取员工信息
		GetEmployeeByName(ctx context.Context, name string) (response TR, err error)
		// HasEmployeeByName 员工名称是否存在
		HasEmployeeByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool
		// HasEmployeeByNo 员工工号是否存在
		HasEmployeeByNo(ctx context.Context, no string, unionMainId int64, excludeIds ...int64) bool
		// GetEmployeeBySession 获取当前登录的员工信息
		GetEmployeeBySession(ctx context.Context) (response TR, err error)
		// QueryEmployeeList 获取员工列表
		QueryEmployeeList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error)
		// CreateEmployee 创建员工信息
		CreateEmployee(ctx context.Context, info *co_model.Employee, bindUser *sys_model.SysUser) (response TR, err error)
		// UpdateEmployee 更新员工信息
		UpdateEmployee(ctx context.Context, info *co_model.UpdateEmployee) (response TR, err error)
		// UpdateEmployeeAvatar 更新员工头像
		UpdateEmployeeAvatar(ctx context.Context, id int64, avatar string) bool
		// DeleteEmployee 删除员工信息
		DeleteEmployee(ctx context.Context, id int64) (bool, error)
		// GetEmployeeDetailById 根据ID获取员工详细信息
		GetEmployeeDetailById(ctx context.Context, id int64) (response TR, err error)
		// GetEmployeeListByRoleId 根据角色ID获取所有所属员工
		GetEmployeeListByRoleId(ctx context.Context, roleId int64) (*base_model.CollectRes[TR], error)
		// SetEmployeeState 设置员工状态
		SetEmployeeState(ctx context.Context, id int64, state int) (bool, error)
		// SetCommissionRate 设置员工提成比例
		SetCommissionRate(ctx context.Context, userId int64, commissionRate int, actionUserId int64) (bool, error)
	}
	ITeam[TR co_model.ITeamRes] interface {
		SetXDao(dao co_dao.XDao)
		// GetTeamById 根据ID获取公司团队信息
		GetTeamById(ctx context.Context, id int64) (TR, error)
		// GetTeamByName 根据Name获取团队信息
		GetTeamByName(ctx context.Context, name string) (TR, error)
		// HasTeamByName 团队名称是否存在
		HasTeamByName(ctx context.Context, name string, unionMainId int64, parentId int64, excludeIds ...int64) bool
		// QueryTeamList 查询团队
		QueryTeamList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error)
		// QueryTeamMemberList 查询所有团队成员记录
		QueryTeamMemberList(ctx context.Context, search *base_model.SearchParams, isExport ...bool) (*base_model.CollectRes[*co_model.TeamMemberRes], error)
		// CreateTeam 创建团队或小组|信息
		CreateTeam(ctx context.Context, info *co_model.Team) (TR, error)
		// UpdateTeam 更新团队或小组|信息
		UpdateTeam(ctx context.Context, info *co_model.Team) (TR, error)
		// QueryTeamListByEmployee 根据员工查询团队
		QueryTeamListByEmployee(ctx context.Context, employeeId int64, unionMainId int64) (*base_model.CollectRes[TR], error)
		// SetTeamMember 设置团队队员或小组组员
		SetTeamMember(ctx context.Context, teamId int64, employeeIds []int64) (api_v1.BoolRes, error)
		// RemoveTeamMember  移除团队队员或小组组员
		RemoveTeamMember(ctx context.Context, teamId int64, employeeIds []int64) (api_v1.BoolRes, error)
		// SetTeamOwner 设置团队或小组的负责人
		SetTeamOwner(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error)
		// SetTeamCaptain 设置团队队长或小组组长
		SetTeamCaptain(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error)
		// DeleteTeam 删除团队
		DeleteTeam(ctx context.Context, teamId int64) (api_v1.BoolRes, error)
		// DeleteTeamMemberByEmployee 删除某个员工的所有团队成员记录
		DeleteTeamMemberByEmployee(ctx context.Context, employeeId int64) (bool, error)
		// GetEmployeeListByTeamId 获取团队成员|列表
		GetEmployeeListByTeamId(ctx context.Context, teamId int64) (*base_model.CollectRes[co_model.IEmployeeRes], error)
		// GetTeamInviteCode 获取团队邀约码
		GetTeamInviteCode(ctx context.Context, teamId, userId int64) (*co_model.TeamInviteCodeRes, error)
		// JoinTeamByInviteCode 扫码邀约码进入团队
		JoinTeamByInviteCode(ctx context.Context, inviteCode string, userId int64) (bool, error)
	}

	IMy interface {
		// GetProfile 获取当前员工及用户信息
		GetProfile(ctx context.Context) (*co_model.MyProfileRes, error)
		// GetCompany 获取当前公司信息
		GetCompany(ctx context.Context) (*co_model.MyCompanyRes, error)
		// GetTeams 获取当前员工团队信息
		GetTeams(ctx context.Context) (res co_model.MyTeamListRes, err error)
		// SetMyMobile 设置我的手机号
		SetMyMobile(ctx context.Context, newMobile string, captcha string, password string) (bool, error)
		// SetMyMail 设置我的邮箱
		SetMyMail(ctx context.Context, oldMail string, newMail string, captcha string, password string) (bool, error)
		// SetMyAvatar 设置我的头像
		SetMyAvatar(ctx context.Context, imageId int64) (bool, error)
		// GetAccountBills 我的账单|列表
		GetAccountBills(ctx context.Context, pagination *base_model.SearchParams) (*co_model.MyAccountBillRes, error)
		// GetAccounts 获取我的财务账号|列表
		GetAccounts(ctx context.Context) (*co_model.FdAccountListRes, error)
		// GetBankCards 获取我的银行卡｜列表
		GetBankCards(ctx context.Context) (*co_model.FdBankCardListRes, error)
		// GetInvoices 获取我的发票抬头|列表
		GetInvoices(ctx context.Context) (*co_model.FdInvoiceListRes, error)
		// UpdateAccount  修改我的财务账号
		UpdateAccount(ctx context.Context, accountId int64, info *co_model.UpdateAccount) (api_v1.BoolRes, error)
		// GetMyCompanyPermissionList 获取我的公司权限列表
		GetMyCompanyPermissionList(ctx context.Context, permissionType *int) (*sys_model.MyPermissionListRes, error)
	}

	IFdAccount[TR co_model.IFdAccountRes] interface {
		// CreateAccount 创建财务账号
		CreateAccount(ctx context.Context, info co_model.FdAccountRegister, userId int64) (response TR, err error)
		// GetAccountById 根据ID获取财务账号
		GetAccountById(ctx context.Context, id int64) (response TR, err error)
		// UpdateAccount 修改财务账号
		UpdateAccount(ctx context.Context, accountId int64, info *co_model.UpdateAccount) (bool, error)
		// UpdateAccountIsEnable 修改财务账号状态（是否启用：0禁用 1启用）
		UpdateAccountIsEnable(ctx context.Context, id int64, isEnabled int, userId int64) (bool, error)
		// HasAccountByName 判断财务账号名是否存在
		HasAccountByName(ctx context.Context, name string) (response TR, err error)
		// UpdateAccountLimitState 修改财务账号的限制状态 （0不限制，1限制支出、2限制收入）
		UpdateAccountLimitState(ctx context.Context, id int64, limitState int, userId int64) (bool, error)
		// SetAccountCurrencyCode 设置财务账号货币单位
		SetAccountCurrencyCode(ctx context.Context, accountId int64, currencyCode string, userId int64) (bool, error)
		// QueryAccountListByUserId 获取指定用户的所有财务账号
		QueryAccountListByUserId(ctx context.Context, userId int64) (*base_model.CollectRes[TR], error)
		// UpdateAccountBalance 修改财务账户余额(上下文, 财务账号id, 需要修改的钱数目, 版本, 收支类型)
		UpdateAccountBalance(ctx context.Context, accountId int64, amount int64, version int, inOutType co_enum.FinanceInOutType, sysSessionUserId int64) (int64, error)
		// GetAccountByUnionUserIdAndCurrencyCode 根据用户union_user_id和货币代码currency_code获取财务账号
		GetAccountByUnionUserIdAndCurrencyCode(ctx context.Context, unionUserId int64, currencyCode string) (response TR, err error)
		// GetAccountByUnionUserIdAndScene 根据union_user_id和业务类型找出财务账号，
		GetAccountByUnionUserIdAndScene(ctx context.Context, unionUserId int64, accountType co_enum.AccountType, sceneType ...co_enum.SceneType) (response TR, err error)
		// GetAccountDetailById 根据财务账号id查询账单金额明细统计记录，如果主体id找不到财务账号的时候就创建财务账号
		GetAccountDetailById(ctx context.Context, id int64) (res *co_model.FdAccountDetailRes, err error)
		// CreateAccountDetail 创建财务账单金额明细统计记录
		CreateAccountDetail(ctx context.Context, info *co_model.FdAccountDetail) (res *co_model.FdAccountDetailRes, err error)
		// Increment 收入
		Increment(ctx context.Context, id int64, amount int) (bool, error)
		// Decrement 支出
		Decrement(ctx context.Context, id int64, amount int) (bool, error)
		// SetAccountAllowExceed 设置财务账号是否允许存在负余额
		SetAccountAllowExceed(ctx context.Context, accountId int64, allowExceed int) (bool, error)
		// QueryDetailByUnionUserIdAndSceneType  获取用户指定业务场景的财务账号金额明细统计记录|列表
		QueryDetailByUnionUserIdAndSceneType(ctx context.Context, unionUserId int64, sceneType co_enum.SceneType) (*base_model.CollectRes[co_model.FdAccountDetailRes], error)
	}
	IFdBankCard[TR co_model.IFdBankCardRes] interface {
		// CreateBankCard 添加银行卡账号
		CreateBankCard(ctx context.Context, info co_model.BankCardRegister, createUser *sys_model.SysUser) (response TR, err error)
		// GetBankCardById 根据银行卡id获取银行卡信息
		GetBankCardById(ctx context.Context, id int64) (response TR, err error)
		// GetBankCardByCardNumber 根据银行卡号获取银行卡
		GetBankCardByCardNumber(ctx context.Context, cardNumber string) (response TR, err error)
		// UpdateBankCardState 修改银行卡状态 (0禁用 1正常)
		UpdateBankCardState(ctx context.Context, bankCardId int64, state int) (bool, error)
		// DeleteBankCardById 删除银行卡 (标记删除: 标记删除的银行卡号，将记录ID的后6位附加到卡号尾部，用下划线隔开,并修改状态)
		DeleteBankCardById(ctx context.Context, bankCardId int64) (bool, error)
		// QueryBankCardListByUserId 根据用户id查询银行卡列表
		QueryBankCardListByUserId(ctx context.Context, userId int64) (*base_model.CollectRes[TR], error)
	}
	IFdCurrency[TR co_model.IFdCurrencyRes] interface {
		// QueryCurrencyList 获取币种列表
		QueryCurrencyList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error)
		// GetCurrencyByCode 根据货币代码查找货币(主键)
		GetCurrencyByCode(ctx context.Context, currencyCode string) (response TR, err error)
		// GetCurrencyByCnName 根据国家查找货币信息
		GetCurrencyByCnName(ctx context.Context, cnName string) (response TR, err error)
	}
	IFdInvoice[TR co_model.IFdInvoiceRes] interface {
		// CreateInvoice 添加发票抬头
		CreateInvoice(ctx context.Context, info co_model.FdInvoiceRegister) (response TR, err error)
		// GetInvoiceById 根据id获取发票
		GetInvoiceById(ctx context.Context, id int64) (response TR, err error)
		// QueryInvoiceList 获取发票抬头列表
		QueryInvoiceList(ctx context.Context, info *base_model.SearchParams, userId int64) (*base_model.CollectRes[TR], error)
		// DeletesFdInvoiceById 删除发票抬头
		DeletesFdInvoiceById(ctx context.Context, invoiceId int64) (bool, error)
		// GetFdInvoiceByTaxId 根据纳税识别号获取发票抬头信息
		GetFdInvoiceByTaxId(ctx context.Context, taxId string) (response TR, err error)
	}
	IFdInvoiceDetail[TR co_model.IFdInvoiceDetailRes] interface {
		// CreateInvoiceDetail 创建发票详情，相当于创建审核列表，审核是人工审核
		CreateInvoiceDetail(ctx context.Context, info co_model.FdInvoiceDetailRegister) (response TR, err error)
		// GetInvoiceDetailById 根据id获取发票详情
		GetInvoiceDetailById(ctx context.Context, id int64) (response TR, err error)
		// MakeInvoiceDetail 开票
		MakeInvoiceDetail(ctx context.Context, invoiceDetailId int64, makeInvoiceDetail co_model.FdMakeInvoiceDetail) (res bool, err error)
		// AuditInvoiceDetail 审核发票
		AuditInvoiceDetail(ctx context.Context, invoiceDetailId int64, auditInfo co_model.FdInvoiceAuditInfo) (bool, error)
		// QueryInvoiceDetailListByInvoiceId 根据发票抬头，获取已开票的发票详情列表
		QueryInvoiceDetailListByInvoiceId(ctx context.Context, invoiceId int64) (*base_model.CollectRes[TR], error)
		// DeleteInvoiceDetail 标记删除发票详情
		DeleteInvoiceDetail(ctx context.Context, id int64) (bool, error)
		// QueryInvoiceDetail 根据限定的条件查询发票列表
		QueryInvoiceDetail(ctx context.Context, info *base_model.SearchParams, userId int64, unionMainId int64) (*base_model.CollectRes[TR], error)
	}
	IFdAccountBills[TR co_model.IFdAccountBillsRes] interface {
		// InstallTradeHook 订阅Hook
		InstallTradeHook(hookKey co_hook.AccountBillHookKey, hookFunc co_hook.AccountBillHookFunc)
		// GetTradeHook 获取Hook
		GetTradeHook() base_hook.BaseHook[co_hook.AccountBillHookKey, co_hook.AccountBillHookFunc]
		// CreateAccountBills 创建财务账单
		CreateAccountBills(ctx context.Context, info co_model.AccountBillsRegister) (bool, error)
		// GetAccountBillsByAccountId  根据财务账号id获取账单
		GetAccountBillsByAccountId(ctx context.Context, accountId int64, pagination *base_model.SearchParams) (*base_model.CollectRes[TR], error)
	}

	IFdRecharge[TR co_model.IFdRechargeRes] interface {
		// AccountRecharge 充值
		AccountRecharge(ctx context.Context, info *co_model.FdRecharge, createUser *sys_model.SysUser) (TR, error)
		// SetAccountRechargeAudit 设置充值记录审核
		SetAccountRechargeAudit(ctx context.Context, id int64, state sys_enum.AuditAction, reply string) (bool, error)
		// GetAccountRechargeById 根据充值记录id获取充值记录
		GetAccountRechargeById(ctx context.Context, id int64) (TR, error)
		// QueryAccountRecharge 获取财务账号充值记录|列表
		QueryAccountRecharge(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error)
	}
)

type IConfig interface {
	GetConfig() *co_model.Config
}

//type ModuleFactory[
//	ITCompanyRes co_model.ICompanyRes,
//	ITEmployeeRes co_model.IEmployeeRes,
//	ITTeamRes co_model.ITeamRes,
//	ITFdAccountRes co_model.IFdAccountRes,
//	ITFdAccountBillsRes co_model.IFdAccountBillsRes,
//	ITFdBankCardRes co_model.IFdBankCardRes,
//	ITFdCurrencyRes co_model.IFdCurrencyRes,
//	ITFdInvoiceRes co_model.IFdInvoiceRes,
//	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
//] struct {
//	NewEmployee func(modules IModules[
//		ITCompanyRes,
//		ITEmployeeRes,
//		ITTeamRes,
//		ITFdAccountRes,
//		ITFdAccountBillsRes,
//		ITFdBankCardRes,
//		ITFdCurrencyRes,
//		ITFdInvoiceRes,
//		ITFdInvoiceDetailRes,
//	]) IEmployee[ITEmployeeRes]
//
//	NewTeam func(modules IModules[
//		ITCompanyRes,
//		ITEmployeeRes,
//		ITTeamRes,
//		ITFdAccountRes,
//		ITFdAccountBillsRes,
//		ITFdBankCardRes,
//		ITFdCurrencyRes,
//		ITFdInvoiceRes,
//		ITFdInvoiceDetailRes,
//	]) ITeam[ITTeamRes, ITEmployeeRes]
//}

type IModuleBase interface {
	IConfig
	My() IMy
	SetI18n(i18n *gi18n.Manager) error
	T(ctx context.Context, content string) string
	Tf(ctx context.Context, format string, values ...interface{}) string
	Dao() *co_dao.XDao
}

type IModules[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
] interface {
	IModuleBase
	Company() ICompany[ITCompanyRes]
	Team() ITeam[ITTeamRes]
	Employee() IEmployee[ITEmployeeRes]
	Account() IFdAccount[ITFdAccountRes]
	AccountBills() IFdAccountBills[ITFdAccountBillRes]
	BankCard() IFdBankCard[ITFdBankCardRes]
	Invoice() IFdInvoice[ITFdInvoiceRes]
	InvoiceDetail() IFdInvoiceDetail[ITFdInvoiceDetailRes]
	Recharge() IFdRecharge[ITFdRechargeRes]
}
