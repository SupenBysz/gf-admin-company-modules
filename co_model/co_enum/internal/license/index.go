package license

type license struct {
	AuthType authType
	State    state
}

var License = license{
	AuthType: AuthType,
	State:    State,
}
