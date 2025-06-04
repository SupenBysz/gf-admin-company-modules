package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type FdInvoiceViewRes struct {
	co_entity.FdInvoiceView
	User      *sys_model.SysUser
	UnionMain *co_entity.CompanyView
}

type FdInvoiceViewListRes base_model.CollectRes[FdInvoiceViewRes]

func (m *FdInvoiceViewRes) Data() *FdInvoiceViewRes {
	return m
}

type IFdInvoiceViewRes interface {
	Data() *FdInvoiceViewRes
}
