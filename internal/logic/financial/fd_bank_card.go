package financial

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_model"
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
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

// 银行卡管理
type sFdBankCard struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewFdBankCard(modules co_interface.IModules, xDao *co_dao.XDao) co_interface.IFdBankCard {
	return &sFdBankCard{
		modules: modules,
		dao:     xDao,
	}
}

// CreateBankCard 添加银行卡账号
func (s *sFdBankCard) CreateBankCard(ctx context.Context, info co_model.BankCardRegister, user *sys_entity.SysUser) (*co_entity.FdBankCard, error) {

	// 判断userid是否存在
	userInfo, err := sys_service.SysUser().GetSysUserById(ctx, info.UserId)
	if err != nil || userInfo == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#User} {#error_Data_Get_Failed}"), sys_dao.SysUser.Table())
	}

	// 判断银行卡是否重复
	bankCard, err := s.GetBankCardByCardNumber(ctx, info.CardNumber)

	if bankCard != nil || err == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard} {#error_AlreadyExist}"), s.dao.FdBankCard.Table())
	}

	bankCardInfo := co_do.FdBankCard{}
	gconv.Struct(info, &bankCardInfo)
	bankCardInfo.Id = idgen.NextId()

	// 当前用户创建的就是自己的银行卡账号
	bankCardInfo.UserId = user.Id

	// 添加银行卡
	_, err = s.dao.FdBankCard.Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(bankCardInfo).Insert()

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard} {#error_Add_Failed}"), s.dao.FdBankCard.Table())
	}

	return s.GetBankCardById(ctx, gconv.Int64(bankCardInfo.Id))
}

// GetBankCardById 根据银行卡id获取银行卡信息
func (s *sFdBankCard) GetBankCardById(ctx context.Context, id int64) (*co_entity.FdBankCard, error) {
	if id == 0 {
		return nil, gerror.New(s.modules.T(ctx, "{#BankCard} {#error_Id_NotNull}"))
	}
	result, err := daoctl.GetByIdWithError[co_entity.FdBankCard](s.dao.FdBankCard.Ctx(ctx), id)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetBankCardById_Failed"), s.dao.FdBankCard.Table())
	}

	return result, nil
}

// GetBankCardByCardNumber 根据银行卡号获取银行卡
func (s *sFdBankCard) GetBankCardByCardNumber(ctx context.Context, cardNumber string) (*co_entity.FdBankCard, error) {
	if cardNumber == "" {
		return nil, gerror.New(s.modules.T(ctx, "error_BankCardNumber_NonNull"))
	}

	bankCard := co_entity.FdBankCard{}

	err := s.dao.FdBankCard.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdBankCard{CardNumber: cardNumber}).Scan(&bankCard)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard} {#error_Data_Get_Failed}"), s.dao.FdBankCard.Table())

	}

	return &bankCard, nil
}

// UpdateBankCardState 修改银行卡状态 (0禁用 1正常)
func (s *sFdBankCard) UpdateBankCardState(ctx context.Context, bankCardId int64, state int) (bool, error) {
	bankCard, err := s.GetBankCardById(ctx, bankCardId)
	if err != nil || bankCard == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard} {#error_NonExist}"), s.dao.FdBankCard.Table())
	}

	// 修改状态
	result, err := s.dao.FdBankCard.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdBankCard{Id: bankCardId}).Update(co_do.FdBankCard{State: state})

	if err != nil || result == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard} {#error_State_Update_Failed}"), s.dao.FdBankCard.Table())
	}

	return true, nil
}

// DeleteBankCardById 删除银行卡 (标记删除: 标记删除的银行卡号，将记录ID的后6位附加到卡号尾部，用下划线隔开,并修改状态)
func (s *sFdBankCard) DeleteBankCardById(ctx context.Context, bankCardId int64) (bool, error) {
	var result sql.Result
	var err error

	bankCard, err := s.GetBankCardById(ctx, bankCardId)
	if err != nil || bankCard == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard} {#error_NonExist}"), s.dao.FdBankCard.Table())
	}

	err = s.dao.FdBankCard.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		bankCardIdLen := len(gconv.String(bankCard.Id))
		subId := gstr.SubStr(gconv.String(bankCard.Id), bankCardIdLen-6, 6)

		// 修改 1.修改状态为禁用0   2. 标记银行卡号：bankcardNum_bankCard.Id的后六位
		newBankCardNum := bankCard.CardNumber + "_" + subId

		result, err = s.dao.FdBankCard.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdBankCard{Id: bankCardId}).Data(co_do.FdBankCard{
			State:      0,
			CardNumber: newBankCardNum,
		}).Update()

		// 删除
		result, err = s.dao.FdBankCard.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdBankCard{Id: bankCardId}).Delete()

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil || result == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard} {#error_Delete_Failed}"), s.dao.FdBankCard.Table())
	}

	return true, nil
}

// QueryBankCardListByUserId 根据用户id查询银行卡列表
func (s *sFdBankCard) QueryBankCardListByUserId(ctx context.Context, userId int64) (*co_model.BankCardListRes, error) {
	if userId == 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#User} {#error_Id_NotNull}"), sys_dao.SysUser.Table())
	}

	result, err := daoctl.Query[co_entity.FdBankCard](s.dao.FdBankCard.Ctx(ctx).Hook(daoctl.CacheHookHandler), &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0), sys_model.FilterInfo{
			Field: s.dao.FdBankCard.Columns().UserId,
			Where: "=",
			Value: userId,
		}),
		Pagination: sys_model.Pagination{
			PageNum:  1,
			PageSize: 20,
		},
	}, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_QueryBankCardList_Failed"), s.dao.FdBankCard.Table())
	}

	return (*co_model.BankCardListRes)(result), nil
}
