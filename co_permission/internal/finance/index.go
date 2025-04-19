package finance

import "github.com/SupenBysz/gf-admin-company-modules/co_interface"

type finance struct {
	PermissionType func(modules co_interface.IConfig) *permissionType[co_interface.IConfig]
}

var Finance = finance{
	PermissionType: PermissionType,
}
