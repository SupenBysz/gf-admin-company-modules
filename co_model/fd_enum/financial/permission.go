package financial

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/util/gconv"
)

type PermissionTypeEnum = *permission.SysPermissionTree

type permissionType struct {
	ViewInvoiceDetail  PermissionTypeEnum
	ViewInvoice        PermissionTypeEnum
	ViewBankCardDetail PermissionTypeEnum
	BankCardList       PermissionTypeEnum
	InvoiceList        PermissionTypeEnum
	InvoiceDetailList  PermissionTypeEnum
	AuditInvoiceDetail PermissionTypeEnum
	MakeInvoiceDetail  PermissionTypeEnum
	CreateInvoice      PermissionTypeEnum
	CreateBankCard     PermissionTypeEnum
	DeleteInvoice      PermissionTypeEnum
	DeleteBankCard     PermissionTypeEnum
	GetAccountBalance  PermissionTypeEnum
}

var (
	PermissionType = permissionType{
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
	}
	permissionTypeMap = gmap.NewStrAnyMapFrom(gconv.Map(PermissionType))
)

// ByForm 通过枚举值取枚举类型
func (e *permissionType) ByForm(code int64) *sys_model.SysPermission {
	return sys_service.SysPermission().PermissionTypeForm(code, permissionTypeMap)
}

//
// type PermissionEnum = *permission.SysPermissionTree
//
// type permissionType[T co_interface.IModules] struct {
//	modules T
//	enumMap *kmap.HashMap[string, PermissionEnum]
//
//	ViewInvoiceDetail  PermissionEnum
//	ViewInvoice        PermissionEnum
//	ViewBankCardDetail PermissionEnum
//	BankCardList       PermissionEnum
//	InvoiceList        PermissionEnum
//	InvoiceDetailList  PermissionEnum
//	AuditInvoiceDetail PermissionEnum
//	MakeInvoiceDetail  PermissionEnum
//	CreateInvoice      PermissionEnum
//	CreateBankCard     PermissionEnum
//	DeleteInvoice      PermissionEnum
//	DeleteBankCard     PermissionEnum
//	GetAccountBalance  PermissionEnum
// }
//
// var (
//	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IModules]]()
//	PermissionType    = func(modules co_interface.IModules) *permissionType[co_interface.IModules] {
//		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IModules]{
//			ViewInvoiceDetail:  permission.New(5953153121845334, "ViewDetail", "查看发票详情", "查看发票详情信息"),
//			ViewInvoice:        permission.New(5953153121845335, "ViewInvoice", "查看发票抬头信息", "查看发票抬头信息"),
//			ViewBankCardDetail: permission.New(5953153121845336, "ViewBankCardDetail", "查看提现账号", "查看提现账号信息"),
//
//			BankCardList:      permission.New(5953153121845337, "BankCardList", "提现账号列表", "查看所有提现账号"),
//			InvoiceList:       permission.New(5953153121845338, "InvoiceList", "发票抬头列表", "查看所有发票抬头"),
//			InvoiceDetailList: permission.New(5953153121845339, "InvoiceDetailList", "发票详情列表", "查看所有发票详情"),
//
//			AuditInvoiceDetail: permission.New(5953153121845340, "AuditInvoiceDetail", "审核发票", "审核发票申请"),
//			MakeInvoiceDetail:  permission.New(5953153121845341, "MakeInvoiceDetail", "开发票", "添加发票详情记录"),
//
//			CreateInvoice:  permission.New(5953153121845342, "CreateInvoice", "添加发票抬头", "添加发票抬头信息"),
//			CreateBankCard: permission.New(5953153121845343, "CreateBankCard", "申请提现账号", "添加提现账号信息"),
//
//			DeleteInvoice:  permission.New(5953153121845344, "DeleteInvoice", "删除发票抬头", "删除发票抬头信息"),
//			DeleteBankCard: permission.New(5953153121845345, "DeleteBankCard", "删除提现账号", "删除提现账号信息"),
//
//			GetAccountBalance: permission.New(5953153121845346, "GetAccountBalance", "查看余额", "查看账号余额"),
//		})
//		for k, v := range gconv.Map(result) {
//			result.enumMap.Set(k, v.(PermissionEnum))
//		}
//		return result
//	}
// )
//
// // ByCode 通过枚举值取枚举类型
// func (e *permissionType[T]) ByCode(identifier string) *sys_entity.SysPermission {
//	v, has := e.enumMap.Search(identifier)
//	if v != nil && has {
//		return v.SysPermission
//	}
//	return nil
// }
