package finance

type finance struct {
	InOutType inOutType
	TradeType tradeType

	AccountType accountType
	SceneType   sceneType

	AllowExceed allowExceed

	RechargeState  rechargeState
	RechargeMethod rechargeMethod
}

var Finance = finance{
	InOutType:      InOutType,
	TradeType:      TradeType,
	AccountType:    AccountType,
	SceneType:      SceneType,
	AllowExceed:    AllowExceed,
	RechargeState:  RechargeState,
	RechargeMethod: RechargeMethod,
}
