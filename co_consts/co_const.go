package co_consts

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/base_permission"
	"strings"
)

type global struct {
	// 默认货币类型
	DefaultCurrency string

	PlatformUserTypeArr []int

	// 客户端配置
	ClientConfig []co_model.ClientConfig
}

var (
	Global = global{
		ClientConfig: []co_model.ClientConfig{},
	}

	PermissionTree []base_permission.IPermission

	FinancePermissionTree []base_permission.IPermission

	ModulesConfigArr = make([]*co_model.Config, 0)
	ModuleArr        = make([]co_interface.IModuleBase, 0)
)

func init() {
	defaultCurrency, _ := g.Cfg().Get(context.Background(), "service.defaultCurrency", "CNY")
	Global.DefaultCurrency = defaultCurrency.String()
	Global.PlatformUserTypeArr = g.Cfg().MustGet(context.Background(), "service.platformUserType", []int{}).Ints()

	for _, clientConfig := range g.Cfg().MustGet(context.Background(), "service.clientConfig").Array() {
		configItem := co_model.ClientConfig{
			AllowSkipLevelCreateCompany:       false,
			CompanyCommissionModel:            co_enum.Common.CompanyCommissionMode.Superior.Code(),
			CompanyCommissionAllocationLevel:  3,
			EmployeeCommissionModel:           co_enum.Common.EmployeeCommissionMode.Superior.Code(),
			EmployeeCommissionLevelMode:       co_enum.Common.EmployeeCommissionLevelMode.Superior.Code(),
			EmployeeCommissionAllocationLevel: 3,
			GroupNameCanRepeated:              false,
			EmployeeNameCanRepeated:           false,
			AutoCreateUserFinanceAccount:      false,
			RegisterBindMemberLevelId:         0,
			DefaultCurrency:                    defaultCurrency.String(),
		}

		err := gconv.Struct(clientConfig, &configItem)

		if err != nil {
			g.Log().Error(context.Background(), err)
		}

		Global.ClientConfig = append(Global.ClientConfig, configItem)
	}
}

func (s global) GetClientConfig(ctx context.Context) (*co_model.ClientConfig, error) {
	xClient := ghttp.RequestFromCtx(ctx).Header.Get("X-CLIENT-ID")

	for _, v := range s.ClientConfig {
		if strings.EqualFold(v.XClientToken, xClient) 	{
			if v.DefaultCurrency == "" {
				v.DefaultCurrency = Global.DefaultCurrency
			}
			return &v, nil
		}
	}

	return nil, gerror.NewCode(gcode.CodeNotAuthorized, "error_client_info_incorrect")
}

func (s global) GetClientConfigByKey(ctx context.Context, key string) (*co_model.ClientConfig, error) {
	for _, v := range s.ClientConfig {
		if strings.EqualFold(v.XClientToken, key) {
			return &v, nil
		}
	}
	return nil, gerror.NewCode(gcode.CodeNotAuthorized, "error_client_info_incorrect")
}
