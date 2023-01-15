package consts

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_module"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type global struct {
	Company co_interface.IModules
}

var (
	PermissionTree []*permission.SysPermissionTree
	Global         = global{
		Company: co_module.NewModules(
			&co_model.Config{
				I18n:                           initI18n(nil),
				AllowEmptyNo:                   true,
				IsCreateDefaultEmployeeAndRole: false,
				HardDeleteWaitAt:               0,
				KeyIndex:                       "Company",
				RoutePrefix:                    "/company",
				StoragePath:                    "./resources/company",
				UserType:                       sys_enum.User.Type.SuperAdmin,
				Identifier: co_model.Identifier{
					Company:  "company",
					Employee: "employee",
					Team:     "team",
				},
			},
			co_dao.Company,
			co_dao.CompanyEmployee,
			co_dao.CompanyTeam,
			co_dao.CompanyTeamMember,
		),
	}
)

func initI18n(i18n *gi18n.Manager) *gi18n.Manager {
	if i18n == nil {
		if i18n == nil {
			i18n = gi18n.New()
			i18n.SetLanguage("zh-CN")
			if err := i18n.SetPath("i18n/company"); err != nil {
				panic(err)
			}
		}
	}
	return i18n
}

func init() {
	PermissionTree = initPermission(&Global.Company)
}

// 初始化权限树
func initPermission(conf *co_interface.IModules) []*permission.SysPermissionTree {
	result := []*permission.SysPermissionTree{
		// 公司
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5947986066667973,
				Name:       Global.Company.T(context.TODO(), "{#CompanyName}"),
				Identifier: Global.Company.GetConfig().Identifier.Company,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Company.PermissionType(Global.Company).Create,
				co_enum.Company.PermissionType(Global.Company).ViewDetail,
				co_enum.Company.PermissionType(Global.Company).List,
				co_enum.Company.PermissionType(Global.Company).Update,
				co_enum.Company.PermissionType(Global.Company).SetLogo,
				co_enum.Company.PermissionType(Global.Company).SetState,
				co_enum.Company.PermissionType(Global.Company).SetAdminUser,
				co_enum.Company.PermissionType(Global.Company).ViewLicense,
				co_enum.Company.PermissionType(Global.Company).AuditLicense,
			},
		},
		// 员工
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       Global.Company.T(context.TODO(), "{#CompanyName}{#EmployeeName}"),
				Identifier: Global.Company.GetConfig().Identifier.Employee,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Employee.PermissionType(Global.Company).ViewDetail,
				co_enum.Employee.PermissionType(Global.Company).MoreDetail,
				co_enum.Employee.PermissionType(Global.Company).List,
				co_enum.Employee.PermissionType(Global.Company).Create,
				co_enum.Employee.PermissionType(Global.Company).Update,
				co_enum.Employee.PermissionType(Global.Company).Delete,
				co_enum.Employee.PermissionType(Global.Company).SetMobile,
				co_enum.Employee.PermissionType(Global.Company).SetAvatar,
				co_enum.Employee.PermissionType(Global.Company).SetState,
				co_enum.Employee.PermissionType(Global.Company).ViewLicense,
				co_enum.Employee.PermissionType(Global.Company).AuditLicense,
				co_enum.Employee.PermissionType(Global.Company).UpdateLicense,
			},
		},
		// 团队
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       Global.Company.T(context.TODO(), "{#CompanyName}{#TeamName}"),
				Identifier: Global.Company.GetConfig().Identifier.Team,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Team.PermissionType(Global.Company).Create,
				co_enum.Team.PermissionType(Global.Company).ViewDetail,
				co_enum.Team.PermissionType(Global.Company).List,
				co_enum.Team.PermissionType(Global.Company).Update,
				co_enum.Team.PermissionType(Global.Company).Delete,
				co_enum.Team.PermissionType(Global.Company).MemberDetail,
				co_enum.Team.PermissionType(Global.Company).SetMember,
				co_enum.Team.PermissionType(Global.Company).SetOwner,
				co_enum.Team.PermissionType(Global.Company).SetCaptain,
			}},
	}
	return result
}
