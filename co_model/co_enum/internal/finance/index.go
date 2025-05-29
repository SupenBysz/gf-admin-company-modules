package finance

type finance struct {
	InOutType           inOutType
	TradeType           tradeType
	TradeState          tradeState
	PaymentState        paymentState
	SpecialPaymentState specialPaymentState

	AccountType accountType
	SceneType   sceneType

	AllowExceed allowExceed

	RechargeState  rechargeState
	RechargeMethod rechargeMethod
}

var Finance = finance{
	InOutType:           InOutType,
	TradeType:           TradeType,
	TradeState:          TradeState,
	AccountType:         AccountType,
	PaymentState:        PaymentState,
	SpecialPaymentState: SpecialPaymentState,
	SceneType:           SceneType,
	AllowExceed:         AllowExceed,
	RechargeState:       RechargeState,
	RechargeMethod:      RechargeMethod,
}
