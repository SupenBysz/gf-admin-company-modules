package co_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

type cEmployee[T co_interface.IModules] struct {
	modules T
}

var Employee = func(modules co_interface.IModules) *cEmployee[co_interface.IModules] {
	return &cEmployee[co_interface.IModules]{
		modules: modules,
	}
}

func (c *cEmployee[T]) GetEmployeeById(ctx context.Context, req *co_v1.GetEmployeeByIdReq) (*co_model.EmployeeRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeRes, error) {
			ret, err := c.modules.Employee().GetEmployeeById(ctx, req.Id)
			return (*co_model.EmployeeRes)(ret), err
		},
		co_enum.Employee.PermissionType(c.modules).ViewDetail,
	)
}

// GetEmployeeDetailById 获取员工详情信息
func (c *cEmployee[T]) GetEmployeeDetailById(ctx context.Context, req *co_v1.GetEmployeeDetailByIdReq) (res *co_model.EmployeeRes, err error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeRes, error) {
			ret, err := c.modules.Employee().GetEmployeeDetailById(ctx, req.Id)
			return (*co_model.EmployeeRes)(ret), err
		},
		co_enum.Employee.PermissionType(c.modules).MoreDetail,
	)
}

// HasEmployeeByName 员工名称是否存在
func (c *cEmployee[T]) HasEmployeeByName(ctx context.Context, req *co_v1.HasEmployeeByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Employee().HasEmployeeByName(ctx, req.Name, req.UnionMainId, req.ExcludeId) == true, nil
		},
	)
}

// HasEmployeeByNo 员工工号是否存在
func (c *cEmployee[T]) HasEmployeeByNo(ctx context.Context, req *co_v1.HasEmployeeByNoReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Employee().HasEmployeeByNo(ctx, req.No, req.UnionMainId, req.ExcludeId) == true, nil
		},
	)
}

// QueryEmployeeList 查询员工列表
func (c *cEmployee[T]) QueryEmployeeList(ctx context.Context, req *co_v1.QueryEmployeeListReq) (*co_model.EmployeeListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeListRes, error) {
			return c.modules.Employee().QueryEmployeeList(ctx, &req.SearchParams)
		},
		co_enum.Employee.PermissionType(c.modules).List,
	)
}

// CreateEmployee 创建员工信息
func (c *cEmployee[T]) CreateEmployee(ctx context.Context, req *co_v1.CreateEmployeeReq) (*co_model.EmployeeRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeRes, error) {
			ret, err := c.modules.Employee().CreateEmployee(ctx, &req.Employee)
			return (*co_model.EmployeeRes)(ret), err
		},
		co_enum.Employee.PermissionType(c.modules).Create,
	)
}

// UpdateEmployee 更新员工信息
func (c *cEmployee[T]) UpdateEmployee(ctx context.Context, req *co_v1.UpdateEmployeeReq) (*co_model.EmployeeRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.EmployeeRes, error) {
			ret, err := c.modules.Employee().UpdateEmployee(ctx, &req.Employee)
			return (*co_model.EmployeeRes)(ret), err
		},
		co_enum.Employee.PermissionType(c.modules).Update,
	)
}

// DeleteEmployee 删除员工信息
func (c *cEmployee[T]) DeleteEmployee(ctx context.Context, req *co_v1.DeleteEmployeeReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Employee().DeleteEmployee(ctx, req.Id)
			return ret == true, err
		},
		co_enum.Employee.PermissionType(c.modules).Delete,
	)
}

// SetEmployeeMobile 设置员工手机号
func (c *cEmployee[T]) SetEmployeeMobile(ctx context.Context, req *co_v1.SetEmployeeMobileReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Employee().SetEmployeeMobile(ctx, req.Mobile, req.Captcha)
			return ret == true, err
		},
		co_enum.Employee.PermissionType(c.modules).SetMobile,
	)
}

// SetEmployeeAvatar 设置员工头像
func (c *cEmployee[T]) SetEmployeeAvatar(ctx context.Context, req *co_v1.SetEmployeeAvatarReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Employee().SetEmployeeAvatar(ctx, req.ImageId)
			return ret == true, err
		},
		co_enum.Employee.PermissionType(c.modules).SetAvatar,
	)
}
