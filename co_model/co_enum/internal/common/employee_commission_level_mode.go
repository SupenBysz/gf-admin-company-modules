package common

import "github.com/kysion/base-library/utility/enum"

// EmployeeCommissionLevelModeEnum 员工提成级别模式
// 仅作用于员工：1师徒/邀请模式、2部门/团队/小组模式、3角色模式
// 注意：一旦选定模式后，在一个统计周期内不能修改，否则可能会引起报表统计错误(但不影响实际收益人的佣金结算)，在下一个统计周期报表才能正常
type EmployeeCommissionLevelModeEnum enum.IEnumCode[int]

type employeeCommissionLevelMode struct {
	Superior        EmployeeCommissionLevelModeEnum
	TradeAmount     EmployeeCommissionLevelModeEnum
	TradeCommission EmployeeCommissionLevelModeEnum
}

var EmployeeCommissionLevelMode = employeeCommissionLevelMode{
	Superior:        enum.New[EmployeeCommissionLevelModeEnum](0, "相较于上级提成比例"),
	TradeAmount:     enum.New[EmployeeCommissionLevelModeEnum](1, "相较于交易金额百分比"),
	TradeCommission: enum.New[EmployeeCommissionLevelModeEnum](2, "相较于交易佣金百分比"),
}

func (e employeeCommissionLevelMode) New(code int, description string) EmployeeCommissionLevelModeEnum {
	if code == EmployeeCommissionLevelMode.Superior.Code() {
		return EmployeeCommissionLevelMode.Superior
	}
	if code == EmployeeCommissionLevelMode.TradeAmount.Code() {
		return EmployeeCommissionLevelMode.TradeAmount
	}

	return enum.New[EmployeeCommissionLevelModeEnum](code, description)
}
