package employee

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
)

type employee struct {
	PermissionType func(modules co_interface.IConfig) *permissionType[co_interface.IConfig]
}

var Employee = employee{
	PermissionType: PermissionType,
}
