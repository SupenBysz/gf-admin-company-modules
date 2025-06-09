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
	OverrideDo     base_interface.DoModel[*co_do.CompanyEmployee] `json:"-"`
	Id             int64                                          `json:"id"           dc:"ID，保持与USERID一致"`
	No             string                                         `json:"no"           v:"max-length:16#工号长度超出限定16字符" dc:"工号"`
	Avatar         string                                         `json:"avatar"       dc:"头像"`
	WorkCardAvatar string                                         `json:"workCardAvatar" orm:"work_card_avatar" description:"工牌头像"`
	Name           string                                         `json:"name"         v:"required|max-length:32#姓名不能为空|姓名长度超出限定32字符" dc:"姓名"`
	Mobile         string                                         `json:"mobile"       v:"phone#手机号校验失败" dc:"手机号"`
	State          int                                            `json:"state"        v:"in:-1,0,1#请选择员工状态" dc:"状态：-1已离职，0待确认，1已入职"`
	UnionMainId    int64                                          `json:"-"            dc:"所属主体"`
	HiredAt        *gtime.Time                                    `json:"hiredAt"      v:"date-format:Y-m-d#入职日期格式错误" dc:"入职日期"`
	Sex            int                                            `json:"sex"          dc:"性别：0未知、1男、2女"`
	Remark         string                                         `json:"remark"       dc:"备注"`
	CountryCode    string                                         `json:"countryCode"  dc:"所属国家编码" v:"max-length:16|in:AF,AX,AL,DZ,AS,AD,AO,AI,AG,AR,AM,AW,AU,AT,AZ,BS,BH,BD,BB,BY,BE,BZ,BJ,BM,BT,BO,BQ,BA,BW,BV,BR,IO,BN,BG,BF,BI,CV,KH,CM,CA,KY,CF,TD,CL,CN,CX,CC,CO,KM,CD,CG,CK,CR,CI,HR,CU,CW,CY,CZ,DK,DJ,DM,DO,EC,EG,SV,GQ,ER,EE,ET,FK,FO,FJ,FI,FR,GF,PF,TF,GA,GM,GE,DE,GH,GI,GR,GL,GD,GP,GU,GT,GG,GN,GW,GY,HT,HM,VA,GN,SV,GQ,HU,IS,IN,ID,IR,IQ,IE,IM,IL,IT,JM,JP,JE,JO,KZ,KE,KI,KP,KR,KW,KG,LA,LV,LB,LS,LR,LY,LI,LT,LU,MO,MK,MG,MW,MY,MV,ML,MT,MH,MQ,MR,MU,YT,MX,FM,MD,MC,MN,ME,MS,MA,MZ,MM,NA,NR,NP,NL,NC,NZ,NI,NE,NG,NU,NF,MP,NO,OM,PK,PW,PS,PA,PG,PY,PE,PH,PN,PL,PT,PR,QA,RE,RO,RU,RW,BL,SH,KN,LC,MF,PM,VC,WS,SM,ST,SA,SN,RS,SC,SL,SG,SX,SK,SI,SB,SO,ZA,GS,SS,ES,LK,SD,SR,SJ,SZ,SE,CH,SY,TW,TJ,TZ,TH,TL,TG,TK,TO,TT,TN,TR,TM,TC,TV,UG,UA,AE,GB,US,UM,UY,UZ,VU,VE,VN,VG,VI,WF,EH,YE,ZM,ZW#所属国家编码最多支持16个字符|国家编码错误"`
	Region         string                                         `json:"region"       dc:"所属地区" v:"max-length:128#所属地区最多支持128个字符"`
	Email          string                                         `json:"email"        dc:"业务邮箱"  v:"email|max-length:64#业务邮箱格式错误|邮箱长度超出限定64字符"`
	CreatedBy	   int64										  `json:"createdBy"    dc:"创建者"`
}

