package financial

import "github.com/kysion/base-library/utility/enum"

// AccountTypeEnum 账户类型：1系统账户、2银行卡、3支付宝、4微信、5云闪付
type AccountTypeEnum enum.IEnumCode[int]

type accountType struct {
    System   AccountTypeEnum
    BankCard AccountTypeEnum
    Alipay   AccountTypeEnum
    WeiXin   AccountTypeEnum
    UnionPay AccountTypeEnum
}

var AccountType = accountType{
    System:   enum.New[AccountTypeEnum](1, "系统账户"),
    BankCard: enum.New[AccountTypeEnum](2, "银行卡"),
    Alipay:   enum.New[AccountTypeEnum](3, "支付宝"),
    WeiXin:   enum.New[AccountTypeEnum](4, "微信"),
    UnionPay: enum.New[AccountTypeEnum](5, "云闪付"),
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

    return enum.New[AccountTypeEnum](code, description)

}
