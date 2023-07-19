package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type License struct {
	IdcardFrontPath           string      `json:"idcardFrontPath"           dc:"身份证头像面照片" v:"required-license#请上传身份证头像面照片"`
	IdcardBackPath            string      `json:"idcardBackPath"            dc:"身份证国徽面照片" v:"required-license#请上传身份证国徽面照片"`
	IdcardNo                  string      `json:"idcardNo"                  dc:"身份证号" v:"required-license|resident-id#请输入身份证号|请输入正确的身份证号"`
	IdcardExpiredDate         *gtime.Time `json:"idcardExpiredDate"         dc:"身份证有效期" v:"date#身份证日期格式错误"`
	IdcardAddress             string      `json:"idcardAddress"             dc:"身份证户籍地址" v:"max-length:128#身份证地址最大支持128个字符"`
	PersonContactName         string      `json:"personContactName"         dc:"法人，必须是自然人" v:"required-license|max-length:16#请输入法人姓名|法人姓名最大支持16个字符"`
	PersonContactMobile       string      `json:"personContactMobile"       dc:"法人，联系电话" v:"required-license|max-length:16#请输入法人联系电话|法人联系电话最大支持16个字符"`
	BusinessLicenseName       string      `json:"businessLicenseName"       dc:"公司全称" v:"required-license|max-length:128#请输入公司全称|公司全称最大支持128个字符"`
	BusinessLicenseAddress    string      `json:"businessLicenseAddress"    dc:"公司地址" v:"required-license|max-length:128#请输入公司地址|公司地址最大支持128个字符"`
	BusinessLicensePath       string      `json:"businessLicensePath"       dc:"营业执照图片地址" v:"required-license#请上传营业执照图片"`
	BusinessLicenseScope      string      `json:"businessLicenseScope"      dc:"经营范围"`
	BusinessLicenseRegCapital string      `json:"businessLicenseRegCapital" dc:"注册资本" v:"max-length:32#注册资本最大支持32字符"`
	BusinessLicenseTermTime   string      `json:"businessLicenseTermTime"   dc:"营业期限" v:"max-length:64#营业期限最大支持64字符"`
	BusinessLicenseOrgCode    string      `json:"businessLicenseOrgCode"    dc:"组织机构代码" v:"max-length:16#组织机构代码最大支持16字符"`
	BusinessLicenseCreditCode string      `json:"businessLicenseCreditCode" dc:"统一社会信用代码" v:"required-license|max-length:32#请输入统一社会信用代码|统一社会信用代码最大支持32个字符"`
	BusinessLicenseLegal      string      `json:"businessLicenseLegal"      dc:"法人" v:"required-license|max-length:64#请输入法人名称|法人名称最大支持64个字符"`
	BusinessLicenseLegalPath  string      `json:"businessLicenseLegalPath"  dc:"法人证照，如果法人不是自然人，则该项必填" v:"max-length:256#法人证照地址最大支持256个字符"`
	Remake                    string      `json:"remake"                    dc:"备注"`
	BrandName                 string      `json:"brandName"                 description:"品牌名称"`
}

type LicenseRes co_entity.License
type LicenseListRes base_model.CollectRes[co_entity.License]

type AuditLicense struct {
	UnionMainId int64 `json:"unionMainId"             dc:"资质审核关联的业务主体ID"`
	LicenseId   int64 `json:"licenseId"             dc:"资质ID"`
	UserId      int64 `json:"userId" dc:"上传资质的userId"`
	License
}
