package co_enum

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/common"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/company"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/employee"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/finance"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/invoice"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/license"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/team"
)

type (
	// AppealState 申诉状态
	AppealState = common.AppealStateEnum
	// CompanyCommissionMode 公司佣金模式
	CompanyCommissionMode = common.CompanyCommissionModeEnum
	// EmployeeCommissionMode 员工绩效提成模式
	EmployeeCommissionMode = common.EmployeeCommissionModeEnum
	// EmployeeCommissionLevelMode 员工提成分配等级模式
	EmployeeCommissionLevelMode = common.EmployeeCommissionLevelModeEnum
	// Audit

	// CompanyState 主体状态
	CompanyState = company.StateEnum
	// EmployeeState 员工状态
	EmployeeState = employee.StateEnum

	// Sex 性别
	Sex = employee.SexEnum

	// FinanceInOutType 财务收/支 类型
	FinanceInOutType = finance.InOutTypeEnum
	// FinanceTradeType 交易类型
	FinanceTradeType = finance.TradeTypeEnum
	// FinanceTradeState 交易状态
	FinanceTradeState = finance.TradeStateEnum
	// AccountType 财务账号类型
	AccountType = finance.AccountTypeEnum
	// PaymentState 支付状态
	PaymentState = finance.PaymentStateEnum
	// SpecialPaymentState 特殊场景支付状态或账单状态
	SpecialPaymentState = finance.SpecialPaymentStateEnum

	// SceneType 场景类型
	SceneType = finance.SceneTypeEnum
	// AllowExceed 是否允许存在负数余额
	AllowExceed = finance.AllowExceedEnum
	// RechargeMethod 充值方式
	RechargeMethod = finance.RechargeMethodEnum
	// RechargeState 充值状态
	RechargeState = finance.RechargeStateEnum

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
	// Common 通用
	Common = common.Common
	// Company 组织主体
	Company = company.Company
	// Employee 员工
	Employee = employee.Employee
	// Team 团队
	Team = team.Team

	// Finance 财务
	Finance = finance.Finance

	// Invoice 发票
	Invoice = invoice.Invoice

	// License 主体资质
	License = license.License
)
