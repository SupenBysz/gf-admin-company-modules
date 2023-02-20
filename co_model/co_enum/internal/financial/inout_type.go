package financial

import "github.com/kysion/base-library/utility/enum"

type InOutTypeEnum enum.IEnumCode[int]

type inOutType struct {
	In  InOutTypeEnum
	Out InOutTypeEnum
}

var InOutType = inOutType{
	In:  enum.New[InOutTypeEnum](1, "收入"),
	Out: enum.New[InOutTypeEnum](2, "支出"),
}

func (e inOutType) New(code int, description string) InOutTypeEnum {
	if (code&InOutType.In.Code()) == InOutType.In.Code() ||
		(code&InOutType.Out.Code()) == InOutType.Out.Code() {
		return enum.New[InOutTypeEnum](code, description)
	}
	panic("uploadEventState: error")
}
