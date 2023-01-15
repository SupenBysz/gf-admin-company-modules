package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/example/internal/consts"
)

func init() {
	consts.Global.Company.SetI18n(nil)
	consts.PermissionTree = initPermission(consts.Global.Company)
}

// 初始化权限树
func initPermission(module co_interface.IModules) []*permission.SysPermissionTree {
	result := []*permission.SysPermissionTree{
		// 公司
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5947986066667973,
				Name:       module.T(context.TODO(), "{#CompanyName}"),
				Identifier: module.GetConfig().Identifier.Company,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Company.PermissionType(module).Create,
				co_enum.Company.PermissionType(module).ViewDetail,
				co_enum.Company.PermissionType(module).List,
				co_enum.Company.PermissionType(module).Update,
				co_enum.Company.PermissionType(module).SetLogo,
				co_enum.Company.PermissionType(module).SetState,
				co_enum.Company.PermissionType(module).SetAdminUser,
				co_enum.Company.PermissionType(module).ViewLicense,
				co_enum.Company.PermissionType(module).AuditLicense,
			},
		},
		// 员工
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       module.T(context.TODO(), "{#CompanyName}{#EmployeeName}"),
				Identifier: module.GetConfig().Identifier.Employee,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Employee.PermissionType(module).ViewDetail,
				co_enum.Employee.PermissionType(module).MoreDetail,
				co_enum.Employee.PermissionType(module).List,
				co_enum.Employee.PermissionType(module).Create,
				co_enum.Employee.PermissionType(module).Update,
				co_enum.Employee.PermissionType(module).Delete,
				co_enum.Employee.PermissionType(module).SetMobile,
				co_enum.Employee.PermissionType(module).SetAvatar,
				co_enum.Employee.PermissionType(module).SetState,
				co_enum.Employee.PermissionType(module).ViewLicense,
				co_enum.Employee.PermissionType(module).AuditLicense,
				co_enum.Employee.PermissionType(module).UpdateLicense,
			},
		},
		// 团队
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       module.T(context.TODO(), "{#CompanyName}{#TeamName}"),
				Identifier: module.GetConfig().Identifier.Team,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Team.PermissionType(module).Create,
				co_enum.Team.PermissionType(module).ViewDetail,
				co_enum.Team.PermissionType(module).List,
				co_enum.Team.PermissionType(module).Update,
				co_enum.Team.PermissionType(module).Delete,
				co_enum.Team.PermissionType(module).MemberDetail,
				co_enum.Team.PermissionType(module).SetMember,
				co_enum.Team.PermissionType(module).SetOwner,
				co_enum.Team.PermissionType(module).SetCaptain,
			}},
	}
	return result
}
