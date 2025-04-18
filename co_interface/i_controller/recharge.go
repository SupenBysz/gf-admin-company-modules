package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type IFdRecharge[TR co_model.IFdRechargeRes] interface {
	GetAccountRechargeById(ctx context.Context, id int64) (TR, error)
	SetAccountRechargeAudit(ctx context.Context, id int64, state sys_enum.AuditAction, reply string) (api_v1.BoolRes, error)
	QueryAccountRecharge(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error)
	AccountRecharge(ctx context.Context, info *co_model.FdRecharge) (TR, error)
}
