// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyTeamMemberView is the golang structure of table co_company_team_member_view for DAO operations like Where/Data.
type CompanyTeamMemberView struct {
	g.Meta       `orm:"table:co_company_team_member_view, do:true"`
	Id           interface{} //
	TeamId       interface{} //
	EmployeeId   interface{} //
	InviteUserId interface{} //
	UnionMainId  interface{} //
	JoinAt       *gtime.Time //
	CompanyType  interface{} //
}
