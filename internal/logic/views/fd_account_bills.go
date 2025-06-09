package views

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
)

type sFdBillsView struct {
}

func init() {
	co_service.RegisterFdBillsView(NewFdBillsView())
}

func NewFdBillsView() co_service.IFdBillsView {
	return &sFdBillsView{}
}

func (s *sFdBillsView)  GetBillsById(ctx context.Context, id int64, makeResource bool) (*co_model.CompanyBillsViewRes, error) {
	data, err := daoctl.GetByIdWithError[co_model.CompanyBillsViewRes](co_dao.FdAccountBillsView.Ctx(ctx), id)

	if data == nil || err != nil {
		return nil, err
	}

	return s.makeBillsMore(ctx, data, makeResource), nil
}

// QueryBillsList 公司账单账单查询
func (s *sFdBillsView) QueryBillsList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.CompanyBillsViewListRes, error) {
	data, err := daoctl.Query[co_model.CompanyBillsViewRes](co_dao.FdAccountBillsView.Ctx(ctx), params, false)

	if data == nil || err != nil {
		return nil, err
	}

	result := &co_model.CompanyBillsViewListRes{
		Records:       data.Records,
		PaginationRes: data.PaginationRes,
	}

	if len(data.Records) > 0 && makeResource {
		for i, record := range data.Records {
			data := s.makeBillsMore(ctx, &record, makeResource)
			result.Records[i] = co_model.CompanyBillsViewRes(*data)
		}
	}

	return result, err
}

func (s *sFdBillsView) makeBillsMore(ctx context.Context, data *co_model.CompanyBillsViewRes, makeResource bool) *co_model.CompanyBillsViewRes {
	if !makeResource {
		return data
	}

	if data.FromUserId > 0 {
		fromUser, _ := sys_service.SysUser().GetSysUserById(ctx, data.FromUserId)
		if fromUser != nil {
			data.FormUser = fromUser
		}
	}
	if data.ToUserId > 0 {
		toUser, _ := sys_service.SysUser().GetSysUserById(ctx, data.ToUserId)
		if toUser != nil {
			data.ToUser = toUser
		}
	}
	if data.FdAccountId > 0 {
		fdAccount := daoctl.GetById[co_do.FdAccount](co_dao.FdAccountView.Ctx(ctx), data.FdAccountId)
		data.FdAccount = fdAccount
	}

	return data
}
