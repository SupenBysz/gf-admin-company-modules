package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/base_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type AccountBillsRegister struct {
	OverrideDo    base_interface.DoModel[co_do.FdAccountBills] `json:"-"`
	FromUserId    int64       `json:"fromUserId"    orm:"from_user_id"   description:"交易发起方UserID，如果是系统则固定为-1"`
	ToUserId      int64       `json:"toUserId"      orm:"to_user_id"     description:"交易对象UserID"`
	FdAccountId   int64       `json:"fdAccountId"   orm:"fd_account_id"  description:"财务账户ID"`
	BeforeBalance int64       `json:"beforeBalance" orm:"before_balance" description:"交易前账户余额"`
	Amount        int64       `json:"amount"        orm:"amount"         description:"交易金额"`
	AfterBalance  int64       `json:"afterBalance"  orm:"after_balance"  description:"交易后账户余额"`
	UnionOrderId  int64       `json:"unionOrderId"  orm:"union_order_id" description:"关联业务订单ID"`
	InOutType     int         `json:"inOutType"     orm:"in_out_type"    description:"收支类型：1收入，2支出"`
	TradeType     int         `json:"tradeType"     orm:"trade_type"     description:"交易类型，1转账、2消费、4退款、8佣金、16保证金、32诚意金、64手续费/服务费、128提现、256充值、512营收，8192其它"`
	TradeAt       *gtime.Time `json:"tradeAt"       orm:"trade_at"       description:"交易时间"`
	Remark        string      `json:"remark"        orm:"remark"         description:"备注信息"`
	TradeState    int         `json:"tradeState"    orm:"trade_state"    description:"交易状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭、512已冻结、1024、已解冻、2048解冻失败"`
	HandlingFee   int64       `json:"handlingFee"   orm:"handling_fee"   description:"手续费，当前记录产生的手续费，如果有的话"`
	ExtJson       string      `json:"extJson"       orm:"ext_json"       description:"扩展数据"`
}

type FdAccountBillsListRes base_model.CollectRes[FdAccountBillsRes]

type FdAccountBillsRes struct {
	co_entity.FdAccountBills
}

func (m *FdAccountBillsRes) Data() *FdAccountBillsRes {
	return m
}

type IFdAccountBillsRes interface {
	Data() *FdAccountBillsRes
}
