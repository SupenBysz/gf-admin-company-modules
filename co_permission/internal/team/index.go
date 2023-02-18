package team

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
)

type team struct {
	PermissionType func(modules co_interface.IModules) *permissionType[co_interface.IModules]
}

var Team = team{
	PermissionType: PermissionType,
}
