// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyTeam is the golang structure of table co_company_team for DAO operations like Where/Data.
type CompanyTeam struct {
	g.Meta            `orm:"table:co_company_team, do:true"`
	Id                interface{} // ID
	Name              interface{} // 团队名称，公司维度下唯一
	OwnerEmployeeId   interface{} // 团队所有者/业务总监/业务经理/团队队长
	CaptainEmployeeId interface{} // 团队队长编号/小组组长
	UnionMainId       interface{} // 所属主体单位ID
	ParentId          interface{} // 父级ID
	Remark            interface{} // 备注
	CreatedAt         *gtime.Time //
	UpdatedAt         *gtime.Time //
	DeletedAt         *gtime.Time //
	Type              interface{} // 类型：0部门，1团队，2小组
	Slogan            interface{} // 口号
	Title             interface{} // 称号
	LogoId            interface{} // LOGO
	CreatedBy         interface{} // 创建者
}
