package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

type hookInfo sys_model.HookEventType[co_model.AccountBillHookFilter, co_model.AccountBillHookFunc]

// 财务账单
type sFdAccountBill struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
	hookArr *garray.Array
}

func NewFdAccountBill(modules co_interface.IModules, xDao *co_dao.XDao) co_interface.IFdAccountBill {
	return &sFdAccountBill{
		modules: modules,
		dao:     xDao,
		hookArr: garray.NewArray(),
	}
}

// InstallHook 安装Hook
func (s *sFdAccountBill) InstallHook(filter co_model.AccountBillHookFilter, hookFunc co_model.AccountBillHookFunc) {
	item := hookInfo{Key: filter, Value: hookFunc}
	s.hookArr.Append(item)
}

// UnInstallHook 卸载Hook
func (s *sFdAccountBill) UnInstallHook(filter co_model.AccountBillHookFilter) {
	newFuncArr := garray.NewArray()
	s.hookArr.Iterator(func(key int, value interface{}) bool {
		item := value.(hookInfo)

		if item.Key.TradeType != filter.TradeType ||
			item.Key.InOutType != filter.InOutType ||
			item.Key.InTransaction != filter.InTransaction {
			newFuncArr.Append(item)
		}

		return true
	})
	s.hookArr = newFuncArr
}

// ClearAllHook 清除Hook
func (s *sFdAccountBill) ClearAllHook() {
	s.hookArr.Clear()
}

// CreateAccountBill 创建财务账单
func (s *sFdAccountBill) CreateAccountBill(ctx context.Context, info co_model.AccountBillRegister) (bool, error) {
	// 判断交易时间是否大于当前系统时间
	now := *gtime.Now()

	if !now.After(info.TradeAt) { // 系统时间是否在交易时间之后
		return false, gerror.New("非法操作！")
	}

	// 交易金额是否为负数
	if info.Amount < 0 {
		return false, gerror.New("交易金额不能是负数")
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
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败"), "", s.dao.FdAccountBill.Table())
	}

	return success == true, err
}

// 修改支付状态

// 接收微信支付宝的支付通知  （账单id，微信的交易id，支付结果）

// 接收微信支付宝的握手

// 待支付：

// HOOK (Hook)

// income 收入
func (s *sFdAccountBill) income(ctx context.Context, info co_model.AccountBillRegister) (bool, error) {
	// 判断接受者是否存在
	toUser, err := sys_service.SysUser().GetSysUserById(ctx, info.ToUserId)
	if err != nil || toUser == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败,交易接收者不存在"), "", s.dao.FdAccountBill.Table())
	}

	// 先通过财务账号id查询账号出来
	account, err := s.modules.Account().GetAccountById(ctx, info.FdAccountId)

	// 判断需要收款的用户是否存在
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败,交易接收账户查询失败"), "", s.dao.FdAccountBill.Table())
	}

	bill := co_model.AccountBillInfo{}

	// 使用乐观锁校验余额，和更新余额
	err = s.dao.FdAccountBill.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 版本
		version := account.Version
		// 余额 = 之前的余额 + 本次交易的金额
		afterBalance := account.Balance + info.Amount

		// 1. 添加一条收入财务账单流水
		info.BeforeBalance = account.Balance
		info.AfterBalance = afterBalance
		gconv.Struct(info, &bill)
		bill.Id = idgen.NextId()

		result, err := s.dao.FdAccountBill.Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(bill)

		if result == nil || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "财务账单创建失败"), "", s.dao.FdAccountBill.Table())
		}

		// 2.修改财务账号的余额
		// 参数：上下文, 财务账号id, 需要修改的钱数目, 查询到的版本, 收支类型
		affected, err := s.modules.Account().UpdateAccountBalance(ctx, account.Id, info.Amount, version, info.InOutType)

		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "更新账户余额失败"), "", s.dao.FdAccountBill.Table())
		}

		s.hookArr.Iterator(func(_ int, v interface{}) bool {
			hook := v.(hookInfo)
			if hook.Key.InTransaction && hook.Key.InOutType == co_enum.Financial.InOutType.In {
				if hook.Key.TradeType.Code()&info.TradeType == info.TradeType {
					hook.Value(ctx, hook.Key, bill)
				}
			}
			return true
		})
		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败"), err.Error(), s.dao.FdAccountBill.Table())
	}

	s.hookArr.Iterator(func(_ int, v interface{}) bool {
		hook := v.(hookInfo)
		if !hook.Key.InTransaction && hook.Key.InOutType == co_enum.Financial.InOutType.In {
			if hook.Key.TradeType.Code()&info.TradeType == info.TradeType {
				hook.Value(ctx, hook.Key, bill)
			}
		}
		return true
	})

	return true, nil
}

