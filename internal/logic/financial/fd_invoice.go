package financial

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/yitter/idgenerator-go/idgen"
	"reflect"
)

type sFdInvoice[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	TR co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
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
		TR,
		ITFdInvoiceDetailRes,
	]
	dao *co_dao.XDao
	// hookArr       []hookInfo
}

func NewFdInvoice[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	TR co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	TR,
	ITFdInvoiceDetailRes,
]) co_interface.IFdInvoice[TR] {
	result := &sFdInvoice[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		TR,
		ITFdInvoiceDetailRes,
	]{
		modules: modules,
		dao:     modules.Dao(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sFdInvoice[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	TR,
	ITFdInvoiceDetailRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.IFdInvoiceRes
	ret = &co_model.FdInvoiceRes{}
	return ret.(TR)
}

// CreateInvoice 添加发票抬头
func (s *sFdInvoice[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	TR,
	ITFdInvoiceDetailRes,
]) CreateInvoice(ctx context.Context, info co_model.FdInvoiceRegister) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 一个公司只能有一个发票抬头
	selectInvoice, err := s.GetFdInvoiceByTaxId(ctx, info.TaxId)
	if err == nil {
		if info.UnionMainId == selectInvoice.Data().UnionMainId {
			return response, gerror.New(s.modules.T(ctx, "error_OneCompany_CanRegisterOnlyOne_Invoice"))
		}
	}

	// 判断审核状态
	if info.State == co_enum.Invoice.AuditType.Reject.Code() && info.AuditReplyMsg == "" {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_AuditReply_MustHasMsg"), s.dao.FdInvoice.Table())
	}

	// 创建发票
	invoice := s.FactoryMakeResponseInstance()
	gconv.Struct(info, invoice.Data())
	invoice.Data().Id = idgen.NextId()
	invoice.Data().AuditUserId = 0
	invoice.Data().State = co_enum.Invoice.AuditType.WaitReview.Code()

	invoice.Data().CreatedBy = sessionUser.Id

	data := kconv.Struct(invoice.Data(), &co_do.FdInvoice{})
	doData, err := info.OverrideDo.MakeDo(*data)
	if err != nil {
		return response, err
	}

	_, err = s.dao.FdInvoice.Ctx(ctx).Data(doData).Insert()
	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Invoice_Create_Failed"), s.dao.FdInvoice.Table())
	}

	return s.GetInvoiceById(ctx, invoice.Data().Id)
}

// GetInvoiceById 根据id获取发票
func (s *sFdInvoice[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	TR,
	ITFdInvoiceDetailRes,
]) GetInvoiceById(ctx context.Context, id int64) (response TR, err error) {
	if id == 0 {
		return response, gerror.New(s.modules.T(ctx, "error_Id_NotNull"))
	}

	result, err := daoctl.GetByIdWithError[TR](s.dao.FdInvoice.Ctx(ctx), id)

	if result == nil || err != nil {
		return response, sys_service.SysLogs().InfoSimple(ctx, err, s.modules.T(ctx, "error_NotNowHas_Invoice"), s.dao.FdInvoice.Table())
	}

	return *result, nil
}

// QueryInvoiceList 获取发票抬头列表
func (s *sFdInvoice[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	TR,
	ITFdInvoiceDetailRes,
]) QueryInvoiceList(ctx context.Context, info *base_model.SearchParams, userId int64) (*base_model.CollectRes[TR], error) {
	newFields := make([]base_model.FilterInfo, 0)
	// 筛选条件强制指定所属用户
	newFields = append(newFields, base_model.FilterInfo{
		Field: s.dao.FdInvoice.Columns().UserId, // type
		Where: "=",
		Value: userId,
	})

	if info != nil {
		// 排除搜索参数中指定的所属用户参数
		for _, field := range info.Filter {
			if field.Field != s.dao.FdInvoice.Columns().UserId {
				newFields = append(newFields, field)
			}
		}
	}
	info.Filter = newFields

	result, err := daoctl.Query[TR](s.dao.FdInvoice.Ctx(ctx), info, false)

	return result, err
}

// DeletesFdInvoiceById 删除发票抬头
func (s *sFdInvoice[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	TR,
	ITFdInvoiceDetailRes,
]) DeletesFdInvoiceById(ctx context.Context, invoiceId int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	var result sql.Result
	var err error

	invoice, err := s.GetInvoiceById(ctx, invoiceId)
	if err != nil || reflect.ValueOf(invoice).IsNil() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Invoice}{#error_NonExist}"), s.dao.FdInvoice.Table())
	}

	err = s.dao.FdInvoice.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 添加删除记录留痕
		result, err = s.dao.FdInvoice.Ctx(ctx).Where(co_do.FdInvoice{Id: invoice.Data().Id}).Update(co_do.FdInvoice{
			DeletedBy: sessionUser.Id,
			DeletedAt: gtime.Now(),
		})

		// 删除
		result, err = s.dao.FdInvoice.Ctx(ctx).Where(co_do.FdInvoice{Id: invoice.Data().Id}).Delete()

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil || result == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Inovice}{#error_Delete_Failed}"), s.dao.FdInvoice.Table())
	}

	return true, nil
}

// GetFdInvoiceByTaxId 根据纳税识别号获取发票抬头信息
func (s *sFdInvoice[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	TR,
	ITFdInvoiceDetailRes,
]) GetFdInvoiceByTaxId(ctx context.Context, taxId string) (response TR, err error) {
	if taxId == "" {
		return response, gerror.New(s.modules.T(ctx, "error_TaxId_NotNull"))
	}
	result := s.FactoryMakeResponseInstance()

	err = s.dao.FdInvoice.Ctx(ctx).Where(co_do.FdInvoice{TaxId: taxId}).Scan(result.Data())
	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Invoice}{#error_Data_Get_Failed}"), s.dao.FdInvoice.Table())
	}

	return result, nil
}
