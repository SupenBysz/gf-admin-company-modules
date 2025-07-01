package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type CompanyBillsView struct {
	co_entity.FdAccountBillsView
	FormUser  *sys_model.SysUser `json:"formUser,omitempty"`
	ToUser    *sys_model.SysUser `json:"toUser,omitempty"`
	FdAccount *co_entity.FdAccount   `json:"fdAccount,omitempty"`
}

type CompanyBillsViewRes CompanyBillsView

type CompanyBillsViewListRes base_model.CollectRes[CompanyBillsViewRes]
