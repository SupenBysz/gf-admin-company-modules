package team

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
)

type team struct {
	PermissionType func(modules co_interface.IConfig) *permissionType[co_interface.IConfig]
}

var Team = team{
	PermissionType: PermissionType,
}
