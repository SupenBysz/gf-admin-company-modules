package common

import "github.com/kysion/base-library/utility/enum"

// EmployeeCommissionModeEnum 员工提成模式
// 仅作用于员工，0不启用提成机制，1:相较于上级，2:相较于相较于交易金额百分比,3:相较于交易佣金百分比
// 注意：一旦选定模式后，在一个统计周期内不能修改，否则会引起统计错误，在下一个统计周期才能正常）
type EmployeeCommissionModeEnum enum.IEnumCode[int]

type employeeCommissionMode struct {
	Superior        EmployeeCommissionModeEnum
	TradeAmount     EmployeeCommissionModeEnum
	TradeCommission EmployeeCommissionModeEnum
}

var EmployeeCommissionMode = employeeCommissionMode{
	Superior:        enum.New[EmployeeCommissionModeEnum](0, "相较于上级提成比例"),
	TradeAmount:     enum.New[EmployeeCommissionModeEnum](1, "相较于交易金额百分比"),
	TradeCommission: enum.New[EmployeeCommissionModeEnum](2, "相较于交易佣金百分比"),
}

func (e employeeCommissionMode) New(code int, description string) EmployeeCommissionModeEnum {
	if code == EmployeeCommissionMode.Superior.Code() {
		return EmployeeCommissionMode.Superior
	}
	if code == EmployeeCommissionMode.TradeAmount.Code() {
		return EmployeeCommissionMode.TradeAmount
	}

	return enum.New[EmployeeCommissionModeEnum](code, description)
}
