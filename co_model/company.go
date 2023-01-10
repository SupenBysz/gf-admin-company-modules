package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type Company struct {
	Id            int64  `json:"id"             description:"ID"`
	Name          string `json:"name"           description:"服务商名称" v:"required|max-length:128#请输入名称|名称最大支持128个字符"`
	ContactName   string `json:"contactName"    description:"商务联系人" v:"required|max-length:16#请输入商务联系人姓名|商务联系人姓名最多支持16个字符"`
	ContactMobile string `json:"contactMobile"  description:"商务联系电话" v:"required|phone|max-length:32#请输入商务联系人电话|商务联系人电话格式错误|商务联系人电话最多支持16个字符"`
	Remark        string `json:"remark"         description:"备注"`
}

type CompanyRes co_entity.Company

type CompanyListRes sys_model.CollectRes[*co_entity.Company]
