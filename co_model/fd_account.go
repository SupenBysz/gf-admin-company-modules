package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/base_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type FdAccountRegister struct {
	Name        string `json:"name" v:"required#请输入财务账号名称"    dc:"账户开户名称，姓名或者是单位全称"`
	UnionUserId int64  `json:"unionUserId"        dc:"关联用户ID"`
	UnionMainId int64  `json:"unionMainId"         dc:"关联主体ID"`

	// 货币标识后期还有积分标识
	CurrencyCode       string `json:"currencyCode"  v:"required|in:USD,HKD,TWD,JPY,CNY,EUR#货币代码不能为空|请输入正确的货币代码"     dc:"货币代码:USD,HKD,TWD,JPY,CNY,EUR"`
	PrecisionOfBalance int    `json:"precisionOfBalance" v:"required#请输入财务账号货币单位精度" dc:"货币单位精度：1:元，10:角，100:分，1000:厘，10000:毫"`

	SceneType     int    `json:"sceneType"          dc:"场景类型：0不限制、1充电佣金收入、"`
	AccountType   int    `json:"accountType"       v:"required|in:1,2,3,4,5,6#请输入正确的货币代码|账户类型限定为：1系统帐户、2银行卡、3支付宝、4微信、5云闪付、6翼支付" dc:"账户类型：1系统账户、2银行卡、3支付宝、4微信、5云闪付、6翼支付"`
	AccountNumber string `json:"accountNumber"      dc:"账户编号，例如银行卡号、支付宝账号、微信账号等对应账户类型的编号"`
	AllowExceed int `json:"allowExceed"        dc:"是否允许存在负余额: 0禁止、1允许" v:"in:0,1#是否允许存在负余额数据校验失败"`
}

type UpdateAccount struct {
	OverrideDo base_interface.DoModel[co_do.FdAccount] `json:"-"`

	Name          *string `json:"name" dc:"账户开户名称，姓名或者是单位全称"`
	AccountType   *int    `json:"accountType"       v:"required|in:1,2,3,4,5,6#请输入正确的货币代码|账户类型限定为：1系统帐户、2银行卡、3支付宝、4微信、5云闪付、6翼支付" dc:"账户类型：1系统账户、2银行卡、3支付宝、4微信、5云闪付、6翼支付"`
	AccountNumber *string `json:"accountNumber"      dc:"账户编号，例如银行卡号、支付宝账号、微信账号等对应账户类型的编号"`
	AllowExceed   *int    `json:"allowExceed"        dc:"是否允许存在负余额: 0禁止、1允许" v:"in:0,1#是否允许存在负余额数据校验失败"`
}

type FdAccountListRes base_model.CollectRes[FdAccountRes]

type FdAccountRes struct {
	co_entity.FdAccount
	Detail *co_entity.FdAccountDetail `json:"detail" dc:"财务账号周期金额统计"`
}

func (m *FdAccountRes) Data() *FdAccountRes {
	return m
}

type IFdAccountRes interface {
	Data() *FdAccountRes
}
