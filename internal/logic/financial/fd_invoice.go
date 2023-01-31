package financial

import (
	"context"
	"database/sql"
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
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

type sFdInvoice struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
	// hookArr       []hookInfo
}

func NewFdInvoice(modules co_interface.IModules, xDao *co_dao.XDao) co_interface.IFdInvoice {
	return &sFdInvoice{
		modules: modules,
		dao:     xDao,
		// hookArr:       make([]hookInfo, 0),
	}
}

// CreateInvoice 添加发票抬头
func (s *sFdInvoice) CreateInvoice(ctx context.Context, info co_model.FdInvoiceRegister) (*co_entity.FdInvoice, error) {
	// 一个公司只能有一个发票抬头
	invoice, _ := s.GetFdInvoiceByTaxId(ctx, info.TaxId)
	if invoice != nil {
		if info.UnionMainId == invoice.UnionMainId {
			return nil, gerror.New("一个主体只能注册一个发票抬头")
		}
	}

	// 判断审核状态
	if info.State == co_enum.Invoice.AuditType.Reject.Code() && info.AuditReplayMsg == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", s.dao.FdInvoice.Table())
	}

	// 创建发票
	data := co_entity.FdInvoice{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()
	data.AuditUserId = 0

	data.State = co_enum.Invoice.AuditType.WaitReview.Code()

	_, err := s.dao.FdInvoice.Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(data).Insert()
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "创建发票记录失败", s.dao.FdInvoice.Table())
	}

	return s.GetInvoiceById(ctx, data.Id)
}

// GetInvoiceById 根据id获取发票
func (s *sFdInvoice) GetInvoiceById(ctx context.Context, id int64) (*co_entity.FdInvoice, error) {
	if id == 0 {
		return nil, gerror.New("id不能为空")
	}

	result, err := daoctl.GetByIdWithError[co_entity.FdInvoice](s.dao.FdInvoice.Ctx(ctx).Hook(daoctl.CacheHookHandler), id)

	if result == nil || err != nil {
		return nil, sys_service.SysLogs().InfoSimple(ctx, err, "当前没有发票抬头记录", s.dao.FdInvoice.Table())
	}

	return result, nil
}

// QueryInvoiceList 获取发票抬头列表
func (s *sFdInvoice) QueryInvoiceList(ctx context.Context, info *sys_model.SearchParams, userId int64) (*co_model.FdInvoiceListRes, error) {
	newFields := make([]sys_model.FilterInfo, 0)
	// 筛选条件强制指定所属用户
	newFields = append(newFields, sys_model.FilterInfo{
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

	result, err := daoctl.Query[co_entity.FdInvoice](s.dao.FdInvoice.Ctx(ctx).Hook(daoctl.CacheHookHandler), info, false)

	return (*co_model.FdInvoiceListRes)(result), err
}

// DeletesFdInvoiceById 删除发票抬头
func (s *sFdInvoice) DeletesFdInvoiceById(ctx context.Context, invoiceId int64) (bool, error) {
	var result sql.Result
	var err error

	invoice, err := s.GetInvoiceById(ctx, invoiceId)
	if err != nil || invoice == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "发票抬头不存在", s.dao.FdInvoice.Table())
	}

	err = s.dao.FdInvoice.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除
		result, err = s.dao.FdInvoice.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdInvoice{Id: invoice.Id}).Delete()

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil || result == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "发票抬头删除失败", s.dao.FdInvoice.Table())
	}

	return true, nil
}

// GetFdInvoiceByTaxId 根据纳税识别号获取发票抬头信息
func (s *sFdInvoice) GetFdInvoiceByTaxId(ctx context.Context, taxId string) (*co_entity.FdInvoice, error) {
	if taxId == "" {
		return nil, gerror.New("纳税识别号不能为空")
	}
	result := co_entity.FdInvoice{}

	err := s.dao.FdInvoice.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.FdInvoice{TaxId: taxId}).Scan(&result)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "发票抬头信息获取失败", s.dao.FdInvoice.Table())
	}

	return &result, nil
}
