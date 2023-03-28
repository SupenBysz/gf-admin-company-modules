package controller

import "github.com/SupenBysz/gf-admin-company-modules/co_model"

type ModuleController[
	TICompanyRes co_model.ICompanyRes,
	TIEmployeeRes co_model.IEmployeeRes,
	TITeamRes co_model.ITeamRes,
] struct {
	Company  *CompanyController[TICompanyRes]
	Employee *EmployeeController[TIEmployeeRes]
	Team     *TeamController[TITeamRes]
	My       *MyController
}
