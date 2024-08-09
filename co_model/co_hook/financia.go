package co_hook

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

type AccountBillHookKey struct {
	InOutType     co_enum.FinancialInOutType // 收支类型
	TradeType     co_enum.FinancialTradeType // 交易类型
	InTransaction bool                       // 是否在事务中
}

type AccountBillHookFunc sys_model.HookFunc[AccountBillHookKey, co_model.IFdAccountBillRes]
