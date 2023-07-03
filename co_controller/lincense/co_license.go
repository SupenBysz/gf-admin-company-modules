package lincense

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_license_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
)

// License 合作伙伴资质
var License = cLicense{}

type cLicense struct{}

// GetLicenseById 根据ID获取主体资质|信息
func (c *cLicense) GetLicenseById(ctx context.Context, req *co_license_v1.GetLicenseByIdReq) (*co_license_v1.LicenseRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, co_permission.License.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result, err := co_service.License().GetLicenseById(ctx, req.Id)
	return (*co_license_v1.LicenseRes)(result), err
}

// QueryLicenseList 查询主体认证|列表
func (c *cLicense) QueryLicenseList(ctx context.Context, req *co_license_v1.QueryLicenseListReq) (*co_license_v1.LicenseListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, co_permission.License.PermissionType.List); has != true {
		return nil, err
	}

	result, err := co_service.License().QueryLicenseList(ctx, req.SearchParams)

	if err != nil {
		return nil, err
	}

	return (*co_license_v1.LicenseListRes)(result), err
}

// CreateLicense 新增主体认证|信息
// func (c *cLicense) CreateLicense(ctx context.Context, req *sys_api.CreateLicenseReq) (*sys_api.LicenseRes, error) {
//	result, err := sys_service.License().CreateLicense(ctx, req.License, req.OperatorId)
//	return (*sys_api.LicenseRes)(result), err
// }

// // UpdateLicense 更新主体资质|信息
// func (c *cLicense) UpdateLicense(ctx context.Context, req *sys_api.UpdateLicenseReq) (*sys_api.LicenseRes, error) {
// 	result, err := sys_service.License().UpdateLicense(ctx, req.License, req.Id)
// 	return (*sys_api.LicenseRes)(result), err
// }

// // SetLicenseState 设置主体信息状态
// func (c *cLicense) SetLicenseState(ctx context.Context, req *sys_api.SetLicenseStateReq) (api_sys_api.BoolRes, error) {
//	result, err := sys_service.License().SetLicenseState(ctx, req.Id, req.State)
//	return result == true, err
// }

// DeleteLicense 软删除主体资质
// func (c *cLicense) DeleteLicense(ctx context.Context, req *sys_api.DeleteLicenseReq) (api_sys_api.BoolRes, error) {
// 	result, err := sys_service.License().DeleteLicense(ctx, req.Id, true)
// 	return result == true, err
// }
