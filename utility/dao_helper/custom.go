package dao_helper

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/utility/kmap"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type CustomDao[T any] struct {
	co_interface.IDao[T]
	conf *co_model.Config
}

var (
	daoMap = kmap.New[string, interface{}]()
)

func NewDao[T any, TC co_interface.IDao[T]](conf *co_model.Config, dao TC) TC {
	var value interface{} = dao
	v, ok := value.(*CustomDao[T])
	if ok {
		v.conf = conf
	}
	return dao
}

func NewDaoMap[T any, TC co_interface.IDao[T]](conf *co_model.Config, dao TC) co_interface.IDao[T] {
	value, found := daoMap.Search(conf.KeyIndex)
	if value != nil && found {
		v, ok := value.(co_interface.IDao[T])
		if ok {
			return v
		} else {
			panic("数据访问对象类型校验失败")
		}
	}

	result := &CustomDao[T]{
		IDao: dao,
		conf: conf,
	}

	daoMap.Set(dao.Table(), result)

	return result
}

// DB 检索并返回当前DAO的底层原始数据库管理对象。
func (d *CustomDao[T]) DB() gdb.DB {
	if d.conf.DB != nil {
		return d.conf.DB
	}
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

func (d *CustomDao[T]) Dao() co_interface.IDao[T] {
	return d.IDao
}

// Transaction 使用函数f包装事务逻辑。
// 回滚事务，如果函数f返回非nil错误，则返回该错误。
// 如果函数f返回nil，则提交事务并返回nil。
//
// 注意，你不应该在函数f中提交或回滚事务, 因为它由这个函数自动处理。
func (d *CustomDao[T]) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return d.Ctx(ctx).Transaction(ctx, f)
}
