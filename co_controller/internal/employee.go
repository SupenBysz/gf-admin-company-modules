package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

type EmployeeController struct {
	modules co_interface.IModules
}

var Employee = func(modules co_interface.IModules) i_controller.IEmployee {
	return &EmployeeController{
		modules: modules,
	}
}

func (c *EmployeeController) GetModules() co_interface.IModules {
	return c.modules
}

func (c *EmployeeController) GetEmployeeById(ctx context.Context, req *co_company_api.GetEmployeeByIdReq) (*co_model.EmployeeRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeRes, error) {
			return c.modules.Employee().GetEmployeeById(ctx, req.Id)
		},
		co_enum.Employee.PermissionType(c.modules).ViewDetail,
	)
}

// GetEmployeeDetailById 获取员工详情信息
func (c *EmployeeController) GetEmployeeDetailById(ctx context.Context, req *co_company_api.GetEmployeeDetailByIdReq) (res *co_model.EmployeeRes, err error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeRes, error) {
			return c.modules.Employee().GetEmployeeDetailById(ctx, req.Id)
		},
		co_enum.Employee.PermissionType(c.modules).MoreDetail,
	)
}

// HasEmployeeByName 员工名称是否存在
func (c *EmployeeController) HasEmployeeByName(ctx context.Context, req *co_company_api.HasEmployeeByNameReq) (api_v1.BoolRes, error) {
	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Employee().HasEmployeeByName(ctx, req.Name, unionMainId, req.ExcludeId) == true, nil
		},
	)
}

// HasEmployeeByNo 员工工号是否存在
func (c *EmployeeController) HasEmployeeByNo(ctx context.Context, req *co_company_api.HasEmployeeByNoReq) (api_v1.BoolRes, error) {
	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Employee().HasEmployeeByNo(ctx, req.No, unionMainId, req.ExcludeId) == true, nil
		},
	)
}

// QueryEmployeeList 查询员工列表
func (c *EmployeeController) QueryEmployeeList(ctx context.Context, req *co_company_api.QueryEmployeeListReq) (*co_model.EmployeeListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeListRes, error) {
			return c.modules.Employee().QueryEmployeeList(ctx, &req.SearchParams)
		},
		co_enum.Employee.PermissionType(c.modules).List,
	)
}

// CreateEmployee 创建员工信息
func (c *EmployeeController) CreateEmployee(ctx context.Context, req *co_company_api.CreateEmployeeReq) (*co_model.EmployeeRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeRes, error) {
			ret, err := c.modules.Employee().CreateEmployee(ctx, &req.Employee)
			return (*co_model.EmployeeRes)(ret), err
		},
		co_enum.Employee.PermissionType(c.modules).Create,
	)
}

// UpdateEmployee 更新员工信息
func (c *EmployeeController) UpdateEmployee(ctx context.Context, req *co_company_api.UpdateEmployeeReq) (*co_model.EmployeeRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeRes, error) {
			ret, err := c.modules.Employee().UpdateEmployee(ctx, &req.Employee)
			return (*co_model.EmployeeRes)(ret), err
		},
		co_enum.Employee.PermissionType(c.modules).Update,
	)
}

// DeleteEmployee 删除员工信息
func (c *EmployeeController) DeleteEmployee(ctx context.Context, req *co_company_api.DeleteEmployeeReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Employee().DeleteEmployee(ctx, req.Id)
			return ret == true, err
		},
		co_enum.Employee.PermissionType(c.modules).Delete,
	)
}

// SetEmployeeMobile 设置员工手机号
func (c *EmployeeController) SetEmployeeMobile(ctx context.Context, req *co_company_api.SetEmployeeMobileReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Employee().SetEmployeeMobile(ctx, req.Mobile, req.Captcha)
			return ret == true, err
		},
		co_enum.Employee.PermissionType(c.modules).SetMobile,
	)
}

// GetEmployeeListByRoleId 根据角色ID获取所有所属员工
func (c *EmployeeController) GetEmployeeListByRoleId(ctx context.Context, req *co_company_api.GetEmployeeListByRoleIdReq) (*co_model.EmployeeListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeListRes, error) {
			return c.modules.Employee().GetEmployeeListByRoleId(ctx, req.Id)
		},
		co_enum.Employee.PermissionType(c.modules).ViewDetail,
	)
}