type UpdateEmployee struct {
	OverrideDo     base_interface.DoModel[*co_do.CompanyEmployee] `json:"-"`
	Id             int64                                          `json:"id"           dc:"ID，保持与USERID一致" v:"required#请输入员工ID"`
	No             *string                                        `json:"no"           v:"max-length:16#工号长度超出限定16字符" dc:"工号"`
	Avatar         string                                         `json:"avatar"       dc:"头像"`
	WorkCardAvatar string                                         `json:"workCardAvatar" orm:"work_card_avatar" description:"工牌头像"`
	Name           *string                                        `json:"name"         v:"required|max-length:32#姓名不能为空|姓名长度超出限定32字符" dc:"姓名"`
	Mobile         *string                                        `json:"mobile"       v:"phone#手机号校验失败" dc:"手机号"`
	State          *int                                           `json:"state"        v:"in:-1,0,1#请选择员工状态" dc:"状态：-1已离职，0待确认，1已入职"`
	HiredAt        *gtime.Time                                    `json:"hiredAt"      v:"date-format:Y-m-d#入职日期格式错误" dc:"入职日期"`
	Sex            *int                                           `json:"sex"          dc:"性别：0未知、1男、2女"`
	Remark         *string                                        `json:"remark"       dc:"备注"`
	CountryCode    *string                                        `json:"countryCode"  dc:"所属国家编码" v:"max-length:16|in:AF,AX,AL,DZ,AS,AD,AO,AI,AG,AR,AM,AW,AU,AT,AZ,BS,BH,BD,BB,BY,BE,BZ,BJ,BM,BT,BO,BQ,BA,BW,BV,BR,IO,BN,BG,BF,BI,CV,KH,CM,CA,KY,CF,TD,CL,CN,CX,CC,CO,KM,CD,CG,CK,CR,CI,HR,CU,CW,CY,CZ,DK,DJ,DM,DO,EC,EG,SV,GQ,ER,EE,ET,FK,FO,FJ,FI,FR,GF,PF,TF,GA,GM,GE,DE,GH,GI,GR,GL,GD,GP,GU,GT,GG,GN,GW,GY,HT,HM,VA,GN,SV,GQ,HU,IS,IN,ID,IR,IQ,IE,IM,IL,IT,JM,JP,JE,JO,KZ,KE,KI,KP,KR,KW,KG,LA,LV,LB,LS,LR,LY,LI,LT,LU,MO,MK,MG,MW,MY,MV,ML,MT,MH,MQ,MR,MU,YT,MX,FM,MD,MC,MN,ME,MS,MA,MZ,MM,NA,NR,NP,NL,NC,NZ,NI,NE,NG,NU,NF,MP,NO,OM,PK,PW,PS,PA,PG,PY,PE,PH,PN,PL,PT,PR,QA,RE,RO,RU,RW,BL,SH,KN,LC,MF,PM,VC,WS,SM,ST,SA,SN,RS,SC,SL,SG,SX,SK,SI,SB,SO,ZA,GS,SS,ES,LK,SD,SR,SJ,SZ,SE,CH,SY,TW,TJ,TZ,TH,TL,TG,TK,TO,TT,TN,TR,TM,TC,TV,UG,UA,AE,GB,US,UM,UY,UZ,VU,VE,VN,VG,VI,WF,EH,YE,ZM,ZW#所属国家编码最多支持16个字符|国家编码错误"`
	Region         *string                                        `json:"region"       dc:"所属地区" v:"max-length:128#所属地区最多支持128个字符"`
	Email          *string                                        `json:"email"        dc:"业务邮箱"  v:"email|max-length:64#业务邮箱格式错误|邮箱长度超出限定64字符"`
	UpdatedBy	   int64										  `json:"updatedBy"    dc:"更新者"`
}

type EmployeeUser struct {
	g.Meta   `orm:"table:sys_user"`
	Id       int64  `json:"id"        dc:"ID，保持与USERID一致"`
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
	if data == nil || reflect.ValueOf(data).Type() != reflect.ValueOf(m.TeamList).Type() {
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
