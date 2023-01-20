package co_company_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type GetEmployeeByIdReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type HasEmployeeByNameReq struct {
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionMainId int64  `json:"unionMainId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的员工ID"`
}

type HasEmployeeByNoReq struct {
	No          string `json:"no" dc:"工号"`
	UnionMainId int64  `json:"unionMainId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的员工ID"`
}

type QueryEmployeeListReq struct {
	sys_model.SearchParams
}

type CreateEmployeeReq struct {
	co_model.Employee
}

type UpdateEmployeeReq struct {
	co_model.Employee
}

type DeleteEmployeeReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type SetEmployeeMobileReq struct {
	Mobile  int64  `json:"mobile" v:"required|phone#请数据手机号|手机号错误" dc:"手机号"`
	Captcha string `json:"captcha" v:"required#请输入手机验证码"`
}

type SetEmployeeAvatarReq struct {
	ImageId int64 `json:"imageId" dc:"头像ID"`
}

type GetEmployeeDetailByIdReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type GetEmployeeListByRoleIdReq struct {
	RoleId int64 `json:"id" v:"required#ID校验失败" dc:"角色ID"`
}
