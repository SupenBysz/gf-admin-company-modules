package consts

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
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
			func(conf *co_model.Config) {
				// 模块初始化逻辑
				PermissionTree = initPermission(conf)
			},
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

func initPermission(conf *co_model.Config) []*permission.SysPermissionTree {
	result := []*permission.SysPermissionTree{
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5947986066667973,
				Name:       conf.I18n.T(context.TODO(), "CompanyName"),
				Identifier: conf.Identifier.Company,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{},
		},
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       conf.I18n.T(context.TODO(), "{#CompanyName}{#EmployeeName}"),
				Identifier: conf.Identifier.Employee,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{},
		},
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       conf.I18n.T(context.TODO(), "{#CompanyName}{#TeamName}"),
				Identifier: conf.Identifier.Team,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{},
		},
	}
	return result
}
