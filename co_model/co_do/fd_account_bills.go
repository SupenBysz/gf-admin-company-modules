// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FdAccountBills is the golang structure of table co_fd_account_bills for DAO operations like Where/Data.
type FdAccountBills struct {
	g.Meta        `orm:"table:co_fd_account_bills, do:true"`
	Id            interface{} // ID
	FromUserId    interface{} // 交易发起方UserID，如果是系统则固定为-1
	ToUserId      interface{} // 交易对象UserID
	FdAccountId   interface{} // 财务账户ID
	BeforeBalance interface{} // 交易前账户余额
	Amount        interface{} // 交易金额
	AfterBalance  interface{} // 交易后账户余额
	UnionOrderId  interface{} // 关联业务订单ID
	UnionMainId   interface{} // 关联单位ID
	InOutType     interface{} // 收支类型：1收入，2支出
	TradeType     interface{} // 交易类型，1转账、2消费、4退款、8佣金、16保证金、32诚意金、64手续费/服务费、128提现、256充值、512营收，8192其它
	TradeAt       *gtime.Time // 交易时间
	Remark        interface{} // 备注信息
	TradeState    interface{} // 交易状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭
	HandlingFee   interface{} // 手续费，当前记录产生的手续费，如果有的话
	ExtJson       interface{} // 扩展数据
	DeletedAt     *gtime.Time //
	DeletedBy     interface{} //
	CreatedAt     *gtime.Time //
	CreatedBy     interface{} //
}
