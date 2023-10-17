package financial

import "github.com/kysion/base-library/utility/enum"

// 是否允许存在负数余额：0禁止、1允许 （可支配金额，是否允许超出账户余额）

type AllowExceedEnum enum.IEnumCode[int]

type allowExceed struct {
	Disabled AllowExceedEnum
	Allow    AllowExceedEnum
}

var AllowExceed = allowExceed{
	Disabled: enum.New[AllowExceedEnum](0, "禁止"),
	Allow:    enum.New[AllowExceedEnum](1, "允许"),
}

func (e allowExceed) New(code int, description string) AllowExceedEnum {
	if (code&AllowExceed.Disabled.Code()) == AllowExceed.Disabled.Code() ||
		(code&AllowExceed.Allow.Code()) == AllowExceed.Allow.Code() {
		return enum.New[AllowExceedEnum](code, description)
	}
	panic("uploadEventState: error")
}
