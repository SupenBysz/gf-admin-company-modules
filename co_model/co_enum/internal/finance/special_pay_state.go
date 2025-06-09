package finance

import "github.com/kysion/base-library/utility/enum"

// SpecialPaymentStateEnum 特殊业务场景状态枚举
type SpecialPaymentStateEnum enum.IEnumCode[int]

type specialPaymentState struct {
	PreAuthorized      SpecialPaymentStateEnum
	PreAuthCapture     SpecialPaymentStateEnum
	PreAuthRelease     SpecialPaymentStateEnum
	Splitting          SpecialPaymentStateEnum
	Disputed           SpecialPaymentStateEnum
	Settled            SpecialPaymentStateEnum
	Settling           SpecialPaymentStateEnum
	Reconciling        SpecialPaymentStateEnum
	ForeignExchange    SpecialPaymentStateEnum
	CustomsDeclaration SpecialPaymentStateEnum
	QrCodeScanned      SpecialPaymentStateEnum
	PosTransaction     SpecialPaymentStateEnum
}

var SpecialPaymentState = specialPaymentState{
	PreAuthorized:      enum.New[SpecialPaymentStateEnum](1, "预授权"),
	PreAuthCapture:     enum.New[SpecialPaymentStateEnum](2, "预授权完成"),
	PreAuthRelease:     enum.New[SpecialPaymentStateEnum](4, "预授权撤销"),
	Splitting:          enum.New[SpecialPaymentStateEnum](8, "分账中"),
	Disputed:           enum.New[SpecialPaymentStateEnum](16, "争议中"),
	Settled:            enum.New[SpecialPaymentStateEnum](32, "已结算"),
	Settling:           enum.New[SpecialPaymentStateEnum](64, "清算中"),
	Reconciling:        enum.New[SpecialPaymentStateEnum](128, "对账中"),
	ForeignExchange:    enum.New[SpecialPaymentStateEnum](256, "外汇处理中"),
	CustomsDeclaration: enum.New[SpecialPaymentStateEnum](512, "海关申报中"),
	QrCodeScanned:      enum.New[SpecialPaymentStateEnum](1024, "扫码未支付"),
	PosTransaction:     enum.New[SpecialPaymentStateEnum](2048, "POS机交易中"),
}

func (e specialPaymentState) New(code int, description string) SpecialPaymentStateEnum {
	if code == SpecialPaymentState.PreAuthorized.Code() {
		return SpecialPaymentState.PreAuthorized
	}
	if code == SpecialPaymentState.PreAuthCapture.Code() {
		return SpecialPaymentState.PreAuthCapture
	}
	if code == SpecialPaymentState.PreAuthRelease.Code() {
		return SpecialPaymentState.PreAuthRelease
	}
	if code == SpecialPaymentState.Splitting.Code() {
		return SpecialPaymentState.Splitting
	}
	if code == SpecialPaymentState.Disputed.Code() {
		return SpecialPaymentState.Disputed
	}
	if code == SpecialPaymentState.Settled.Code() {
		return SpecialPaymentState.Settled
	}
	if code == SpecialPaymentState.Settling.Code() {
		return SpecialPaymentState.Settling
	}
	if code == SpecialPaymentState.Reconciling.Code() {
		return SpecialPaymentState.Reconciling
	}
	if code == SpecialPaymentState.ForeignExchange.Code() {
		return SpecialPaymentState.ForeignExchange
	}
	if code == SpecialPaymentState.CustomsDeclaration.Code() {
		return SpecialPaymentState.CustomsDeclaration
	}
	if code == SpecialPaymentState.QrCodeScanned.Code() {
		return SpecialPaymentState.QrCodeScanned
	}
	if code == SpecialPaymentState.PosTransaction.Code() {
		return SpecialPaymentState.PosTransaction
	}

	return enum.New[SpecialPaymentStateEnum](code, description)
}
