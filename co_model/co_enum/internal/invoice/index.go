package co_enum_invoice

type invoice struct {
	AuditType auditType
	State     state
	MakeType  makeType
}

var Invoice = invoice{
	AuditType: AuditType,
	State:     State,
	MakeType:  MakeType,
}
