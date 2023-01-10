package company

import (
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/kmap"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
)

type PermissionEnum = *permission.SysPermissionTree

type permissionType[T co_interface.IModules] struct {
	modules       T
	enumMap       *kmap.HashMap[string, PermissionEnum]
	ViewDetail    PermissionEnum
	ViewMobile    PermissionEnum
	Create        PermissionEnum
	Update        PermissionEnum
	List          PermissionEnum
	SetLogo       PermissionEnum
	SetState      PermissionEnum
	SetAdminUser  PermissionEnum
	ViewLicense   PermissionEnum
	AuditLicense  PermissionEnum
	UpdateLicense PermissionEnum
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IModules]]()
	PermissionType    = func(modules co_interface.IModules) *permissionType[co_interface.IModules] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IModules]{
			modules:       modules,
			enumMap:       kmap.New[string, PermissionEnum](),
			ViewDetail:    permission.NewInIdentifier("ViewDetail", "查看明细", ""),
			Create:        permission.NewInIdentifier("Create", "新增", ""),
			Update:        permission.NewInIdentifier("Update", "更新", ""),
			List:          permission.NewInIdentifier("List", "列表", ""),
			SetLogo:       permission.NewInIdentifier("SetLogo", "设置LOGO", ""),
			SetState:      permission.NewInIdentifier("SetState", "设置状态", ""),
			SetAdminUser:  permission.NewInIdentifier("SetAdminUser", "设置管理员", ""),
			ViewLicense:   permission.NewInIdentifier("ViewLicense", "查看资质认证信息", ""),
			AuditLicense:  permission.NewInIdentifier("AuditLicense", "审核资质认证信息", ""),
			UpdateLicense: permission.NewInIdentifier("UpdateLicense", "更新资质认证信息", ""),
		})

		for k, v := range gconv.Map(result) {
			result.enumMap.Set(k, v.(PermissionEnum))
		}
		return result
	}
)

// ByCode 通过枚举值取枚举类型
func (e *permissionType[T]) ByCode(identifier string) *sys_entity.SysPermission {
	v, has := e.enumMap.Search(identifier)
	if v != nil && has {
		return v.SysPermission
	}
	return nil
}
