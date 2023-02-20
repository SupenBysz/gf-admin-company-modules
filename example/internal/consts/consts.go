package consts

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_module"
)

type global struct {
	Modules co_interface.IModules
}

var (
	PermissionTree []*sys_model.SysPermissionTree

	// FinancialPermissionTree 财务服务权限树 (可选)
	FinancialPermissionTree []*sys_model.SysPermissionTree

	Global = global{
		Modules: co_module.NewModules(
			&co_model.Config{
				AllowEmptyNo:                   true,
				IsCreateDefaultEmployeeAndRole: false,
				HardDeleteWaitAt:               0,
				KeyIndex:                       "Company",
				RoutePrefix:                    "/company",
				StoragePath:                    "./resources/company",
				UserType:                       sys_enum.User.Type.SuperAdmin,
				Identifier: co_model.Identifier{
					Company:         "company",
					Employee:        "employee",
					Team:            "team",
					FdAccount:       "fdAccount",
					FdAccountBill:   "fdAccountBill",
					FdInvoice:       "fdInvoice",
					FdInvoiceDetail: "fdInvoiceDetail",
					FdBankCard:      "fdBankCard",
				},
			},
			&co_dao.XDao{ // 以下为业务层实例化dao模型，如果不是使用默认模型时需要将自定义dao模型作为参数传进去，相同属性前缀需要配合使用不能拆开应用
				Company:  co_dao.NewCompany(),
				Employee: co_dao.NewCompanyEmployee(),

				Team:       co_dao.NewCompanyTeam(),
				TeamMember: co_dao.NewCompanyTeamMember(),

				FdAccount:       co_dao.NewFdAccount(),
				FdAccountBill:   co_dao.NewFdAccountBill(),
				FdInvoice:       co_dao.NewFdInvoice(),
				FdInvoiceDetail: co_dao.NewFdInvoiceDetail(),
				FdCurrency:      co_dao.NewFdCurrency(),
				FdBankCard:      co_dao.NewFdBankCard(),
			},
		),
	}
)
