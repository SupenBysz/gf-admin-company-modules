package financial

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/kmap"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/gogf/gf/v2/util/gconv"
)

type Permission = *permission.SysPermissionTree

type permissionType[T co_interface.IModules] struct {
	modules T
	enumMap *kmap.HashMap[string, Permission]

	ViewInvoiceDetail  Permission
	ViewInvoice        Permission
	ViewBankCardDetail Permission
	BankCardList       Permission
	InvoiceList        Permission
	InvoiceDetailList  Permission
	AuditInvoiceDetail Permission
	MakeInvoiceDetail  Permission
	CreateInvoice      Permission
	CreateBankCard     Permission
	DeleteInvoice      Permission
	DeleteBankCard     Permission
	GetAccountBalance  Permission
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IModules]]()
	PermissionType    = func(modules co_interface.IModules) *permissionType[co_interface.IModules] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IModules]{
			enumMap:            kmap.New[string, Permission](),
			ViewInvoiceDetail:  permission.New(5953153121845334, "ViewDetail", "查看发票详情", "查看发票详情信息"),
			ViewInvoice:        permission.New(5953153121845335, "ViewInvoice", "查看发票抬头信息", "查看发票抬头信息"),
			ViewBankCardDetail: permission.New(5953153121845336, "ViewBankCardDetail", "查看提现账号", "查看提现账号信息"),

			BankCardList:      permission.New(5953153121845337, "BankCardList", "提现账号列表", "查看所有提现账号"),
			InvoiceList:       permission.New(5953153121845338, "InvoiceList", "发票抬头列表", "查看所有发票抬头"),
			InvoiceDetailList: permission.New(5953153121845339, "InvoiceDetailList", "发票详情列表", "查看所有发票详情"),

			AuditInvoiceDetail: permission.New(5953153121845340, "AuditInvoiceDetail", "审核发票", "审核发票申请"),
			MakeInvoiceDetail:  permission.New(5953153121845341, "MakeInvoiceDetail", "开发票", "添加发票详情记录"),

			CreateInvoice:  permission.New(5953153121845342, "CreateInvoice", "添加发票抬头", "添加发票抬头信息"),
			CreateBankCard: permission.New(5953153121845343, "CreateBankCard", "申请提现账号", "添加提现账号信息"),

			DeleteInvoice:  permission.New(5953153121845344, "DeleteInvoice", "删除发票抬头", "删除发票抬头信息"),
			DeleteBankCard: permission.New(5953153121845345, "DeleteBankCard", "删除提现账号", "删除提现账号信息"),

			GetAccountBalance: permission.New(5953153121845346, "GetAccountBalance", "查看余额", "查看账号余额"),
		})
		for k, v := range gconv.Map(result) {
			result.enumMap.Set(k, v.(Permission))
		}
		return result
	}
)

// ByCode 通过枚举值取枚举类型
func (e *permissionType[T]) ByCode(identifier string) *sys_entity.SysPermission {
	v, has := e.enumMap.Search(identifier)
	if v != nil && has {
		return v.SysPermission
	}
	return nil
}
