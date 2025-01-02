package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type License struct {
	IdCardFrontPath        string      `json:"idCardFrontPath"           dc:"身份证头像面照片" v:"required-license#请上传身份证头像面照片"`
	IdCardBackPath         string      `json:"idCardBackPath"            dc:"身份证国徽面照片" v:"required-license#请上传身份证国徽面照片"`
	IdCardNo               string      `json:"idCardNo"                  dc:"身份证号" v:"required-license|resident-id#请输入身份证号|请输入正确的身份证号"`
	IdCardExpiredDate      *gtime.Time `json:"idCardExpiredDate"         dc:"身份证有效期" v:"date#身份证日期格式错误"`
	IdCardIssuingAuthority *gtime.Time `json:"idCardIssuingAuthority"    dc:"身份证签发机关"`
	IdCardAddress          string      `json:"idCardAddress"             dc:"身份证户籍地址" v:"max-length:128#身份证地址最大支持128个字符"`
	PersonContactName      string      `json:"personContactName"         dc:"法人，必须是自然人" v:"required-license|max-length:16#请输入法人姓名|法人姓名最大支持16个字符"`
	PersonContactMobile    string      `json:"personContactMobile"       dc:"法人，联系电话" v:"required-license|max-length:16#请输入法人联系电话|法人联系电话最大支持16个字符"`

	BusinessLicenseName        string `json:"businessLicenseName"        dc:"公司全称" v:"required-license|max-length:128#请输入公司全称|公司全称最大支持128个字符"`
	BusinessLicenseAddress     string `json:"businessLicenseAddress"     dc:"公司地址" v:"required-license|max-length:128#请输入公司地址|公司地址最大支持128个字符"`
	BusinessLicensePath        string `json:"businessLicensePath"        dc:"营业执照图片地址" v:"required-license#请上传营业执照图片"`
	BusinessLicenseScope       string `json:"businessLicenseScope"       dc:"经营范围"`
	BusinessLicenseRegCapital  string `json:"businessLicenseRegCapital"  dc:"注册资本" v:"max-length:32#注册资本最大支持32字符"`
	BusinessLicenseTermTime    string `json:"businessLicenseTermTime"    dc:"营业期限" v:"max-length:64#营业期限最大支持64字符"`
	BusinessLicenseOrgCode     string `json:"businessLicenseOrgCode"     dc:"组织机构代码" v:"max-length:16#组织机构代码最大支持16字符"`
	BusinessLicenseCreditCode  string `json:"businessLicenseCreditCode"  dc:"统一社会信用代码" v:"required-license|max-length:32#请输入统一社会信用代码|统一社会信用代码最大支持32个字符"`
	BusinessLicenseLegal       string `json:"businessLicenseLegal"       dc:"法人" v:"required-license|max-length:64#请输入法人名称|法人名称最大支持64个字符"`
	BusinessLicenseLegalMobile string `json:"businessLicenseLegalMobile" dc:"法人联系电话"`
	BusinessLicenseLegalPath   string `json:"businessLicenseLegalPath"   dc:"法人证照，如果法人不是自然人，则该项必填" v:"max-length:256#法人证照地址最大支持256个字符"`
	Remark                     string `json:"remark"                     dc:"备注"`
	BrandName                  string `json:"brandName"                  dc:"品牌名称"`
	ServerMobile               string `json:"serverMobile"               dc:"服务电话"`

	State         int              `json:"state"           dc:"状态：0失效、1正常" v:"in:0,1#状态错误"`
	AuthType      int              `json:"authType"        dc:"主体资质认证类型："`
	DoorPictures  []AttachPictures `json:"doorPictures"    dc:"门头照列表"`
	OtherPictures []AttachPictures `json:"otherPictures"   dc:"其他资质图片列表"`
	Summary       string           `json:"summary"         dc:"概述"`
}

type AttachPictures struct {
	Title string `json:"title" dc:"标题"`
	Id    string `json:"id"    dc:"id" v:"required#请输入门头照图片"`
	Desc  string `json:"desc"  dc:"描述"`
	Url   string `json:"url"   dc:"图片地址"`
	Size  int64  `json:"size"  dc:"图片大小"`
	Ext   string `json:"ext"   dc:"图片后缀"`
}

type LicenseRes co_entity.License
type LicenseListRes base_model.CollectRes[co_entity.License]

type AuditLicense struct {
	UnionMainId int64             `json:"unionMainId"    dc:"资质审核关联的业务主体ID"`
	LicenseId   int64             `json:"licenseId"      dc:"资质ID"`
	UserId      int64             `json:"userId"         dc:"上传资质的userId"`
	Summary     string            `json:"summary"        dc:"概述"`
	UserType    sys_enum.UserType `json:"userType"       dc:"用户类型"`
	License
}
