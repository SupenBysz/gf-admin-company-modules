package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-company-modules/base_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/kconv"
	"reflect"
)

type Employee struct {
	OverrideDo  base_interface.DoModel[co_do.CompanyEmployee] `json:"-"`
	Id          int64                                         `json:"id"           dc:"ID，保持与USERID一致"`
	No          string                                        `json:"no"           v:"max-length:16#工号长度超出限定16字符" dc:"工号"`
	Avatar      string                                        `json:"avatar"       dc:"头像"`
	Name        string                                        `json:"name"         v:"required|max-length:16#名称不能为空|工号长度超出限定16字符" dc:"姓名"`
	Mobile      string                                        `json:"mobile"       v:"phone#手机号校验失败" dc:"手机号"`
	State       int                                           `json:"state"        v:"in:-1,0,1#请选择员工状态" dc:"状态：-1已离职，0待确认，1已入职"`
	UnionMainId int64                                         `json:"-"  dc:"所属主体"`
	HiredAt     *gtime.Time                                   `json:"hiredAt"      v:"date-format:Y-m-d#入职日期格式错误" dc:"入职日期"`
	Sex         int                                           `json:"sex"          dc:"性别：0未知、1男、2女"`
	Remark      string                                        `json:"remark"       description:"备注"`
}

type UpdateEmployee struct {
	OverrideDo base_interface.DoModel[co_do.CompanyEmployee] `json:"-"`
	Id         int64                                         `json:"id"           dc:"ID，保持与USERID一致" v:"required#请输入员工ID"`
	No         *string                                       `json:"no"           v:"max-length:16#工号长度超出限定16字符" dc:"工号"`
	Name       *string                                       `json:"name"         v:"required|max-length:16#名称不能为空|工号长度超出限定16字符" dc:"姓名"`
	Mobile     *string                                       `json:"mobile"       v:"phone#手机号校验失败" dc:"手机号"`
	State      *int                                          `json:"state"        v:"in:-1,0,1#请选择员工状态" dc:"状态：-1已离职，0待确认，1已入职"`
	HiredAt    *gtime.Time                                   `json:"hiredAt"      v:"date-format:Y-m-d#入职日期格式错误" dc:"入职日期"`
	Sex        *int                                          `json:"sex"          dc:"性别：0未知、1男、2女"`
	Remark     *string                                       `json:"remark"       description:"备注"`
}

type EmployeeUser struct {
	g.Meta   `orm:"table:sys_user"`
	Id       int64  `json:"id"           dc:"ID，保持与USERID一致"`
	Username string `json:"username"  dc:"账号"`
	State    int    `json:"state"     dc:"状态：0未激活、1正常、-1封号、-2异常、-3已注销"`
	Type     int    `json:"type"      dc:"用户类型，0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心"`
}

type EmployeeRes struct {
	co_entity.CompanyEmployee
	User     *EmployeeUser             ` json:"user"`
	Detail   *sys_entity.SysUserDetail ` json:"detail"`
	TeamList []Team                    `json:"teamList"`
}

func (m *EmployeeRes) SetUser(user interface{}) {
	if user == nil || reflect.ValueOf(user).Type() != reflect.ValueOf(m.User).Type() {
		return
	}
	kconv.Struct(user, &m.User)
}

func (m *EmployeeRes) SetTeamList(data interface{}) {
	if data == nil || reflect.ValueOf(data).Type() != reflect.ValueOf(m).Type() {
		return
	}

	kconv.Struct(data, &m.TeamList)
}

type EmployeeListRes base_model.CollectRes[*EmployeeRes]

func (m *EmployeeRes) Data() *EmployeeRes {
	return m
}

func (m *EmployeeRes) NewTeamList() interface{} {
	teamMemberItems := make([]*co_entity.CompanyTeamMember, 0)

	return teamMemberItems
}

type IEmployeeRes interface {
	NewTeamList() interface{}
	Data() *EmployeeRes
	SetUser(user interface{})
	SetTeamList(teamList interface{})
}

// type EmployeeListRes base_model.CollectRes[*EmployeeRes]
