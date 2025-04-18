package finance

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

// 货币类型管理
type sFdCurrency struct {
}

func NewFdCurrency() co_service.IFdCurrency {
	return &sFdCurrency{}
}

// GetCurrencyByCode 根据货币代码查找货币(主键)
func (s *sFdCurrency) GetCurrencyByCode(ctx context.Context, currencyCode string) (response *co_model.FdCurrencyRes, err error) {
	if currencyCode == "" {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_CurrencyCode_NotNull", co_dao.FdCurrency.Table())
	}

	err = co_dao.FdCurrency.Ctx(ctx).Where(co_do.FdCurrency{CurrencyCode: currencyCode}).Scan(response)

	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, "{#Currency}{#error_Data_Get_Failed}", co_dao.FdCurrency.Table())
	}

	return response, nil
}

// QueryCurrencyList 获取币种列表
func (s *sFdCurrency) QueryCurrencyList(ctx context.Context, search *base_model.SearchParams) (*co_model.FdCurrencyListRes, error) {
	result, err := daoctl.Query[co_model.FdCurrencyRes](co_dao.FdCurrency.Ctx(ctx), search, true)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return kconv.Struct(result, &co_model.FdCurrencyListRes{}), nil
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_Data_Get_Failed}", co_dao.FdCurrency.Table())
	}

	return kconv.Struct(result, &co_model.FdCurrencyListRes{}), nil
}

// GetCurrencyByCnName 根据国家查找货币信息
func (s *sFdCurrency) GetCurrencyByCnName(ctx context.Context, cnName string) (response *co_model.FdCurrencyRes, err error) {
	if cnName == "" {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_CurrencyCnName_NotNull", co_dao.FdCurrency.Table())
	}

	err = co_dao.FdCurrency.Ctx(ctx).Where(co_do.FdCurrency{CnName: cnName}).Scan(response)
	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, "{#Currency}{#error_Data_Get_Failed}", co_dao.FdCurrency.Table())
	}

	return response, nil
}
