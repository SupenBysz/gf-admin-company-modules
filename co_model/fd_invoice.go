package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type FdInvoiceRegister struct {
	Name          string      `json:"name"           v:"required#请输入发票抬头名称" dc:"发票抬头名称"`
	TaxId         string      `json:"taxId"          v:"required#请输入纳税识别号" dc:"纳税识别号"`
	Addr          string      `json:"addr"           dc:"发票收件地址，限纸质"`
	Email         string      `json:"email"          dc:"发票收件邮箱，限电子发票"`
	UserId        int64       `json:"userId"         dc:"申请人UserID"`
	AuditUserId   int64       `json:"auditUserId"    dc:"审核人UserID"`
	AuditReplyMsg string      `json:"auditReplyMsg" dc:"审核回复，仅审核不通过时才有值"`
	AuditAt       *gtime.Time `json:"auditAt"        dc:"审核时间"`
	State         int         `json:"state"          dc:"状态：0待审核、1已通过、-1不通过"`
	UnionMainId   int64       `json:"unionMainId"    dc:"主体ID：运营商ID、服务商ID、商户ID、消费者ID"`
}

type FdInvoiceRes struct {
	co_entity.FdInvoice
}

type FdInvoiceListRes base_model.CollectRes[FdInvoiceRes]

func (m *FdInvoiceRes) Data() *FdInvoiceRes {
	return m
}

type IFdInvoiceRes interface {
	Data() *FdInvoiceRes
}
