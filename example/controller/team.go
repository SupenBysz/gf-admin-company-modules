package controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/utility/kconv"
)

type TeamController[
	ITTeamRes co_model.ITeamRes,
] struct {
	i_controller.ITeam[ITTeamRes]
}

func Team[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) *TeamController[ITTeamRes] {
	return &TeamController[ITTeamRes]{
		ITeam: co_controller.Team(modules),
	}
}

func (c *TeamController[ITTeamRes]) GetTeamById(ctx context.Context, req *co_v1.GetTeamByIdReq) (ITTeamRes, error) {
	return c.ITeam.GetTeamById(ctx, &req.GetTeamByIdReq)
}

func (c *TeamController[ITTeamRes]) HasTeamByName(ctx context.Context, req *co_v1.HasTeamByNameReq) (api_v1.BoolRes, error) {
	return c.ITeam.HasTeamByName(ctx, &req.HasTeamByNameReq)
}

func (c *TeamController[ITTeamRes]) QueryTeamList(ctx context.Context, req *co_v1.QueryTeamListReq) (*api_v1.MapRes, error) {
	ret, err := c.ITeam.QueryTeamList(ctx, &req.QueryTeamListReq)
	return kconv.Struct(ret, &api_v1.MapRes{}), err
}

func (c *TeamController[ITTeamRes]) CreateTeam(ctx context.Context, req *co_v1.CreateTeamReq) (ITTeamRes, error) {
	return c.ITeam.CreateTeam(ctx, &req.CreateTeamReq)
}

func (c *TeamController[ITTeamRes]) UpdateTeam(ctx context.Context, req *co_v1.UpdateTeamReq) (ITTeamRes, error) {
	return c.ITeam.UpdateTeam(ctx, &req.UpdateTeamReq)
}

func (c *TeamController[ITTeamRes]) DeleteTeam(ctx context.Context, req *co_v1.DeleteTeamReq) (api_v1.BoolRes, error) {
	return c.ITeam.DeleteTeam(ctx, &req.DeleteTeamReq)
}

func (c *TeamController[ITTeamRes]) QueryTeamListByEmployee(ctx context.Context, req *co_v1.QueryTeamListByEmployeeReq) (*api_v1.MapRes, error) {
	ret, err := c.ITeam.QueryTeamListByEmployee(ctx, &req.QueryTeamListByEmployeeReq)
	return kconv.Struct(ret, &api_v1.MapRes{}), err
}

func (c *TeamController[ITTeamRes]) SetTeamMember(ctx context.Context, req *co_v1.SetTeamMemberReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamMember(ctx, &req.SetTeamMemberReq)
}

func (c *TeamController[ITTeamRes]) SetTeamOwner(ctx context.Context, req *co_v1.SetTeamOwnerReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamOwner(ctx, &req.SetTeamOwnerReq)
}

func (c *TeamController[ITTeamRes]) SetTeamCaptain(ctx context.Context, req *co_v1.SetTeamCaptainReq) (api_v1.BoolRes, error) {
	return c.ITeam.SetTeamCaptain(ctx, &req.SetTeamCaptainReq)
}

func (c *TeamController[ITTeamRes]) GetEmployeeListByTeamId(ctx context.Context, req *co_v1.GetEmployeeListByTeamIdReq) (*api_v1.MapRes, error) {
	ret, err := c.ITeam.GetEmployeeListByTeamId(ctx, &req.GetEmployeeListByTeamIdReq)

	return kconv.Struct(ret, &api_v1.MapRes{}), err
}
