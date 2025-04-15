package internal

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_module"
	"github.com/kysion/base-library/utility/base_permission"
)

type Global struct {
	co_interface.IModules[
		*co_model.CompanyRes,
		*co_model.EmployeeRes,
		*co_model.TeamRes,
		*co_model.FdAccountRes,
		*co_model.FdAccountBillsRes,
		*co_model.FdBankCardRes,
		*co_model.FdCurrencyRes,
		*co_model.FdInvoiceRes,
		*co_model.FdInvoiceDetailRes,
	]

	PermissionTree []base_permission.IPermission

	// FinancePermissionTree 财务服务权限树 (可选)
	FinancePermissionTree []base_permission.IPermission
}

var global *Global

func Modules() *Global {
	if global != nil {
		return global
	}

	global = &Global{
		IModules: co_module.NewModules[
			*co_model.CompanyRes,
			*co_model.EmployeeRes,
			*co_model.TeamRes,
			*co_model.FdAccountRes,
			*co_model.FdAccountBillsRes,
			*co_model.FdBankCardRes,
			*co_model.FdCurrencyRes,
			*co_model.FdInvoiceRes,
			*co_model.FdInvoiceDetailRes,
		](
			&co_model.Config{
				AllowEmptyNo:                   true,
				IsCreateDefaultEmployeeAndRole: false,
				HardDeleteWaitAt:               0,
				KeyIndex:                       "Company",
				I18nName:                       "company",
				RoutePrefix:                    "/company",
				StoragePath:                    "resource/company",
				UserType:                       sys_enum.User.Type.SuperAdmin, // 业务层用户类型需自定义
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
				FdAccountBill:   co_dao.NewFdAccountBills(),
				FdInvoice:       co_dao.NewFdInvoice(),
				FdInvoiceDetail: co_dao.NewFdInvoiceDetail(),
				FdCurrency:      co_dao.NewFdCurrency(),
				FdBankCard:      co_dao.NewFdBankCard(),
			},
		),
	}

	return global
}
