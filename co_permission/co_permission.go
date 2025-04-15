package co_permission

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/company"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/employee"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/finance"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/liceense"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission/internal/team"
)

type (
	CompanyPermissionType  = company.Permission
	EmployeePermissionType = employee.Permission
	TeamPermissionType     = team.Permission
	FinancePermissionType  = finance.Permission

	LicensePermissionType = liceense.PermissionTypeEnum
)

var (
	Company  = company.Company
	Employee = employee.Employee
	Team     = team.Team
	Finance  = finance.Finance

	License = liceense.License
)
