package finance

import "github.com/kysion/base-library/utility/enum"

type RechargeMethodEnum enum.IEnumCode[int]

type rechargeMethod struct {
	BankCard RechargeMethodEnum
	Alipay   RechargeMethodEnum
	WeChat   RechargeMethodEnum
	CloudPay RechargeMethodEnum
	Cash     RechargeMethodEnum
	Other    RechargeMethodEnum
}

var RechargeMethod = rechargeMethod{
	BankCard: enum.New[RechargeMethodEnum](1, "银行卡"),
	Alipay:   enum.New[RechargeMethodEnum](2, "支付宝"),
	WeChat:   enum.New[RechargeMethodEnum](3, "微信"),
	CloudPay: enum.New[RechargeMethodEnum](4, "云闪付"),
	Cash:     enum.New[RechargeMethodEnum](5, "线下现金"),
	Other:    enum.New[RechargeMethodEnum](6, "其他"),
}

func (e rechargeMethod) New(code int, description string) RechargeMethodEnum {
	switch code {
	case e.BankCard.Code():
		return e.BankCard
	case e.Alipay.Code():
		return e.Alipay
	case e.WeChat.Code():
		return e.WeChat
	case e.CloudPay.Code():
		return e.CloudPay
	case e.Cash.Code():
		return e.Cash
	case e.Other.Code():
		return e.Other
	default:
		panic("RechargeMethod: error")
	}
}
