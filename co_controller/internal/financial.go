package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

// FinancialController 财务服务控制器
type FinancialController struct {
	i_controller.IMy
	modules co_interface.IModules
}

// Financial 财务服务
var Financial = func(modules co_interface.IModules) i_controller.IFinancial {
	return &FinancialController{
		modules: modules,
	}
}

// GetAccountBalance 查看账户余额
func (c *FinancialController) GetAccountBalance(ctx context.Context, req *co_company_api.GetAccountBalanceReq) (api_v1.Int64Res, error) {

	return funs.CheckPermission(ctx,
		func() (api_v1.Int64Res, error) {
			ret, err := c.modules.Account().GetAccountById(ctx, req.AccountId)
			if err != nil {
				return 0, err
			}
			return (api_v1.Int64Res)(ret.Balance), err
		},
		co_enum.Financial.PermissionType.GetAccountBalance,
	)
}

// InvoiceRegister 添加发票抬头
func (c *FinancialController) InvoiceRegister(ctx context.Context, req *co_company_api.CreateInvoiceReq) (*co_model.FdInvoiceInfoRes, error) {
	// 给userID和UnionMainId赋值
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	req.UserId = user.Id
	req.UnionMainId = user.UnionMainId

	return funs.CheckPermission(ctx,
		func() (*co_model.FdInvoiceInfoRes, error) {
			ret, err := c.modules.Invoice().CreateInvoice(ctx, req.FdInvoiceRegister)
			return (*co_model.FdInvoiceInfoRes)(ret), err
		},
		co_enum.Financial.PermissionType.CreateInvoice,
	)
}

// QueryInvoice 获取我的发票抬头列表
func (c *FinancialController) QueryInvoice(ctx context.Context, req *co_company_api.QueryInvoiceReq) (*co_model.FdInvoiceListRes, error) {
	// 权限判断
	return funs.CheckPermission(ctx,
		func() (*co_model.FdInvoiceListRes, error) {
			return c.modules.Invoice().QueryInvoiceList(ctx, &req.SearchParams, req.UserId)
		},
		co_enum.Financial.PermissionType.ViewInvoice,
	)

}

// DeletesFdInvoiceById 删除发票抬头
func (c *FinancialController) DeletesFdInvoiceById(ctx context.Context, req *co_company_api.DeleteInvoiceByIdReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Invoice().DeletesFdInvoiceById(ctx, req.InvoiceId)
			return ret == true, err
		},
		co_enum.Financial.PermissionType.DeleteInvoice,
	)
}

// InvoiceDetailRegister 申请开发票
func (c *FinancialController) InvoiceDetailRegister(ctx context.Context, req *co_company_api.CreateInvoiceDetailReq) (*co_model.FdInvoiceDetailInfoRes, error) {
	ret, err := c.modules.InvoiceDetail().CreateInvoiceDetail(ctx, req.FdInvoiceDetailRegister)

	return (*co_model.FdInvoiceDetailInfoRes)(ret), err
}

// QueryInvoiceDetailList 获取发票详情列表
func (c *FinancialController) QueryInvoiceDetailList(ctx context.Context, req *co_company_api.QueryInvoiceDetailListReq) (*co_model.FdInvoiceDetailListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.FdInvoiceDetailListRes, error) {
			ret, err := c.modules.InvoiceDetail().QueryInvoiceDetail(ctx, &req.SearchParams, req.UserId, req.UnionMainId)
			return ret, err
		},
		co_enum.Financial.PermissionType.ViewInvoiceDetail,
	)
}

// MakeInvoiceDetailReq 开发票
func (c *FinancialController) MakeInvoiceDetailReq(ctx context.Context, req *co_company_api.MakeInvoiceDetailReq) (api_v1.BoolRes, error) {
	ret, err := c.modules.InvoiceDetail().MakeInvoiceDetail(ctx, req.InvoiceDetailId, req.FdMakeInvoiceDetail)
	return ret == true, err
}

// AuditInvoiceDetail 审核发票
func (c *FinancialController) AuditInvoiceDetail(ctx context.Context, req *co_company_api.AuditInvoiceDetailReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.InvoiceDetail().AuditInvoiceDetail(ctx, req.InvoiceDetailId, req.AuditInfo)
			return ret == true, err
		},
		co_enum.Financial.PermissionType.AuditInvoiceDetail,
	)
}

// BankCardRegister 申请提现账号
func (c *FinancialController) BankCardRegister(ctx context.Context, req *co_company_api.BankCardRegisterReq) (*co_model.BankCardInfoRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	return funs.CheckPermission(ctx,
		func() (*co_model.BankCardInfoRes, error) {
			ret, err := c.modules.BankCard().CreateBankCard(ctx, req.BankCardRegister, &user.SysUser)
			return (*co_model.BankCardInfoRes)(ret), err
		},
		co_enum.Financial.PermissionType.CreateBankCard,
	)
}

// DeleteBankCard 删除提现账号
func (c *FinancialController) DeleteBankCard(ctx context.Context, req *co_company_api.DeleteBankCardReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.BankCard().DeleteBankCardById(ctx, req.BankCardId)
			return ret == true, err
		},
		co_enum.Financial.PermissionType.DeleteBankCard,
	)
}

// QueryBankCardList 获取用户的银行卡列表
func (c *FinancialController) QueryBankCardList(ctx context.Context, req *co_company_api.QueryBankCardListReq) (*co_model.BankCardListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.BankCardListRes, error) {
			ret, err := c.modules.BankCard().QueryBankCardListByUserId(ctx, req.UserId)
			return ret, err
		},
		co_enum.Financial.PermissionType.ViewBankCardDetail,
	)
}
