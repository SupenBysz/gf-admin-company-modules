package invoice

import "github.com/kysion/base-library/utility/enum"

type TypeEnum enum.IEnumCode[int]

type makeType struct {
	Normal       TypeEnum
	Special      TypeEnum
	professional TypeEnum
}

var MakeType = makeType{
	Normal:       enum.New[TypeEnum](1, "普通发票"),
	Special:      enum.New[TypeEnum](2, "增值税专用发票"),
	professional: enum.New[TypeEnum](3, "专业发票"),
}

func (e makeType) New(code int, description string) TypeEnum {
	if (code&MakeType.Normal.Code()) == MakeType.Normal.Code() ||
		(code&MakeType.Special.Code()) == MakeType.Special.Code() ||
		(code&MakeType.professional.Code()) == MakeType.professional.Code() {
		return enum.New[TypeEnum](code, description)
	} else {
		panic("kyInvoice.Type.New: error")
	}
}
