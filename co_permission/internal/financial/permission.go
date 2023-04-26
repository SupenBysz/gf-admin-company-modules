package financial

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/kmap"
)

type Permission = *sys_model.SysPermissionTree

type permissionType[T co_interface.IConfig] struct {
	modules T
	enumMap *kmap.HashMap[string, Permission]

	ViewInvoiceDetail   Permission
	ViewInvoice         Permission
	ViewBankCardDetail  Permission
	BankCardList        Permission
	InvoiceList         Permission
	InvoiceDetailList   Permission
	AuditInvoiceDetail  Permission
	MakeInvoiceDetail   Permission
	CreateInvoice       Permission
	CreateBankCard      Permission
	DeleteInvoice       Permission
	DeleteBankCard      Permission
	GetAccountBalance   Permission
	GetAccountDetail    Permission
	UpdateAccountAmount Permission
	UpdateAccountState  Permission
	UpdateAccountDetail Permission
}

var (
	permissionTypeMap = kmap.New[string, *permissionType[co_interface.IConfig]]()
	PermissionType    = func(modules co_interface.IConfig) *permissionType[co_interface.IConfig] {
		result := permissionTypeMap.GetOrSet(modules.GetConfig().KeyIndex, &permissionType[co_interface.IConfig]{
			enumMap:            kmap.New[string, Permission](),
			ViewInvoiceDetail:  permission.New(5953153121845334, "ViewDetail", "查看发票详情", "查看发票详情信息"),
			ViewInvoice:        permission.New(5953153121845335, "ViewInvoice", "查看发票抬头信息", "查看发票抬头信息"),
			ViewBankCardDetail: permission.New(5953153121845336, "ViewBankCardDetail", "查看银行卡", "查看银行卡信息"),

			BankCardList:      permission.New(5953153121845337, "BankCardList", "银行卡列表", "查看所有银行卡"),
			InvoiceList:       permission.New(5953153121845338, "InvoiceList", "发票抬头列表", "查看所有发票抬头"),
			InvoiceDetailList: permission.New(5953153121845339, "InvoiceDetailList", "发票详情列表", "查看所有发票详情"),

			AuditInvoiceDetail: permission.New(5953153121845340, "AuditInvoiceDetail", "审核发票", "审核发票申请"),
			MakeInvoiceDetail:  permission.New(5953153121845341, "MakeInvoiceDetail", "开发票", "添加发票详情记录"),

			CreateInvoice:  permission.New(5953153121845342, "CreateInvoice", "添加发票抬头", "添加发票抬头信息"),
			CreateBankCard: permission.New(5953153121845343, "CreateBankCard", "添加银行卡", "添加银行卡信息"),

			DeleteInvoice:  permission.New(5953153121845344, "DeleteInvoice", "删除发票抬头", "删除发票抬头信息"),
			DeleteBankCard: permission.New(5953153121845345, "DeleteBankCard", "删除银行卡", "删除银行卡信息"),

			GetAccountBalance: permission.New(5953153121845346, "GetAccountBalance", "查看余额", "查看账号余额"),

			GetAccountDetail:    permission.New(5953153121849347, "GetAccountDetail", "查看财务账号详情", "查看财务账号金额明细"),
			UpdateAccountAmount: permission.New(5953153121849348, "UpdateAccountAmount", "修改财务金额", "修改财务账号金额明细"),
			UpdateAccountState:  permission.New(5953153121849349, "UpdateAccountState", "修改财务账号状态", "修改财务账号状态"),
			UpdateAccountDetail: permission.New(5953153121898321, "UpdateAccountDetail", "修改财务账号", "修改财务账号详情"),
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
