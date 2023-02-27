package invoice

type invoice struct {
	AuditType  auditType
	State      state
	MakeType   makeType
	BelongType belongType
}

var Invoice = invoice{
	AuditType:  AuditType,
	State:      State,
	MakeType:   MakeType,
	BelongType: BelongType,
}
