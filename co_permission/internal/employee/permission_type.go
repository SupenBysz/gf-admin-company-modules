package employee

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/kmap"
)

type Permission = *sys_model.SysPermissionTree

type permissionType[T co_interface.IModules] struct {
	modules       T
	enumMap       *kmap.HashMap[string, Permission]
	ViewDetail    Permission
	MoreDetail    Permission
	List          Permission
	Create        Permission
	Update        Permission
	Delete        Permission
	SetMobile     Permission
	SetAvatar     Permission
	SetState      Permission
	ViewLicense   Permission
	AuditLicense  Permission
	UpdateLicense Permission
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IModules]]()
	PermissionType    = func(modules co_interface.IModules) *permissionType[co_interface.IModules] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IModules]{
			modules:       modules,
			enumMap:       kmap.New[string, Permission](),
			ViewDetail:    permission.NewInIdentifier("ViewDetail", "详情", "查看员工详情"),
			MoreDetail:    permission.NewInIdentifier("MoreDetail", "更多详情", "查看员工更多详情含手机号等"),
			List:          permission.NewInIdentifier("List", "列表", "查看员工列表"),
			Create:        permission.NewInIdentifier("Create", "新增", "新增员工信息"),
			Update:        permission.NewInIdentifier("Update", "更新", "更新员工信息"),
			Delete:        permission.NewInIdentifier("Delete", "删除", "删除员工信息"),
			SetMobile:     permission.NewInIdentifier("SetMobile", "设置手机号", "修改员工手机号"),
			SetAvatar:     permission.NewInIdentifier("SetAvatar", "设置头像", "设置员工头像"),
			SetState:      permission.NewInIdentifier("SetState", "设置状态", "设置员工任职状态"),
			ViewLicense:   permission.NewInIdentifier("ViewLicense", "查看认证信息", "查看员工认证信息"),
			AuditLicense:  permission.NewInIdentifier("AuditLicense", "审核认证信息", "审核员工认证信息"),
			UpdateLicense: permission.NewInIdentifier("UpdateLicense", "更新认证信息", "更新员工认证信息"),
		})

		for k, v := range gconv.Map(result) {
			result.enumMap.Set(k, v.(Permission))
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
