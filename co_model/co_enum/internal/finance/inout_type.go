package finance

import "github.com/kysion/base-library/utility/enum"

type InOutTypeEnum enum.IEnumCode[int]

type inOutType struct {
	Auto InOutTypeEnum
	In   InOutTypeEnum
	Out  InOutTypeEnum
}

var InOutType = inOutType{
	Auto: enum.New[InOutTypeEnum](0, "自动：正数收入，负数支出"),
	In:   enum.New[InOutTypeEnum](1, "收入"),
	Out:  enum.New[InOutTypeEnum](2, "支出"),
}

func (e inOutType) New(code int, description string) InOutTypeEnum {
	if (code&InOutType.Auto.Code()) == InOutType.Auto.Code() ||
		(code&InOutType.In.Code()) == InOutType.In.Code() ||
		(code&InOutType.Out.Code()) == InOutType.Out.Code() {
		return enum.New[InOutTypeEnum](code, description)
	}
	panic("uploadEventState: error")
}
