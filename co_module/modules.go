package co_module

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/financial"

	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/internal/boot"
	"github.com/SupenBysz/gf-admin-company-modules/internal/logic/company"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/text/gstr"
)

type Modules struct {
	conf          *co_model.Config
	company       co_interface.ICompany
	team          co_interface.ITeam
	employee      co_interface.IEmployee
	my            co_interface.IMy
	account       co_interface.IFdAccount
	accountBill   co_interface.IFdAccountBill
	bankCard      co_interface.IFdBankCard
	currency      co_interface.IFdCurrency
	invoice       co_interface.IFdInvoice
	invoiceDetail co_interface.IFdInvoiceDetail

	i18n *gi18n.Manager
	xDao *co_dao.XDao
}

func (m *Modules) My() co_interface.IMy {
	return m.my
}

func (m *Modules) Company() co_interface.ICompany {
	return m.company
}

func (m *Modules) Team() co_interface.ITeam {
	return m.team
}

func (m *Modules) Employee() co_interface.IEmployee {
	return m.employee
}

func (m *Modules) Account() co_interface.IFdAccount {
	return m.account
}

func (m *Modules) AccountBill() co_interface.IFdAccountBill {
	return m.accountBill
}

func (m *Modules) BankCard() co_interface.IFdBankCard {
	return m.bankCard
}

func (m *Modules) Currency() co_interface.IFdCurrency {
	return m.currency
}

func (m *Modules) Invoice() co_interface.IFdInvoice {
	return m.invoice
}

func (m *Modules) InvoiceDetail() co_interface.IFdInvoiceDetail {
	return m.invoiceDetail
}

func (m *Modules) GetConfig() *co_model.Config {
	return m.conf
}

func (m *Modules) T(ctx context.Context, content string) string {
	return m.i18n.Translate(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.i18n.TranslateFormat(ctx, format, values...)
}

func (m *Modules) SetI18n(i18n *gi18n.Manager) error {
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

func (m *Modules) Dao() *co_dao.XDao {
	return m.xDao
}

func NewModules(
	conf *co_model.Config,
	xDao *co_dao.XDao,
) *Modules {
	module := &Modules{
		conf: conf,
		xDao: xDao,
	}

	// ??????????????????????????????
	module.SetI18n(nil)

	module.company = company.NewCompany(module)
	module.employee = company.NewEmployee(module)
	module.team = company.NewTeam(module)
	module.my = company.NewMy(module)
	module.account = financial.NewFdAccount(module)
	module.accountBill = financial.NewFdAccountBill(module)
	module.bankCard = financial.NewFdBankCard(module)
	module.currency = financial.NewFdCurrency(module)
	module.invoice = financial.NewFdInvoice(module)

	module.invoiceDetail = financial.NewFdInvoiceDetail(module)

	// ?????????????????????
	co_consts.PermissionTree = append(co_consts.PermissionTree, boot.InitPermission(module)...)

	return module
}
