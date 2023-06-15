package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type Company struct {
	Id            int64  `json:"id"             dc:"ID"`
	Name          string `json:"name"           dc:"服务商名称" v:"required|max-length:128#请输入名称|名称最大支持128个字符"`
	ContactName   string `json:"contactName"    dc:"商务联系人" v:"required|max-length:16#请输入商务联系人姓名|商务联系人姓名最多支持16个字符"`
	ContactMobile string `json:"contactMobile"  dc:"商务联系电话" v:"required-if:id,0|phone|max-length:32#请输入商务联系人电话|商务联系人电话格式错误|商务联系人电话最多支持16个字符"`
	Remark        string `json:"remark"         dc:"备注"`
	Address       string `json:"address"       dc:"地址，主体资质审核通过后，会通过实际地址覆盖掉该地址"`
}

type CompanyRes struct {
	*co_entity.Company
	AdminUser *EmployeeRes `json:"adminUser"`
}

type CompanyListRes base_model.CollectRes[CompanyRes]

func (m *CompanyRes) Data() *CompanyRes {
	return m
}

type ICompanyRes interface {
	Data() *CompanyRes
}
