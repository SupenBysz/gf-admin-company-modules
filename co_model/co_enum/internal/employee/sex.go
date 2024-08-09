package employee

import "github.com/kysion/base-library/utility/enum"

//用户性别: 0未知、1男、2女

type SexEnum enum.IEnumCode[int]

type userSex struct {
	Unknown SexEnum
	Woman   SexEnum
	Man     SexEnum
}

var Sex = userSex{
	Unknown: enum.New[SexEnum](0, "未知"),
	Man:     enum.New[SexEnum](1, "男"),
	Woman:   enum.New[SexEnum](2, "女"),
}

func (e userSex) New(code int) SexEnum {
	if code == Sex.Unknown.Code() {
		return Sex.Unknown
	}

	if code == Sex.Man.Code() {
		return Sex.Man
	}

	if code == Sex.Woman.Code() {
		return Sex.Woman
	}

	panic("User.Sex.New: error")
}
