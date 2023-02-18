package co_enum

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/company"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/employee"
	invoice "github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/invoice"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum/internal/team"
)

type (
	CompanyPermissionType = company.PermissionEnum

	EmployeePermissionType = employee.PermissionEnum
	EmployeeState          = employee.StateEnum

	TeamPermissionType = team.PermissionEnum
)

var (
	Company  = company.Company
	Employee = employee.Employee
	Team     = team.Team
	Invoice  = invoice.Invoice
)
