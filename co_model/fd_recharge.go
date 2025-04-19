package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/base_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type FdRecharge struct {
	OverrideDo base_interface.DoModel[co_do.FdRecharge] `json:"-"`
	co_entity.FdRecharge
	CreatedAt *gtime.Time `json:"-"      orm:"created_at"       description:"记录创建时间，即充值请求提交时间"`
	UpdatedAt *gtime.Time `json:"-"      orm:"updated_at"       description:"记录最后更新时间，每次记录状态等信息变更时更新"`
	DeletedAt *gtime.Time `json:"-"      orm:"deleted_at"       description:"逻辑删除时间，用于软删除，非真正物理删除，便于数据追溯和恢复"`
}

type FdRechargeRes struct {
	co_entity.FdRecharge
}

type FdRechargeListRes base_model.CollectRes[FdRechargeRes]

func (m *FdRechargeRes) Data() *FdRechargeRes {
	return m
}

type IFdRechargeRes interface {
	Data() *FdRechargeRes
}

type FdRechargeViewRes struct {
	co_entity.FdRechargeView
}

type FdRechargeViewListRes base_model.CollectRes[FdRechargeViewRes]
