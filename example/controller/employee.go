package controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type EmployeeController struct {
	i_controller.IEmployee
}

var Employee = func(modules co_interface.IModules) *EmployeeController {
	return &EmployeeController{
		co_controller.Employee(modules),
	}
}

func (c *EmployeeController) GetModules() co_interface.IModules {
	return c.IEmployee.GetModules()
}

func (c *EmployeeController) GetEmployeeById(ctx context.Context, req *co_v1.GetEmployeeByIdReq) (*co_model.EmployeeRes, error) {
	return c.IEmployee.GetEmployeeById(ctx, &req.GetEmployeeByIdReq)
}

// GetEmployeeDetailById 获取员工详情信息
func (c *EmployeeController) GetEmployeeDetailById(ctx context.Context, req *co_v1.GetEmployeeDetailByIdReq) (res *co_model.EmployeeRes, err error) {
	return c.IEmployee.GetEmployeeDetailById(ctx, &req.GetEmployeeDetailByIdReq)
}

// HasEmployeeByName 员工名称是否存在
func (c *EmployeeController) HasEmployeeByName(ctx context.Context, req *co_v1.HasEmployeeByNameReq) (api_v1.BoolRes, error) {
	return c.IEmployee.HasEmployeeByName(ctx, &req.HasEmployeeByNameReq)
}

// HasEmployeeByNo 员工工号是否存在
func (c *EmployeeController) HasEmployeeByNo(ctx context.Context, req *co_v1.HasEmployeeByNoReq) (api_v1.BoolRes, error) {
	return c.IEmployee.HasEmployeeByNo(ctx, &req.HasEmployeeByNoReq)
}

// QueryEmployeeList 查询员工列表
func (c *EmployeeController) QueryEmployeeList(ctx context.Context, req *co_v1.QueryEmployeeListReq) (*co_model.EmployeeListRes, error) {
	return c.IEmployee.QueryEmployeeList(ctx, &req.QueryEmployeeListReq)
}

// CreateEmployee 创建员工信息
func (c *EmployeeController) CreateEmployee(ctx context.Context, req *co_v1.CreateEmployeeReq) (*co_model.EmployeeRes, error) {
	return c.IEmployee.CreateEmployee(ctx, &req.CreateEmployeeReq)
}

// UpdateEmployee 更新员工信息
func (c *EmployeeController) UpdateEmployee(ctx context.Context, req *co_v1.UpdateEmployeeReq) (*co_model.EmployeeRes, error) {
	return c.IEmployee.UpdateEmployee(ctx, &req.UpdateEmployeeReq)
}

// DeleteEmployee 删除员工信息
func (c *EmployeeController) DeleteEmployee(ctx context.Context, req *co_v1.DeleteEmployeeReq) (api_v1.BoolRes, error) {
	return c.IEmployee.DeleteEmployee(ctx, &req.DeleteEmployeeReq)
}

// SetEmployeeMobile 设置员工手机号
func (c *EmployeeController) SetEmployeeMobile(ctx context.Context, req *co_v1.SetEmployeeMobileReq) (api_v1.BoolRes, error) {
	return c.IEmployee.SetEmployeeMobile(ctx, &req.SetEmployeeMobileReq)
}

func (c *EmployeeController) GetEmployeeListByRoleId(ctx context.Context, req *co_v1.GetEmployeeListByRoleIdReq) (*co_model.EmployeeListRes, error) {
	return c.IEmployee.GetEmployeeListByRoleId(ctx, &req.GetEmployeeListByRoleIdReq)
}
