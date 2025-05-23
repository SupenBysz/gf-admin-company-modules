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
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/kconv"
)

// FinanceController 财务服务控制器
type FinanceController[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
] struct {
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]
	dao co_dao.XDao
}

// Finance 财务服务
func Finance[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) i_controller.IFinance[
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
] {
	return &FinanceController[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]{
		modules: modules,
		dao:     *modules.Dao(),
	}
}

// GetAccountBalance 查看账户余额
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountBalance(ctx context.Context, req *co_company_api.GetAccountBalanceReq) (api_v1.Int64Res, error) {

	return funs.CheckPermission(ctx,
		func() (api_v1.Int64Res, error) {
			ret, err := c.modules.Account().GetAccountById(ctx, req.AccountId)
			if err != nil {
				return 0, err
			}
			return (api_v1.Int64Res)(ret.Data().Balance), err
		},
		co_permission.Finance.PermissionType(c.modules).GetAccountBalance,
	)
}

// InvoiceRegister 添加发票抬头
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) InvoiceRegister(ctx context.Context, req *co_company_api.CreateInvoiceReq) (ITFdInvoiceRes, error) {
	// 给userID和UnionMainId赋值
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	req.UserId = user.Id
	req.UnionMainId = user.UnionMainId

	return funs.CheckPermission(ctx,
		func() (ITFdInvoiceRes, error) {
			ret, err := c.modules.Invoice().CreateInvoice(ctx, req.FdInvoiceRegister)
			return ret, err
		},
		co_permission.Finance.PermissionType(c.modules).CreateInvoice,
	)
}

// QueryInvoice 获取我的发票抬头列表
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryInvoice(ctx context.Context, req *co_company_api.QueryInvoiceReq) (*base_model.CollectRes[ITFdInvoiceRes], error) {
	// 权限判断
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[ITFdInvoiceRes], error) {
			return c.modules.Invoice().QueryInvoiceList(ctx, &req.SearchParams, req.UserId)
		},
		co_permission.Finance.PermissionType(c.modules).ViewInvoice,
	)

}

// DeletesFdInvoiceById 删除发票抬头
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) DeletesFdInvoiceById(ctx context.Context, req *co_company_api.DeleteInvoiceByIdReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Invoice().DeletesFdInvoiceById(ctx, req.InvoiceId)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).DeleteInvoice,
	)
}

// InvoiceDetailRegister 申请开发票
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) InvoiceDetailRegister(ctx context.Context, req *co_company_api.CreateInvoiceDetailReq) (ITFdInvoiceDetailRes, error) {
	ret, err := c.modules.InvoiceDetail().CreateInvoiceDetail(ctx, req.FdInvoiceDetailRegister)
	return ret, err
}

// QueryInvoiceDetailList 获取发票详情列表
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryInvoiceDetailList(ctx context.Context, req *co_company_api.QueryInvoiceDetailListReq) (*base_model.CollectRes[ITFdInvoiceDetailRes], error) {
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[ITFdInvoiceDetailRes], error) {
			ret, err := c.modules.InvoiceDetail().QueryInvoiceDetail(ctx, &req.SearchParams, req.UserId, req.UnionMainId)
			return ret, err
		},
		co_permission.Finance.PermissionType(c.modules).ViewInvoiceDetail,
	)
}

// MakeInvoiceDetailReq 开发票
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) MakeInvoiceDetailReq(ctx context.Context, req *co_company_api.MakeInvoiceDetailReq) (api_v1.BoolRes, error) {
	ret, err := c.modules.InvoiceDetail().MakeInvoiceDetail(ctx, req.InvoiceDetailId, req.FdMakeInvoiceDetail)
	return ret == true, err
}

// AuditInvoiceDetail 审核发票
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) AuditInvoiceDetail(ctx context.Context, req *co_company_api.AuditInvoiceDetailReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.InvoiceDetail().AuditInvoiceDetail(ctx, req.InvoiceDetailId, req.AuditInfo)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).AuditInvoiceDetail,
	)
}

// BankCardRegister 申请提现账号
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) BankCardRegister(ctx context.Context, req *co_company_api.BankCardRegisterReq) (ITFdBankCardRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	return funs.CheckPermission(ctx,
		func() (ITFdBankCardRes, error) {
			ret, err := c.modules.BankCard().CreateBankCard(ctx, req.BankCardRegister, &user.SysUser)
			return ret, err
		},
		co_permission.Finance.PermissionType(c.modules).CreateBankCard,
	)
}

// DeleteBankCard 删除提现账号
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) DeleteBankCard(ctx context.Context, req *co_company_api.DeleteBankCardReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.BankCard().DeleteBankCardById(ctx, req.BankCardId)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).DeleteBankCard,
	)
}

// QueryBankCardList 获取用户的银行卡列表
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryBankCardList(ctx context.Context, req *co_company_api.QueryBankCardListReq) (*base_model.CollectRes[ITFdBankCardRes], error) {
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[ITFdBankCardRes], error) {
			ret, err := c.modules.BankCard().QueryBankCardListByUserId(ctx, req.UserId)
			return ret, err
		},
		co_permission.Finance.PermissionType(c.modules).ViewBankCardDetail,
	)
}

