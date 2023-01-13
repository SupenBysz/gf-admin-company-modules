package boot

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/permission"

	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_module"
)

func init() {
	company := co_module.NewModules(&co_model.Config{
		I18n:                           nil,
		AllowEmptyNo:                   true,
		IsCreateDefaultEmployeeAndRole: false,
		HardDeleteWaitAt:               0,
		CompanyName:                    "公司",
		KeyIndex:                       "Company",
		RoutePrefix:                    "/company",
		StoragePath:                    "./resources/company",
		UserType:                       sys_enum.User.Type.SuperAdmin,
		Identifier: co_model.Identifier{
			Company:  "company",
			Employee: "employee",
			Team:     "team",
		},
	})
	company.GetConfig().I18n = InitI18n(company.GetConfig().I18n)
	co_consts.Global.Company = company
	InitPermissionTree(company.GetConfig())
}

func InitI18n(i18n *gi18n.Manager) *gi18n.Manager {
	if i18n == nil {
		i18n = gi18n.New()
		i18n.SetPath("i18n/company")
		i18n.SetLanguage("zh-CN")
	}

	return i18n
}

func InitPermissionTree(conf *co_model.Config) {
	co_consts.PermissionTree = []*permission.SysPermissionTree{
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5947986066667973,
				Name:       conf.I18n.T(context.TODO(), "permission.Company.Name"),
				Identifier: conf.Identifier.Company,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看用户，查看某个用户登录账户
				sys_enum.User.PermissionType.ViewDetail,
				// 用户列表，查看所有用户
				sys_enum.User.PermissionType.List,
				// 重置密码，重置某个用户的登录密码
				sys_enum.User.PermissionType.ResetPassword,
				// 设置状态，设置某个用户的状态
				sys_enum.User.PermissionType.SetState,
				// 修改密码，修改自己的登录密码
				sys_enum.User.PermissionType.ChangePassword,
			},
		},
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       conf.CompanyName + "员工",
				Identifier: conf.Identifier.Employee,
				Type:       1,
				IsShow:     1,
			},
		},
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       conf.CompanyName + "团队",
				Identifier: conf.Identifier.Team,
				Type:       1,
				IsShow:     1,
			},
		},
	}
}
