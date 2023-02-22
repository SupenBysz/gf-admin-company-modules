package co_module

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/internal/boot"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/company"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/financial"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/text/gstr"
)

type Modules[
	TCompanyRes co_model.ICompanyRes,
	TEmployeeRes co_model.IEmployeeRes,
	TTeamRes co_model.ITeamRes,
	TFdAccountRes co_model.IFdAccountRes,
	TFdAccountBillRes co_model.IFdAccountBillRes,
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
	accountBill   co_interface.IFdAccountBill[TFdAccountBillRes]
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
	TFdBankCardRes,
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
]) AccountBill() co_interface.IFdAccountBill[TFdAccountBillRes] {
	return m.accountBill
}

func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillRes,
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
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
	TFdAccountBillRes,
	TFdBankCardRes,
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
]) T(ctx context.Context, content string) string {
	return m.i18n.Translate(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules[
	TCompanyRes,
	TEmployeeRes,
	TTeamRes,
	TFdAccountRes,
	TFdAccountBillRes,
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
	TFdAccountBillRes,
	TFdBankCardRes,
	TFdCurrencyRes,
	TFdInvoiceRes,
	TFdInvoiceDetailRes,
]) SetI18n(i18n *gi18n.Manager) error {
	if i18n == nil {
		i18n = gi18n.New()
		i18n.SetLanguage("zh-CN")
		err := i18n.SetPath("i18n/" + gstr.ToLower(m.conf.KeyIndex))
		if err != nil {
			return err
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
	TFdAccountBillRes,
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
	ITFdAccountBillRes co_model.IFdAccountBillRes,
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
	ITFdAccountBillRes,
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
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		conf: conf,
		xDao: xDao,
	}

	response = module

	// 初始化默认多语言对象
	module.SetI18n(nil)

	module.company = company.NewCompany(response)
	module.employee = company.NewEmployee(response)
	module.team = company.NewTeam(response)
	module.my = company.NewMy(response)
	module.account = financial.NewFdAccount(response)
	module.accountBill = financial.NewFdAccountBill(response)
	module.bankCard = financial.NewFdBankCard(response)
	module.currency = financial.NewFdCurrency(response)
	module.invoice = financial.NewFdInvoice(response)
	module.invoiceDetail = financial.NewFdInvoiceDetail(response)

	// 权限树追加权限
	co_consts.PermissionTree = append(co_consts.PermissionTree, boot.InitPermission(response)...)

	return module
}
