package finance

import (
	"context"
	"fmt"

	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/kconv"

	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_hook"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"

	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// sFdAccountBills 财务账单
type sFdAccountBills[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	TR co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		TR,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]
	dao     *co_dao.XDao
	hookArr base_hook.BaseHook[co_hook.AccountBillHookKey, co_hook.AccountBillHookFunc]
}

func NewFdAccountBills[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	TR co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) co_interface.IFdAccountBills[TR] {
	result := &sFdAccountBills[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		TR,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]{
		modules: modules,
		dao:     modules.Dao(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

// InstallTradeHook 订阅Hook
func (s *sFdAccountBills[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) InstallTradeHook(hookKey co_hook.AccountBillHookKey, hookFunc co_hook.AccountBillHookFunc) {
	s.hookArr.InstallHook(hookKey, hookFunc)
}

// GetTradeHook 获取Hook
func (s *sFdAccountBills[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetTradeHook() base_hook.BaseHook[co_hook.AccountBillHookKey, co_hook.AccountBillHookFunc] {
	return s.hookArr
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sFdAccountBills[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) FactoryMakeResponseInstance() TR {
	instance, _ := base_funs.CreateGenericInstance[TR]()
	return instance
}

// CreateAccountBills 创建财务账单
func (s *sFdAccountBills[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) CreateAccountBills(ctx context.Context, info co_model.AccountBillsRegister) (bool, error) {
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
	if info.InOutType == co_enum.Finance.InOutType.In.Code() {
		success, err = s.income(ctx, info)
	} else if info.InOutType == co_enum.Finance.InOutType.Out.Code() {
		success, err = s.spending(ctx, info)
	}

	if success == false {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Transaction_Failed"), s.dao.FdAccountBills.Table())
	}

	return success == true, err
}

// 修改支付状态

// 接收微信支付宝的支付通知  （账单id，微信的交易id，支付结果）

// 接收微信支付宝的握手

// 待支付：

// HOOK (Hook)

// income 收入
func (s *sFdAccountBills[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) income(ctx context.Context, info co_model.AccountBillsRegister) (bool, error) {
	//sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断接受者是否存在
	toUser, err := sys_service.SysUser().GetSysUserById(ctx, info.ToUserId)
	if err != nil || toUser == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Transaction_Failed}{#error_Transaction_ToUser_NoExist}"), s.dao.FdAccountBills.Table())
	}

	// 先通过财务账号id查询账号出来
	account, err := s.modules.Account().GetAccountById(ctx, info.FdAccountId)

	// 判断需要收款的用户是否存在
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Transaction_Failed}{#error_ToUserAccount_NoExist}"), s.dao.FdAccountBills.Table())
	}

	bill, _ := base_funs.CreateGenericInstance[TR]()

	// 使用乐观锁校验余额，和更新余额
	err = s.dao.FdAccountBills.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
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
		//bill.Data().CreatedBy = sessionUser.Id
		bill.Data().CreatedBy = info.ToUserId

		data := kconv.Struct(bill.Data(), &co_do.FdAccountBills{})

		// 重载Do模型
		doData, err := info.OverrideDo.DoFactory(*data)
		if err != nil {
			return err
		}

		result, err := s.dao.FdAccountBills.Ctx(ctx).Insert(doData)

		if result == nil || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#AccountBill}{#error_Create_Failed}"), s.dao.FdAccountBills.Table())
		}

		// 2.修改财务账号的余额
		// 参数：上下文, 财务账号id, 需要修改的钱数目, 查询到的版本, 收支类型
		affected, err := s.modules.Account().UpdateAccountBalance(ctx, account.Data().Id, info.Amount, version, co_enum.Finance.InOutType.New(info.InOutType, ""), info.FromUserId)

		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_AccountBalance_Update_Failed"), s.dao.FdAccountBills.Table())
		}

		// 3.修改财务账号金额明细统计
		increment, err := s.modules.Account().Increment(ctx, account.Data().Id, gconv.Int(info.Amount))

		if increment == false || err != nil {
			return err
		}

		// 如果类型是保证金则增加冻结金额
		if info.TradeType == co_enum.Finance.TradeType.SecurityDeposit.Code() && info.TradeState == co_enum.Finance.TradeState.Unfrozen.Code() {
			increment, err = s.modules.Account().DecrementFrozenAmount(ctx, account.Data().Id, info.Amount)
		}

		if increment == false || err != nil {
			return err
		}

		_ = g.Try(ctx, func(ctx context.Context) {
			s.hookArr.Iterator(func(key co_hook.AccountBillHookKey, value co_hook.AccountBillHookFunc) {
				if key.InTransaction && key.InOutType == co_enum.Finance.InOutType.In {
					if key.TradeType.Code()&info.TradeType == info.TradeType {
						_ = value(ctx, key, bill)
					}
				}
			})
		})

		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Transaction_Failed"), s.dao.FdAccountBills.Table())
	}

	_ = g.Try(ctx, func(ctx context.Context) {
		s.hookArr.Iterator(func(key co_hook.AccountBillHookKey, value co_hook.AccountBillHookFunc) {
			// 在事务中 && 订阅key是收入类型的
			if !key.InTransaction && key.InOutType == co_enum.Finance.InOutType.In {
				// 订阅的交易类型一致
				if key.TradeType.Code()&info.TradeType == info.TradeType {
					_ = value(ctx, key, bill)
				}

			}
		})
	})

	return true, nil
}

// spending 支出
func (s *sFdAccountBills[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) spending(ctx context.Context, info co_model.AccountBillsRegister) (bool, error) {
	//sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 先通过财务账号id查询账号出来
	account, err := s.modules.Account().GetAccountById(ctx, info.FdAccountId)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Transaction_Failed}{#error_ToUserAccount_NoExist}"), s.dao.FdAccountBills.Table())
	}

	bill := s.FactoryMakeResponseInstance()

	// 使用乐观锁校验余额，和更新余额
	err = s.dao.FdAccountBills.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 版本
		version := account.Data().Version
		// 余额 = 之前的余额 - 本次交易的金额
		afterBalance := account.Data().Balance - info.Amount
		// 判断余额是否足够
		// 检验机制：余额足够 ||  账号类型是系统账户&&场景不限&&允许存在负数余额 || 允许存在负数余额  （相当于只校验了Allow参数，条件2直接忽略了）
		if account.Data().Balance >= info.Amount ||
			(((account.Data().AccountType & co_enum.Finance.AccountType.System.Code()) == account.Data().AccountType&co_enum.Finance.AccountType.System.Code()) &&
				(account.Data().SceneType&co_enum.Finance.SceneType.UnLimit.Code()) == co_enum.Finance.SceneType.UnLimit.Code() && account.Data().AllowExceed == co_enum.Finance.AllowExceed.Allow.Code()) ||
			account.Data().AllowExceed == co_enum.Finance.AllowExceed.Allow.Code() {
			// 1. 添加一条财务账单流水
			info.BeforeBalance = account.Data().Balance
			info.AfterBalance = afterBalance

			gconv.Struct(info, bill.Data())
			bill.Data().Id = idgen.NextId()
			bill.Data().CreatedAt = gtime.Now()
			//bill.Data().CreatedBy = sessionUser.Id  // TODO 后续解开
			bill.Data().CreatedBy = info.FromUserId

			data := kconv.Struct(bill.Data(), &co_do.FdAccountBills{})

			// 重载Do模型
			doData, err := info.OverrideDo.DoFactory(*data)
			if err != nil {
				return err
			}

			result, err := s.dao.FdAccountBills.Ctx(ctx).Insert(doData)

			if result == nil || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_AccountBalance_Update_Failed"), s.dao.FdAccountBills.Table())
			}

			// 2.修改财务账号的余额
			// 参数：上下文, 财务账号id, 需要修改的钱数目, 查询到的版本, 收支类型
			affected, err := s.modules.Account().UpdateAccountBalance(ctx, account.Data().Id, info.Amount, version, co_enum.Finance.InOutType.New(info.InOutType, ""), info.FromUserId)

			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_AccountBalance_Update_Failed"), s.dao.FdAccountBills.Table())
			}
			// 3.修改财务账号金额明细统计
			decrement, err := s.modules.Account().Decrement(ctx, account.Data().Id, gconv.Int(info.Amount))

			if decrement == false || err != nil {
				return err
			}

			// 如果类型是保证金则增加冻结金额
			if info.TradeType == co_enum.Finance.TradeType.SecurityDeposit.Code() {
				decrement, err = s.modules.Account().IncrementFrozenAmount(ctx, account.Data().Id, info.Amount)
			}

			if decrement == false || err != nil {
				return err
			}

			g.Try(ctx, func(ctx context.Context) {
				s.hookArr.Iterator(func(key co_hook.AccountBillHookKey, value co_hook.AccountBillHookFunc) {
					// 在事务中 && 订阅key是收入类型的
					if key.InTransaction && key.InOutType == co_enum.Finance.InOutType.Out {
						// 订阅的交易类型一致
						if key.TradeType.Code()&info.TradeType == info.TradeType {
							value(ctx, key, bill)
						}
					}
				})
			})
		} else {
			return gerror.New(s.modules.T(ctx, "error_Transaction_FromUser_InsufficientBalance"))
		}

		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Transaction_Failed"), s.dao.FdAccountBills.Table())
	}

	g.Try(ctx, func(ctx context.Context) {
		s.hookArr.Iterator(func(key co_hook.AccountBillHookKey, value co_hook.AccountBillHookFunc) {
			if !key.InTransaction && key.InOutType == co_enum.Finance.InOutType.Out {
				if key.TradeType.Code()&info.TradeType == info.TradeType {
					value(ctx, key, bill)
				}
			}
		})
	})

	return true, nil
}

// GetAccountBillsByAccountId  根据财务账号id获取账单
func (s *sFdAccountBills[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	TR,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountBillsByAccountId(ctx context.Context, accountId int64, searchParams *base_model.SearchParams) (*base_model.CollectRes[TR], error) {
	if accountId == 0 {
		return nil, gerror.New(s.modules.T(ctx, "error_AccountId_NonZero"))
	}

	//if pagination == nil {
	//	pagination = &base_model.Pagination{
	//		PageNum:  1,
	//		PageSize: 20,
	//	}
	//}

	result, err := daoctl.Query[TR](s.dao.FdAccountBills.Ctx(ctx).
		Where(s.dao.FdAccountBills.Columns().FdAccountId, accountId).
		OrderDesc(s.dao.FdAccountBills.Columns().CreatedAt), searchParams, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#AccountBill}{#error_Data_Get_Failed}"), s.dao.FdAccountBills.Table())
	}

	return result, nil
}
