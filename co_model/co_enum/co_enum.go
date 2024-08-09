package co_enum

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/company"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/employee"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/financial"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/invoice"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/license"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/team"
)

type (
	// CompanyState 主体状态
	CompanyState = company.StateEnum
	// EmployeeState 员工状态
	EmployeeState = employee.StateEnum

	// Sex 性别
	Sex = employee.SexEnum

	// FinancialInOutType 财务收/支 类型
	FinancialInOutType = financial.InOutTypeEnum
	// FinancialTradeType 交易类型
	FinancialTradeType = financial.TradeTypeEnum
	// AccountType 财务账号类型
	AccountType = financial.AccountTypeEnum
	// SceneType 场景类型
	SceneType = financial.SceneTypeEnum
	// AllowExceed 是否允许存在负数余额
	AllowExceed = financial.AllowExceedEnum

	// InvoiceAuditType 发票审核状态
	InvoiceAuditType = invoice.AuditTypeEnum
	// InvoiceState 发票状态
	InvoiceState = invoice.StateEnum
	// InvoiceType  发票类型
	InvoiceType = invoice.TypeEnum
	// InvoiceBelongType 发票拥有者类型
	InvoiceBelongType = invoice.BelongTypeEnum

	// LicenseState 主体资质状态
	LicenseState = license.StateEnum
	// LicenseAuthType 主体资质认证类型
	LicenseAuthType = license.AuthTypeEnum
)

var (
	// Company 组织主体
	Company = company.Company
	// Employee 员工
	Employee = employee.Employee
	// Team 团队
	Team = team.Team

	// Financial 财务
	Financial = financial.Financial

	// Invoice 发票
	Invoice = invoice.Invoice

	// License 主体资质
	License = license.License
)
