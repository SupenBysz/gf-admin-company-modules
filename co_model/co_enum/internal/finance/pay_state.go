package finance

import "github.com/kysion/base-library/utility/enum"

// PaymentStateEnum 核心支付状态枚举
type PaymentStateEnum enum.IEnumCode[int]

type paymentState struct {
	WaitPay      PaymentStateEnum
	Processing   PaymentStateEnum
	Paid         PaymentStateEnum
	Failed       PaymentStateEnum
	Refunding    PaymentStateEnum
	Refunded     PaymentStateEnum
	RefundFailed PaymentStateEnum
	Cancelled    PaymentStateEnum
}

var PaymentState = paymentState{
	WaitPay:      enum.New[PaymentStateEnum](1, "待支付"),
	Processing:   enum.New[PaymentStateEnum](2, "支付中"),
	Paid:         enum.New[PaymentStateEnum](4, "已支付"),
	Failed:       enum.New[PaymentStateEnum](8, "支付失败"),
	Refunding:    enum.New[PaymentStateEnum](16, "退款中"),
	Refunded:     enum.New[PaymentStateEnum](32, "已退款"),
	RefundFailed: enum.New[PaymentStateEnum](64, "退款失败"),
	Cancelled:    enum.New[PaymentStateEnum](128, "已取消"),
}

func (e paymentState) New(code int, description string) PaymentStateEnum {
	if code == PaymentState.WaitPay.Code() {
		return PaymentState.WaitPay
	}
	if code == PaymentState.Processing.Code() {
		return PaymentState.Processing
	}
	if code == PaymentState.Paid.Code() {
		return PaymentState.Paid
	}
	if code == PaymentState.Failed.Code() {
		return PaymentState.Failed
	}
	if code == PaymentState.Refunding.Code() {
		return PaymentState.Refunding
	}
	if code == PaymentState.Refunded.Code() {
		return PaymentState.Refunded
	}
	if code == PaymentState.RefundFailed.Code() {
		return PaymentState.RefundFailed
	}
	if code == PaymentState.Cancelled.Code() {
		return PaymentState.Cancelled
	}

	return enum.New[PaymentStateEnum](code, description)
}
