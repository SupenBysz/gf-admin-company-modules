package co_v1

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetEmployeeByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取员工|信息" tags:"员工"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type HasEmployeeByNameReq struct {
	g.Meta      `method:"post" summary:"判断名称是否存在" tags:"员工"`
	Name        string `json:"name" v:"required#名称不能为空" dc:"名称"`
	UnionMainId int64  `json:"unionMainId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的员工ID"`
}

type HasEmployeeByNoReq struct {
	g.Meta      `method:"post" summary:"判断工号是否存在" tags:"员工"`
	No          string `json:"no" dc:"工号"`
	UnionMainId int64  `json:"unionMainId" dc:"关联主体ID"`
	ExcludeId   int64  `json:"excludeId" dc:"要排除的员工ID"`
}

type QueryEmployeeListReq struct {
	g.Meta `method:"post" summary:"查询员工|列表" tags:"员工"`
	sys_model.SearchParams
}

type CreateEmployeeReq struct {
	g.Meta `method:"post" summary:"创建员工|信息" tags:"员工"`
	co_model.Employee
}

type UpdateEmployeeReq struct {
	g.Meta `method:"post" summary:"更新员工|信息" tags:"员工"`
	co_model.Employee
}

type DeleteEmployeeReq struct {
	g.Meta `method:"post" summary:"删除员工|信息" tags:"员工"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}

type SetEmployeeMobileReq struct {
	g.Meta  `method:"post" summary:"删除员工|信息" tags:"员工"`
	Mobile  int64  `json:"id" v:"required|phone#请数据手机号|手机号错误" dc:"手机号"`
	Captcha string `json:"captcha" v:"required#请输入手机验证码"`
}

type SetEmployeeAvatarReq struct {
	g.Meta  `method:"post" summary:"设置头像|信息" tags:"员工"`
	ImageId int64 `json:"imageId" dc:"头像ID"`
}

type GetEmployeeDetailByIdReq struct {
	g.Meta `method:"post" summary:"获取员工详情|信息" tags:"员工"`
	Id     int64 `json:"id" v:"required#ID校验失败" dc:"员工ID"`
}
