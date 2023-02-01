package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

type sFdAccount struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewFdAccount(modules co_interface.IModules, xDao *co_dao.XDao) co_interface.IFdAccount {
	return &sFdAccount{
		modules: modules,
		dao:     xDao,
	}
}

// CreateAccount 创建财务账号
func (s *sFdAccount) CreateAccount(ctx context.Context, info co_model.FdAccountRegister) (*co_entity.FdAccount, error) {
	// 检查指定参数是否为空
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return nil, err
	}

	// 关联用户id是否正确
	user, err := daoctl.GetByIdWithError[sys_entity.SysUser](sys_dao.SysUser.Ctx(ctx), info.UnionUserId)
	if user == nil || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Financial_UnionUserId_Failed"), sys_dao.SysUser.Table())
	}

	// 判断货币代码是否符合标准
	currency, err := s.modules.Currency().GetCurrencyByCurrencyCode(ctx, info.CurrencyCode)
	if err != nil || currency == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Financial_CurrencyCode_Failed"), s.dao.FdCurrency.Table())
	}
	if currency.IsLegalTender != 1 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_PleaseUse_Legal_Currency"), s.dao.FdCurrency.Table())

	}
	// 生产随机id
	data := co_do.FdAccount{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()

	// 插入财务账号
	_, err = s.dao.FdAccount.Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(data)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Account_Save_Failed"), s.dao.FdAccount.Table())
	}

	return s.GetAccountById(ctx, gconv.Int64(data.Id))
}

// GetAccountById 根据ID获取财务账号
func (s *sFdAccount) GetAccountById(ctx context.Context, id int64) (*co_entity.FdAccount, error) {
	if id == 0 {
		return nil, gerror.New(s.modules.T(ctx, "error_AccountId_NonNull"))
	}
	result, err := daoctl.GetByIdWithError[co_entity.FdAccount](s.dao.FdAccount.Ctx(ctx).Hook(daoctl.CacheHookHandler), id)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetAccountById_Failed"), s.dao.FdAccount.Table())
	}

	return result, nil
}

// UpdateAccountIsEnable 修改财务账号状态（是否启用：0禁用 1启用）
func (s *sFdAccount) UpdateAccountIsEnable(ctx context.Context, id int64, isEnabled int64) (bool, error) {
	account, err := daoctl.GetByIdWithError[co_entity.FdAccount](s.dao.FdAccount.Ctx(ctx).Hook(daoctl.CacheHookHandler), id)
	if account == nil || err != nil {
		return false, gerror.New(s.modules.T(ctx, "{#Account}{#error_Data_NotFound}"))
	}

	_, err = s.dao.FdAccount.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdAccount{Id: id}).Update(co_do.FdAccount{IsEnabled: isEnabled})
	if err != nil {
		return false, err
	}

	return true, nil
}

// HasAccountByName 根据账户名查询财务账户
func (s *sFdAccount) HasAccountByName(ctx context.Context, name string) (*co_entity.FdAccount, error) {
	data := &co_entity.FdAccount{}
	err := s.dao.FdAccount.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdAccount{Name: name}).Scan(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateAccountLimitState 修改财务账号的限制状态 （0不限制，1限制支出、2限制收入）
func (s *sFdAccount) UpdateAccountLimitState(ctx context.Context, id int64, limitState int64) (bool, error) {
	_, err := s.dao.FdAccount.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdAccount{Id: id}).Update(co_do.FdAccount{LimitState: limitState})
	if err != nil {
		return false, err
	}

	return true, nil
}

// QueryAccountListByUserId 获取指定用户的所有财务账号
func (s *sFdAccount) QueryAccountListByUserId(ctx context.Context, userId int64) (*co_model.AccountList, error) {
	accountList := co_model.AccountList{}

	if userId == 0 {
		return nil, gerror.New("用户id不能为空")
	}

	err := s.dao.FdAccount.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdAccount{UnionUserId: userId}).Scan(&accountList)

	if err != nil || len(accountList) <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_ThisUser_NotHas_Account"), s.dao.FdAccount.Table())
	}

	return &accountList, nil
}

// UpdateAccountBalance 修改财务账户余额(上下文, 财务账号id, 需要修改的钱数目, 版本, 收支类型)
func (s *sFdAccount) UpdateAccountBalance(ctx context.Context, accountId int64, amount int64, version int, inOutType int) (int64, error) {
	db := s.dao.FdAccount.Ctx(ctx)

	data := co_do.FdAccount{
		Version: gdb.Raw(s.dao.FdAccount.Columns().Version + "+1"),
	}

	if inOutType == 1 { // 收入
		// 余额 = 之前的余额 + 本次交易的余额
		data.Balance = gdb.Raw(s.dao.FdAccount.Columns().Balance + "+" + gconv.String(amount))
	} else if inOutType == 2 { // 支出
		// 余额 = 之前的余额 - 本次交易的余额
		data.Balance = gdb.Raw(s.dao.FdAccount.Columns().Balance + "-" + gconv.String(amount))
	}

	result, err := db.Data(data).Where(co_do.FdAccount{
		Id:      accountId,
		Version: version,
	}).Update()

	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()

	return affected, err
}

// GetAccountByUnionUserIdAndCurrencyCode 根据用户union_user_id和货币代码currency_code获取财务账号
func (s *sFdAccount) GetAccountByUnionUserIdAndCurrencyCode(ctx context.Context, unionUserId int64, currencyCode string) (*co_entity.FdAccount, error) {
	if unionUserId == 0 {
		return nil, gerror.New(s.modules.T(ctx, "error_Account_UnionUserId_NotNull"))
	}

	result := co_entity.FdAccount{}

	// 查找指定用户名下指定货币类型的财务账号
	err := s.dao.FdAccount.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdAccount{
		UnionUserId:  unionUserId,
		CurrencyCode: currencyCode,
	}).Scan(&result)

	return &result, err
}
