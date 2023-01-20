package i_controller

import "github.com/SupenBysz/gf-admin-company-modules/co_interface"

type iModule interface {
	GetModules() co_interface.IModules
}