// spending 支出
func (s *sFdAccountBill) spending(ctx context.Context, info co_model.AccountBillRegister) (bool, error) {
	// 先通过财务账号id查询账号出来
	account, err := s.modules.Account().GetAccountById(ctx, info.FdAccountId)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败,交易接收账户查询失败"), "", s.dao.FdAccountBill.Table())
	}

	bill := co_model.AccountBillInfo{}

	// 使用乐观锁校验余额，和更新余额
	err = s.dao.FdAccountBill.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 版本
		version := account.Version
		// 余额 = 之前的余额 - 本次交易的金额
		afterBalance := account.Balance - info.Amount

		// 判断余额是否足够
		if account.Balance >= info.Amount { // 足够
			// 1. 添加一条财务账单流水
			info.BeforeBalance = account.Balance
			info.AfterBalance = afterBalance

			gconv.Struct(info, &bill)
			bill.Id = idgen.NextId()

			result, err := s.dao.FdAccountBill.Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(bill)

			if result == nil || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "财务账单创建失败"), "", s.dao.FdAccountBill.Table())
			}

			// 2.修改财务账号的余额
			// 参数：上下文, 财务账号id, 需要修改的钱数目, 查询到的版本, 收支类型
			affected, err := s.modules.Account().UpdateAccountBalance(ctx, account.Id, info.Amount, version, info.InOutType)

			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "更新账户余额失败"), "", s.dao.FdAccountBill.Table())
			}

			s.hookArr.Iterator(func(_ int, v interface{}) bool {
				hook := v.(hookInfo)
				// 判断收支类型
				if hook.Key.InTransaction && hook.Key.InOutType == co_enum.Financial.InOutType.Out {
					// 判断交易类型
					if hook.Key.TradeType.Code()&info.TradeType == info.TradeType {
						hook.Value(ctx, hook.Key, bill)
					}
				}
				return true
			})

		} else {
			return gerror.New("交易发起者的账户余额不足")
		}

		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败"), err.Error(), s.dao.FdAccountBill.Table())
	}

	s.hookArr.Iterator(func(_ int, v interface{}) bool {
		hook := v.(hookInfo)
		if !hook.Key.InTransaction && hook.Key.InOutType == co_enum.Financial.InOutType.Out {
			if hook.Key.TradeType.Code()&info.TradeType == info.TradeType {
				hook.Value(ctx, hook.Key, bill)
			}
		}
		return true
	})

	return true, nil
}

// GetAccountBillByAccountId  根据财务账号id获取账单
func (s *sFdAccountBill) GetAccountBillByAccountId(ctx context.Context, accountId int64, pagination *sys_model.Pagination) (*co_model.AccountBillListRes, error) {
	if accountId == 0 {
		return nil, gerror.New("财务账号不能为0！")
	}

	if pagination == nil {
		pagination = &sys_model.Pagination{
			PageNum:  1,
			PageSize: 20,
		}
	}

	result, err := daoctl.Query[co_entity.FdAccountBill](s.dao.FdAccountBill.Ctx(ctx).Hook(daoctl.CacheHookHandler), &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0), sys_model.FilterInfo{
			Field: s.dao.FdAccountBill.Columns().FdAccountId,
			Where: "=",
			Value: accountId,
		}),
		OrderBy: append(make([]sys_model.OrderBy, 0), sys_model.OrderBy{
			Field: s.dao.FdAccountBill.Columns().CreatedAt,
			Sort:  "asc",
		}),
		Pagination: *pagination,
	}, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "账单查询失败", s.dao.FdAccountBill.Table())
	}

	return (*co_model.AccountBillListRes)(result), nil
}