package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
)

// InitPermission 初始化权限树
func InitPermission(module co_interface.IModules) []*permission.SysPermissionTree {
	result := []*permission.SysPermissionTree{
		// 公司
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5947986066667973,
				Name:       module.T(context.TODO(), "{#CompanyName}"),
				Identifier: module.GetConfig().Identifier.Company,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Company.PermissionType(module).Create,
				co_enum.Company.PermissionType(module).ViewDetail,
				co_enum.Company.PermissionType(module).List,
				co_enum.Company.PermissionType(module).Update,
				co_enum.Company.PermissionType(module).SetLogo,
				co_enum.Company.PermissionType(module).SetState,
				co_enum.Company.PermissionType(module).SetAdminUser,
				co_enum.Company.PermissionType(module).ViewLicense,
				co_enum.Company.PermissionType(module).AuditLicense,
			},
		},
		// 员工
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408325,
				Name:       module.T(context.TODO(), "{#CompanyName}{#EmployeeName}"),
				Identifier: module.GetConfig().Identifier.Employee,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Employee.PermissionType(module).ViewDetail,
				co_enum.Employee.PermissionType(module).MoreDetail,
				co_enum.Employee.PermissionType(module).List,
				co_enum.Employee.PermissionType(module).Create,
				co_enum.Employee.PermissionType(module).Update,
				co_enum.Employee.PermissionType(module).Delete,
				co_enum.Employee.PermissionType(module).SetMobile,
				co_enum.Employee.PermissionType(module).SetAvatar,
				co_enum.Employee.PermissionType(module).SetState,
				co_enum.Employee.PermissionType(module).ViewLicense,
				co_enum.Employee.PermissionType(module).AuditLicense,
				co_enum.Employee.PermissionType(module).UpdateLicense,
			},
		},
		// 团队
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948221667408326,
				Name:       module.T(context.TODO(), "{#CompanyName}{#TeamName}"),
				Identifier: module.GetConfig().Identifier.Team,
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				co_enum.Team.PermissionType(module).Create,
				co_enum.Team.PermissionType(module).ViewDetail,
				co_enum.Team.PermissionType(module).List,
				co_enum.Team.PermissionType(module).Update,
				co_enum.Team.PermissionType(module).Delete,
				co_enum.Team.PermissionType(module).MemberDetail,
				co_enum.Team.PermissionType(module).SetMember,
				co_enum.Team.PermissionType(module).SetOwner,
				co_enum.Team.PermissionType(module).SetCaptain,
			},
		},
		// 财务服务权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5953153121845349,
				Name:       "财务服务",
				Identifier: "Financial",
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看发票详情，查看发票详情信息
				co_enum.Financial.PermissionType.ViewInvoiceDetail,
				// 查看发票抬头信息，查看发票抬头信息
				co_enum.Financial.PermissionType.ViewInvoice,
				// 查看提现账号，查看提现账号信息
				co_enum.Financial.PermissionType.ViewBankCardDetail,
				// 提现账号列表，查看所有提现账号
				co_enum.Financial.PermissionType.BankCardList,
				// 发票抬头列表，查看所有发票抬头
				co_enum.Financial.PermissionType.InvoiceList,
				// 发票详情列表，查看所有发票详情
				co_enum.Financial.PermissionType.InvoiceDetailList,
				// 审核发票，审核发票申请
				co_enum.Financial.PermissionType.AuditInvoiceDetail,
				// 开发票，添加发票详情记录
				co_enum.Financial.PermissionType.MakeInvoiceDetail,
				// 添加发票抬头，添加发票抬头信息
				co_enum.Financial.PermissionType.CreateInvoice,
				//申请提现账号，添加提现账号信息
				co_enum.Financial.PermissionType.CreateBankCard,
				// 删除发票抬头，删除发票抬头信息
				co_enum.Financial.PermissionType.DeleteInvoice,
				//  删除提现账号，删除提现账号信息
				co_enum.Financial.PermissionType.DeleteBankCard,
				// 查看余额，查看账号余额
				co_enum.Financial.PermissionType.GetAccountBalance,
			},
		},
	}
	return result
}
