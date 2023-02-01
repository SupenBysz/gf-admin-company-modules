package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

// 发票详情
type sFdInvoiceDetail struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewFdInvoiceDetail(modules co_interface.IModules, xDao *co_dao.XDao) co_interface.IFdInvoiceDetail {
	return &sFdInvoiceDetail{
		modules: modules,
		dao:     xDao,
	}
}

// CreateInvoiceDetail 创建发票详情，相当于创建审核列表，审核是人工审核
func (s *sFdInvoiceDetail) CreateInvoiceDetail(ctx context.Context, info co_model.FdInvoiceDetailRegister) (*co_entity.FdInvoiceDetail, error) {
	// 检查指定参数是否为空
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return nil, err
	}

	// 获取发票详情

	// 创建发票审核详情记录
	data := co_entity.FdInvoiceDetail{}
	gconv.Struct(info, &data)

	data.Id = idgen.NextId()
	// 设置审核状态为待审核
	data.State = co_enum.Invoice.AuditType.WaitReview.Code()

	result, err := s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(data).Insert()

	if err != nil || result == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#InvoiceDetail} {#error_Create_Failed}"), s.dao.FdInvoiceDetail.Table())
	}

	return s.GetInvoiceDetailById(ctx, data.Id)
}

// GetInvoiceDetailById 根据id获取发票详情
func (s *sFdInvoiceDetail) GetInvoiceDetailById(ctx context.Context, id int64) (*co_entity.FdInvoiceDetail, error) {
	if id == 0 {
		return nil, gerror.New(s.modules.T(ctx, "{#InvoiceDetail} {#error_Id_NotNull}"))
	}

	result, err := daoctl.GetByIdWithError[co_entity.FdInvoiceDetail](s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler), id)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetInvoiceDetailById_Failed"), s.dao.FdInvoiceDetail.Table())
	}

	invoiceDetail := co_entity.FdInvoiceDetail{}
	gconv.Struct(result, &invoiceDetail)

	return &invoiceDetail, nil
}

// MakeInvoiceDetail 开票
func (s *sFdInvoiceDetail) MakeInvoiceDetail(ctx context.Context, invoiceDetailId int64, makeInvoiceDetail co_model.FdMakeInvoiceDetail) (res bool, err error) {
	invoiceDetailInfo, err := s.GetInvoiceDetailById(ctx, invoiceDetailId)
	if err != nil || invoiceDetailInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetailId_Error"), s.dao.FdInvoiceDetail.Table())
	}

	// 校验状态是否为待开票
	if invoiceDetailInfo.State != co_enum.Invoice.State.WaitForInvoice.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_InvoiceDetail_Create_Failed} {#error_InvoiceState_Error}"), s.dao.FdInvoiceDetail.Table())
	}

	// 判断是纸质发票还是电子发票 然后添加审核过后的数据
	if makeInvoiceDetail.Type == 1 { // 电子发票
		_, err = s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).OmitNilData().Data(co_do.FdInvoiceDetail{
			MakeType:   makeInvoiceDetail.MakeType,
			MakeUserId: makeInvoiceDetail.MakeUserId,
			Email:      makeInvoiceDetail.Email,
			State:      co_enum.Invoice.State.Success.Code(),
			MakeAt:     gtime.Now(),
		}).Where(co_do.FdInvoiceDetail{
			Id: invoiceDetailInfo.Id,
		}).Update()

	} else if makeInvoiceDetail.Type == 2 { // 纸质发票
		_, err = s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).OmitNilData().Data(co_do.FdInvoiceDetail{
			MakeType:      makeInvoiceDetail.MakeType,
			MakeUserId:    makeInvoiceDetail.MakeUserId,
			CourierName:   makeInvoiceDetail.CourierName,
			CourierNumber: makeInvoiceDetail.CourierNumber,
			State:         co_enum.Invoice.State.Success.Code(),
			MakeAt:        gtime.Now(),
		}).Where(co_do.FdInvoiceDetail{
			Id: invoiceDetailInfo.Id,
		}).Update()

	} else {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_InvoiceDetail_Create_Failed} {#error_InvoiceState_Error}"), s.dao.FdInvoiceDetail.Table())
	}

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetail_Update_Failed"), s.dao.FdInvoiceDetail.Table())
	}

	return true, nil
}

