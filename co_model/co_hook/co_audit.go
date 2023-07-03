package co_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

type AuditHookFunc func(ctx context.Context, state co_enum.AuditEvent, info co_entity.Audit) error
type AuditHookInfo struct {
	Key      co_enum.AuditEvent
	Value    AuditHookFunc
	Category int `json:"category" dc:"业务类别"`
}
