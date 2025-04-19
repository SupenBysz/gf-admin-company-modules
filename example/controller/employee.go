package controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/utility/kconv"
)

type EmployeeController[TIRes co_model.IEmployeeRes] struct {
	i_controller.IEmployee[TIRes]
}

func Employee[
	ITCompanyRes co_model.ICompanyRes,
	TIRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillsRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) *EmployeeController[TIRes] {
	return &EmployeeController[TIRes]{
		IEmployee: co_controller.Employee(modules),
	}
}

func (c *EmployeeController[TIRes]) GetEmployeeById(ctx context.Context, req *co_v1.GetEmployeeByIdReq) (TIRes, error) {
	return c.IEmployee.GetEmployeeById(ctx, &req.GetEmployeeByIdReq)
}

// GetEmployeeDetailById 获取员工详情信息
func (c *EmployeeController[TIRes]) GetEmployeeDetailById(ctx context.Context, req *co_v1.GetEmployeeDetailByIdReq) (TIRes, error) {
	return c.IEmployee.GetEmployeeDetailById(ctx, &req.GetEmployeeDetailByIdReq)
}

// HasEmployeeByName 员工名称是否存在
func (c *EmployeeController[TIRes]) HasEmployeeByName(ctx context.Context, req *co_v1.HasEmployeeByNameReq) (api_v1.BoolRes, error) {
	return c.IEmployee.HasEmployeeByName(ctx, &req.HasEmployeeByNameReq)
}

// HasEmployeeByNo 员工工号是否存在
func (c *EmployeeController[TIRes]) HasEmployeeByNo(ctx context.Context, req *co_v1.HasEmployeeByNoReq) (api_v1.BoolRes, error) {
	return c.IEmployee.HasEmployeeByNo(ctx, &req.HasEmployeeByNoReq)
}

// QueryEmployeeList 查询员工列表
func (c *EmployeeController[TIRes]) QueryEmployeeList(ctx context.Context, req *co_v1.QueryEmployeeListReq) (*co_model.EmployeeListRes, error) {
	ret, err := c.IEmployee.QueryEmployeeList(ctx, &req.QueryEmployeeListReq)
	return kconv.Struct(ret, &co_model.EmployeeListRes{}), err
}

// CreateEmployee 创建员工信息
func (c *EmployeeController[TIRes]) CreateEmployee(ctx context.Context, req *co_v1.CreateEmployeeReq) (TIRes, error) {
	return c.IEmployee.CreateEmployee(ctx, &req.CreateEmployeeReq)
}

// UpdateEmployee 更新员工信息
func (c *EmployeeController[TIRes]) UpdateEmployee(ctx context.Context, req *co_v1.UpdateEmployeeReq) (TIRes, error) {
	return c.IEmployee.UpdateEmployee(ctx, &req.UpdateEmployeeReq)
}

// DeleteEmployee 删除员工信息
func (c *EmployeeController[TIRes]) DeleteEmployee(ctx context.Context, req *co_v1.DeleteEmployeeReq) (api_v1.BoolRes, error) {
	return c.IEmployee.DeleteEmployee(ctx, &req.DeleteEmployeeReq)
}

func (c *EmployeeController[TIRes]) GetEmployeeListByRoleId(ctx context.Context, req *co_v1.GetEmployeeListByRoleIdReq) (*co_model.EmployeeListRes, error) {
	ret, err := c.IEmployee.GetEmployeeListByRoleId(ctx, &req.GetEmployeeListByRoleIdReq)
	return kconv.Struct(ret, &co_model.EmployeeListRes{}), err
}

// SetEmployeeRoles 设置员工角色
func (c *EmployeeController[TIRes]) SetEmployeeRoles(ctx context.Context, req *co_v1.SetEmployeeRolesReq) (api_v1.BoolRes, error) {
	return c.IEmployee.SetEmployeeRoles(ctx, &req.SetEmployeeRolesReq)
}

// SetEmployeeState 设置员工状态
func (c *EmployeeController[TIRes]) SetEmployeeState(ctx context.Context, req *co_v1.SetEmployeeStateReq) (api_v1.BoolRes, error) {
	ret, err := c.IEmployee.SetEmployeeState(ctx, &req.SetEmployeeStateReq)
	return ret == true, err
}