// AuditInvoiceDetail 审核发票
func (s *sFdInvoiceDetail) AuditInvoiceDetail(ctx context.Context, invoiceDetailId int64, auditInfo co_model.FdInvoiceAuditInfo) (bool, error) {
	// 审核行仅允许 co_enum_invoice.State.WaitForInvoice 和 co_enum_invoice.State.Failure 待开票、开票失败
	if auditInfo.State != co_enum.Invoice.State.WaitForInvoice.Code() && auditInfo.State != co_enum.Invoice.State.Failure.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_InvoiceDetail_AuditState_Error"), s.dao.FdInvoiceDetail.Table())
	}

	if auditInfo.State == co_enum.Invoice.State.Failure.Code() && auditInfo.ReplyMsg == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_AuditReplay_MustHasMsg"), s.dao.FdInvoiceDetail.Table())
	}

	invoiceDetailInfo, err := s.GetInvoiceDetailById(ctx, invoiceDetailId)
	if err != nil || invoiceDetailInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_InvoiceDetailId_Error"), s.dao.FdInvoiceDetail.Table())
	}

	// 代表已审过的
	if invoiceDetailInfo.State > co_enum.Invoice.State.WaitAudit.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_InvoiceDetail_RepeatSubmit"), s.dao.FdInvoiceDetail.Table())
	}

	// 添加审核过后的数据
	_, err = s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).OmitNilData().Data(co_do.FdInvoiceDetail{
		AuditUserId:   auditInfo.AuditUserId,
		AuditReplyMsg: auditInfo.ReplyMsg,
		State:         auditInfo.State,
		AuditAt:       gtime.Now(),
	}).Where(co_do.FdInvoiceDetail{
		Id: invoiceDetailInfo.Id,
	}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetail_Update_Failed"), s.dao.FdInvoiceDetail.Table())
	}

	return true, nil
}

// QueryInvoiceDetailListByInvoiceId 根据发票抬头，获取已开票的发票详情列表
func (s *sFdInvoiceDetail) QueryInvoiceDetailListByInvoiceId(ctx context.Context, invoiceId int64) (*co_model.FdInvoiceDetailListRes, error) {
	result, err := daoctl.Query[co_entity.FdInvoiceDetail](s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler), &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0), sys_model.FilterInfo{
			Field: s.dao.FdInvoiceDetail.Columns().FdInvoiceId,
			Where: "=",
			Value: invoiceId,
		}),
		Pagination: sys_model.Pagination{
			PageNum:  1,
			PageSize: -1,
		},
	}, false)

	return (*co_model.FdInvoiceDetailListRes)(result), err
}

// DeleteInvoiceDetail 标记删除发票详情
func (s *sFdInvoiceDetail) DeleteInvoiceDetail(ctx context.Context, id int64) (bool, error) {
	// 判断是否存在该发票
	invoice, err := s.GetInvoiceDetailById(ctx, id)
	if err != nil || invoice == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_InvoiceDetailId_Error"), s.dao.FdInvoiceDetail.Table())
	}

	err = s.dao.FdInvoiceDetail.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 状态修改为已撤消，
		_, err = s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).
			Where(co_do.FdInvoiceDetail{Id: id}).
			Update(co_do.FdInvoiceDetail{State: co_enum.Invoice.State.Cancel.Code()})
		if err != nil {
			return err
		}

		// 删除
		_, err = s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdInvoiceDetail{Id: id}).Delete()

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
func (s *sFdInvoiceDetail) QueryInvoiceDetail(ctx context.Context, info *sys_model.SearchParams, userId int64, unionMainId int64) (*co_model.FdInvoiceDetailListRes, error) {
	newFields := make([]sys_model.FilterInfo, 0)
	// 筛选条件强制指定所属用户
	if unionMainId != 0 {
		newFields = append(newFields, sys_model.FilterInfo{
			Field: s.dao.FdInvoiceDetail.Columns().UnionMainId, // type
			Where: "=",
			Value: unionMainId,
		})
	}

	if userId != 0 {
		newFields = append(newFields, sys_model.FilterInfo{
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

	result, err := daoctl.Query[co_entity.FdInvoiceDetail](s.dao.FdInvoiceDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler), info, false)

	return (*co_model.FdInvoiceDetailListRes)(result), err
}
