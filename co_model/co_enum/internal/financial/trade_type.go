package financial

import "github.com/kysion/base-library/utility/enum"

type TradeTypeEnum enum.IEnumCode[int]

type tradeType struct {
	Transfer        TradeTypeEnum
	Consumption     TradeTypeEnum
	Refund          TradeTypeEnum
	Commission      TradeTypeEnum
	SecurityDeposit TradeTypeEnum
	EarnestMoney    TradeTypeEnum
	ServiceCharge   TradeTypeEnum
	CashWithdrawal  TradeTypeEnum
	Recharge        TradeTypeEnum
	OperatingIncome TradeTypeEnum
	Other           TradeTypeEnum
}

var TradeType = tradeType{
	Transfer:        enum.New[TradeTypeEnum](1, "转账"),
	Consumption:     enum.New[TradeTypeEnum](2, "消费"),
	Refund:          enum.New[TradeTypeEnum](4, "退款"),
	Commission:      enum.New[TradeTypeEnum](8, "佣金"),
	SecurityDeposit: enum.New[TradeTypeEnum](16, "保证金"),
	EarnestMoney:    enum.New[TradeTypeEnum](32, "诚意金"),
	ServiceCharge:   enum.New[TradeTypeEnum](64, "手续费/服务费"),
	CashWithdrawal:  enum.New[TradeTypeEnum](128, "提现"),
	Recharge:        enum.New[TradeTypeEnum](256, "充值"),
	OperatingIncome: enum.New[TradeTypeEnum](512, "营收"),
	Other:           enum.New[TradeTypeEnum](8192, "其它"),
}

func (e tradeType) New(code int, description string) TradeTypeEnum {
	if (code&TradeType.Transfer.Code()) == TradeType.Transfer.Code() ||
		(code&TradeType.Consumption.Code()) == TradeType.Consumption.Code() ||
		(code&TradeType.Refund.Code()) == TradeType.Refund.Code() ||
		(code&TradeType.Commission.Code()) == TradeType.Commission.Code() ||
		(code&TradeType.SecurityDeposit.Code()) == TradeType.SecurityDeposit.Code() ||
		(code&TradeType.EarnestMoney.Code()) == TradeType.EarnestMoney.Code() ||
		(code&TradeType.ServiceCharge.Code()) == TradeType.ServiceCharge.Code() ||
		(code&TradeType.CashWithdrawal.Code()) == TradeType.CashWithdrawal.Code() ||
		(code&TradeType.Recharge.Code()) == TradeType.Recharge.Code() ||
		(code&TradeType.OperatingIncome.Code()) == TradeType.OperatingIncome.Code() ||
		(code&TradeType.Other.Code()) == TradeType.Other.Code() {
		return enum.New[TradeTypeEnum](code, description)
	}

	return enum.New[TradeTypeEnum](code, description)

	//panic("kyFinancial.TradeType.New: error")
}
