package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/base_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/utility/kconv"
	"reflect"
)

type Team struct {
	OverrideDo        base_interface.DoModel[co_do.CompanyTeam]       `json:"-"`
	TeamMemberDo      base_interface.DoModel[co_do.CompanyTeamMember] `json:"-"`
	Id                int64                                           `json:"id"                dc:"ID"`
	Name              string                                          `json:"name"              v:"required|max-length:128#名称不能为空|名称长度超128字符出限定范围" dc:"团队名称，公司维度下唯一"`
	OwnerEmployeeId   int64                                           `json:"ownerEmployeeId"   dc:"团队所有者/业务总监/业务经理/团队队长"`
	CaptainEmployeeId int64                                           `json:"captainEmployeeId" dc:"团队队长编号/小组组长"`
	ParentId          int64                                           `json:"parentId" dc:"团队或小组父级ID"`
	Remark            string                                          `json:"remark"            dc:"备注"`
}

type TeamRes struct {
	co_entity.CompanyTeam
	Owner     *EmployeeRes `json:"owner" dc:"团队所有者/业务总监/业务经理/团队队长"`
	Captain   *EmployeeRes `json:"captain" dc:"团队队长编号/小组组长"`
	UnionMain *CompanyRes  `json:"unionMain" dc:"关联主体"`
	Parent    *TeamRes     `json:"parent" dc:"团队或小组父级ID"`
}

// 这些从配置加载
//type TeamInvite struct {
//	ExpireAt       *gtime.Time `json:"expireAt"       description:"邀约码的过期失效" `
//	ActivateNumber int         `json:"activateNumber" description:"邀约码的激活次数限制" dc:"邀约码激活总次数"`
//}

type TeamInviteCodeRes struct {
	Team      co_entity.CompanyTeam `json:"team" dc:"团队信息"`
	InviteRes *sys_model.InviteRes  `json:"inviteRes" dc:"邀约信息"`
}

func (m *TeamRes) Data() *TeamRes {
	return m
}

func (m *TeamRes) SetOwner(employee interface{}) {
	if employee == nil || reflect.ValueOf(employee).Type() != reflect.ValueOf(m.Owner).Type() {
		return
	}
	kconv.Struct(employee, &m.Owner)
}

func (m *TeamRes) SetCaptain(employee interface{}) {
	if employee == nil || reflect.ValueOf(employee).Type() != reflect.ValueOf(m.Captain).Type() {
		return
	}
	kconv.Struct(employee, &m.Captain)
}

func (m *TeamRes) SetUnionMain(unionMain interface{}) {
	if unionMain == nil || reflect.ValueOf(unionMain).Type() != reflect.ValueOf(m.UnionMain).Type() {
		return
	}
	kconv.Struct(unionMain, &m.UnionMain)
}

func (m *TeamRes) SetParentTeam(parent interface{}) {
	if parent == nil || reflect.ValueOf(parent).Type() != reflect.ValueOf(m.Parent).Type() {
		return
	}
	kconv.Struct(parent, &m.Parent)
}

type ITeamRes interface {
	Data() *TeamRes
	SetOwner(employee interface{})
	SetCaptain(employee interface{})
	SetUnionMain(unionMain interface{})
	SetParentTeam(parent interface{})
}

type TeamMemberRes struct {
	co_entity.CompanyTeamMember
	Employee   *EmployeeRes `json:"employee"   dc:"成员"`
	InviteUser *EmployeeRes `json:"inviteUser" dc:"邀约人"`
	UnionMain  *CompanyRes  `json:"unionMain"  dc:"关联主体"`
}
