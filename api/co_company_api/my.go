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
	Mobile   string `json:"mobile" v:"required|phone#请数据手机号|手机号错误" dc:"手机号，注意：此手机号不是用于登陆的手机号，通常属于工作的手机号联系方式"`
	Captcha  string `json:"captcha" v:"required#请输入手机验证码"`
	Password string `json:"password" v:"required#请输入登陆账号密码" dc:"登录密码"`
}

type SetMailReq struct {
	OldMail  string `json:"oldMail" v:"email|max-length:64#邮箱账号格式错误|邮箱长度超出限定64字符" dc:"原邮箱，首次设置原邮箱地址可为空"`
	NewMail  string `json:"newMail" v:"required|email|max-length:64#请输入新邮箱账号|邮箱账号格式错误|邮箱长度超出限定64字符" dc:"新邮箱，注意：此邮箱不是用于登陆的邮箱，通常属于工作的邮箱联系方式"`
	Captcha  string `json:"captcha" v:"required#请输入邮箱验证码"`
	Password string `json:"password" v:"required#请输入登陆账号密码" dc:"登录密码" v:"min-length:6#密码最短为6位"`
}

type GetAccountBillsReq struct {
	base_model.SearchParams
}

type GetAccountsReq struct{}

type GetBankCardsReq struct{}

type GetInvoicesReq struct{}

type UpdateAccountReq struct {
	AccountId int64 `json:"accountId" dc:"具体需要修改的财务账号id" v:"required#财务账号id不能为空"`
	co_model.UpdateAccount
}

type GetMyCompanyPermissionListReq struct {
	PermissionType *int `json:"permissionType" dc:"过滤权限类型：0不限，1API接口，2菜单"`
}
