package finance

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

type sFdRecharge[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	TTFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	TR co_model.IFdRechargeRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		TTFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		TR,
	]
	dao *co_dao.XDao
}

func NewFdRecharge[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	TFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	TR co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) co_interface.IFdRecharge[TR] {
	result := &sFdRecharge[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		TFdInvoiceRes,
		ITFdInvoiceDetailRes,
		TR,
	]{
		modules: modules,
		dao:     modules.Dao(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sFdRecharge[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.IFdRechargeRes
	ret = &co_model.FdRechargeRes{}
	return ret.(TR)
}

func (s *sFdRecharge[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) GetAccountRechargeById(ctx context.Context, id int64) (TR, error) {
	var ret TR

	if id == 0 {
		return ret, gerror.New(s.modules.T(ctx, "error_Financial_AccountId_Failed"))
	}

	result, err := daoctl.GetByIdWithError[TR](s.modules.Dao().FdRecharge.Ctx(ctx), id)

	if err != nil {
		return ret, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Data_Get_Failed}"), s.dao.FdRecharge.Table())
	}

	return *result, nil
}

func (s *sFdRecharge[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) SetAccountRechargeAudit(ctx context.Context, id int64, state sys_enum.AuditAction, reply string) (bool, error) {
	if id == 0 {
		return false, gerror.New(s.modules.T(ctx, "error_Financial_AccountId_Failed"))
	}

	info, err := daoctl.GetByIdWithError[co_model.FdRecharge](s.modules.Dao().FdRecharge.Ctx(ctx), id)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Data_Get_Failed}"), s.dao.FdRecharge.Table())
	}

	if !base_funs.Contains([]int{
		sys_enum.Audit.AuditState.WaitReview.Code(),
		sys_enum.Audit.AuditState.PendingReview.Code(),
		sys_enum.Audit.AuditState.Reviewing.Code(),
	}, info.AuditState) {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#error_Financial_AccountRecharge_Audit_Failed}"), s.dao.FdRecharge.Table())
	}

	if state == sys_enum.Audit.AuditState.Reject && reply == "" {
		return false, gerror.New(s.modules.T(ctx, "error_Financial_AccountRecharge_Audit_Failed"))
	}

	data := co_do.FdRecharge{
		State:      base_funs.If(state == sys_enum.Audit.Action.Approve, sys_enum.Audit.AuditState.Reject, sys_enum.Audit.AuditState.Approve),
		AuditState: state,
		AuditReply: reply,
		UpdatedAt:  gtime.Now(),
	}

	err = s.dao.FdRecharge.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := s.dao.FdRecharge.Ctx(ctx).Where(s.dao.FdRecharge.Columns().Id, id).
			Data(data).Update()

		return err
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Data_Update_Failed}"), s.dao.FdRecharge.Table())
	}

	return true, nil
}

func (s *sFdRecharge[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) QueryAccountRecharge(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error) {
	// 过滤UnionMainId字段查询条件
	search = s.modules.Company().FilterUnionMainId(ctx, search)

	isExport := false
	if ctx.Value("isExport") == nil {
		r := g.RequestFromCtx(ctx)
		isExport = r.GetForm("isExport", false).Bool()
	} else {
		isExport = gconv.Bool(ctx.Value("isExport"))
	}

	result, err := daoctl.Query[TR](s.dao.FdRecharge.Ctx(ctx), search, isExport)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &base_model.CollectRes[TR]{}, nil
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Data_Get_Failed}"), s.dao.FdRecharge.Table())
	}

	return result, nil
}

func (s *sFdRecharge[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TFdInvoiceRes,
	ITFdInvoiceDetailRes,
	TR,
]) AccountRecharge(ctx context.Context, info *co_model.FdRecharge, createUser *sys_model.SysUser) (TR, error) {
	var ret TR

	if info.AccountId == 0 || info.CurrencyCode == "" || info.Amount <= 0 || info.RechargeMethod <= 0 || info.UserId <= 0 || createUser == nil {
		return ret, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#error_Parameter}"), s.dao.FdAccount.Table())
	}

	if info.Amount <= 0 {
		return ret, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#error_Parameter}"), s.dao.FdAccount.Table())
	}

	accountInfo, err := daoctl.GetByIdWithError[co_entity.FdAccount](s.dao.FdAccount.Ctx(ctx), info.AccountId)

	if accountInfo == nil || err != nil {
		return ret, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Data_NotFound}"), s.dao.FdAccount.Table())
	}

	employee, _ := s.modules.Employee().GetEmployeeById(ctx, info.UserId)

	data := kconv.Struct(info.FdRecharge, &co_do.FdRecharge{})
	data.Id = idgen.NextId()
	data.State = co_enum.Finance.RechargeState.Pending.Code()
	data.AuditState = sys_enum.Audit.AuditState.WaitReview.Code()
	data.RechargeMethod = info.RechargeMethod
	data.Username = createUser.Username
	data.UserId = createUser.Id
	data.CurrencyCode = accountInfo.CurrencyCode
	data.IpAddress = ghttp.RequestFromCtx(ctx).GetClientIp()
	data.UserAgent = ghttp.RequestFromCtx(ctx).Request.UserAgent()
	data.CreatedAt = gtime.Now()
	data.UpdatedAt = gtime.Now()
	data.PaymentAt = gtime.Now()

	if employee.Data() != nil {
		data.UnionMainId = employee.Data().UnionMainId
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		if info.RechargeMethod == co_enum.Finance.RechargeMethod.WeChat.Code() || info.RechargeMethod == co_enum.Finance.RechargeMethod.Alipay.Code() || info.RechargeMethod == co_enum.Finance.RechargeMethod.CloudPay.Code() {
			data.State = co_enum.Finance.RechargeState.Processing.Code()
		} else {
			data.State = co_enum.Finance.RechargeState.Awaiting.Code()
		}

		newData, err := info.OverrideDo.DoFactory(data)

		if err != nil {
			return err
		}

		affected, err := daoctl.InsertWithError(s.dao.FdRecharge.Ctx(ctx), newData)

		if affected == 0 || err != nil {
			return err
		}

		if info.ScreenshotId > 0 {
			// 从数据库中获取
			file, _ := sys_service.File().GetAnyFileById(ctx, gconv.Int64(info.ScreenshotId), "")

			// 如果文件不存在，则代表文件还在缓存中
			if file == nil {
				file, _ = sys_service.File().GetFileById(ctx, gconv.Int64(info.ScreenshotId), "")

				if file != nil {
					uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
					fileRecord, _ := sys_service.File().SaveFile(ctx, uploadPath+"/recharge/screenshot/"+gconv.String(info.UserId)+"/"+gconv.String(info.ScreenshotId)+file.Ext, file)
					data.ScreenshotId = fileRecord.Id
				}
			}
		}

		err = info.OverrideDo.DoSaved(data, newData)

		return err
	}); err != nil {
		return ret, err
	}

	return s.GetAccountRechargeById(ctx, gconv.Int64(data.Id))
}
