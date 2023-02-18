package team

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/kmap"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/gogf/gf/v2/util/gconv"
)

type Permission = *permission.SysPermissionTree

type permissionType[T co_interface.IModules] struct {
	modules      T
	enumMap      *kmap.HashMap[string, Permission]
	ViewDetail   Permission
	List         Permission
	ViewMobile   Permission
	Create       Permission
	Update       Permission
	Delete       Permission
	MemberDetail Permission
	SetMember    Permission
	SetOwner     Permission
	SetCaptain   Permission
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IModules]]()
	PermissionType    = func(modules co_interface.IModules) *permissionType[co_interface.IModules] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IModules]{
			modules:      modules,
			enumMap:      kmap.New[string, Permission](),
			ViewDetail:   permission.NewInIdentifier("ViewDetail", "详情", "查看团队详情"),
			List:         permission.NewInIdentifier("List", "列表", "查看团队列表"),
			Create:       permission.NewInIdentifier("Create", "新增", "新增团队信息"),
			Update:       permission.NewInIdentifier("Update", "更新", "更新团队信息"),
			Delete:       permission.NewInIdentifier("Delete", "删除", "删除团队信息"),
			MemberDetail: permission.NewInIdentifier("MemberDetail", "成员详情", "查看团队成员详情"),
			SetMember:    permission.NewInIdentifier("SetMember", "设置成员", "设置团队成员"),
			SetOwner:     permission.NewInIdentifier("SetOwner", "设置管理人", "设置团队管理人，可以不是团队成员"),
			SetCaptain:   permission.NewInIdentifier("SetCaptain", "设置队长或组长", "设置团队队长或小组组长，必须是团队成员"),
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
