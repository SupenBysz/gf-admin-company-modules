package invoice

import "github.com/kysion/base-library/utility/enum"

type BelongTypeEnum enum.IEnumCode[int]

type belongType struct {
	OnSelf  BelongTypeEnum
	Subject BelongTypeEnum
}

var BelongType = belongType{
	OnSelf:  enum.New[BelongTypeEnum](1, "个人"),
	Subject: enum.New[BelongTypeEnum](2, "主体"),
}

func (e belongType) New(code int, description string) BelongTypeEnum {
	if (code&BelongType.OnSelf.Code()) == BelongType.OnSelf.Code() ||
		(code&BelongType.Subject.Code()) == BelongType.Subject.Code() {
		return enum.New[BelongTypeEnum](code, description)
	} else {
		panic("kyBelong.Type.New: error")
	}
}
