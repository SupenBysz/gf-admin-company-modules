package sys_enum_audit

type audit struct {
	Event eventState
}

var Audit = audit{
	Event: Event,
}
