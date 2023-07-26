package company

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/base-library/utility/kmap"

	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
)

type Permission = base_permission.IPermission

type permissionType[T co_interface.IConfig] struct {
	modules       T
	enumMap       *kmap.HashMap[string, Permission]
	ViewDetail    Permission
	ViewMobile    Permission
	Create        Permission
	Update        Permission
	List          Permission
	SetLogo       Permission
	SetState      Permission
	SetAdminUser  Permission
	ViewLicense   Permission
	AuditLicense  Permission
	UpdateLicense Permission
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IConfig]]()
	PermissionType    = func(modules co_interface.IConfig) *permissionType[co_interface.IConfig] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IConfig]{
			modules:       modules,
			enumMap:       kmap.New[string, Permission](),
			ViewDetail:    base_permission.NewInIdentifier("ViewDetail", "查看明细", ""),
			ViewMobile:    base_permission.NewInIdentifier("ViewMobile", "查看手机号", ""),
			Create:        base_permission.NewInIdentifier("Create", "新增", ""),
			Update:        base_permission.NewInIdentifier("Update", "更新", ""),
			List:          base_permission.NewInIdentifier("List", "列表", ""),
			SetLogo:       base_permission.NewInIdentifier("SetLogo", "设置LOGO", ""),
			SetState:      base_permission.NewInIdentifier("SetState", "设置状态", ""),
			SetAdminUser:  base_permission.NewInIdentifier("SetAdminUser", "设置管理员", ""),
			ViewLicense:   base_permission.NewInIdentifier("ViewLicense", "查看认证信息", "查看公司认证信息"),
			AuditLicense:  base_permission.NewInIdentifier("AuditLicense", "审核认证信息", "审核公司认证信息"),
			UpdateLicense: base_permission.NewInIdentifier("UpdateLicense", "更新认证信息", "更新公司认证信息"),
		})

		for k, v := range gconv.Map(result) {
			result.enumMap.Set(k, v.(Permission))
		}
		return result
	}
)

// ByCode 通过枚举值取枚举类型
func (e *permissionType[T]) ByCode(identifier string) base_permission.IPermission {
	v, has := e.enumMap.Search(identifier)
	if v != nil && has {
		return v
	}
	return nil
}
