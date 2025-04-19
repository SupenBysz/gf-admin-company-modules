package co_company_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type GetAccountRechargeByIdReq struct {
	Id int64 `json:"id" v:"required#ID校验失败" dc:"充值记录ID"`
}

type SetAccountRechargeAuditReq struct {
	Id    int64                `json:"id" v:"required#ID校验失败" dc:"充值记录ID"`
	State sys_enum.AuditAction `json:"state" v:"required#请输入审核状态" dc:"审核状态"`
	Reply string               `json:"reply" dc:"审核回复"`
}

type QueryAccountRechargeReq struct {
	base_model.SearchParams
}

type AccountRecharge struct {
	AccountId int64                `json:"id" v:"required#ID校验失败" dc:"资金账户ID"`
	Info      *co_model.FdRecharge `json:"info,omitempty"`
}
