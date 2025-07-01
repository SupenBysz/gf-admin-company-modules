package finance

import "github.com/kysion/base-library/utility/enum"

// TradeStateEnum 交易状态枚举
// 数值采用2的幂次方，支持按位组合
type TradeStateEnum interface {
	enum.IEnumCode[int]
}

type tradeState struct {
	None           TradeStateEnum
	WaitPay        TradeStateEnum
	Paying         TradeStateEnum
	Paid           TradeStateEnum
	Cancel         TradeStateEnum
	Completed      TradeStateEnum
	Refunding      TradeStateEnum
	Refunded       TradeStateEnum
	PaymentTimeout TradeStateEnum
	Closed         TradeStateEnum
	Frozen         TradeStateEnum // 新增：已冻结
	Unfrozen       TradeStateEnum // 新增：已解冻
	UnfreezeFailed TradeStateEnum // 新增：解冻失败

	// 状态映射表，用于快速查找
	codeMap map[int]TradeStateEnum
}

var TradeState = func() tradeState {
	ts := tradeState{
		None:           enum.New[TradeStateEnum](0, "无"),
		WaitPay:        enum.New[TradeStateEnum](1, "待支付"),
		Paying:         enum.New[TradeStateEnum](2, "支付中"),
		Paid:           enum.New[TradeStateEnum](4, "已支付"),
		Cancel:         enum.New[TradeStateEnum](8, "取消支付"),
		Completed:      enum.New[TradeStateEnum](16, "交易完成"),
		Refunding:      enum.New[TradeStateEnum](32, "退款中"),
		Refunded:       enum.New[TradeStateEnum](64, "已退款"),
		PaymentTimeout: enum.New[TradeStateEnum](128, "支付超时"),
		Closed:         enum.New[TradeStateEnum](256, "已关闭"),
		Frozen:         enum.New[TradeStateEnum](512, "已冻结"),
		Unfrozen:       enum.New[TradeStateEnum](1024, "已解冻"),
		UnfreezeFailed: enum.New[TradeStateEnum](2048, "解冻失败"),
		codeMap:        make(map[int]TradeStateEnum),
	}

	// 初始化映射表
	states := []TradeStateEnum{
		ts.None,
		ts.WaitPay, ts.Paying, ts.Paid, ts.Cancel,
		ts.Completed, ts.Refunding, ts.Refunded,
		ts.PaymentTimeout, ts.Closed, ts.Frozen,
		ts.Unfrozen, ts.UnfreezeFailed,
	}

	for _, state := range states {
		ts.codeMap[state.Code()] = state
	}

	return ts
}()

// New 根据代码获取交易状态，不存在则创建新的（谨慎使用）
func (e tradeState) New(code int, description string) TradeStateEnum {
	if state, exists := e.codeMap[code]; exists {
		return state
	}

	// 谨慎：通常不应该创建新的状态，这里保留原逻辑
	return enum.New[TradeStateEnum](code, description)
}

// Get 根据代码获取交易状态，不存在返回 nil
func (e tradeState) Get(code int) TradeStateEnum {
	return e.codeMap[code]
}
