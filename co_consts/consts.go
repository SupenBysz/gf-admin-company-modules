package co_consts

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_module"
)

type global struct {
	Company co_interface.IModules
}

var (
	Global = global{
		Company: co_module.NewModules(&co_model.Config{
			I18n:                           nil,
			AllowEmptyNo:                   true,
			IsCreateDefaultEmployeeAndRole: false,
			HardDeleteWaitAt:               0,
			CompanyName:                    "公司",
			KeyIndex:                       "Company",
			RoutePrefix:                    "/company",
			StoragePath:                    "/company",
			UserType:                       sys_enum.User.Type.Operator,
			TableName:                      co_model.TableName{},
			Identifier:                     co_model.Identifier{},
		}),
	}
	// PermissionTree 权限信息定义
	PermissionTree []*permission.SysPermissionTree
)
