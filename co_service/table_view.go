// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package co_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type (
	IFdRechargeView interface {
		// QueryAccountRecharge 查询充值记录列表
		QueryAccountRecharge(ctx context.Context, search *base_model.SearchParams) (*co_model.FdRechargeViewListRes, error)
		// GetAccountRechargeById 根据ID获取充值记录
		GetAccountRechargeById(ctx context.Context, id int64) (*co_model.FdRechargeViewRes, error)
	}
)

var (
	localFdRechargeView IFdRechargeView
)

func FdRechargeView() IFdRechargeView {
	if localFdRechargeView == nil {
		panic("implement not found for interface IFdRechargeView, forgot register?")
	}
	return localFdRechargeView
}

func RegisterFdRechargeView(i IFdRechargeView) {
	localFdRechargeView = i
}
