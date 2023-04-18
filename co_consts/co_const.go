package co_consts

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type global struct {
	// 默认货币类型
	DefaultCurrency string
}

var (
	Global = global{}

	PermissionTree []*sys_model.SysPermissionTree

	FinancialPermissionTree []*sys_model.SysPermissionTree
)

func init() {
	defaultCurrency, _ := g.Cfg().Get(context.Background(), "service.defaultCurrency")
	Global.DefaultCurrency = defaultCurrency.String()
}
