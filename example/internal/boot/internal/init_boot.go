package internal

import (
	"github.com/SupenBysz/gf-admin-company-modules/example/internal/consts"
	"github.com/SupenBysz/gf-admin-company-modules/internal/boot"
)

func init() {
	consts.Global.Modules.SetI18n(nil)
	consts.PermissionTree = boot.InitPermission(consts.Global.Modules)

	// 初始化财务服务权限树 (可选)
	// consts.FinancialPermissionTree = boot.InitFinancialPermission()
}
