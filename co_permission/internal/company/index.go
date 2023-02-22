package company

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
)

type company struct {
	PermissionType func(modules co_interface.IConfig) *permissionType[co_interface.IConfig]
}

var Company = company{
	PermissionType: PermissionType,
}
