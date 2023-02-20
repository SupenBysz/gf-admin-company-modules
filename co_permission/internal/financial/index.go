package financial

import "github.com/SupenBysz/gf-admin-company-modules/co_interface"

type financial struct {
	PermissionType func(modules co_interface.IModules) *permissionType[co_interface.IModules]
}

var Financial = financial{
	PermissionType: PermissionType,
}
