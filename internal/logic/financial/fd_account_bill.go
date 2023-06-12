package financial

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"

	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_hook"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

// 财务账单
type sFdAccountBill[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	TR co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		TR,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
	dao     *co_dao.XDao
	hookArr base_hook.BaseHook[co_hook.AccountBillHookFilter, co_hook.AccountBillHookFunc]
	account co_interface.IFdAccount[ITFdAccountRes]
}

func NewFdAccountBill[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	TR co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) co_interface.IFdAccountBill[TR] {
	result := &sFdAccountBill[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		TR,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		modules: modules,
		dao:     modules.Dao(),
		account: modules.Account(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

// InstallTradeHook 订阅Hook
func (s *sFdAccountBill[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) InstallTradeHook(hookKey co_hook.AccountBillHookFilter, hookFunc co_hook.AccountBillHookFunc) {
	s.hookArr.InstallHook(hookKey, hookFunc)
}

// GetTradeHook 获取Hook
func (s *sFdAccountBill[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetTradeHook() base_hook.BaseHook[co_hook.AccountBillHookFilter, co_hook.AccountBillHookFunc] {
	return s.hookArr
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sFdAccountBill[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.IFdAccountBillRes
	ret = &co_model.FdAccountBillRes{}
	return ret.(TR)
}

// CreateAccountBill 创建财务账单
func (s *sFdAccountBill[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) CreateAccountBill(ctx context.Context, info co_model.AccountBillRegister) (bool, error) {
	// 判断交易时间是否大于当前系统时间
	now := gtime.Now()

	if now.Format("Y-m-d H:i:s") < info.TradeAt.Format("Y-m-d H:i:s") { // 系统时间是否在交易时间之后
		fmt.Println(now.Format("Y-m-d H:i:s"))
		fmt.Println(info.TradeAt.Format("Y-m-d H:i:s"))

		return false, gerror.New(s.modules.T(ctx, "error_Legal_Operation"))
	}

	// 交易金额是否为负数
	if info.Amount < 0 {
		return false, gerror.New(s.modules.T(ctx, "error_AccountAmount_NonNegative"))
	}

	var success bool
	var err error

	// 判读收支类型  收入/支出
	if info.InOutType == 1 {
		success, err = s.income(ctx, info)
	} else if info.InOutType == 2 {
		success, err = s.spending(ctx, info)
	}

	if success == false {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Transaction_Failed"), s.dao.FdAccountBill.Table())
	}

	return success == true, err
}

// 修改支付状态

// 接收微信支付宝的支付通知  （账单id，微信的交易id，支付结果）

// 接收微信支付宝的握手

// 待支付：

// HOOK (Hook)

// income 收入
func (s *sFdAccountBill[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) income(ctx context.Context, info co_model.AccountBillRegister) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断接受者是否存在
	toUser, err := sys_service.SysUser().GetSysUserById(ctx, info.ToUserId)
	if err != nil || toUser == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Transaction_Failed}{#error_Transaction_ToUser_NoExist}"), s.dao.FdAccountBill.Table())
	}

	// 先通过财务账号id查询账号出来
	account, err := s.account.GetAccountById(ctx, info.FdAccountId)

	// 判断需要收款的用户是否存在
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Transaction_Failed}{#error_ToUserAccount_NoExist}"), s.dao.FdAccountBill.Table())
	}

	bill := s.FactoryMakeResponseInstance()

	// 使用乐观锁校验余额，和更新余额
	err = s.dao.FdAccountBill.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 版本
		version := account.Data().Version
		// 余额 = 之前的余额 + 本次交易的金额
		afterBalance := account.Data().Balance + info.Amount

		// 1. 添加一条收入财务账单流水
		info.BeforeBalance = account.Data().Balance
		info.AfterBalance = afterBalance
		gconv.Struct(info, bill.Data())
		bill.Data().Id = idgen.NextId()
		bill.Data().CreatedAt = gtime.Now()
		bill.Data().CreatedBy = sessionUser.Id

		result, err := s.dao.FdAccountBill.Ctx(ctx).Insert(bill)

		if result == nil || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#AccountBill}{#error_Create_Failed}"), s.dao.FdAccountBill.Table())
		}

		// 2.修改财务账号的余额
		// 参数：上下文, 财务账号id, 需要修改的钱数目, 查询到的版本, 收支类型
		affected, err := s.account.UpdateAccountBalance(ctx, account.Data().Id, info.Amount, version, info.InOutType)

		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_AccountBalance_Update_Failed"), s.dao.FdAccountBill.Table())
		}

		// 3.修改财务账号金额明细统计
		increment, err := s.account.Increment(ctx, account.Data().Id, gconv.Int(info.Amount))

		if increment == false || err != nil {
			return err
		}

		s.hookArr.Iterator(func(key co_hook.AccountBillHookFilter, value co_hook.AccountBillHookFunc) {
			if key.InTransaction && key.InOutType == co_enum.Financial.InOutType.In {
				if key.TradeType.Code()&info.TradeType == info.TradeType {
					value(ctx, key, bill)
				}
			}
		})
		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Transaction_Failed"), s.dao.FdAccountBill.Table())
	}

	s.hookArr.Iterator(func(key co_hook.AccountBillHookFilter, value co_hook.AccountBillHookFunc) {
		if !key.InTransaction && key.InOutType == co_enum.Financial.InOutType.In {
			if key.TradeType.Code()&info.TradeType == info.TradeType {
				value(ctx, key, bill)
			}

		}
	})

	return true, nil
}

// spending 支出
func (s *sFdAccountBill[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) spending(ctx context.Context, info co_model.AccountBillRegister) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 先通过财务账号id查询账号出来
	account, err := s.account.GetAccountById(ctx, info.FdAccountId)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Transaction_Failed}{#error_ToUserAccount_NoExist}"), s.dao.FdAccountBill.Table())
	}

	bill := s.FactoryMakeResponseInstance()

	// 使用乐观锁校验余额，和更新余额
	err = s.dao.FdAccountBill.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 版本
		version := account.Data().Version
		// 余额 = 之前的余额 - 本次交易的金额
		afterBalance := account.Data().Balance - info.Amount
		// 判断余额是否足够   // 余额足够、账号类型是系统账户，场景不限
		if account.Data().Balance >= info.Amount ||
			((account.Data().AccountType&co_enum.Financial.AccountType.System.Code()) == account.Data().AccountType&co_enum.Financial.AccountType.System.Code()) &&
				(account.Data().SceneType&co_enum.Financial.SceneType.UnLimit.Code()) == co_enum.Financial.SceneType.UnLimit.Code() {
			// 1. 添加一条财务账单流水
			info.BeforeBalance = account.Data().Balance
			info.AfterBalance = afterBalance

			gconv.Struct(info, bill.Data())
			bill.Data().Id = idgen.NextId()
			bill.Data().CreatedAt = gtime.Now()
			bill.Data().CreatedBy = sessionUser.Id

			result, err := s.dao.FdAccountBill.Ctx(ctx).Insert(bill)

			if result == nil || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_AccountBalance_Update_Failed"), s.dao.FdAccountBill.Table())
			}

			// 2.修改财务账号的余额
			// 参数：上下文, 财务账号id, 需要修改的钱数目, 查询到的版本, 收支类型
			affected, err := s.account.UpdateAccountBalance(ctx, account.Data().Id, info.Amount, version, info.InOutType)

			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_AccountBalance_Update_Failed"), s.dao.FdAccountBill.Table())
			}
			// 3.修改财务账号金额明细统计
			decrement, err := s.account.Decrement(ctx, account.Data().Id, gconv.Int(info.Amount))

			if decrement == false || err != nil {
				return err
			}

			s.hookArr.Iterator(func(key co_hook.AccountBillHookFilter, value co_hook.AccountBillHookFunc) {
				if key.InTransaction && key.InOutType == co_enum.Financial.InOutType.Out {
					if key.TradeType.Code()&info.TradeType == info.TradeType {
						value(ctx, key, bill)
					}
				}
			})
		} else {
			return gerror.New(s.modules.T(ctx, "error_Transaction_FromUser_InsufficientBalance"))
		}

		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Transaction_Failed"), s.dao.FdAccountBill.Table())
	}

	s.hookArr.Iterator(func(key co_hook.AccountBillHookFilter, value co_hook.AccountBillHookFunc) {
		if !key.InTransaction && key.InOutType == co_enum.Financial.InOutType.Out {
			if key.TradeType.Code()&info.TradeType == info.TradeType {
				value(ctx, key, bill)
			}
		}
	})

	return true, nil
}

// GetAccountBillByAccountId  根据财务账号id获取账单
func (s *sFdAccountBill[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccountBillByAccountId(ctx context.Context, accountId int64, pagination *base_model.Pagination) (*base_model.CollectRes[TR], error) {
	if accountId == 0 {
		return nil, gerror.New(s.modules.T(ctx, "error_AccountId_NonZero"))
	}

	if pagination == nil {
		pagination = &base_model.Pagination{
			PageNum:  1,
			PageSize: 20,
		}
	}

	result, err := daoctl.Query[TR](s.dao.FdAccountBill.Ctx(ctx), &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
			Field: s.dao.FdAccountBill.Columns().FdAccountId,
			Where: "=",
			Value: accountId,
		}),
		OrderBy: append(make([]base_model.OrderBy, 0), base_model.OrderBy{
			Field: s.dao.FdAccountBill.Columns().CreatedAt,
			Sort:  "asc",
		}),
		Pagination: *pagination,
	}, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#AccountBill}{#error_Data_Get_Failed}"), s.dao.FdAccountBill.Table())
	}

	return result, nil
}
