package co_permission

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/company"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/employee"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/financial"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/team"
)

type (
	CompanyPermissionType   = company.Permission
	EmployeePermissionType  = employee.Permission
	TeamPermissionType      = team.Permission
	FinancialPermissionType = financial.Permission
)

var (
	Company   = company.Company
	Employee  = employee.Employee
	Team      = team.Team
	Financial = financial.Financial
)
