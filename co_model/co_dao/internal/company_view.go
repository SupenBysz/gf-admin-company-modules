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

// CompanyViewDao is the data access object for table co_company_view.
type CompanyViewDao struct {
	dao_interface.IDao
	table       string             // table is the underlying table name of the DAO.
	group       string             // group is the database configuration group name of current DAO.
	columns     CompanyViewColumns // columns contains all the column names of Table for convenient usage.
	daoConfig   *dao_interface.DaoConfig
	ignoreCache bool
	exWhereArr  []string
}

// CompanyViewColumns defines and stores column names for table co_company_view.
type CompanyViewColumns struct {
	Id             string //
	Name           string //
	ContactName    string //
	ContactMobile  string //
	UserId         string //
	State          string //
	Remark         string //
	CreatedBy      string //
	CreatedAt      string //
	UpdatedBy      string //
	UpdatedAt      string //
	DeletedBy      string //
	DeletedAt      string //
	ParentId       string //
	Address        string //
	LicenseId      string //
	LicenseState   string //
	LogoId         string //
	StartLevel     string //
	CountryCode    string //
	Region         string //
	Score          string //
	CommissionRate string //
	CompanyType    string //
}

// companyViewColumns holds the columns for table co_company_view.
var companyViewColumns = CompanyViewColumns{
	Id:             "id",
	Name:           "name",
	ContactName:    "contact_name",
	ContactMobile:  "contact_mobile",
	UserId:         "user_id",
	State:          "state",
	Remark:         "remark",
	CreatedBy:      "created_by",
	CreatedAt:      "created_at",
	UpdatedBy:      "updated_by",
	UpdatedAt:      "updated_at",
	DeletedBy:      "deleted_by",
	DeletedAt:      "deleted_at",
	ParentId:       "parent_id",
	Address:        "address",
	LicenseId:      "license_id",
	LicenseState:   "license_state",
	LogoId:         "logo_id",
	StartLevel:     "start_level",
	CountryCode:    "country_code",
	Region:         "region",
	Score:          "score",
	CommissionRate: "commission_rate",
	CompanyType:    "company_type",
}

// NewCompanyViewDao creates and returns a new DAO object for table data access.
func NewCompanyViewDao(proxy ...dao_interface.IDao) *CompanyViewDao {
	var dao *CompanyViewDao
	if len(proxy) > 0 {
		dao = &CompanyViewDao{
			group:       proxy[0].Group(),
			table:       proxy[0].Table(),
			columns:     companyViewColumns,
			daoConfig:   proxy[0].DaoConfig(context.Background()),
			IDao:        proxy[0].DaoConfig(context.Background()).Dao,
			ignoreCache: proxy[0].DaoConfig(context.Background()).IsIgnoreCache(),
			exWhereArr:  proxy[0].DaoConfig(context.Background()).Dao.GetExtWhereKeys(),
		}

		return dao
	}

	return &CompanyViewDao{
		group:   "default",
		table:   "co_company_view",
		columns: companyViewColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanyViewDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanyViewDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanyViewDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *CompanyViewDao) Columns() CompanyViewColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanyViewDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *CompanyViewDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) *dao_interface.DaoConfig {
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
func (dao *CompanyViewDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

func (dao *CompanyViewDao) GetExtWhereKeys() []string {
	return dao.exWhereArr
}

func (dao *CompanyViewDao) IsIgnoreCache() bool {
	return dao.ignoreCache
}

func (dao *CompanyViewDao) IgnoreCache() dao_interface.IDao {
	dao.ignoreCache = true
	return dao
}
func (dao *CompanyViewDao) IgnoreExtModel(whereKey ...string) dao_interface.IDao {
	dao.exWhereArr = append(dao.exWhereArr, whereKey...)
	return dao
}
