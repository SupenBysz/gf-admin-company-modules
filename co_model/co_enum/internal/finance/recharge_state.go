package finance

import "github.com/kysion/base-library/utility/enum"

type RechargeStateEnum enum.IEnumCode[int]

type rechargeState struct {
	Pending    RechargeStateEnum
	Processing RechargeStateEnum
	Paid       RechargeStateEnum
	Partially  RechargeStateEnum
	Failed     RechargeStateEnum
	Cancelled  RechargeStateEnum
	Awaiting   RechargeStateEnum
}

var RechargeState = rechargeState{
	Pending:    enum.New[RechargeStateEnum](0, "待处理"),
	Processing: enum.New[RechargeStateEnum](1, "处理中"),
	Paid:       enum.New[RechargeStateEnum](2, "已支付"),
	Partially:  enum.New[RechargeStateEnum](3, "部分成功"),
	Failed:     enum.New[RechargeStateEnum](4, "失败"),
	Cancelled:  enum.New[RechargeStateEnum](5, "已取消"),
	Awaiting:   enum.New[RechargeStateEnum](6, "待确认"),
}

func (e rechargeState) New(code int, description string) RechargeStateEnum {
	switch code {
	case e.Pending.Code():
		return e.Pending
	case e.Processing.Code():
		return e.Processing
	case e.Paid.Code():
		return e.Paid
	case e.Partially.Code():
		return e.Partially
	case e.Failed.Code():
		return e.Failed
	case e.Cancelled.Code():
		return e.Cancelled
	case e.Awaiting.Code():
		return e.Awaiting
	default:
		panic("RechargeState: error")
	}
}
