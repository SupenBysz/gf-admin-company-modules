package invoice

import "github.com/kysion/base-library/utility/enum"

type AuditTypeEnum enum.IEnumCode[int]

type auditType struct {
	Reject     AuditTypeEnum
	WaitReview AuditTypeEnum
	Approve    AuditTypeEnum
}

var AuditType = auditType{
	WaitReview: enum.New[AuditTypeEnum](0, "待审核"),
	Approve:    enum.New[AuditTypeEnum](1, "通过"),
	Reject:     enum.New[AuditTypeEnum](-1, "不通过"),
}

func (e auditType) New(code int, description string) AuditTypeEnum {
	if (code&AuditType.Reject.Code()) == AuditType.Reject.Code() ||
		(code&AuditType.WaitReview.Code()) == AuditType.WaitReview.Code() ||
		(code&AuditType.Approve.Code()) == AuditType.Approve.Code() {
		return enum.New[AuditTypeEnum](code, description)
	} else {
		panic("kyAudit.Action.New: error")
	}
}
