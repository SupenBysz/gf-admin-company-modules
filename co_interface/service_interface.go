package co_interface

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type IDao interface {
	DB() gdb.DB
	Table() string
	Group() string
	Ctx(ctx context.Context) *gdb.Model
	Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error)
}

type (
	ICompany interface {
		GetCompanyById(ctx context.Context, id int64) (*co_entity.Company, error)
		GetCompanyByName(ctx context.Context, name string) (*co_entity.Company, error)
		HasCompanyByName(ctx context.Context, name string, excludeIds ...int64) bool
		QueryCompanyList(ctx context.Context, filter *sys_model.SearchParams) (*co_model.CompanyListRes, error)
		CreateCompany(ctx context.Context, info *co_model.Company) (*co_entity.Company, error)
		UpdateCompany(ctx context.Context, info *co_model.Company) (*co_entity.Company, error)
		GetCompanyDetail(ctx context.Context, id int64) (*co_entity.Company, error)
	}
	IEmployee interface {
		GetEmployeeById(ctx context.Context, id int64) (*co_entity.CompanyEmployee, error)
		GetEmployeeByName(ctx context.Context, name string) (*co_entity.CompanyEmployee, error)
		HasEmployeeByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool
		HasEmployeeByNo(ctx context.Context, no string, unionMainId int64, excludeIds ...int64) bool
		GetEmployeeBySession(ctx context.Context) (*co_entity.CompanyEmployee, error)
		QueryEmployeeList(ctx context.Context, search *sys_model.SearchParams) (*co_model.EmployeeListRes, error)
		CreateEmployee(ctx context.Context, info *co_model.Employee) (*co_entity.CompanyEmployee, error)
		UpdateEmployee(ctx context.Context, info *co_model.Employee) (*co_entity.CompanyEmployee, error)
		DeleteEmployee(ctx context.Context, id int64) (bool, error)
		SetEmployeeMobile(ctx context.Context, newMobile int64, captcha string) (bool, error)
		SetEmployeeAvatar(ctx context.Context, imageId int64) (bool, error)
		GetEmployeeDetailById(ctx context.Context, id int64) (*co_entity.CompanyEmployee, error)
		GetEmployeeListByRoleId(ctx context.Context, roleId int64) (*co_model.EmployeeListRes, error)
	}
	ITeam interface {
		GetTeamById(ctx context.Context, id int64) (*co_entity.CompanyTeam, error)
		GetTeamByName(ctx context.Context, name string) (*co_entity.CompanyTeam, error)
		HasTeamByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool
		QueryTeamList(ctx context.Context, search *sys_model.SearchParams) (*co_model.TeamListRes, error)
		QueryTeamMemberList(ctx context.Context, search *sys_model.SearchParams) (*co_model.TeamMemberListRes, error)
		CreateTeam(ctx context.Context, info *co_model.Team) (*co_entity.CompanyTeam, error)
		UpdateTeam(ctx context.Context, id int64, name string, remark string) (*co_entity.CompanyTeam, error)
		GetTeamMemberList(ctx context.Context, id int64) (*co_model.EmployeeListRes, error)
		QueryTeamListByEmployee(ctx context.Context, employeeId int64, unionMainId int64) (*co_model.TeamListRes, error)
		SetTeamMember(ctx context.Context, teamId int64, employeeIds []int64) (api_v1.BoolRes, error)
		SetTeamOwner(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error)
		SetTeamCaptain(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error)
		DeleteTeam(ctx context.Context, teamId int64) (api_v1.BoolRes, error)
		DeleteTeamMemberByEmployee(ctx context.Context, employeeId int64) (bool, error)
	}
	IMy interface {
		GetProfile(ctx context.Context) (*co_model.MyProfileRes, error)
		GetCompany(ctx context.Context) (*co_model.MyCompanyRes, error)
		GetTeams(ctx context.Context) (res co_model.MyTeamListRes, err error)
	}
)

type IModules interface {
	Company() ICompany
	Team() ITeam
	Employee() IEmployee
	GetConfig() *co_model.Config
	My() IMy
	SetI18n(i18n *gi18n.Manager) error
	T(ctx context.Context, content string) string
	Tf(ctx context.Context, format string, values ...interface{}) string
}
