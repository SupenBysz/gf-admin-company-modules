package co_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
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
	return funs.ProxyFunc1[*co_model.EmployeeRes](
		ctx, req.Id,
		c.modules.Employee().GetEmployeeById, nil,
		co_enum.Employee.PermissionType(c.modules).ViewDetail,
	)
}

// GetEmployeeDetailById 获取员工详情信息
func (c *cEmployee[T]) GetEmployeeDetailById(ctx context.Context, req *co_v1.GetEmployeeDetailByIdReq) (res *co_model.EmployeeRes, err error) {
	return funs.ProxyFunc1[*co_model.EmployeeRes](
		ctx, req.Id,
		c.modules.Employee().GetEmployeeDetailById, nil,
		co_enum.Employee.PermissionType(c.modules).MoreDetail,
	)
}

// HasEmployeeByName 员工名称是否存在
func (c *cEmployee[T]) HasEmployeeByName(ctx context.Context, req *co_v1.HasEmployeeByNameReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc[api_v1.BoolRes](
		ctx,
		func(ctx context.Context) (api_v1.BoolRes, error) {
			return c.modules.Employee().HasEmployeeByName(ctx, req.Name, req.UnionMainId, req.ExcludeId) == true, nil
		}, false,
	)
}

// HasEmployeeByNo 员工工号是否存在
func (c *cEmployee[T]) HasEmployeeByNo(ctx context.Context, req *co_v1.HasEmployeeByNoReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc[api_v1.BoolRes](
		ctx,
		func(ctx context.Context) (api_v1.BoolRes, error) {
			return c.modules.Employee().HasEmployeeByName(ctx, req.No, req.UnionMainId, req.ExcludeId) == true, nil
		}, false,
	)
}

// QueryEmployeeList 查询员工列表
func (c *cEmployee[T]) QueryEmployeeList(ctx context.Context, req *co_v1.QueryEmployeeListReq) (*co_model.EmployeeListRes, error) {
	return funs.ProxyFunc1[*co_model.EmployeeListRes](
		ctx, &req.SearchParams,
		c.modules.Employee().QueryEmployeeList,
		&co_model.EmployeeListRes{
			PaginationRes: sys_model.PaginationRes{
				Pagination: req.Pagination,
				PageTotal:  0,
				Total:      0,
			},
		},
		co_enum.Employee.PermissionType(c.modules).List,
	)
}

// CreateEmployee 创建员工信息
func (c *cEmployee[T]) CreateEmployee(ctx context.Context, req *co_v1.CreateEmployeeReq) (*co_model.EmployeeRes, error) {
	return funs.ProxyFunc1[*co_model.EmployeeRes](
		ctx, &req.Employee,
		c.modules.Employee().CreateEmployee, nil,
		co_enum.Employee.PermissionType(c.modules).Create,
	)
}

// UpdateEmployee 更新员工信息
func (c *cEmployee[T]) UpdateEmployee(ctx context.Context, req *co_v1.UpdateEmployeeReq) (*co_model.EmployeeRes, error) {
	return funs.ProxyFunc1[*co_model.EmployeeRes](
		ctx, &req.Employee,
		c.modules.Employee().UpdateEmployee, nil,
		co_enum.Employee.PermissionType(c.modules).Update,
	)
}

// DeleteEmployee 删除员工信息
func (c *cEmployee[T]) DeleteEmployee(ctx context.Context, req *co_v1.DeleteEmployeeReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc1[api_v1.BoolRes](
		ctx, req.Id,
		c.modules.Employee().DeleteEmployee, false,
		co_enum.Employee.PermissionType(c.modules).Delete,
	)
}

// SetEmployeeMobile 设置员工手机号
func (c *cEmployee[T]) SetEmployeeMobile(ctx context.Context, req *co_v1.SetEmployeeMobileReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc2[api_v1.BoolRes](
		ctx, req.Mobile, req.Captcha,
		c.modules.Employee().SetEmployeeMobile, false,
		co_enum.Employee.PermissionType(c.modules).SetMobile,
	)
}

// SetEmployeeAvatar 设置员工头像
func (c *cEmployee[T]) SetEmployeeAvatar(ctx context.Context, req *co_v1.SetEmployeeAvatarReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc1[api_v1.BoolRes](
		ctx, req.ImageId,
		c.modules.Employee().SetEmployeeAvatar, false,
		co_enum.Employee.PermissionType(c.modules).SetAvatar,
	)
}
