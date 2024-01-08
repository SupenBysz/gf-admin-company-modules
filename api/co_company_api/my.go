package co_company_api

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type GetProfileReq struct {
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type GetCompanyReq struct{}

type GetTeamsReq struct {
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type SetAvatarReq struct {
	ImageId int64 `json:"imageId" dc:"头像ID"`
}

type SetMobileReq struct {
	Mobile   string `json:"mobile" v:"required|phone#请数据手机号|手机号错误" dc:"手机号"`
	Captcha  string `json:"captcha" v:"required#请输入手机验证码"`
	Password string `json:"password" v:"required#请输入账号密码" dc:"登录密码"`
}

type GetAccountBillsReq struct {
	base_model.Pagination
}

type GetAccountsReq struct{}

type GetBankCardsReq struct{}

type GetInvoicesReq struct{}

type UpdateAccountReq struct {
	AccountId int64 `json:"accountId" dc:"具体需要修改的财务账号id" v:"required#财务账号id不能为空"`
	co_model.UpdateAccount
}
