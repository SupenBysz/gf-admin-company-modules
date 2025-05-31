package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/co_consts"

	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type cModuleConf struct{}

var ModuleConf = cModuleConf{}

func (c *cModuleConf) CompanyType(ctx context.Context, req *co_v1.ModuleConfInfoReq) (*co_model.ModuleConfInfoRes, error) {
	clientConfig, err := sys_consts.Global.GetClientConfig(ctx)

	if err != nil {
		return nil, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 校验登录类型
	if !clientConfig.AllowLoginUserTypeArr.Contains(sessionUser.Type) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_user_type_mismatch", sys_dao.SysUser.Table())
	}

	for _, moduleConfig := range co_consts.ModulesConfigArr {
		if moduleConfig.UserType.Code() == sessionUser.Type {
			return &co_model.ModuleConfInfoRes{
				ModuleType: moduleConfig.UserType.Code(),
				ModulePath: moduleConfig.RoutePrefix,
			}, nil
		}
	}

	return &co_model.ModuleConfInfoRes{
		ModuleType: 0,
		ModulePath: "",
	}, nil
}
