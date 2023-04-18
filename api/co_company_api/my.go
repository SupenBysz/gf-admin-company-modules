package co_company_api

type GetProfileReq struct {
}

type GetCompanyReq struct {
}

type GetTeamsReq struct {
}

type SetAvatarReq struct {
	ImageId int64 `json:"imageId" dc:"头像ID"`
}

type SetMobileReq struct {
	Mobile   string `json:"mobile" v:"required|phone#请数据手机号|手机号错误" dc:"手机号"`
	Captcha  string `json:"captcha" v:"required#请输入手机验证码"`
	Password string `json:"password" v:"required#请输入账号密码" dc:"登录密码"`
}
