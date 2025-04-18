package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/base_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type FdRecharge struct {
	OverrideDo     base_interface.DoModel[co_do.FdRecharge] `json:"-"`
	Id             int                                      `json:"id"             orm:"id"               description:""`
	UserId         int64                                    `json:"userId"         orm:"user_id"          description:"用户ID"`
	CurrencyCode   string                                   `json:"currencyCode"   orm:"currency_code"    description:"币种"`
	Amount         float64                                  `json:"amount"         orm:"amount"           description:"金额"`
	RechargeMethod int                                      `json:"rechargeMethod" orm:"recharge_method"  description:"方式：1手动冲正、2银行卡、3支付宝、4微信、5云闪付、6翼支付"`
	RechargeTime   *gtime.Time                              `json:"rechargeTime"   orm:"recharge_time"    description:"充值时间"`
	Status         string                                   `json:"status"         orm:"status"           description:"状态:0待处理、1已完成、2已取消、4失败"`
	PaymentOrderId string                                   `json:"paymentOrderId" orm:"payment_order_id" description:"外部支付订单，即一般为第三方支付平台生成的订单号"`
	BillsId        string                                   `json:"billsId"        orm:"bills_id"         description:"交易流水号，一般用于后续的队长和查询"`
	AuditState     int                                      `json:"auditState"     orm:"audit_state"      description:"审核状态：0待审核，1通过、2不通过"`
	AccountId      int64                                    `json:"accountId"      orm:"account_id"       description:"财务账户"`
	Remark         string                                   `json:"remark"         orm:"remark"           description:"备注"`
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
