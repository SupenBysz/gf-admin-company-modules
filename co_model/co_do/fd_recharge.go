// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FdRecharge is the golang structure of table co_fd_recharge for DAO operations like Where/Data.
type FdRecharge struct {
	g.Meta         `orm:"table:co_fd_recharge, do:true"`
	Id             interface{} //
	UserId         interface{} // 用户ID，关联用户表主键，用于标识充值所属用户
	Username       interface{} // 用户账号，冗余存储方便快速查询用户相关充值记录，无需每次都关联用户表
	CurrencyCode   interface{} // 货币代码，如CNY（人民币）、USD（美元）等
	Amount         interface{} // 充值金额，小数点后保留2位，且金额需大于0
	RechargeMethod interface{} // 充值方式：1 - 银行卡；2 - 支付宝；3 - 微信；4 - 云闪付；5 - ApplePay；6 - PayPal；7 - AmazonPay；8 - 线下现金（若有对应业务）；9 - 区块链钱包；100 - 其他（可进一步在备注说明）
	PaymentAt      *gtime.Time // 充值支付时间，记录用户实际支付成功的时间
	PaymentOrderNo interface{} // 外部支付订单号，第三方支付平台生成的订单编号
	TransactionNo  interface{} // 交易流水号，系统内部生成，用于后续对账和查询
	State          interface{} // 充值状态：0 - 待处理；1 - 处理中；2 - 已支付；3 - 部分成功；4 - 失败；5 - 已取消；6 - 待确认
	AuditState     interface{} // 审核状态：0 - 待审核；1 - 审核通过；2 - 审核不通过；3 - 审核中（人工复审）；4 - 补充资料待审核
	AuditReply     interface{} // 审核意见，审核人员填写审核通过或不通过的原因等
	IpAddress      interface{} // 用户发起充值请求时的IP地址
	UserAgent      interface{} // 用户使用的设备和浏览器信息
	UnionMainId    interface{} //
	AccountId      interface{} // 财务账户
	Remark         interface{} // 备注，可记录一些特殊情况或额外信息
	ScreenshotId   interface{} // 充值截图
	CreatedAt      *gtime.Time // 记录创建时间，即充值请求提交时间
	UpdatedAt      *gtime.Time // 记录最后更新时间，每次记录状态等信息变更时更新
	DeletedAt      *gtime.Time // 逻辑删除时间，用于软删除，非真正物理删除，便于数据追溯和恢复
}
