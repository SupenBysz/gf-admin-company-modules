package financial

type financial struct {
	InOutType inOutType
	TradeType tradeType

	AccountType accountType
	SceneType   sceneType

	AllowExceed allowExceed
}

var Financial = financial{
	InOutType:   InOutType,
	TradeType:   TradeType,
	AccountType: AccountType,
	SceneType:   SceneType,
	AllowExceed: AllowExceed,
}
