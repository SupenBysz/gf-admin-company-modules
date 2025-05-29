package co_consts

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/utility/base_permission"
)

type global struct {
	// 默认货币类型
	DefaultCurrency string

	PlatformUserTypeArr []int

	// 是否允许主体下的不同团队内的小组同名
	GroupNameCanRepeated bool

	// 是否允许主体下的员工同名
	EmployeeNameCanRepeated bool

	// 是否自动创建员工财务账户
	AutoCreateUserFinanceAccount bool

	// 是否允许跨级创建子公司单位
	AllowSkipLevelCreateCompany bool

	CommissionModel co_enum.CommissionMode
}

var (
	Global = global{}

	PermissionTree []base_permission.IPermission

	FinancePermissionTree []base_permission.IPermission
)

func init() {
	defaultCurrency, _ := g.Cfg().Get(context.Background(), "service.defaultCurrency")
	Global.DefaultCurrency = defaultCurrency.String()
	Global.GroupNameCanRepeated = g.Cfg().MustGet(context.Background(), "service.groupNameCanRepeated", false).Bool()
	Global.EmployeeNameCanRepeated = g.Cfg().MustGet(context.Background(), "service.employeeNameCanRepeated", true).Bool()
	Global.AutoCreateUserFinanceAccount = g.Cfg().MustGet(context.Background(), "service.autoCreateUserFinanceAccount", true).Bool()
	Global.AllowSkipLevelCreateCompany = g.Cfg().MustGet(context.Background(), "service.allowSkipLevelCreateCompany", false).Bool()
	Global.PlatformUserTypeArr = g.Cfg().MustGet(context.Background(), "service.platformUserType", []int{}).Ints()
	Global.CommissionModel = co_enum.Common.CommissionMode.New(g.Cfg().MustGet(context.Background(), "service.commissionModel", co_enum.Common.CommissionMode.Superior.Code()).Int(), "")
}
