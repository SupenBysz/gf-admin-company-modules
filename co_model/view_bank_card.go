package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type FdBankCardView struct {
	co_entity.FdBankCardView
}

type FdBankCardViewRes struct {
	FdBankCardView `json:"employeeView"`
	User           *sys_model.SysUser `json:"user" dc:"用户"`
}

type FdBankCardViewListRes base_model.CollectRes[FdBankCardViewRes]
