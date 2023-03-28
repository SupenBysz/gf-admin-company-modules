package financial

type financial struct {
	InOutType inOutType
	TradeType tradeType

	AccountType accountType
	SceneType   sceneType
}

var Financial = financial{
	InOutType:   InOutType,
	TradeType:   TradeType,
	AccountType: AccountType,
	SceneType:   SceneType,
}
