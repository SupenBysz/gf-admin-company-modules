package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type BankCardRegister struct {
	BankName      string      `json:"bankName"      v:"required#请输入银行卡名称" dc:"银行名称"`
	CardType      int         `json:"cardType"      v:"required|in:1,2#请输入银行类型" dc:"银行卡类型：1借记卡，2储蓄卡"`
	CardNumber    string      `json:"cardNumber"    v:"required|bank-card#请输入银行卡号|银行卡号错误" dc:"银行卡号"`
	ExpiredAt     *gtime.Time `json:"expiredAt"     dc:"有效期"`
	HolderName    string      `json:"holderName"    v:"required#请输入银行卡开户姓名" dc:"银行卡开户名"`
	UserId        int64       `json:"userId"        dc:"银行卡所属用户id，表示属于谁"`
	BankOfAccount string      `json:"bankOfAccount" dc:"开户行"`
	State         int         `json:"state"         dc:"状态：0禁用，1正常"`
	Remark        string      `json:"remark"        dc:"备注信息"`
}

type BankCardInfoRes co_entity.FdBankCard

type BankCardListRes base_model.CollectRes[co_entity.FdBankCard]
