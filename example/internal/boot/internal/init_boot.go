package internal

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/example/internal/consts"
	"github.com/SupenBysz/gf-admin-company-modules/internal/boot"
)

func init() {
	// 注册自定义参数校验规则
	boot.InitCustomRules()

	// 设置国际化语言
	_ = consts.Global.SetI18n(nil)

	// 初始化权限树
	consts.Global.PermissionTree = append(boot.InitPermission(consts.Global.IModules.(co_interface.IModules[
		*co_model.CompanyRes,
		*co_model.EmployeeRes,
		*co_model.TeamRes,
		*co_model.FdAccountRes,
		*co_model.FdAccountBillsRes,
		*co_model.FdBankCardRes,
		*co_model.FdInvoiceRes,
		*co_model.FdInvoiceDetailRes,
		*co_model.FdRechargeRes,
	])), boot.InitAuditAndLicensePermission()...)

	// 导入财务服务权限树
	consts.Global.FinancePermissionTree = boot.InitFinancePermission(consts.Global.IModules.(co_interface.IModules[
		*co_model.CompanyRes,
		*co_model.EmployeeRes,
		*co_model.TeamRes,
		*co_model.FdAccountRes,
		*co_model.FdAccountBillsRes,
		*co_model.FdBankCardRes,
		*co_model.FdInvoiceRes,
		*co_model.FdInvoiceDetailRes,
		*co_model.FdRechargeRes,
	]))
}
