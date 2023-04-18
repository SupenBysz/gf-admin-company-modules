package co_v1

import (
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/gogf/gf/v2/frame/g"
)

type GetEmployeeByIdReq struct {
	g.Meta `method:"post" summary:"根据ID获取员工|信息" tags:"员工"`
	co_company_api.GetEmployeeByIdReq
}

type HasEmployeeByNameReq struct {
	g.Meta `method:"post" summary:"判断名称是否存在" tags:"员工"`
	co_company_api.HasEmployeeByNameReq
}

type HasEmployeeByNoReq struct {
	g.Meta `method:"post" summary:"判断工号是否存在" tags:"员工"`
	co_company_api.HasEmployeeByNoReq
}

type QueryEmployeeListReq struct {
	g.Meta `method:"post" summary:"查询员工|列表" tags:"员工"`
	co_company_api.QueryEmployeeListReq
}

type CreateEmployeeReq struct {
	g.Meta `method:"post" summary:"创建员工|信息" tags:"员工"`
	co_company_api.CreateEmployeeReq
}

type UpdateEmployeeReq struct {
	g.Meta `method:"post" summary:"更新员工|信息" tags:"员工"`
	co_company_api.UpdateEmployeeReq
}

type DeleteEmployeeReq struct {
	g.Meta `method:"post" summary:"删除员工|信息" tags:"员工"`
	co_company_api.DeleteEmployeeReq
}

type GetEmployeeDetailByIdReq struct {
	g.Meta `method:"post" summary:"获取员工详情|信息" tags:"员工"`
	co_company_api.GetEmployeeDetailByIdReq
}

type GetEmployeeListByRoleIdReq struct {
	g.Meta `method:"post" summary:"根据角色ID获取所有所属员工|列表" tags:"员工"`
	co_company_api.GetEmployeeListByRoleIdReq
}

type SetEmployeeRolesReq struct {
	g.Meta `method:"post" summary:"设置员工角色" dc:"设置员工所属角色" tags:"员工"`
	co_company_api.SetEmployeeRolesReq
}

type SetEmployeeStateReq struct {
	g.Meta `method:"post" summary:"设置员工状态" dc:"设置员工状态" tags:"员工"`
	co_company_api.SetEmployeeStateReq
}
