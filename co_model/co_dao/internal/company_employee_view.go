// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CompanyEmployeeViewDao is the data access object for table co_company_employee_view.
type CompanyEmployeeViewDao struct {
	dao_interface.IDao
	table       string                     // table is the underlying table name of the DAO.
	group       string                     // group is the database configuration group name of current DAO.
	columns     CompanyEmployeeViewColumns // columns contains all the column names of Table for convenient usage.
	daoConfig   *dao_interface.DaoConfig
	ignoreCache bool
	exWhereArr  []string
}

// CompanyEmployeeViewColumns defines and stores column names for table co_company_employee_view.
type CompanyEmployeeViewColumns struct {
	Id             string //
	No             string //
	Avatar         string //
	Name           string //
	Mobile         string //
	UnionMainId    string //
	State          string //
	LastActiveIp   string //
	HiredAt        string //
	CreatedBy      string //
	CreatedAt      string //
	UpdatedBy      string //
	UpdatedAt      string //
	DeletedBy      string //
	DeletedAt      string //
	Sex            string //
	Remark         string //
	CountryCode    string //
	Region         string //
	Email          string //
	WeixinAccount  string //
	Address        string //
	WorkCardAvatar string //
	CommissionRate string //
	CompanyType    string //
}

// companyEmployeeViewColumns holds the columns for table co_company_employee_view.
var companyEmployeeViewColumns = CompanyEmployeeViewColumns{
	Id:             "id",
	No:             "no",
	Avatar:         "avatar",
	Name:           "name",
	Mobile:         "mobile",
	UnionMainId:    "union_main_id",
	State:          "state",
	LastActiveIp:   "last_active_ip",
	HiredAt:        "hired_at",
	CreatedBy:      "created_by",
	CreatedAt:      "created_at",
	UpdatedBy:      "updated_by",
	UpdatedAt:      "updated_at",
	DeletedBy:      "deleted_by",
	DeletedAt:      "deleted_at",
	Sex:            "sex",
	Remark:         "remark",
	CountryCode:    "country_code",
	Region:         "region",
	Email:          "email",
	WeixinAccount:  "weixin_account",
	Address:        "address",
	WorkCardAvatar: "work_card_avatar",
	CommissionRate: "commission_rate",
	CompanyType:    "company_type",
}

// NewCompanyEmployeeViewDao creates and returns a new DAO object for table data access.
func NewCompanyEmployeeViewDao(proxy ...dao_interface.IDao) *CompanyEmployeeViewDao {
	var dao *CompanyEmployeeViewDao
	if len(proxy) > 0 {
		dao = &CompanyEmployeeViewDao{
			group:       proxy[0].Group(),
			table:       proxy[0].Table(),
			columns:     companyEmployeeViewColumns,
			daoConfig:   proxy[0].DaoConfig(context.Background()),
			IDao:        proxy[0].DaoConfig(context.Background()).Dao,
			ignoreCache: proxy[0].DaoConfig(context.Background()).IsIgnoreCache(),
			exWhereArr:  proxy[0].DaoConfig(context.Background()).Dao.GetExtWhereKeys(),
		}

		return dao
	}

	return &CompanyEmployeeViewDao{
		group:   "default",
		table:   "co_company_employee_view",
		columns: companyEmployeeViewColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanyEmployeeViewDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanyEmployeeViewDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanyEmployeeViewDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *CompanyEmployeeViewDao) Columns() CompanyEmployeeViewColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanyEmployeeViewDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *CompanyEmployeeViewDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) *dao_interface.DaoConfig {
	//if dao.daoConfig != nil && len(dao.exWhereArr) == 0 {
	//	return dao.daoConfig
	//}

	var daoConfig = daoctl.NewDaoConfig(ctx, dao, cacheOption...)
	dao.daoConfig = &daoConfig

	if len(dao.exWhereArr) > 0 {
		daoConfig.IgnoreExtModel(dao.exWhereArr...)
		dao.exWhereArr = []string{}

	}

	if dao.ignoreCache {
		daoConfig.IgnoreCache()
	}

	return dao.daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CompanyEmployeeViewDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

func (dao *CompanyEmployeeViewDao) GetExtWhereKeys() []string {
	return dao.exWhereArr
}

func (dao *CompanyEmployeeViewDao) IsIgnoreCache() bool {
	return dao.ignoreCache
}

func (dao *CompanyEmployeeViewDao) IgnoreCache() dao_interface.IDao {
	dao.ignoreCache = true
	return dao
}
func (dao *CompanyEmployeeViewDao) IgnoreExtModel(whereKey ...string) dao_interface.IDao {
	dao.exWhereArr = append(dao.exWhereArr, whereKey...)
	return dao
}
