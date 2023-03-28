package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type IEmployee[TIRes co_model.IEmployeeRes] interface {
	// GetEmployeeById 根据id获取员工信息
	GetEmployeeById(ctx context.Context, req *co_company_api.GetEmployeeByIdReq) (TIRes, error)

	// GetEmployeeDetailById 获取员工详情信息
	GetEmployeeDetailById(ctx context.Context, req *co_company_api.GetEmployeeDetailByIdReq) (res TIRes, err error)

	// HasEmployeeByName 员工名称是否存在
	HasEmployeeByName(ctx context.Context, req *co_company_api.HasEmployeeByNameReq) (api_v1.BoolRes, error)

	// HasEmployeeByNo 员工工号是否存在
	HasEmployeeByNo(ctx context.Context, req *co_company_api.HasEmployeeByNoReq) (api_v1.BoolRes, error)

	// QueryEmployeeList 查询员工列表
	QueryEmployeeList(ctx context.Context, req *co_company_api.QueryEmployeeListReq) (*base_model.CollectRes[TIRes], error)

	// CreateEmployee 创建员工信息
	CreateEmployee(ctx context.Context, req *co_company_api.CreateEmployeeReq) (TIRes, error)

	// UpdateEmployee 更新员工信息
	UpdateEmployee(ctx context.Context, req *co_company_api.UpdateEmployeeReq) (TIRes, error)

	// DeleteEmployee 删除员工信息
	DeleteEmployee(ctx context.Context, req *co_company_api.DeleteEmployeeReq) (api_v1.BoolRes, error)

	// GetEmployeeListByRoleId 根据角色ID获取所有所属员工列表
	GetEmployeeListByRoleId(ctx context.Context, req *co_company_api.GetEmployeeListByRoleIdReq) (*base_model.CollectRes[TIRes], error)

	// GetEmployeeListByTeamId 获取团队成员|列表
	GetEmployeeListByTeamId(ctx context.Context, req *co_company_api.GetEmployeeListByTeamId) (*base_model.CollectRes[TIRes], error)
}
