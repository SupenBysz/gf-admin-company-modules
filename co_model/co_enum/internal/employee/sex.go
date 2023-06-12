package employee

import "github.com/kysion/base-library/utility/enum"

type SexEnum enum.IEnumCode[int]

type sex struct {
	In  SexEnum
	Out SexEnum
}

var Sex = sex{
	In:  enum.New[SexEnum](0, "女"),
	Out: enum.New[SexEnum](1, "男"),
}

func (e sex) New(code int, description string) SexEnum {
	// 0 进行运算都是true
	if code == Sex.In.Code() {
		return e.In
	}
	if code == Sex.Out.Code() {
		return e.Out
	}
	return enum.New[SexEnum](code, description)
}
