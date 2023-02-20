package employee

import "github.com/kysion/base-library/utility/enum"

type StateEnum enum.IEnumCode[int]

type state struct {
	Canceled  StateEnum
	Quit      StateEnum
	WaitAudit StateEnum
	Normal    StateEnum
}

var State = state{
	Canceled:  enum.New[StateEnum](-2, "已注销"),
	Quit:      enum.New[StateEnum](-1, "已离职"),
	WaitAudit: enum.New[StateEnum](0, "待认证"),
	Normal:    enum.New[StateEnum](1, "已入职"),
}

func (e *state) New(code int, description string) StateEnum {
	if (code&State.Canceled.Code()) == State.Canceled.Code() ||
		(code&State.Quit.Code()) == State.Quit.Code() ||
		(code&State.WaitAudit.Code()) == State.WaitAudit.Code() ||
		(code&State.Normal.Code()) == State.Normal.Code() {
		return enum.New[StateEnum](code, description)
	}
	panic("User.State.New: error")
}
