package common

type common struct {
	AppealState    appealState
	CommissionMode commissionMode
}

var Common = common{
	AppealState:    AppealState,
	CommissionMode: CommissionMode,
}
