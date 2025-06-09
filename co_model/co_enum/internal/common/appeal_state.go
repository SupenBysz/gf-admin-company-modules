package common

import "github.com/kysion/base-library/utility/enum"

// AppealStateEnum 申诉状态枚举
type AppealStateEnum enum.IEnumCode[int]

type appealState struct {
	None       AppealStateEnum
	Submitted  AppealStateEnum
	Processing AppealStateEnum
	Resolved   AppealStateEnum
}

var AppealState = appealState{
	None:       enum.New[AppealStateEnum](0, "无"),
	Submitted:  enum.New[AppealStateEnum](1, "已提交/申诉中"),
	Processing: enum.New[AppealStateEnum](2, "处理中"),
	Resolved:   enum.New[AppealStateEnum](4, "已处理"),
}

func (e appealState) New(code int, description string) AppealStateEnum {
	if code == AppealState.None.Code() {
		return AppealState.None
	}
	if code == AppealState.Submitted.Code() {
		return AppealState.Submitted
	}
	if code == AppealState.Processing.Code() {
		return AppealState.Processing
	}
	if code == AppealState.Resolved.Code() {
		return AppealState.Resolved
	}

	return enum.New[AppealStateEnum](code, description)
}
