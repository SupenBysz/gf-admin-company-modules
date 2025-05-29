package finance

import "github.com/kysion/base-library/utility/enum"

// TradeStateEnum 交易状态：1待支付、2支付中、4已支付、8取消支付、16交易完成、32退款中、64已退款、128支付超时、256已关闭
type TradeStateEnum enum.IEnumCode[int]

type tradeState struct {
	WaitPay        TradeStateEnum
	Paying         TradeStateEnum
	Paid           TradeStateEnum
	Cancel         TradeStateEnum
	Completed      TradeStateEnum
	Refunding      TradeStateEnum
	Refunded       TradeStateEnum
	PaymentTimeout TradeStateEnum
	Closed         TradeStateEnum
}

var TradeState = tradeState{
	WaitPay:        enum.New[TradeStateEnum](0, "待支付"),
	Paying:         enum.New[TradeStateEnum](1, "支付中"),
	Paid:           enum.New[TradeStateEnum](2, "已支付"),
	Cancel:         enum.New[TradeStateEnum](4, "取消支付"),
	Completed:      enum.New[TradeStateEnum](8, "交易完成"),
	Refunding:      enum.New[TradeStateEnum](16, "退款中"),
	Refunded:       enum.New[TradeStateEnum](32, "已退款"),
	PaymentTimeout: enum.New[TradeStateEnum](64, "支付超时"),
	Closed:         enum.New[TradeStateEnum](128, "已关闭"),
}

func (e tradeState) New(code int, description string) TradeStateEnum {
	if code == TradeState.WaitPay.Code() {
		return TradeState.WaitPay
	}
	if code == TradeState.Paying.Code() {
		return TradeState.Paying
	}
	if code == TradeState.Paid.Code() {
		return TradeState.Paid
	}
	if code == TradeState.Cancel.Code() {
		return TradeState.Cancel
	}
	if code == TradeState.Completed.Code() {
		return TradeState.Completed
	}
	if code == TradeState.Refunding.Code() {
		return TradeState.Refunding
	}
	if code == TradeState.Refunded.Code() {
		return TradeState.Refunded
	}
	if code == TradeState.PaymentTimeout.Code() {
		return TradeState.PaymentTimeout
	}
	if code == TradeState.Closed.Code() {
		return TradeState.Closed
	}

	return enum.New[TradeStateEnum](code, description)
}
