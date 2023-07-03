package sys_enum_audit

type audit struct {
	Action action
	Event  eventState
}

var Audit = audit{
	Action: Action,
	Event:  Event,
}
