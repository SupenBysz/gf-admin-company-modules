package employee

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
)

type employee struct {
	PermissionType func(modules co_interface.IModules) *permissionType[co_interface.IModules]
}

var Employee = employee{
	PermissionType: PermissionType,
}
