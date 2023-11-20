package liceense

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
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
		ViewDetail: permission.NewInIdentifier("ViewDetail", "查看主体资质信息", "查看某条主体资质信息"),
		List:       permission.NewInIdentifier("List", "主体资质列表", "查看所有主体资质信息"),
		Update:     permission.NewInIdentifier("Update", "更新资质审核信息", "更新某条资质审核信息"),
		SetState:   permission.NewInIdentifier("SetState", "设置主体资质状态", "设置某主体资质认证状态"),
		Create:     permission.NewInIdentifier("Create", "创建主体资质", "创建主体资质信息"),
	}
	permissionTypeMap = gmap.NewStrAnyMapFrom(gconv.Map(PermissionType))
)
