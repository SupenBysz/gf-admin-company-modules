package invoice

import "github.com/kysion/base-library/utility/enum"

type StateEnum enum.IEnumCode[int]

type state struct {
	WaitAudit      StateEnum
	WaitForInvoice StateEnum
	Failure        StateEnum
	Success        StateEnum
	Cancel         StateEnum
}

var State = state{
	WaitAudit:      enum.New[StateEnum](1, "待审核"),
	WaitForInvoice: enum.New[StateEnum](2, "待开票"),
	Failure:        enum.New[StateEnum](4, "开票失败"),
	Success:        enum.New[StateEnum](8, "已开票"),
	Cancel:         enum.New[StateEnum](16, "已撤销"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.WaitAudit.Code()) == State.WaitAudit.Code() ||
		(code&State.WaitForInvoice.Code()) == State.WaitForInvoice.Code() ||
		(code&State.Failure.Code()) == State.Failure.Code() ||
		(code&State.Success.Code()) == State.Success.Code() ||
		(code&State.Cancel.Code()) == State.Cancel.Code() {
		return enum.New[StateEnum](code, description)
	} else {
		panic("kyInvoice.State.New: error")
	}
}
