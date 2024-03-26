package employee

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/base-library/utility/kmap"
)

type Permission = base_permission.IPermission

type permissionType[T co_interface.IConfig] struct {
	modules       T
	enumMap       *kmap.HashMap[string, Permission]
	ViewDetail    Permission
	MoreDetail    Permission
	ViewMobile    Permission
	List          Permission
	Create        Permission
	Update        Permission
	Delete        Permission
	SetMobile     Permission
	SetAvatar     Permission
	SetRoles      Permission
	SetState      Permission
	ViewLicense   Permission
	AuditLicense  Permission
	UpdateLicense Permission
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IConfig]]()
	PermissionType    = func(modules co_interface.IConfig) *permissionType[co_interface.IConfig] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IConfig]{
			modules: modules,
			enumMap: kmap.New[string, Permission](),
			// ViewDetail1:   base_permission.New("232424342423534","ViewDetail", "详情", "查看员工详情"),
			ViewDetail:    base_permission.NewInIdentifier("ViewDetail", "详情", "查看员工详情"),
			MoreDetail:    base_permission.NewInIdentifier("MoreDetail", "更多详情", "查看员工更多详情含手机号等"),
			ViewMobile:    base_permission.NewInIdentifier("ViewMobile", "查看手机号", ""),
			List:          base_permission.NewInIdentifier("List", "列表", "查看员工列表"),
			Create:        base_permission.NewInIdentifier("Create", "新增", "新增员工信息"),
			Update:        base_permission.NewInIdentifier("Update", "更新", "更新员工信息"),
			Delete:        base_permission.NewInIdentifier("Delete", "删除", "删除员工信息"),
			SetMobile:     base_permission.NewInIdentifier("SetMobile", "设置手机号", "修改员工手机号"),
			SetAvatar:     base_permission.NewInIdentifier("SetAvatar", "设置头像", "设置员工头像"),
			SetRoles:      base_permission.NewInIdentifier("SetRoles", "设置员工角色", "设置员工角色"),
			SetState:      base_permission.NewInIdentifier("SetState", "设置状态", "设置员工任职状态"),
			ViewLicense:   base_permission.NewInIdentifier("ViewLicense", "查看认证信息", "查看员工认证信息"),
			AuditLicense:  base_permission.NewInIdentifier("AuditLicense", "审核认证信息", "审核员工认证信息"),
			UpdateLicense: base_permission.NewInIdentifier("UpdateLicense", "更新认证信息", "更新员工认证信息"),
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
