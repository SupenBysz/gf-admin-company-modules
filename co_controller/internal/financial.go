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
	//
	//i_controller.IFinancial
	//modules T
}

// Financial 财务服务
var Financial = func(modules co_interface.IModules) i_controller.IFinancial {
	return &FinancialController{
		modules: modules,
	}
}

// GetAccountBalance 查看账户余额
func (c *FinancialController) GetAccountBalance(ctx context.Context, req *co_company_api.GetAccountBalanceReq) (api_v1.Int64Res, error) {
	// 权限判断
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, co_enum.Financial.PermissionType.GetAccountBalance); has != true {
	//	return 0, err
	//}
	//
	//account, err := pro_service.ProFdAccount().GetAccountById(ctx, req.AccountId)

	//	return (api_v1.Int64Res)(account.Balance), err

	return funs.CheckPermission(ctx,
		func() (api_v1.Int64Res, error) {
			ret, err := c.modules.Account().GetAccountById(ctx, req.AccountId)
			return (api_v1.Int64Res)(ret.Balance), err
		},
		co_enum.Financial.PermissionType.GetAccountBalance,
	)
}

// InvoiceRegister 添加发票抬头
func (c *FinancialController[T]) InvoiceRegister(ctx context.Context, req *co_company_api.CreateInvoiceReq) (*co_model.FdInvoiceInfoRes, error) {
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
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, co_enum.Financial.PermissionType.InvoiceList); has != true {
	//	return nil, err
	//}
	//
	//result, err := pro_service.ProFdInvoice().QueryInvoiceList(ctx, &req.SearchParams, req.UserId)
	//
	//return result, err

	return funs.CheckPermission(ctx,
		func() (*co_model.FdInvoiceListRes, error) {
			return c.modules.Invoice().QueryInvoiceList(ctx, &req.SearchParams, req.UserId)
		},
		co_enum.Financial.PermissionType.GetAccountBalance,
		// 		co_enum.Financial.PermissionType(c.modules).GetAccountBalance,
	)

}

// DeletesFdInvoiceById 删除发票抬头
func (c *FinancialController) DeletesFdInvoiceById(ctx context.Context, req *co_company_api.DeleteInvoiceByIdReq) (api_v1.BoolRes, error) {
	// 权限判断
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, co_enum.Financial.PermissionType.DeleteInvoice); has != true {
	//	return false, err
	//}
	//
	//result, err := pro_service.ProFdInvoice().DeletesFdInvoiceById(ctx, req.InvoiceId)
	//return result == true, err

	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Invoice().DeletesFdInvoiceById(ctx, req.InvoiceId)
			return ret == true, err
		},
	)
}

// InvoiceDetailRegister 申请开发票
func (c *FinancialController[T]) InvoiceDetailRegister(ctx context.Context, req *co_company_api.CreateInvoiceDetailReq) (*co_model.FdInvoiceDetailInfoRes, error) {
	// userID和unionMainId赋值
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	req.UserId, req.UnionMainId = user.Id, user.UnionMainId

	ret, err := c.modules.InvoiceDetail().CreateInvoiceDetail(ctx, req.FdInvoiceDetailRegister)

	return (*co_model.FdInvoiceDetailInfoRes)(ret), err
}

// QueryInvoiceDetailList 获取发票详情列表
func (c *FinancialController) QueryInvoiceDetailList(ctx context.Context, req *co_company_api.QueryInvoiceDetailListReq) (*co_model.FdInvoiceDetailListRes, error) {
	// 权限判断
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, co_enum.Financial.PermissionType.InvoiceDetailList); has != true {
	//	return nil, err
	//}
	//
	//result, err := pro_service.ProFdInvoiceDetail().QueryInvoiceDetail(ctx, &req.SearchParams, req.UserId, req.UnionMainId)
	//
	//return result, err

	return funs.CheckPermission(ctx,
		func() (*co_model.FdInvoiceDetailListRes, error) {
			ret, err := c.modules.InvoiceDetail().QueryInvoiceDetail(ctx, &req.SearchParams, req.UserId, req.UnionMainId)
			return ret, err
		},
	)
}

// MakeInvoiceDetailReq 开发票
func (c *FinancialController) MakeInvoiceDetailReq(ctx context.Context, req *co_company_api.MakeInvoiceDetailReq) (api_v1.BoolRes, error) {
	ret, err := c.modules.InvoiceDetail().MakeInvoiceDetail(ctx, req.InvoiceDetailId, req.FdMakeInvoiceDetail)
	return ret == true, err
}

// AuditInvoiceDetail 审核发票
func (c *FinancialController) AuditInvoiceDetail(ctx context.Context, req *co_company_api.AuditInvoiceDetailReq) (api_v1.BoolRes, error) {
	// 权限判断
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, co_enum.Financial.PermissionType.AuditInvoiceDetail); has != true {
	//	return false, err
	//}
	//
	//result, err := pro_service.ProFdInvoiceDetail().AuditInvoiceDetail(ctx, req.InvoiceDetailId, req.AuditInfo)
	//return result == true, err

	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.InvoiceDetail().AuditInvoiceDetail(ctx, req.InvoiceDetailId, req.AuditInfo)
			return ret == true, err
		},
	)
}

// BankCardRegister 申请提现账号
func (c *FinancialController) BankCardRegister(ctx context.Context, req *co_company_api.BankCardRegisterReq) (*co_model.BankCardInfoRes, error) {
	// 权限判断
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, co_enum.Financial.PermissionType.CreateBankCard); has != true {
	//	return nil, err
	//}
	//
	//user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	//
	//card, err := s.CreateBankCard(ctx, req.BankCardRegister, &user.SysUser)
	//
	//return (*co_model.BankCardInfoRes)(card), err

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	return funs.CheckPermission(ctx,
		func() (*co_model.BankCardInfoRes, error) {
			ret, err := c.modules.BankCard().CreateBankCard(ctx, req.BankCardRegister, &user.SysUser)
			return (*co_model.BankCardInfoRes)(ret), err
		},
	)
}

// DeleteBankCard 删除提现账号
func (c *FinancialController) DeleteBankCard(ctx context.Context, req *co_company_api.DeleteBankCardReq) (api_v1.BoolRes, error) {
	// 权限判断
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, co_enum.Financial.PermissionType.DeleteBankCard); has != true {
	//	return false, err
	//}
	//
	//result, err := pro_service.ProFdBankCard().DeleteBankCardById(ctx, req.BankCardId)
	//return result == true, err

	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.BankCard().DeleteBankCardById(ctx, req.BankCardId)
			return ret == true, err
		},
	)
}

// QueryBankCardList 获取用户的银行卡列表
func (c *FinancialController) QueryBankCardList(ctx context.Context, req *co_company_api.QueryBankCardListReq) (*co_model.BankCardListRes, error) {
	// 权限判断
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, co_enum.Financial.PermissionType.BankCardList); has != true {
	//	return nil, err
	//}
	//
	//result, err := pro_service.ProFdBankCard().QueryBankCardListByUserId(ctx, req.UserId)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return result, nil

	return funs.CheckPermission(ctx,
		func() (*co_model.BankCardListRes, error) {
			ret, err := c.modules.BankCard().QueryBankCardListByUserId(ctx, req.UserId)
			return ret, err
		},
	)
}
