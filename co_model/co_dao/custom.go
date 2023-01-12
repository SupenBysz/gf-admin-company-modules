package co_dao

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/utility/kmap"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type IDao[T any] interface {
	DB() gdb.DB
	Table() string
	Columns() T
	Group() string
	Ctx(ctx context.Context) *gdb.Model
	Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error)
}

type CustomDao[T any] struct {
	IDao[T]
	conf *co_model.Config
}

var (
	daoMap = kmap.New[string, interface{}]()
)

func NewDao[T any, TC IDao[T]](conf *co_model.Config, dao TC) IDao[T] {
	value, found := daoMap.Search(conf.KeyIndex)
	if value != nil && found {
		v, ok := value.(IDao[T])
		if ok {
			return v
		} else {
			panic("数据访问对象类型校验失败")
		}
	}

	// if dao.Table() != conf.TableName.Company &&
	// 	gstr.HasSuffix(dao.Table(), "employee") &&
	// 	gstr.HasSuffix(dao.Table(), "team") &&
	// 	gstr.HasSuffix(dao.Table(), "team_member") {
	// 	panic("数据访问对象表名校验失败")
	// }

	result := &CustomDao[T]{
		IDao: dao,
		conf: conf,
	}

	daoMap.Set(dao.Table(), result)

	return result
}

// DB 检索并返回当前DAO的底层原始数据库管理对象。
func (d *CustomDao[T]) DB() gdb.DB {
	return g.DB(d.Group())
}

// Table 返回当前DAO的表名。
func (d *CustomDao[T]) Table() string {
	return d.IDao.Table()
}

// Columns 返回当前DAO的所有列名。
func (d *CustomDao[T]) Columns() T {
	return d.IDao.Columns()
}

// Group 返回当前DAO数据库的配置组名。
func (d *CustomDao[T]) Group() string {
	return d.IDao.Group()
}

// Ctx 创建并返回当前DAO的模型，它自动设置当前操作的上下文。
func (d *CustomDao[T]) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.Table()).Safe().Ctx(ctx)
}

func (d *CustomDao[T]) Dao() IDao[T] {
	return d.IDao
}

// Transaction 使用函数f包装事务逻辑。
// 回滚事务，如果函数f返回非nil错误，则返回该错误。
// 如果函数f返回nil，则提交事务并返回nil。
//
// 注意，你不应该在函数f中提交或回滚事务, 因为它由这个函数自动处理。
func (d *CustomDao[T]) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return d.Ctx(ctx).Transaction(ctx, f)
}
