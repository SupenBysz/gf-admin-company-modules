package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/SupenBysz/gf-admin-company-modules/utility/co_rules"
	"github.com/kysion/base-library/utility/base_permission"
)

// InitCustomRules 注册自定义参数校验规则
func InitCustomRules() {
	// 注册资质自定义规则
	co_rules.RequiredLicense()
	InitializePermissionFactory()
}

func InitializePermissionFactory() {
	base_permission.InitializePermissionFactory(func() base_permission.IPermission {
		return &sys_model.SysPermissionTree{
			SysPermission: &sys_entity.SysPermission{},
		}
	})
}

// InitPermission 初始化权限树
func InitPermission[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](module co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) []base_permission.IPermission {
	InitializePermissionFactory()

	result := []base_permission.IPermission{
		// 公司
		// 导入权限的时候判断的是标识符号，所以不用担心下一次启动导入id不同的相同权限
		base_permission.New(idgen.NextId(), module.GetConfig().Identifier.Company, module.T(context.Background(), "CompanyName")).
			SetMatchMode(1).
			SetType(1).
			SetIsShow(1).
			SetItems([]base_permission.IPermission{
				co_permission.Company.PermissionType(module).Create,
				co_permission.Company.PermissionType(module).ViewDetail,
				co_permission.Company.PermissionType(module).List,
				co_permission.Company.PermissionType(module).Update,
				co_permission.Company.PermissionType(module).SetLogo,
				co_permission.Company.PermissionType(module).SetState,
				co_permission.Company.PermissionType(module).SetAdminUser,
				co_permission.Company.PermissionType(module).ViewLicense,
				co_permission.Company.PermissionType(module).AuditLicense,

				// 员工
				// 之前都是固定ID，后面换成了随机ID，所以避免获取的时候ID不固定，所以根据标识符进行匹配1 （是否固定ID，看构建权限信息是用的base_permission.New，还是base_permission.NewInIdentifier ）
				base_permission.New(idgen.NextId(), module.GetConfig().Identifier.Employee, "员工管理").
					SetMatchMode(1).
					SetType(1).
					SetIsShow(1).
					SetItems([]base_permission.IPermission{
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
					}),

				// 团队
				base_permission.New(idgen.NextId(), module.GetConfig().Identifier.Team, "团队管理").
					SetType(1).
					SetIsShow(1).
					SetMatchMode(1).
					SetItems([]base_permission.IPermission{
						co_permission.Team.PermissionType(module).Create,
						co_permission.Team.PermissionType(module).ViewDetail,
						co_permission.Team.PermissionType(module).List,
						co_permission.Team.PermissionType(module).Update,
						co_permission.Team.PermissionType(module).Delete,
						co_permission.Team.PermissionType(module).MemberDetail,
						co_permission.Team.PermissionType(module).SetMember,
						co_permission.Team.PermissionType(module).SetOwner,
						co_permission.Team.PermissionType(module).SetCaptain,
					}),
			}),
	}
	// sms短信

	// oss

	//result = append(result, licensePermission...)

	return result
}

// InitFinancePermission 初始化财务服务权限树
func InitFinancePermission[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](module co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) []base_permission.IPermission {
	InitializePermissionFactory()

	result := []base_permission.IPermission{
		// 财务服务权限树
		base_permission.New(idgen.NextId(), "Finance", "财务").
			SetType(1).
			SetIsShow(1).
			SetItems([]base_permission.IPermission{
				base_permission.New(idgen.NextId(), "Invoice", "发票").
					SetType(1).
					SetIsShow(1).
					SetItems([]base_permission.IPermission{
						// 查看发票详情，查看发票详情信息
						co_permission.Finance.PermissionType(module).ViewInvoiceDetail,
						// 查看发票抬头信息，查看发票抬头信息
						co_permission.Finance.PermissionType(module).ViewInvoice,
						// 发票抬头列表，查看所有发票抬头
						co_permission.Finance.PermissionType(module).InvoiceList,
						// 发票详情列表，查看所有发票详情
						co_permission.Finance.PermissionType(module).InvoiceDetailList,
						// 审核发票，审核发票申请
						co_permission.Finance.PermissionType(module).AuditInvoiceDetail,
						// 开发票，添加发票详情记录
						co_permission.Finance.PermissionType(module).MakeInvoiceDetail,
						// 添加发票抬头，添加发票抬头信息
						co_permission.Finance.PermissionType(module).CreateInvoice,
						// 删除发票抬头，删除发票抬头信息
						co_permission.Finance.PermissionType(module).DeleteInvoice,
					}),

				base_permission.New(idgen.NextId(), "BankCard", "银行卡").
					SetType(1).
					SetIsShow(1).
					SetItems([]base_permission.IPermission{
						// 查看提现账号，查看银行卡账号信息
						co_permission.Finance.PermissionType(module).ViewBankCardDetail,
						// 提现账号列表，查看所有银行卡
						co_permission.Finance.PermissionType(module).BankCardList,
						// 申请提现账号，添加银行卡信息
						co_permission.Finance.PermissionType(module).CreateBankCard,
						//  删除提现账号，删除银行卡信息
						co_permission.Finance.PermissionType(module).DeleteBankCard,
					},
					),

				base_permission.New(idgen.NextId(), "Account", "财务账号").
					SetType(1).
					SetIsShow(1).
					SetItems([]base_permission.IPermission{
						// 查看余额，查看账号余额
						co_permission.Finance.PermissionType(module).GetAccountBalance,

						// 查看财务账号金额明细
						co_permission.Finance.PermissionType(module).GetAccountDetail,
					}),
			}),
	}
	return result
}

func InitAuditAndLicensePermission() []base_permission.IPermission {
	InitializePermissionFactory()

	result := []base_permission.IPermission{
		// 资质权限树
		base_permission.New(idgen.NextId(), "License", "公司资质管理").
			SetId(idgen.NextId()).
			SetName("公司资质管理").
			SetIdentifier("License").
			SetType(1).
			SetIsShow(1).
			SetMatchMode(1).
			SetItems([]base_permission.IPermission{
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
			}),
	}

	return result
}
