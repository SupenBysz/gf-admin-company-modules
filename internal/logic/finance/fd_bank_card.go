package finance

import (
	"context"
	"database/sql"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/utility/kconv"

	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"

	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 银行卡管理
type sFdBankCard[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillsRes co_model.IFdAccountBillsRes,
	TR co_model.IFdBankCardRes,
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
		ITFdAccountBillsRes,
		TR,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]
	dao *co_dao.XDao
}

func NewFdBankCard[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillsRes co_model.IFdAccountBillsRes,
	TR co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) co_interface.IFdBankCard[TR] {
	result := &sFdBankCard[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillsRes,
		TR,
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

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sFdBankCard[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.IFdBankCardRes
	ret = &co_model.FdBankCardRes{}
	return ret.(TR)
}

// CreateBankCard 添加银行卡账号
func (s *sFdBankCard[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) CreateBankCard(ctx context.Context, info co_model.BankCardRegister, createUser *sys_model.SysUser) (response TR, err error) {

	//sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断userid是否存在
	if info.UserId != 0 {
		userInfo, err := sys_service.SysUser().GetSysUserById(ctx, info.UserId)
		if err != nil || userInfo == nil {
			return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#User}{#error_Data_Get_Failed}"), sys_dao.SysUser.Table())
		}
	}
	// 判断银行卡是否重复
	_, err = s.GetBankCardByCardNumber(ctx, info.CardNumber)

	if err == nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard}{#error_AlreadyExist}"), s.dao.FdBankCard.Table())
	}

	bankCardInfo := s.FactoryMakeResponseInstance()
	gconv.Struct(info, bankCardInfo.Data())
	bankCardInfo.Data().Id = idgen.NextId()

	// 当前用户创建的就是自己的银行卡账号
	//bankCardInfo.Data().UserId = user.Id
	bankCardInfo.Data().UserId = info.UserId
	// 默认状态正常
	bankCardInfo.Data().State = 1
	bankCardInfo.Data().CreatedAt = gtime.Now()
	bankCardInfo.Data().CreatedBy = createUser.Id

	data := kconv.Struct(bankCardInfo.Data(), &co_do.FdBankCard{})

	// 重载Do模型
	doData, err := info.OverrideDo.DoFactory(*data)
	if err != nil {
		return response, err
	}

	// 添加银行卡
	_, err = s.dao.FdBankCard.Ctx(ctx).Data(doData).Insert()

	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard}{#error_Add_Failed}"), s.dao.FdBankCard.Table())
	}

	return s.GetBankCardById(ctx, gconv.Int64(bankCardInfo.Data().Id))
}

// GetBankCardById 根据银行卡id获取银行卡信息
func (s *sFdBankCard[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetBankCardById(ctx context.Context, id int64) (response TR, err error) {
	if id == 0 {
		return response, gerror.New(s.modules.T(ctx, "{#BankCard}{#error_Id_NotNull}"))
	}
	data, err := daoctl.GetByIdWithError[TR](s.dao.FdBankCard.Ctx(ctx), id)

	if err != nil || data == nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetBankCardById_Failed"), s.dao.FdBankCard.Table())
	}

	return *data, nil
}

// GetBankCardByCardNumber 根据银行卡号获取银行卡
func (s *sFdBankCard[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetBankCardByCardNumber(ctx context.Context, cardNumber string) (response TR, err error) {
	if cardNumber == "" {
		return response, gerror.New(s.modules.T(ctx, "error_BankCardNumber_NonNull"))
	}

	response = s.FactoryMakeResponseInstance()

	err = s.dao.FdBankCard.Ctx(ctx).Where(co_do.FdBankCard{CardNumber: cardNumber}).Scan(response.Data())
	if err != nil {
		var ret TR
		return ret, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard}{#error_Data_Get_Failed}"), s.dao.FdBankCard.Table())

	}

	return response, nil
}

// UpdateBankCardState 修改银行卡状态 (0禁用 1正常)
func (s *sFdBankCard[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateBankCardState(ctx context.Context, bankCardId int64, state int) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	_, err := s.GetBankCardById(ctx, bankCardId)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard}{#error_NonExist}"), s.dao.FdBankCard.Table())
	}

	// 修改状态
	result, err := s.dao.FdBankCard.Ctx(ctx).Where(co_do.FdBankCard{Id: bankCardId}).Update(co_do.FdBankCard{
		State:     state,
		UpdatedBy: sessionUser.Id,
		UpdatedAt: gtime.Now(),
	})

	if err != nil || result == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard}{#error_State_Update_Failed}"), s.dao.FdBankCard.Table())
	}

	return true, nil
}

// DeleteBankCardById 删除银行卡 (标记删除: 标记删除的银行卡号，将记录ID的后6位附加到卡号尾部，用下划线隔开,并修改状态)
func (s *sFdBankCard[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) DeleteBankCardById(ctx context.Context, bankCardId int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	var result sql.Result
	var err error

	bankCard, err := s.GetBankCardById(ctx, bankCardId)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard}{#error_NonExist}"), s.dao.FdBankCard.Table())
	}

	err = s.dao.FdBankCard.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		bankCardIdLen := len(gconv.String(bankCard.Data().Id))
		subId := gstr.SubStr(gconv.String(bankCard.Data().Id), bankCardIdLen-6, 6)

		// 修改 1.修改状态为禁用0   2. 标记银行卡号：bankcardNum_bankCard.Id的后六位
		newBankCardNum := bankCard.Data().CardNumber + "_" + subId

		result, err = s.dao.FdBankCard.Ctx(ctx).Where(co_do.FdBankCard{Id: bankCardId}).Data(co_do.FdBankCard{
			State:      0,
			CardNumber: newBankCardNum,
			DeletedBy:  sessionUser.Id,
			UpdatedBy:  sessionUser.Id,
		}).Update()

		// 删除
		result, err = s.dao.FdBankCard.Ctx(ctx).Where(co_do.FdBankCard{Id: bankCardId}).Delete()

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil || result == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#BankCard}{#error_Delete_Failed}"), s.dao.FdBankCard.Table())
	}

	return true, nil
}

// QueryBankCardListByUserId 根据用户id查询银行卡列表
func (s *sFdBankCard[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillsRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryBankCardListByUserId(ctx context.Context, userId int64) (*base_model.CollectRes[TR], error) {
	if userId == 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#User}{#error_Id_NotNull}"), sys_dao.SysUser.Table())
	}

	result, err := daoctl.Query[TR](s.dao.FdBankCard.Ctx(ctx), &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
			Field: s.dao.FdBankCard.Columns().UserId,
			Where: "=",
			Value: userId,
		}),
		Pagination: base_model.Pagination{
			PageNum:  1,
			PageSize: 20,
		},
	}, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_QueryBankCardList_Failed"), s.dao.FdBankCard.Table())
	}

	return result, nil
}
