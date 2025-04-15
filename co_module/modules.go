package co_module

import (
	"context"

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
	TFdCurrencyRes co_model.IFdCurrencyRes,
	TFdInvoiceRes co_model.IFdInvoiceRes,
	TFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	co_interface.IConfig
	conf          *co_model.Config
	company       co_interface.ICompany[TCompanyRes]
	employee      co_interface.IEmployee[TEmployeeRes]
	team          co_interface.ITeam[TTeamRes]
	my            co_interface.IMy
	account       co_interface.IFdAccount[TFdAccountRes]
	accountBills  co_interface.IFdAccountBills[TFdAccountBillsRes]
	bankCard      co_interface.IFdBankCard[TFdBankCardRes]
	currency      co_interface.IFdCurrency[TFdCurrencyRes]
	invoice       co_interface.IFdInvoice[TFdInvoiceRes]
	invoiceDetail co_interface.IFdInvoiceDetail[TFdInvoiceDetailRes]

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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
]) Currency() co_interface.IFdCurrency[TFdCurrencyRes] {
	return m.currency
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillsRes,
	TFdBankCardRes,
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
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
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
]) Dao() *co_dao.XDao {
	return m.xDao
}

func NewModules[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillsRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
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
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) {
	module := &Modules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillsRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
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
	module.currency = finance.NewFdCurrency(response)
	module.invoice = finance.NewFdInvoice(response)
	module.invoiceDetail = finance.NewFdInvoiceDetail(response)

	// 权限树追加权限
	co_consts.PermissionTree = append(co_consts.PermissionTree, boot.InitPermission(response)...)
	co_consts.FinancePermissionTree = append(co_consts.FinancePermissionTree, boot.InitFinancePermission(response)...)

	return module
}
