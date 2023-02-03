package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type IEmployee interface {
	iModule
	// GetEmployeeById 根据id获取员工信息
	GetEmployeeById(ctx context.Context, req *co_company_api.GetEmployeeByIdReq) (*co_model.EmployeeRes, error)

	// GetEmployeeDetailById 获取员工详情信息
	GetEmployeeDetailById(ctx context.Context, req *co_company_api.GetEmployeeDetailByIdReq) (res *co_model.EmployeeRes, err error)

	// HasEmployeeByName 员工名称是否存在
	HasEmployeeByName(ctx context.Context, req *co_company_api.HasEmployeeByNameReq) (api_v1.BoolRes, error)

	// HasEmployeeByNo 员工工号是否存在
	HasEmployeeByNo(ctx context.Context, req *co_company_api.HasEmployeeByNoReq) (api_v1.BoolRes, error)

	// QueryEmployeeList 查询员工列表
	QueryEmployeeList(ctx context.Context, req *co_company_api.QueryEmployeeListReq) (*co_model.EmployeeListRes, error)

	// CreateEmployee 创建员工信息
	CreateEmployee(ctx context.Context, req *co_company_api.CreateEmployeeReq) (*co_model.EmployeeRes, error)

	// UpdateEmployee 更新员工信息
	UpdateEmployee(ctx context.Context, req *co_company_api.UpdateEmployeeReq) (*co_model.EmployeeRes, error)

	// DeleteEmployee 删除员工信息
	DeleteEmployee(ctx context.Context, req *co_company_api.DeleteEmployeeReq) (api_v1.BoolRes, error)

	// SetEmployeeMobile 设置员工手机号
	SetEmployeeMobile(ctx context.Context, req *co_company_api.SetEmployeeMobileReq) (api_v1.BoolRes, error)

	// GetEmployeeListByRoleId 根据角色ID获取所有所属员工列表
	GetEmployeeListByRoleId(ctx context.Context, req *co_company_api.GetEmployeeListByRoleIdReq) (*co_model.EmployeeListRes, error)
}
