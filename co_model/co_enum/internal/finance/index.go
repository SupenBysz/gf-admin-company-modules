package finance

type finance struct {
	InOutType  inOutType
	TradeType  tradeType
	TradeState tradeState

	AccountType accountType
	SceneType   sceneType

	AllowExceed allowExceed

	RechargeState  rechargeState
	RechargeMethod rechargeMethod
}

var Finance = finance{
	InOutType:      InOutType,
	TradeType:      TradeType,
	TradeState:     TradeState,
	AccountType:    AccountType,
	SceneType:      SceneType,
	AllowExceed:    AllowExceed,
	RechargeState:  RechargeState,
	RechargeMethod: RechargeMethod,
}
