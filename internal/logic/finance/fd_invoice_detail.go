package finance

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"

	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// 发票详情
type sFdInvoiceDetail[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	TR co_model.IFdInvoiceDetailRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		TR,
	]
	dao *co_dao.XDao
}

func NewFdInvoiceDetail[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	TR co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) co_interface.IFdInvoiceDetail[TR] {
	result := &sFdInvoiceDetail[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		TR,
	]{
		modules: modules,
		dao:     modules.Dao(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sFdInvoiceDetail[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.IFdInvoiceDetailRes
	ret = &co_model.FdInvoiceDetailRes{}
	return ret.(TR)
}

// CreateInvoiceDetail 创建发票详情，相当于创建审核列表，审核是人工审核
func (s *sFdInvoiceDetail[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) CreateInvoiceDetail(ctx context.Context, info co_model.FdInvoiceDetailRegister) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 检查指定参数是否为空
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return response, err
	}

	// 获取发票详情

	// 创建发票审核详情记录
	data := s.FactoryMakeResponseInstance()
	gconv.Struct(info, &data)

	data.Data().Id = idgen.NextId()
	// 设置审核状态为待审核
	data.Data().State = co_enum.Invoice.AuditType.WaitReview.Code()
	data.Data().CreatedBy = sessionUser.Id

	newData := kconv.Struct(data.Data(), &co_do.FdInvoiceDetail{})

	// 重载Do模型
	doData, err := info.OverrideDo.DoFactory(*newData)
	if err != nil {
		return response, err
	}

	result, err := s.dao.FdInvoiceDetail.Ctx(ctx).Data(doData).Insert()

	if err != nil || result == nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#InvoiceDetail}{#error_Create_Failed}"), s.dao.FdInvoiceDetail.Table())
	}

	return s.GetInvoiceDetailById(ctx, data.Data().Id)
}

// GetInvoiceDetailById 根据id获取发票详情
func (s *sFdInvoiceDetail[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) GetInvoiceDetailById(ctx context.Context, id int64) (response TR, err error) {
	if id == 0 {
		return response, gerror.New(s.modules.T(ctx, "{#InvoiceDetail}{#error_Id_NotNull}"))
	}

	result, err := daoctl.GetByIdWithError[TR](s.dao.FdInvoiceDetail.Ctx(ctx), id)

	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetInvoiceDetailById_Failed"), s.dao.FdInvoiceDetail.Table())
	}

	return *result, nil
}

// MakeInvoiceDetail 开票
func (s *sFdInvoiceDetail[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) MakeInvoiceDetail(ctx context.Context, invoiceDetailId int64, makeInvoiceDetail co_model.FdMakeInvoiceDetail) (res bool, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	invoiceDetailInfo, err := s.GetInvoiceDetailById(ctx, invoiceDetailId)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetailId_Error"), s.dao.FdInvoiceDetail.Table())
	}

	// 校验状态是否为待开票
	if invoiceDetailInfo.Data().State != co_enum.Invoice.State.WaitForInvoice.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_InvoiceDetail_Create_Failed}{#error_InvoiceState_Error}"), s.dao.FdInvoiceDetail.Table())
	}

	// 判断是纸质发票还是电子发票 然后添加审核过后的数据
	if makeInvoiceDetail.Type == 1 { // 电子发票
		_, err = s.dao.FdInvoiceDetail.Ctx(ctx).OmitNilData().Data(co_do.FdInvoiceDetail{
			MakeType:   makeInvoiceDetail.MakeType,
			MakeUserId: makeInvoiceDetail.MakeUserId,
			Email:      makeInvoiceDetail.Email,
			State:      co_enum.Invoice.State.Success.Code(),
			MakeAt:     gtime.Now(),
			UpdatedBy:  sessionUser.Id,
		}).Where(co_do.FdInvoiceDetail{
			Id: invoiceDetailInfo.Data().Id,
		}).Update()

	} else if makeInvoiceDetail.Type == 2 { // 纸质发票
		_, err = s.dao.FdInvoiceDetail.Ctx(ctx).OmitNilData().Data(co_do.FdInvoiceDetail{
			MakeType:      makeInvoiceDetail.MakeType,
			MakeUserId:    makeInvoiceDetail.MakeUserId,
			CourierName:   makeInvoiceDetail.CourierName,
			CourierNumber: makeInvoiceDetail.CourierNumber,
			State:         co_enum.Invoice.State.Success.Code(),
			MakeAt:        gtime.Now(),
			UpdatedBy:     sessionUser.Id,
		}).Where(co_do.FdInvoiceDetail{
			Id: invoiceDetailInfo.Data().Id,
		}).Update()

	} else {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_InvoiceDetail_Create_Failed}{#error_InvoiceState_Error}"), s.dao.FdInvoiceDetail.Table())
	}

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetail_Update_Failed"), s.dao.FdInvoiceDetail.Table())
	}

	return true, nil
}

