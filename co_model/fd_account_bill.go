package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
)

type AccountBillRegister struct {
	FromUserId    int64       `json:"fromUserId"     v:"required#请输入交易发起方用户id" dc:"交易发起方UserID，如果是系统则固定为-1"`
	ToUserId      int64       `json:"toUserId"       v:"required#请输入交易接收方用户id"  dc:"交易对象UserID"`
	FdAccountId   int64       `json:"fdAccountId"    v:"required#财务账号ID不能为空" dc:"财务账户ID"`
	BeforeBalance int64       `json:"beforeBalance"  v:"required#交易前余额不能为空" dc:"交易前账户余额"`
	Amount        int64       `json:"amount"         v:"required#交易金额不能为空" dc:"交易金额"`
	AfterBalance  int64       `json:"afterBalance"   v:"required#交易后账号余额不能为空" dc:"交易后账户余额"`
	UnionOrderId  int64       `json:"unionOrderId"   dc:"关联业务订单ID"`
	InOutType     int         `json:"inOutType"      v:"required|in:1,2#请输入1.收入 2.支出" dc:"收支类型：1收入，2支出"`
	TradeType     int         `json:"tradeType"      v:"required|in:1,2,4,8,16,32,64,128,256,512,8192#交易类型错误" dc:"交易类型，1转账、2消费、4退款、8佣金、16保证金、32诚意金、64手续费/服务费、128提现、256充值、512营收，8192其它"`
	TradeAt       *gtime.Time `json:"tradeAt"        v:"required#交易时间不能为空"   dc:"交易时间"`
	Remark        string      `json:"remark"         dc:"备注信息"`
	TradeState    int         `json:"tradeState"     v:"required|in:1,2,4,8,16,32,64,128,256#交易状态错误"  dc:"交易状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭"`
}

type FdAccountBillRes struct {
	co_entity.FdAccountBill
}

func (m *FdAccountBillRes) Data() *FdAccountBillRes {
	return m
}

type IFdAccountBillRes interface {
	Data() *FdAccountBillRes
}
