package company

import "github.com/kysion/base-library/utility/enum"

type StateEnum enum.IEnumCode[int]

type state struct {
	Inactive StateEnum
	Disabled StateEnum
	Normal   StateEnum
}

var State = state{
	Inactive: enum.New[StateEnum](-2, "未激活"),
	Disabled: enum.New[StateEnum](0, "停用"),
	Normal:   enum.New[StateEnum](1, "正常"),
}

func (e *state) New(code int, description string) StateEnum {
	if (code&State.Inactive.Code()) == State.Inactive.Code() ||
		(code&State.Disabled.Code()) == State.Disabled.Code() ||
		(code&State.Normal.Code()) == State.Normal.Code() {
		return enum.New[StateEnum](code, description)
	}
	panic("Company.State.New: error")
}
