// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FdAccountBills is the golang structure for table fd_account_bills.
type FdAccountBills struct {
	Id            int64       `json:"id"            orm:"id"             description:"ID"`
	FromUserId    int64       `json:"fromUserId"    orm:"from_user_id"   description:"交易发起方UserID，如果是系统则固定为-1"`
	ToUserId      int64       `json:"toUserId"      orm:"to_user_id"     description:"交易对象UserID"`
	FdAccountId   int64       `json:"fdAccountId"   orm:"fd_account_id"  description:"财务账户ID"`
	BeforeBalance int64       `json:"beforeBalance" orm:"before_balance" description:"交易前账户余额"`
	Amount        int64       `json:"amount"        orm:"amount"         description:"交易金额"`
	AfterBalance  int64       `json:"afterBalance"  orm:"after_balance"  description:"交易后账户余额"`
	UnionOrderId  int64       `json:"unionOrderId"  orm:"union_order_id" description:"关联业务订单ID"`
	UnionMainId   int64       `json:"unionMainId"   orm:"union_main_id"  description:"关联单位ID"`
	InOutType     int         `json:"inOutType"     orm:"in_out_type"    description:"收支类型：1收入，2支出"`
	TradeType     int         `json:"tradeType"     orm:"trade_type"     description:"交易类型，1转账、2消费、4退款、8佣金、16保证金、32诚意金、64手续费/服务费、128提现、256充值、512营收，8192其它"`
	TradeAt       *gtime.Time `json:"tradeAt"       orm:"trade_at"       description:"交易时间"`
	Remark        string      `json:"remark"        orm:"remark"         description:"备注信息"`
	TradeState    int         `json:"tradeState"    orm:"trade_state"    description:"交易状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭"`
	HandlingFee   int64       `json:"handlingFee"   orm:"handling_fee"   description:"手续费，当前记录产生的手续费，如果有的话"`
	ExtJson       string      `json:"extJson"       orm:"ext_json"       description:"扩展数据"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:""`
	DeletedBy     int64       `json:"deletedBy"     orm:"deleted_by"     description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:""`
}
