package consts

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_module"
	"github.com/SupenBysz/gf-admin-company-modules/example/controller"
)

type global struct {
	Modules    co_interface.IModules
	Controller *controller.ModuleController
}

var (
	PermissionTree []*permission.SysPermissionTree
	Global         = global{
		Modules: co_module.NewModules(
			&co_model.Config{
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
			&co_dao.XDao{
				Company:    co_dao.NewCompany(&co_dao.Company{}),
				Team:       co_dao.NewCompanyTeam(&co_dao.CompanyTeam{}),
				Employee:   co_dao.NewCompanyEmployee(&co_dao.CompanyEmployee{}),
				TeamMember: co_dao.NewCompanyTeamMember(&co_dao.CompanyTeamMember{}),
			},
		),
	}
)

func init() {
	Global.Controller = &controller.ModuleController{
		Company:  controller.Company(Global.Modules),
		Employee: controller.Employee(Global.Modules),
		Team:     controller.Team(Global.Modules),
		My:       controller.My(Global.Modules),
	}
}
