package co_enum

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/company"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/employee"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/financial"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/invoice"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/team"
)

type (
	EmployeeState = employee.StateEnum

	FinancialInOutType = financial.InOutTypeEnum
	FinancialTradeType = financial.TradeTypeEnum

	InvoiceAuditType = invoice.AuditTypeEnum
	InvoiceState     = invoice.StateEnum
	InvoiceType      = invoice.TypeEnum
)

var (
	Company   = company.Company
	Employee  = employee.Employee
	Team      = team.Team
	Financial = financial.Financial
	Invoice   = invoice.Invoice
)
