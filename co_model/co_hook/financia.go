package co_hook

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/fd_enum"
)

type AccountBillHookFilter struct {
	InOutType     fd_enum.FinancialInOutType
	TradeType     fd_enum.FinancialTradeType
	InTransaction bool
}

type AccountBillHookFunc sys_model.HookFunc[AccountBillHookFilter, co_model.AccountBillInfo]
