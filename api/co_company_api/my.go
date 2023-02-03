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
