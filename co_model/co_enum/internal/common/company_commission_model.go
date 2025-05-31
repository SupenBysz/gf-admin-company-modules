package common

import "github.com/kysion/base-library/utility/enum"

// CompanyCommissionModeEnum 公司佣金模式
// 仅作用于公司超管财务账户，即公司财务账户，一般场景为机构、代理、子公司、加盟等关系业务场景
// 即仅作用于公司主体，0:不启用佣金机制，1：相对上级佣金百分比，2：相对交易金额百分比,3:相较于交易佣金百分比。
// 注意：一旦选定模式后，在一个统计周期内不能修改，否则会引起统计错误(但不影响实际收益人的佣金结算)，在下一个统计周期报表才能正常
type CompanyCommissionModeEnum enum.IEnumCode[int]

type companyCommissionMode struct {
	Superior    CompanyCommissionModeEnum
	TradeAmount CompanyCommissionModeEnum
}

var CompanyCommissionMode = companyCommissionMode{
	Superior:    enum.New[CompanyCommissionModeEnum](0, "相对于上级佣金收益"),
	TradeAmount: enum.New[CompanyCommissionModeEnum](1, "相对于交易金额"),
}

func (e companyCommissionMode) New(code int, description string) CompanyCommissionModeEnum {
	if code == CompanyCommissionMode.Superior.Code() {
		return CompanyCommissionMode.Superior
	}
	if code == CompanyCommissionMode.TradeAmount.Code() {
		return CompanyCommissionMode.TradeAmount
	}

	return enum.New[CompanyCommissionModeEnum](code, description)
}
