package company

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
)

type company struct {
	PermissionType func(modules co_interface.IModules) *permissionType[co_interface.IModules]
}

var Company = company{
	PermissionType: PermissionType,
}
