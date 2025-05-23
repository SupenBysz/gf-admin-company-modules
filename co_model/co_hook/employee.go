package co_hook

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

type EmployeeHookFilter struct {
	InOutType     co_enum.FinanceInOutType
	TradeType     co_enum.FinanceTradeType
	InTransaction bool
}

type EmployeeHookFunc sys_model.HookFunc[EmployeeHookFilter, co_model.EmployeeRes]
