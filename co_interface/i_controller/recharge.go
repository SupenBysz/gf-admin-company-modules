package i_controller

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type IFdRecharge[TR co_model.IFdRechargeRes] interface {
	GetAccountRechargeById(ctx context.Context, req *co_company_api.GetAccountRechargeByIdReq) (TR, error)
	SetAccountRechargeAudit(ctx context.Context, req *co_company_api.SetAccountRechargeAuditReq) (api_v1.BoolRes, error)
	QueryAccountRecharge(ctx context.Context, req *co_company_api.QueryAccountRechargeReq) (*base_model.CollectRes[TR], error)
	AccountRecharge(ctx context.Context, req *co_company_api.AccountRecharge) (TR, error)
}
