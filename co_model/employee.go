package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Employee struct {
	Id          int64       `json:"id"           description:"ID，保持与USERID一致"`
	No          string      `json:"no"           v:"max-length:16#工号长度超出限定16字符" description:"工号"`
	Avatar      string      `json:"avatar"       description:"头像"`
	Name        string      `json:"name"         v:"required|max-length:16#名称不能为空|工号长度超出限定16字符" description:"姓名"`
	Mobile      string      `json:"mobile"       v:"phone#手机号校验失败" description:"手机号"`
	State       int         `json:"state"        v:"in:-2,-1,0,1#请选择员工状态" description:"状态： -2已注销，-1已离职，0待认证，1已入职"`
	UnionMainId int64       `json:"unionMainId"  v:"required|请选择员工归属主体" description:"所属主体"`
	HiredAt     *gtime.Time `json:"hiredAt"      v:"date-format:Y-m-d#入职日期" description:"入职日期"`
}

type EmployeeUser struct {
	g.Meta   `orm:"table:sys_user"`
	Id       int64  `json:"id"           description:"ID，保持与USERID一致"`
	Username string `json:"username"  description:"账号"`
	State    int    `json:"state"     description:"状态：0未激活、1正常、-1封号、-2异常、-3已注销"`
}

type EmployeeRes struct {
	co_entity.CompanyEmployee
	User   EmployeeUser             `orm:"with:id" json:"user"`
	Detail sys_entity.SysUserDetail `orm:"with:id" json:"detail"`
}

type EmployeeListRes sys_model.CollectRes[*EmployeeRes]
