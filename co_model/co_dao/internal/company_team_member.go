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

// CompanyTeamMemberDao is the data access object for table co_company_team_member.
type CompanyTeamMemberDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns CompanyTeamMemberColumns // columns contains all the column names of Table for convenient usage.
}

// CompanyTeamMemberColumns defines and stores column names for table co_company_team_member.
type CompanyTeamMemberColumns struct {
	Id           string // ID
	TeamId       string // 团队ID
	EmployeeId   string // 成员ID
	InviteUserId string // 邀约人ID
	UnionMainId  string // 关联主体ID
	JoinAt       string // 加入时间
}

// companyTeamMemberColumns holds the columns for table co_company_team_member.
var companyTeamMemberColumns = CompanyTeamMemberColumns{
	Id:           "id",
	TeamId:       "team_id",
	EmployeeId:   "employee_id",
	InviteUserId: "invite_user_id",
	UnionMainId:  "union_main_id",
	JoinAt:       "join_at",
}

// NewCompanyTeamMemberDao creates and returns a new DAO object for table data access.
func NewCompanyTeamMemberDao(proxy ...dao_interface.IDao) *CompanyTeamMemberDao {
	var dao *CompanyTeamMemberDao
	if len(proxy) > 0 {
		dao = &CompanyTeamMemberDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: companyTeamMemberColumns,
		}
		return dao
	}

	return &CompanyTeamMemberDao{
		group:   "default",
		table:   "co_company_team_member",
		columns: companyTeamMemberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanyTeamMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanyTeamMemberDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanyTeamMemberDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *CompanyTeamMemberDao) Columns() CompanyTeamMemberColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanyTeamMemberDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *CompanyTeamMemberDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CompanyTeamMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
