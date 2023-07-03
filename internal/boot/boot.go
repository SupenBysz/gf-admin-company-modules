package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/SupenBysz/gf-admin-company-modules/utility/co_rules"
	"github.com/yitter/idgenerator-go/idgen"
)

// InitCustomRules 注册自定义参数校验规则
func InitCustomRules() {
	// 注册资质自定义规则
	co_rules.RequiredLicense()
}

// InitPermission 初始化权限树
func InitPermission[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](module co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) []*sys_model.SysPermissionTree {
	result := []*sys_model.SysPermissionTree{
		// 公司
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         idgen.NextId(), // 导入权限的时候判断的是标识符号，所以不用担心下一次启动导入id不同的相同权限
				Name:       module.T(context.TODO(), "{#CompanyName}"),
				Identifier: module.GetConfig().Identifier.Company,
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				co_permission.Company.PermissionType(module).Create,
				co_permission.Company.PermissionType(module).ViewDetail,
				co_permission.Company.PermissionType(module).List,
				co_permission.Company.PermissionType(module).Update,
				co_permission.Company.PermissionType(module).SetLogo,
				co_permission.Company.PermissionType(module).SetState,
				co_permission.Company.PermissionType(module).SetAdminUser,
				co_permission.Company.PermissionType(module).ViewLicense,
				co_permission.Company.PermissionType(module).AuditLicense,
			},
		},
		// 员工
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         idgen.NextId(),
				Name:       module.T(context.TODO(), "{#CompanyName}{#EmployeeName}"),
				Identifier: module.GetConfig().Identifier.Employee,
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				co_permission.Employee.PermissionType(module).ViewDetail,
				co_permission.Employee.PermissionType(module).MoreDetail,
				co_permission.Employee.PermissionType(module).List,
				co_permission.Employee.PermissionType(module).Create,
				co_permission.Employee.PermissionType(module).Update,
				co_permission.Employee.PermissionType(module).Delete,
				co_permission.Employee.PermissionType(module).SetMobile,
				co_permission.Employee.PermissionType(module).SetAvatar,
				co_permission.Employee.PermissionType(module).SetState,
				co_permission.Employee.PermissionType(module).ViewLicense,
				co_permission.Employee.PermissionType(module).AuditLicense,
				co_permission.Employee.PermissionType(module).UpdateLicense,
			},
		},
		// 团队
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         idgen.NextId(),
				Name:       module.T(context.TODO(), "{#CompanyName}{#TeamName}"),
				Identifier: module.GetConfig().Identifier.Team,
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				co_permission.Team.PermissionType(module).Create,
				co_permission.Team.PermissionType(module).ViewDetail,
				co_permission.Team.PermissionType(module).List,
				co_permission.Team.PermissionType(module).Update,
				co_permission.Team.PermissionType(module).Delete,
				co_permission.Team.PermissionType(module).MemberDetail,
				co_permission.Team.PermissionType(module).SetMember,
				co_permission.Team.PermissionType(module).SetOwner,
				co_permission.Team.PermissionType(module).SetCaptain,
			},
		},
		// sms短信

		// oss
	}

	// 添加资质和审核权限树
	licensePermission := initAuditAndLicensePermission()
	result = append(result, licensePermission...)

	return result
}

// InitFinancialPermission 初始化财务服务权限树
func InitFinancialPermission[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](module co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) []*sys_model.SysPermissionTree {
	result := []*sys_model.SysPermissionTree{
		// 财务服务权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         idgen.NextId(),
				Name:       "财务",
				Identifier: "Financial",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				{
					SysPermission: &sys_entity.SysPermission{
						Id:         idgen.NextId(),
						Name:       "发票",
						Identifier: "Invoice",
						Type:       1,
						IsShow:     1,
					},
					Children: []*sys_model.SysPermissionTree{
						// 查看发票详情，查看发票详情信息
						co_permission.Financial.PermissionType(module).ViewInvoiceDetail,
						// 查看发票抬头信息，查看发票抬头信息
						co_permission.Financial.PermissionType(module).ViewInvoice,
						// 发票抬头列表，查看所有发票抬头
						co_permission.Financial.PermissionType(module).InvoiceList,
						// 发票详情列表，查看所有发票详情
						co_permission.Financial.PermissionType(module).InvoiceDetailList,
						// 审核发票，审核发票申请
						co_permission.Financial.PermissionType(module).AuditInvoiceDetail,
						// 开发票，添加发票详情记录
						co_permission.Financial.PermissionType(module).MakeInvoiceDetail,
						// 添加发票抬头，添加发票抬头信息
						co_permission.Financial.PermissionType(module).CreateInvoice,
						// 删除发票抬头，删除发票抬头信息
						co_permission.Financial.PermissionType(module).DeleteInvoice,
					},
				},
				{
					SysPermission: &sys_entity.SysPermission{
						Id:         idgen.NextId(),
						Name:       "银行卡",
						Identifier: "BankCard",
						Type:       1,
						IsShow:     1,
					},
					Children: []*sys_model.SysPermissionTree{
						// 查看提现账号，查看银行卡账号信息
						co_permission.Financial.PermissionType(module).ViewBankCardDetail,
						// 提现账号列表，查看所有银行卡
						co_permission.Financial.PermissionType(module).BankCardList,
						// 申请提现账号，添加银行卡信息
						co_permission.Financial.PermissionType(module).CreateBankCard,
						//  删除提现账号，删除银行卡信息
						co_permission.Financial.PermissionType(module).DeleteBankCard,
					},
				},
				{
					SysPermission: &sys_entity.SysPermission{
						Id:         idgen.NextId(),
						Name:       "财务账号",
						Identifier: "Account",
						Type:       1,
						IsShow:     1,
					},
					Children: []*sys_model.SysPermissionTree{
						// 查看余额，查看账号余额
						co_permission.Financial.PermissionType(module).GetAccountBalance,

						// 查看财务账号金额明细
						co_permission.Financial.PermissionType(module).GetAccountDetail,
					},
				},
			},
		},
	}
	return result
}

func initAuditAndLicensePermission() []*sys_model.SysPermissionTree {
	result := []*sys_model.SysPermissionTree{

		// 资质权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         idgen.NextId(),
				Name:       "资质",
				Identifier: "License",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看资质信息，查看某条资质信息
				co_permission.License.PermissionType.ViewDetail,
				// 资质列表，查看所有资质信息
				co_permission.License.PermissionType.List,
				// 更新资质审核信息，更新某条资质审核信息
				co_permission.License.PermissionType.Update,
				// 创建资质，创建资质信息
				co_permission.License.PermissionType.Create,
				// 设置资质状态，设置某资质认证状态
				co_permission.License.PermissionType.SetState,
			},
		},
		// 审核管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         idgen.NextId(),
				Name:       "审核管理",
				Identifier: "Audit",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看审核信息，查看某条资质审核信息
				co_permission.Audit.PermissionType.ViewDetail,
				// 资质审核列表，查看所有资质审核
				co_permission.Audit.PermissionType.List,
				// 更新资质审核信息，更新某条资质审核信息
				co_permission.Audit.PermissionType.Update,
			},
		},
	}
	return result
}
