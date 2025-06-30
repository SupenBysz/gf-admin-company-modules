package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type FdAccountView struct {
	co_entity.FdAccountView
}
type FdAccountDetailView struct {
	co_entity.FdAccountDetailView
}

type FdAccountViewRes struct {
	Id int64 `json:"id"                 orm:"id"                   description:""`
	FdAccountView
	//FdAccountDetailView
	User      *sys_model.SysUser             `json:"user" dc:"用户"`
	Employee  *co_entity.CompanyEmployeeView `json:"employee" dc:"职工/员工/会员/队员"`
	UnionMain *co_entity.CompanyView         `json:"unionMain" dc:"所属单位"`
}

type FdAccountViewListRes base_model.CollectRes[FdAccountViewRes]
