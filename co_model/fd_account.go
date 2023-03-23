package co_model

import "github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"

type FdAccountRegister struct {
	Name           string `json:"name" v:"required#请输入财务账号名称"    dc:"账户名称"`
	UnionLicenseId int64  `json:"unionLicenseId"  dc:"关联资质ID，大于0时必须保值与 union_user_id 关联得上"`
	UnionUserId    int64  `json:"unionUserId"     v:"required#请输入财务账号关联的用户id"    dc:"关联用户ID"`
	// 货币标识后期还有积分标识
	CurrencyCode       string `json:"currencyCode"  v:"required|in:USD,HKD,TWD,JPY,CNY,EUR#请输入正确的货币代码"     dc:"货币代码:USD,HKD,TWD,JPY,CNY,EUR"`
	IsEnabled          int    `json:"isEnabled"          dc:"是否启用：1启用，0禁用"`
	LimitState         int    `json:"limitState"         dc:"限制状态：0不限制，1限制支出、2限制收入"`
	PrecisionOfBalance int    `json:"precisionOfBalance" v:"required#请输入财务账号货币单位精度" dc:"货币单位精度：1:元，10:角，100:分，1000:厘，10000:毫"`
	Version            int    `json:"version"            description:"乐观锁所需数据版本字段"`

	SceneType     int    `json:"sceneType"          description:"场景类型：0不限制、1充电佣金收入、"`
	AccountType   int    `json:"accountType"        description:"账户类型：1系统账户、2银行卡、3支付宝、4微信、5云闪付、6翼支付"`
	AccountNumber string `json:"accountNumber"      description:"账户编号，例如银行卡号、支付宝账号、微信账号等对应账户类型的编号"`
}

type FdAccountRes struct {
	co_entity.FdAccount
}

func (m *FdAccountRes) Data() *FdAccountRes {
	return m
}

type IFdAccountRes interface {
	Data() *FdAccountRes
}
