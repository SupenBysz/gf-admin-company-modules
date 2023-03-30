package i_controller

import (
    "context"
    "github.com/SupenBysz/gf-admin-community/api_v1"
    "github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
    "github.com/SupenBysz/gf-admin-company-modules/co_model"
    "github.com/kysion/base-library/base_model"
)

type IFinancial[
ITCompanyRes co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdCurrencyRes co_model.IFdCurrencyRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] interface {
    // GetAccountBalance 查看账户余额
    GetAccountBalance(ctx context.Context, req *co_company_api.GetAccountBalanceReq) (api_v1.Int64Res, error)

    // InvoiceRegister 添加发票抬头
    InvoiceRegister(ctx context.Context, req *co_company_api.CreateInvoiceReq) (ITFdInvoiceRes, error)

    // QueryInvoice 获取我的发票抬头列表
    QueryInvoice(ctx context.Context, req *co_company_api.QueryInvoiceReq) (*base_model.CollectRes[ITFdInvoiceRes], error)

    // DeletesFdInvoiceById 删除发票抬头
    DeletesFdInvoiceById(ctx context.Context, req *co_company_api.DeleteInvoiceByIdReq) (api_v1.BoolRes, error)

    // InvoiceDetailRegister 申请开发票
    InvoiceDetailRegister(ctx context.Context, req *co_company_api.CreateInvoiceDetailReq) (ITFdInvoiceDetailRes, error)

    // QueryInvoiceDetailList 获取发票详情列表
    QueryInvoiceDetailList(ctx context.Context, req *co_company_api.QueryInvoiceDetailListReq) (*base_model.CollectRes[ITFdInvoiceDetailRes], error)

    // MakeInvoiceDetailReq 开发票
    MakeInvoiceDetailReq(ctx context.Context, req *co_company_api.MakeInvoiceDetailReq) (api_v1.BoolRes, error)

    // AuditInvoiceDetail 审核发票
    AuditInvoiceDetail(ctx context.Context, req *co_company_api.AuditInvoiceDetailReq) (api_v1.BoolRes, error)

    // BankCardRegister 申请提现账号
    BankCardRegister(ctx context.Context, req *co_company_api.BankCardRegisterReq) (ITFdBankCardRes, error)

    // DeleteBankCard 删除提现账号
    DeleteBankCard(ctx context.Context, req *co_company_api.DeleteBankCardReq) (api_v1.BoolRes, error)

    // QueryBankCardList 获取用户的银行卡列表
    QueryBankCardList(ctx context.Context, req *co_company_api.QueryBankCardListReq) (*base_model.CollectRes[ITFdBankCardRes], error)

    // GetAccountDetail 查看财务账号明细
    GetAccountDetail(ctx context.Context, req *co_company_api.GetAccountDetailReq) (ITFdAccountRes, error)

    // UpdateAccountIsEnabled 修改财务账号启用状态
    UpdateAccountIsEnabled(ctx context.Context, req *co_company_api.UpdateAccountIsEnabledReq) (api_v1.BoolRes, error)

    // UpdateAccountLimitState 修改财务账号限制状态
    UpdateAccountLimitState(ctx context.Context, req *co_company_api.UpdateAccountLimitStateReq) (api_v1.BoolRes, error)

    // CreateAccountDetail 创建财务账单金额明细统计记录  只能被动创建

    // GetAccountDetailById 根据财务账号id查询账单金额明细统计记录
    GetAccountDetailById(ctx context.Context, req *co_company_api.GetAccountDetailByAccountIdReq) (*co_model.FdAccountDetailRes, error)

    // Increment 收入
    Increment(ctx context.Context, req *co_company_api.IncrementReq) (api_v1.BoolRes, error)

    // Decrement 支出
    Decrement(ctx context.Context, req *co_company_api.DecrementReq) (api_v1.BoolRes, error)
}
