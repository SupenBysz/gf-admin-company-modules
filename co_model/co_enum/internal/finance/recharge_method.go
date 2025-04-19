package finance

import "github.com/kysion/base-library/utility/enum"

type RechargeMethodEnum enum.IEnumCode[int]

type rechargeMethod struct {
	BankCard  RechargeMethodEnum
	Alipay    RechargeMethodEnum
	WeChat    RechargeMethodEnum
	CloudPay  RechargeMethodEnum
	ApplePay  RechargeMethodEnum
	PayPal    RechargeMethodEnum
	AmazonPay RechargeMethodEnum
	Cash      RechargeMethodEnum
	Dapp      RechargeMethodEnum
	Other     RechargeMethodEnum
}

var RechargeMethod = rechargeMethod{
	BankCard:  enum.New[RechargeMethodEnum](1, "银行卡"),
	Alipay:    enum.New[RechargeMethodEnum](2, "支付宝"),
	WeChat:    enum.New[RechargeMethodEnum](3, "微信"),
	CloudPay:  enum.New[RechargeMethodEnum](4, "云闪付"),
	ApplePay:  enum.New[RechargeMethodEnum](5, "Apple Pay"),
	PayPal:    enum.New[RechargeMethodEnum](6, "PayPal"),
	AmazonPay: enum.New[RechargeMethodEnum](7, "亚马逊支付"),
	Cash:      enum.New[RechargeMethodEnum](8, "线下现金"),
	Dapp:      enum.New[RechargeMethodEnum](9, "区块链钱包"),
	Other:     enum.New[RechargeMethodEnum](100, "其他"),
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
	case e.ApplePay.Code():
		return e.ApplePay
	case e.PayPal.Code():
		return e.PayPal
	case e.AmazonPay.Code():
		return e.AmazonPay
	case e.Cash.Code():
		return e.Cash
	case e.Dapp.Code():
		return e.Dapp
	case e.Other.Code():
		return e.Other
	default:
		panic("RechargeMethod: error")
	}
}