// GetAccountDetail 查看财务账号明细
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountDetail(ctx context.Context, req *co_company_api.GetAccountDetailReq) (ITFdAccountRes, error) {
	return funs.CheckPermission(ctx,
		func() (ITFdAccountRes, error) {
			ret, err := c.modules.Account().GetAccountById(c.makeMore(ctx), req.AccountId)
			return ret, err
		},
		co_permission.Finance.PermissionType(c.modules).GetAccountDetail,
	)
}

// UpdateAccountIsEnabled 修改财务账号启用状态
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateAccountIsEnabled(ctx context.Context, req *co_company_api.UpdateAccountIsEnabledReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			user := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := c.modules.Account().UpdateAccountIsEnable(ctx, req.AccountId, req.IsEnabled, user.Id)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).UpdateAccountState,
	)
}

// UpdateAccountLimitState 修改财务账号限制状态
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateAccountLimitState(ctx context.Context, req *co_company_api.UpdateAccountLimitStateReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			user := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := c.modules.Account().UpdateAccountLimitState(ctx, req.AccountId, req.LimitState, user.Id)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).UpdateAccountState,
	)
}

// SetAccountCurrencyCode 设置财务账号货币单位
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetAccountCurrencyCode(ctx context.Context, req *co_company_api.SetAccountCurrencyCodeReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			user := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := c.modules.Account().SetAccountCurrencyCode(ctx, req.AccountId, req.CurrencyCode, user.Id)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).SetAccountCurrencyCode,
	)
}

// UpdateAccountBalance 财务账号金额冲正
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateAccountBalance(ctx context.Context, req *co_company_api.UpdateAccountBalanceReq) (api_v1.Int64Res, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.Int64Res, error) {
			user := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := c.modules.Account().UpdateAccountBalance(ctx, req.AccountId, req.Amount, -1, co_enum.Finance.InOutType.Auto, user.Id)
			return api_v1.Int64Res(ret), err
		},
		co_permission.Finance.PermissionType(c.modules).UpdateAccountBalance,
	)
}

// GetAccountDetailById 根据财务账号id查询账单金额明细统计记录
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountDetailById(ctx context.Context, req *co_company_api.GetAccountDetailByAccountIdReq) (*co_model.FdAccountDetailRes, error) {
	return funs.CheckPermission(ctx,
		func() (*co_model.FdAccountDetailRes, error) {
			ret, err := c.modules.Account().GetAccountDetailById(ctx, req.AccountId)
			return ret, err
		},
		co_permission.Finance.PermissionType(c.modules).GetAccountDetail,
	)
}

// Increment 收入
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Increment(ctx context.Context, req *co_company_api.IncrementReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Account().Increment(ctx, req.AccountId, req.Amount)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).UpdateAccountAmount,
	)
}

// Decrement 支出
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) Decrement(ctx context.Context, req *co_company_api.DecrementReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Account().Decrement(ctx, req.AccountId, req.Amount)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).UpdateAccountAmount,
	)
}

// SetAccountAllowExceed 设置财务账号是否允许存在负余额
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetAccountAllowExceed(ctx context.Context, req *co_company_api.SetAccountAllowExceedReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Account().SetAccountAllowExceed(ctx, req.AccountId, req.AllowExceed)
			return ret == true, err
		},
		co_permission.Finance.PermissionType(c.modules).UpdateAccountAmount,
	)
}

// QueryAccountBills  根据财务账号id查询账单
func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryAccountBills(ctx context.Context, req *co_company_api.QueryAccountBillsReq) (*base_model.CollectRes[ITFdAccountBillRes], error) {
	return c.modules.AccountBills().GetAccountBillsByAccountId(ctx, req.AccountId, &req.SearchParams)
}

func (c *FinanceController[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) makeMore(ctx context.Context) context.Context {
	include := &garray.StrArray{}
	if ctx.Value("include") == nil {
		r := g.RequestFromCtx(ctx)
		array := r.GetForm("include").Array()
		arr := kconv.Struct(array, &[]string{})
		include = garray.NewStrArrayFrom(*arr)
	} else {
		array := ctx.Value("include")
		arr := kconv.Struct(array, &[]string{})
		include = garray.NewStrArrayFrom(*arr)
	}

	if include.Contains("*") {
		ctx = base_funs.AttrBuilder[ITFdAccountRes, ITFdAccountRes](ctx, "id")
		ctx = base_funs.AttrBuilder[ITFdAccountRes, co_entity.FdAccountDetail](ctx, "id")
	}

	if include.Contains("detail") {
		ctx = base_funs.AttrBuilder[ITFdAccountRes, ITFdAccountRes](ctx, "id")
	}

	ctx = base_funs.AttrBuilder[ITFdAccountRes, co_entity.FdAccountDetail](ctx, "id")

	// 因为需要附加公共模块user的数据，所以也要添加有关sys_user的附加数据订阅

	return ctx
}
