package financial

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/yitter/idgenerator-go/idgen"
)

type sFdInvoice struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
	// hookArr       []hookInfo
}

func NewFdInvoice(modules co_interface.IModules) co_interface.IFdInvoice {
	return &sFdInvoice{
		modules: modules,
		dao:     modules.Dao(),
		// hookArr:       make([]hookInfo, 0),
	}
}

// CreateInvoice 添加发票抬头
func (s *sFdInvoice) CreateInvoice(ctx context.Context, info co_model.FdInvoiceRegister) (*co_entity.FdInvoice, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 一个公司只能有一个发票抬头
	invoice, _ := s.GetFdInvoiceByTaxId(ctx, info.TaxId)
	if invoice != nil {
		if info.UnionMainId == invoice.UnionMainId {
			return nil, gerror.New(s.modules.T(ctx, "error_OneCompany_CanRegisterOnlyOne_Invoice"))
		}
	}

	// 判断审核状态
	if info.State == co_enum.Invoice.AuditType.Reject.Code() && info.AuditReplayMsg == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_AuditReplay_MustHasMsg"), s.dao.FdInvoice.Table())
	}

	// 创建发票
	data := co_entity.FdInvoice{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()
	data.AuditUserId = 0
	data.State = co_enum.Invoice.AuditType.WaitReview.Code()

	data.CreatedBy = sessionUser.Id

	_, err := s.dao.FdInvoice.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Invoice_Create_Failed"), s.dao.FdInvoice.Table())
	}

	return s.GetInvoiceById(ctx, data.Id)
}

// GetInvoiceById 根据id获取发票
func (s *sFdInvoice) GetInvoiceById(ctx context.Context, id int64) (*co_entity.FdInvoice, error) {
	if id == 0 {
		return nil, gerror.New(s.modules.T(ctx, "error_Id_NotNull"))
	}

	result, err := daoctl.GetByIdWithError[co_entity.FdInvoice](s.dao.FdInvoice.Ctx(ctx), id)

	if result == nil || err != nil {
		return nil, sys_service.SysLogs().InfoSimple(ctx, err, s.modules.T(ctx, "error_NotNowHas_Invoice"), s.dao.FdInvoice.Table())
	}

	return result, nil
}

// QueryInvoiceList 获取发票抬头列表
func (s *sFdInvoice) QueryInvoiceList(ctx context.Context, info *base_model.SearchParams, userId int64) (*co_model.FdInvoiceListRes, error) {
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

	result, err := daoctl.Query[co_entity.FdInvoice](s.dao.FdInvoice.Ctx(ctx), info, false)

	return (*co_model.FdInvoiceListRes)(result), err
}

// DeletesFdInvoiceById 删除发票抬头
func (s *sFdInvoice) DeletesFdInvoiceById(ctx context.Context, invoiceId int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	var result sql.Result
	var err error

	invoice, err := s.GetInvoiceById(ctx, invoiceId)
	if err != nil || invoice == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Invoice}{#error_NonExist}"), s.dao.FdInvoice.Table())
	}

	err = s.dao.FdInvoice.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 添加删除记录留痕
		result, err = s.dao.FdInvoice.Ctx(ctx).Where(co_do.FdInvoice{Id: invoice.Id}).Update(co_do.FdInvoice{
			DeletedBy: sessionUser.Id,
			DeletedAt: gtime.Now(),
		})

		// 删除
		result, err = s.dao.FdInvoice.Ctx(ctx).Where(co_do.FdInvoice{Id: invoice.Id}).Delete()

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
func (s *sFdInvoice) GetFdInvoiceByTaxId(ctx context.Context, taxId string) (*co_entity.FdInvoice, error) {
	if taxId == "" {
		return nil, gerror.New(s.modules.T(ctx, "error_TaxId_NotNull"))
	}
	result := co_entity.FdInvoice{}

	err := s.dao.FdInvoice.Ctx(ctx).Where(co_do.FdInvoice{TaxId: taxId}).Scan(&result)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Invoice}{#error_Data_Get_Failed}"), s.dao.FdInvoice.Table())
	}

	return &result, nil
}
