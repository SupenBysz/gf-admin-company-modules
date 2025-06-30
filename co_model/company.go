package co_model

import (
	"reflect"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/base_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/kconv"
)

type Company struct {
	OverrideDo base_interface.DoModel[co_do.Company] `json:"-"`
	// 兼容用法
	Employee       base_interface.DoModel[*Employee] `json:"-"`
	Id             int64                             `json:"id"             dc:"ID"`
	Name           *string                           `json:"name"           dc:"服务商名称" v:"required|max-length:128#请输入名称|名称最大支持128个字符"`
	ContactName    *string                           `json:"contactName"    dc:"商务联系人" v:"required|max-length:16#请输入商务联系人姓名|商务联系人姓名最多支持16个字符"`
	ContactMobile  *string                           `json:"contactMobile"  dc:"商务联系电话" v:"required-if:id,0|phone|max-length:32#请输入商务联系人电话|商务联系人电话格式错误|商务联系人电话最多支持16个字符"`
	Remark         *string                           `json:"remark"         dc:"备注"`
	Address        *string                           `json:"address"        dc:"地址，主体资质审核通过后，会通过实际地址覆盖掉该地址"`
	LicenseId      *int64                            `json:"licenseId"      dc:"主体资质id"`
	LicenseState   *int                              `json:"licenseState"   dc:"主体状态,和主体资质状态保持一致"`
	CountryCode    *string                           `json:"countryCode"    dc:"所属国家编码" v:"max-length:16|in:AF,AX,AL,DZ,AS,AD,AO,AI,AG,AR,AM,AW,AU,AT,AZ,BS,BH,BD,BB,BY,BE,BZ,BJ,BM,BT,BO,BQ,BA,BW,BV,BR,IO,BN,BG,BF,BI,CV,KH,CM,CA,KY,CF,TD,CL,CN,CX,CC,CO,KM,CD,CG,CK,CR,CI,HR,CU,CW,CY,CZ,DK,DJ,DM,DO,EC,EG,SV,GQ,ER,EE,ET,FK,FO,FJ,FI,FR,GF,PF,TF,GA,GM,GE,DE,GH,GI,GR,GL,GD,GP,GU,GT,GG,GN,GW,GY,HT,HM,VA,GN,SV,GQ,HU,IS,IN,ID,IR,IQ,IE,IM,IL,IT,JM,JP,JE,JO,KZ,KE,KI,KP,KR,KW,KG,LA,LV,LB,LS,LR,LY,LI,LT,LU,MO,MK,MG,MW,MY,MV,ML,MT,MH,MQ,MR,MU,YT,MX,FM,MD,MC,MN,ME,MS,MA,MZ,MM,NA,NR,NP,NL,NC,NZ,NI,NE,NG,NU,NF,MP,NO,OM,PK,PW,PS,PA,PG,PY,PE,PH,PN,PL,PT,PR,QA,RE,RO,RU,RW,BL,SH,KN,LC,MF,PM,VC,WS,SM,ST,SA,SN,RS,SC,SL,SG,SX,SK,SI,SB,SO,ZA,GS,SS,ES,LK,SD,SR,SJ,SZ,SE,CH,SY,TJ,TZ,TH,TL,TG,TK,TO,TT,TN,TR,TM,TC,TV,UG,UA,AE,GB,US,UM,UY,UZ,VU,VE,VN,VG,VI,WF,EH,YE,ZM,ZW#所属国家编码最多支持16个字符|国家编码错误"`
	Region         *string                           `json:"region"         dc:"所属地区" v:"max-length:128#所属地区最多支持128个字符"`
	IsRegister     bool                              `json:"-"              dc:"是否公开注册的行为"`
	State          *int                              `json:"state"          dc:"状态：0未启用，1正常"`
	Score          *int                              `json:"score"          dc:"综合服务分"`
	LogoId         *int64                            `json:"logoId"         dc:"logo id"`
	StarLevel      *int                              `json:"starLevel"      dc:"星级"`
	ParentId       int64                             `json:"parentId"       dc:"父级ID"`
	CommissionRate int                               `json:"commissionRate" dc:"佣金率，如果开启会员权益模块，且佣金率有冲突，则该值优先级高会员权益模块；规则：该值不能大于上级佣金"`
}

type CompanyRes struct {
	*co_entity.CompanyView
	AdminUser *EmployeeRes `json:"adminUser"`
}

type CompanyListRes base_model.CollectRes[CompanyRes]

func (m *CompanyRes) Data() *CompanyRes {
	return m
}

func (m *CompanyRes) SetAdminUser(employee interface{}) {
	if employee == nil || reflect.ValueOf(employee).Type() != reflect.ValueOf(m.AdminUser).Type() {
		return
	}
	kconv.Struct(employee, &m.AdminUser)
}

type ICompanyRes interface {
	Data() *CompanyRes
	SetAdminUser(employee interface{})
}

type IRes[T interface{}] interface {
	Data() *T
	SetAdminUser(sysUser *sys_model.SysUser)
}
