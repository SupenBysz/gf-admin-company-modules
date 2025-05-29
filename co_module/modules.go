package co_module

import (
	"context"
	"github.com/kysion/base-library/utility/base_funs"

	"github.com/gogf/gf/v2/os/gfile"

	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/internal/boot"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/company"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/finance"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type Modules[
	TCompanyRes co_model.ICompanyRes,
	TEmployeeRes co_model.IEmployeeRes,
	TTeamRes co_model.ITeamRes,
	TFdAccountRes co_model.IFdAccountRes,
	TFdAccountBillsRes co_model.IFdAccountBillsRes,
	TFdBankCardRes co_model.IFdBankCardRes,
	TFdInvoiceRes co_model.IFdInvoiceRes,
	TFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
] struct {
	co_interface.IModuleBase
	co_interface.IModules[
		co_model.ICompanyRes,
		co_model.IEmployeeRes,
		co_model.ITeamRes,
		co_model.IFdAccountRes,
		co_model.IFdAccountBillsRes,
		co_model.IFdBankCardRes,
		co_model.IFdInvoiceRes,
		co_model.IFdInvoiceDetailRes,
		co_model.IFdRechargeRes,
	]

	co_interface.IConfig
	conf          *co_model.Config
	company       co_interface.ICompany[TCompanyRes]
	employee      co_interface.IEmployee[TEmployeeRes]
	team          co_interface.ITeam[TTeamRes]
	my            co_interface.IMy
	account       co_interface.IFdAccount[TFdAccountRes]
	accountBills  co_interface.IFdAccountBills[TFdAccountBillsRes]
	bankCard      co_interface.IFdBankCard[TFdBankCardRes]
	invoice       co_interface.IFdInvoice[TFdInvoiceRes]
	invoiceDetail co_interface.IFdInvoiceDetail[TFdInvoiceDetailRes]
	rechargeRes   co_interface.IFdRecharge[ITFdRechargeRes]

	i18n *gi18n.Manager
	xDao *co_dao.XDao
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) My() co_interface.IMy {
	return m.my
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Company() co_interface.ICompany[TCompanyRes] {
	return m.company
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Employee() co_interface.IEmployee[TEmployeeRes] {
	return m.employee
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Team() co_interface.ITeam[TTeamRes] {
	return m.team
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Account() co_interface.IFdAccount[TFdAccountRes] {
	return m.account
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) AccountBills() co_interface.IFdAccountBills[TFdAccountBillsRes] {
	return m.accountBills
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) BankCard() co_interface.IFdBankCard[TFdBankCardRes] {
	return m.bankCard
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Invoice() co_interface.IFdInvoice[TFdInvoiceRes] {
	return m.invoice
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) InvoiceDetail() co_interface.IFdInvoiceDetail[TFdInvoiceDetailRes] {
	return m.invoiceDetail
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Recharge() co_interface.IFdRecharge[ITFdRechargeRes] {
	return m.rechargeRes
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetConfig() *co_model.Config {
	return m.conf
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) T(ctx context.Context, content string) string {
	data := m.i18n.Translate(gi18n.WithLanguage(context.TODO(), "zh-CN"), content)

	return data
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.i18n.TranslateFormat(ctx, format, values...)
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetI18n(i18n *gi18n.Manager) error {
	if i18n == nil {
		i18n = gi18n.New()

		if gfile.IsDir("./i18n/" + m.conf.I18nName) {
			err := i18n.SetPath("./i18n/" + m.conf.I18nName)
			if err != nil {
				return err
			}
		}
	}

	m.i18n = i18n
	return nil
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Dao() *co_dao.XDao {
	return m.xDao
}

var modules []*Modules[
	co_model.ICompanyRes,
	co_model.IEmployeeRes,
	co_model.ITeamRes,
	co_model.IFdAccountRes,
	co_model.IFdAccountBillsRes,
	co_model.IFdBankCardRes,
	co_model.IFdInvoiceRes,
	co_model.IFdInvoiceDetailRes,
	co_model.IFdRechargeRes,
]

func GetModules(predicate func(conf co_model.Config) bool) *Modules[
	co_model.ICompanyRes,
	co_model.IEmployeeRes,
	co_model.ITeamRes,
	co_model.IFdAccountRes,
	co_model.IFdAccountBillsRes,
	co_model.IFdBankCardRes,
	co_model.IFdInvoiceRes,
	co_model.IFdInvoiceDetailRes,
	co_model.IFdRechargeRes] {
	_, oldModule, _ := base_funs.FindInSlice(modules, func(item *Modules[
		co_model.ICompanyRes,
		co_model.IEmployeeRes,
		co_model.ITeamRes,
		co_model.IFdAccountRes,
		co_model.IFdAccountBillsRes,
		co_model.IFdBankCardRes,
		co_model.IFdInvoiceRes,
		co_model.IFdInvoiceDetailRes,
		co_model.IFdRechargeRes,
	]) bool {
		return predicate(*item.GetConfig())
	})

	return oldModule
}

func NewModules[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillsRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](
	conf *co_model.Config,
	xDao *co_dao.XDao,
) (response co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) {
	module := &Modules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillsRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]{
		conf: conf,
		xDao: xDao,
	}

	co_model.ModulesConfigArr = append(co_model.ModulesConfigArr, conf)

	response = module

	// 初始化默认多语言对象
	_ = module.SetI18n(nil)

	module.company = company.NewCompany(response)
	module.employee = company.NewEmployee(response)
	module.team = company.NewTeam(response)
	module.my = company.NewMy(response)
	module.account = finance.NewFdAccount(response)
	module.accountBills = finance.NewFdAccountBills(response)
	module.bankCard = finance.NewFdBankCard(response)
	module.invoice = finance.NewFdInvoice(response)
	module.invoiceDetail = finance.NewFdInvoiceDetail(response)
	module.rechargeRes = finance.NewFdRecharge(response)

	// 权限树追加权限
	co_consts.PermissionTree = append(co_consts.PermissionTree, boot.InitPermission(response)...)
	co_consts.FinancePermissionTree = append(co_consts.FinancePermissionTree, boot.InitFinancePermission(response)...)

	{
		//_, oldModule, _ := base_funs.FindInSlice(modules, func(item *Modules[
		//	co_model.ICompanyRes,
		//	co_model.IEmployeeRes,
		//	co_model.ITeamRes,
		//	co_model.IFdAccountRes,
		//	co_model.IFdAccountBillsRes,
		//	co_model.IFdBankCardRes,
		//	co_model.IFdInvoiceRes,
		//	co_model.IFdInvoiceDetailRes,
		//	co_model.IFdRechargeRes,
		//]) bool {
		//	return item.conf.Identifier != module.conf.Identifier
		//})

		//if oldModule == nil {
		//
		//	var iModule interface{} = module
		//
		//	if module.IModules == nil {
		//		module.IModules = iModule.(*Modules[
		//			co_model.ICompanyRes,
		//			co_model.IEmployeeRes,
		//			co_model.ITeamRes,
		//			co_model.IFdAccountRes,
		//			co_model.IFdAccountBillsRes,
		//			co_model.IFdBankCardRes,
		//			co_model.IFdInvoiceRes,
		//			co_model.IFdInvoiceDetailRes,
		//			co_model.IFdRechargeRes,
		//		])
		//	}
		//
		//	modules = append(modules, module.IModules)
		//}
	}

	return module
}
