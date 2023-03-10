package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

// 货币类型管理
type sFdCurrency struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewFdCurrency(modules co_interface.IModules) co_interface.IFdCurrency {
	return &sFdCurrency{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetCurrencyByCurrencyCode 根据货币代码查找货币(主键)
func (s *sFdCurrency) GetCurrencyByCurrencyCode(ctx context.Context, currencyCode string) (*co_entity.FdCurrency, error) {
	if currencyCode == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_CurrencyCode_NotNull"), s.dao.FdCurrency.Table())
	}

	result := &co_entity.FdCurrency{}

	err := s.dao.FdCurrency.Ctx(ctx).Where(co_do.FdCurrency{CurrencyCode: currencyCode}).Scan(result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Currency}{#error_Data_Get_Failed}"), s.dao.FdCurrency.Table())
	}

	return result, nil
}

// GetCurrencyByCnName 根据国家查找货币信息
func (s *sFdCurrency) GetCurrencyByCnName(ctx context.Context, cnName string) (*co_entity.FdCurrency, error) {
	if cnName == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_CurrencyCnName_NotNull"), s.dao.FdCurrency.Table())
	}

	result := &co_entity.FdCurrency{}

	err := s.dao.FdCurrency.Ctx(ctx).Where(co_do.FdCurrency{CnName: cnName}).Scan(result)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Currency}{#error_Data_Get_Failed}"), s.dao.FdCurrency.Table())
	}

	return result, nil
}