// AuditInvoiceDetail 审核发票
func (s *sFdInvoiceDetail[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) AuditInvoiceDetail(ctx context.Context, invoiceDetailId int64, auditInfo co_model.FdInvoiceAuditInfo) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 审核行仅允许 co_enum_invoice.State.WaitForInvoice 和 co_enum_invoice.State.Failure 待开票、开票失败
	if auditInfo.State != co_enum.Invoice.State.WaitForInvoice.Code() && auditInfo.State != co_enum.Invoice.State.Failure.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_InvoiceDetail_AuditState_Error"), s.dao.FdInvoiceDetail.Table())
	}

	if auditInfo.State == co_enum.Invoice.State.Failure.Code() && auditInfo.ReplyMsg == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_AuditReply_MustHasMsg"), s.dao.FdInvoiceDetail.Table())
	}

	invoiceDetailInfo, err := s.GetInvoiceDetailById(ctx, invoiceDetailId)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_InvoiceDetailId_Error"), s.dao.FdInvoiceDetail.Table())
	}

	// 代表已审过的
	if invoiceDetailInfo.Data().State > co_enum.Invoice.State.WaitAudit.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_InvoiceDetail_RepeatSubmit"), s.dao.FdInvoiceDetail.Table())
	}

	// 添加审核过后的数据
	_, err = s.dao.FdInvoiceDetail.Ctx(ctx).OmitNilData().Data(co_do.FdInvoiceDetail{
		AuditUserId:   auditInfo.AuditUserId,
		AuditReplyMsg: auditInfo.ReplyMsg,
		State:         auditInfo.State,
		AuditAt:       gtime.Now(),
		UpdatedBy:     sessionUser.Id,
	}).Where(co_do.FdInvoiceDetail{
		Id: invoiceDetailInfo.Data().Id,
	}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetail_Update_Failed"), s.dao.FdInvoiceDetail.Table())
	}

	return true, nil
}

// QueryInvoiceDetailListByInvoiceId 根据发票抬头，获取已开票的发票详情列表
func (s *sFdInvoiceDetail[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) QueryInvoiceDetailListByInvoiceId(ctx context.Context, invoiceId int64) (*base_model.CollectRes[TR], error) {
	result, err := daoctl.Query[TR](s.dao.FdInvoiceDetail.Ctx(ctx), &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
			Field: s.dao.FdInvoiceDetail.Columns().FdInvoiceId,
			Where: "=",
			Value: invoiceId,
		}),
		Pagination: base_model.Pagination{
			PageNum:  1,
			PageSize: -1,
		},
	}, false)

	return result, err
}

// DeleteInvoiceDetail 标记删除发票详情
func (s *sFdInvoiceDetail[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) DeleteInvoiceDetail(ctx context.Context, id int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断是否存在该发票
	_, err := s.GetInvoiceDetailById(ctx, id)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetailId_Error"), s.dao.FdInvoiceDetail.Table())
	}

	err = s.dao.FdInvoiceDetail.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 状态修改为已撤消，
		_, err = s.dao.FdInvoiceDetail.Ctx(ctx).
			Where(co_do.FdInvoiceDetail{Id: id}).
			Update(co_do.FdInvoiceDetail{
				State:     co_enum.Invoice.State.Cancel.Code(),
				UpdatedBy: sessionUser.Id,
				DeletedBy: sessionUser.Id,
				DeletedAt: gtime.Now(),
			})
		if err != nil {
			return err
		}

		// 删除
		_, err = s.dao.FdInvoiceDetail.Ctx(ctx).Where(co_do.FdInvoiceDetail{Id: id}).Delete()

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetail_Delete_Failed"), s.dao.FdInvoiceDetail.Table())
	}

	return err == nil, err
}

// QueryInvoiceDetail 根据限定的条件查询发票列表
func (s *sFdInvoiceDetail[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	TR,
]) QueryInvoiceDetail(ctx context.Context, info *base_model.SearchParams, userId int64, unionMainId int64) (*base_model.CollectRes[TR], error) {
	newFields := make([]base_model.FilterInfo, 0)
	// 筛选条件强制指定所属用户
	if unionMainId != 0 {
		newFields = append(newFields, base_model.FilterInfo{
			Field: s.dao.FdInvoiceDetail.Columns().UnionMainId, // type
			Where: "=",
			Value: unionMainId,
		})
	}

	if userId != 0 {
		newFields = append(newFields, base_model.FilterInfo{
			Field: s.dao.FdInvoiceDetail.Columns().UserId,
			Where: "=",
			Value: userId,
		})
	}

	if info != nil {
		for _, field := range info.Filter {
			if field.Field != s.dao.FdInvoiceDetail.Columns().UserId {
				newFields = append(newFields, field)
			}
		}
	}

	info.Filter = newFields

	// 如果没有查询条件，那就会查询出来所有
	if unionMainId == 0 {
		info = s.modules.Company().FilterUnionMainId(ctx, info)
	}

	if userId == 0 {
		info.Filter = append(info.Filter, base_model.FilterInfo{
			Field: s.dao.FdInvoiceDetail.Columns().UserId,
			Where: "=",
			Value: sys_service.SysSession().Get(ctx).JwtClaimsUser.Id,
		})
	}

	result, err := daoctl.Query[TR](s.dao.FdInvoiceDetail.Ctx(ctx), info, false)

	return result, err
}
