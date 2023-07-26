package liceense

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail PermissionTypeEnum
	List       PermissionTypeEnum
	Update     PermissionTypeEnum
	SetState   PermissionTypeEnum
	Create     PermissionTypeEnum
}

var (
	PermissionType = permissionType{
		ViewDetail: base_permission.New(5953153121845328, "ViewDetail", "查看主体信息", "查看某条主体信息"),
		List:       base_permission.New(5953153121845329, "List", "主体列表", "查看所有主体信息"),
		Update:     base_permission.New(5953153121845330, "Update", "更新资质审核信息", "更新某条资质审核信息"),
		SetState:   base_permission.New(5953153121845331, "SetState", "设置主体状态", "设置某主体认证状态"),
		Create:     base_permission.New(5953153121845332, "Create", "创建主体", "创建主体信息"),
	}
	permissionTypeMap = gmap.NewStrAnyMapFrom(gconv.Map(PermissionType))
)
