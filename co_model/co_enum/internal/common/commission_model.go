package common

import "github.com/kysion/base-library/utility/enum"

// CommissionModeEnum 申诉状态枚举
type CommissionModeEnum enum.IEnumCode[int]

type commissionMode struct {
	Superior    CommissionModeEnum
	TradeAmount CommissionModeEnum
}

var CommissionMode = commissionMode{
	Superior:    enum.New[CommissionModeEnum](0, "相对于上级佣金收益"),
	TradeAmount: enum.New[CommissionModeEnum](1, "相对于交易金额"),
}

func (e commissionMode) New(code int, description string) CommissionModeEnum {
	if code == CommissionMode.Superior.Code() {
		return CommissionMode.Superior
	}
	if code == CommissionMode.TradeAmount.Code() {
		return CommissionMode.TradeAmount
	}

	return enum.New[CommissionModeEnum](code, description)
}
