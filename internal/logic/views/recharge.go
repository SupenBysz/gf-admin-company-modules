package views

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SupenBysz/gf-admin-company-modules/co_service"

	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
)

type sFdRechargeView struct {
}

func init() {
	co_service.RegisterFdRechargeView(NewFdRechargeView())
}

func NewFdRechargeView() co_service.IFdRechargeView {
	return &sFdRechargeView{}
}

// QueryAccountRecharge 查询充值记录列表
func (s *sFdRechargeView) QueryAccountRecharge(ctx context.Context, search *base_model.SearchParams) (*co_model.FdRechargeViewListRes, error) {
	isExport := false
	if ctx.Value("isExport") == nil {
		r := g.RequestFromCtx(ctx)
		isExport = r.GetForm("isExport", false).Bool()
	} else {
		isExport = gconv.Bool(ctx.Value("isExport"))
	}

	data, err := daoctl.Query[co_model.FdRechargeViewRes](co_dao.FdRechargeView.Ctx(ctx).OrderDesc(co_dao.FdRechargeView.Columns().PaymentAt), search, isExport)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &co_model.FdRechargeViewListRes{}, nil
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_Data_Get_Failed}", co_dao.FdRechargeView.Table())
	}

	result := &co_model.FdRechargeViewListRes{}
	_ = gconv.Struct(data, &result)

	return result, nil
}

// GetAccountRechargeById 根据资金账户ID获取充值记录
func (s *sFdRechargeView) GetAccountRechargeById(ctx context.Context, id int64) (*co_model.FdRechargeViewRes, error) {

	if id == 0 {
		return nil, gerror.New("error_Financial_RechargeId_Failed")
	}

	result, err := daoctl.GetByIdWithError[co_model.FdRechargeViewRes](co_dao.FdRechargeView.Ctx(ctx), id)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_Data_Get_Failed}", co_dao.FdRechargeView.Table())
	}

	return result, nil
}

// GetRechargeByAccountId 根据资金账户ID获取充值记录
func (s *sFdRechargeView) GetRechargeByAccountId(ctx context.Context, id int64) (*co_model.FdRechargeViewListRes, error) {
	if id == 0 {
		return nil, gerror.New("error_Financial_AccountId_Failed")
	}

	data, err := daoctl.Query[co_model.FdRechargeViewRes](co_dao.FdRechargeView.Ctx(ctx).
		Where(co_dao.FdRechargeView.Columns().AccountId, id).
		OrderDesc(co_dao.FdRechargeView.Columns().CreatedAt),
		nil,
		true,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_Data_Get_Failed}", co_dao.FdRechargeView.Table())
	}

	result := &co_model.FdRechargeViewListRes{}
	_ = gconv.Struct(data, &result)

	return result, nil
}

// GetRechargeByUserId 根据用户ID获取充值记录
func (s *sFdRechargeView) GetRechargeByUserId(ctx context.Context, id int64) (*co_model.FdRechargeViewListRes, error) {
	if id == 0 {
		return nil, gerror.New("error_Financial_UserId_Failed")
	}

	data, err := daoctl.Query[co_model.FdRechargeViewRes](co_dao.FdRechargeView.Ctx(ctx).
		Where(co_dao.FdRechargeView.Columns().UserId, id).
		OrderDesc(co_dao.FdRechargeView.Columns().CreatedAt),
		nil,
		true,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_Data_Get_Failed}", co_dao.FdRechargeView.Table())
	}

	result := &co_model.FdRechargeViewListRes{}
	_ = gconv.Struct(data, &result)

	return result, nil
}

// GetRechargeByCompanyId 根据公司ID获取充值记录
func (s *sFdRechargeView) GetRechargeByCompanyId(ctx context.Context, id int64) (*co_model.FdRechargeViewListRes, error) {
	if id == 0 {
		return nil, gerror.New("error_Financial_CompanyId_Failed")
	}

	data, err := daoctl.Query[co_model.FdRechargeViewRes](co_dao.FdRechargeView.Ctx(ctx).
		Where(co_dao.FdRechargeView.Columns().UnionMainId, id).
		OrderDesc(co_dao.FdRechargeView.Columns().CreatedAt),
		nil,
		true,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_Data_Get_Failed}", co_dao.FdRechargeView.Table())
	}

	result := &co_model.FdRechargeViewListRes{}
	_ = gconv.Struct(data, &result)

	return result, nil
}
