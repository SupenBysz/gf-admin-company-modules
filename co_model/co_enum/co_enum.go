package co_enum

import (
	audit "github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/audit"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/company"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/employee"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/financial"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/invoice"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/team"
)

type (
	CompanyState = company.StateEnum

	EmployeeState = employee.StateEnum
	Sex           = employee.SexEnum

	FinancialInOutType = financial.InOutTypeEnum
	FinancialTradeType = financial.TradeTypeEnum
	AccountType        = financial.AccountTypeEnum
	SceneType          = financial.SceneTypeEnum
	AllowExceed        = financial.AllowExceedEnum

	InvoiceAuditType  = invoice.AuditTypeEnum
	InvoiceState      = invoice.StateEnum
	InvoiceType       = invoice.TypeEnum
	InvoiceBelongType = invoice.BelongTypeEnum

	AuditAction = audit.ActionEnum
	AuditEvent  audit.EventEnum
)

var (
	Company   = company.Company
	Employee  = employee.Employee
	Team      = team.Team
	Financial = financial.Financial
	Invoice   = invoice.Invoice

	Audit = audit.Audit
	//License = license.License
)
