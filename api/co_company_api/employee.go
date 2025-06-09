package co_company_api

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type GetEmployeeByIdReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`

	//Exclude []string `json:"exclude" dc:"需要附加数据的返回值字段集，如果没有填写，默认添加所有附加数据"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type HasEmployeeByNameReq struct {
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionNameId int64  `json:"unionNameId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的员工ID"`
}

type HasEmployeeByNoReq struct {
	No        string `json:"no" dc:"工号"`
	ExcludeId int64  `json:"excludeId" dc:"要排除的员工ID"`
}

type QueryEmployeeListReq struct {
	base_model.SearchParams
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type CreateEmployeeReq struct {
	co_model.Employee
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type UpdateEmployeeReq struct {
	co_model.UpdateEmployee
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type DeleteEmployeeReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type GetEmployeeDetailByIdReq struct {
	Id      int64    `json:"id" v:"required#ID校验失败" dc:"员工ID"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type GetEmployeeListByRoleIdReq struct {
	Id      int64    `json:"id" v:"required#ID校验失败" dc:"角色ID"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type SetEmployeeRolesReq struct {
	UserId  int64   `json:"userId" v:"required#用户ID校验失败" dc:"用户ID"`
	RoleIds []int64 `json:"roleIds" v:"required#角色IDS校验失败" dc:"角色IDS"`
}

type SetEmployeeStateReq struct {
	Id    int64 `json:"id"       v:"required#ID校验失败"     dc:"ID，保持与USERID一致" `
	State int   `json:"state"        v:"in:-1,0,1#请选择员工状态" dc:"状态：-1已离职，0待确认，1已入职"`
}

type SetEmployeeCommissionRateReq struct {
	UserId         int64 `json:"userId"       v:"required#用户ID校验失败"     dc:"用户ID"`
	CommissionRate int   `json:"commissionRate"        v:"in:0,1#请选择佣金比例" dc:"佣金比例：0不参与，1参与"`
}
