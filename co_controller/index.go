package co_controller

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_controller/internal"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
)

type (
	CompanyController  internal.CompanyController
	EmployeeController internal.EmployeeController
	TeamController     internal.TeamController
	MyController       internal.MyController
)

type CoController struct {
	Company  i_controller.ICompany
	Employee i_controller.IEmployee
	Team     i_controller.ITeam
	My       i_controller.IMy
}

var (
	Company  = internal.Company
	Employee = internal.Employee
	Team     = internal.Team
	My       = internal.My
)
