package finance

import "github.com/kysion/base-library/utility/enum"

// AccountTypeEnum 账户类型：1系统账户、2银行卡、3支付宝、4微信、5云闪付
type AccountTypeEnum enum.IEnumCode[int]

type accountType struct {
	System     AccountTypeEnum
	BankCard   AccountTypeEnum
	Alipay     AccountTypeEnum
	WeiXin     AccountTypeEnum
	UnionPay   AccountTypeEnum
	ApplePay   AccountTypeEnum
	PayPal     AccountTypeEnum
	AmazonPay  AccountTypeEnum
	Cash       AccountTypeEnum
	Blockchain AccountTypeEnum
	Other      AccountTypeEnum
}

var AccountType = accountType{
	System:     enum.New[AccountTypeEnum](0, "系统账户"),
	BankCard:   enum.New[AccountTypeEnum](1, "银行卡"),
	Alipay:     enum.New[AccountTypeEnum](2, "支付宝"),
	WeiXin:     enum.New[AccountTypeEnum](3, "微信"),
	UnionPay:   enum.New[AccountTypeEnum](4, "云闪付"),
	ApplePay:   enum.New[AccountTypeEnum](5, "苹果支付"),
	PayPal:     enum.New[AccountTypeEnum](6, "PayPal"),
	AmazonPay:  enum.New[AccountTypeEnum](7, "AmazonPay"),
	Cash:       enum.New[AccountTypeEnum](8, "现金"),
	Blockchain: enum.New[AccountTypeEnum](9, "区块链"),
	Other:      enum.New[AccountTypeEnum](100, "其他"),
}

func (e accountType) New(code int, description string) AccountTypeEnum {
	if code == AccountType.System.Code() {
		return AccountType.System
	}
	if code == AccountType.BankCard.Code() {
		return AccountType.BankCard
	}
	if code == AccountType.Alipay.Code() {
		return AccountType.Alipay
	}
	if code == AccountType.WeiXin.Code() {
		return AccountType.WeiXin
	}
	if code == AccountType.UnionPay.Code() {
		return AccountType.UnionPay
	}
	if code == AccountType.ApplePay.Code() {
		return AccountType.ApplePay
	}
	if code == AccountType.PayPal.Code() {
		return AccountType.PayPal
	}
	if code == AccountType.AmazonPay.Code() {
		return AccountType.AmazonPay
	}
	if code == AccountType.Cash.Code() {
		return AccountType.Cash
	}
	if code == AccountType.Blockchain.Code() {
		return AccountType.Blockchain
	}
	if code == AccountType.Other.Code() {
		return AccountType.Other
	}

	return enum.New[AccountTypeEnum](code, description)
}
