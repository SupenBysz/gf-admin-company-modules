package co_enum_financial

type financial struct {
	InOutType      inOutType
	TradeType      tradeType
	PermissionType permissionType
}

var Financial = financial{
	InOutType:      InOutType,
	TradeType:      TradeType,
	PermissionType: PermissionType,
}
