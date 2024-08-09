package employee

type employee struct {
	State state
	Sex   userSex
}

var Employee = employee{
	State: State,
	Sex:   Sex,
}
