// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package co_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type (
	ILicense interface {
		// GetAuditData 订阅审核数据获取Hook, 将审核数据渲染成主体资质然后进行输出  (数据处理渲染)
		GetAuditData(ctx context.Context, auditEvent sys_enum.AuditEvent, info *sys_entity.SysAudit) error
		// AuditChange 审核成功的处理逻辑 Hook （数据存储）
		AuditChange(ctx context.Context, auditEvent sys_enum.AuditEvent, info *sys_entity.SysAudit) error
		// GetLicenseById 根据ID获取主体认证|信息
		GetLicenseById(ctx context.Context, id int64) (*co_entity.License, error)
		// QueryLicenseList 查询主体认证|列表
		QueryLicenseList(ctx context.Context, search base_model.SearchParams) (*co_model.LicenseListRes, error)
		// CreateLicense 新增主体资质|信息
		CreateLicense(ctx context.Context, info co_model.License) (*co_entity.License, error)
		// UpdateLicense 更新主体认证，如果是已经通过的认证，需要重新认证通过后才生效|信息
		UpdateLicense(ctx context.Context, info co_model.License, id int64) (*co_entity.License, error)
		// GetLicenseByLatestAuditId 获取最新的审核记录Id获取资质信息
		GetLicenseByLatestAuditId(ctx context.Context, auditId int64) *co_entity.License
		// SetLicenseState 设置主体信息状态
		SetLicenseState(ctx context.Context, id int64, state int) (bool, error)
		// SetLicenseAuditNumber 设置主体神审核编号
		SetLicenseAuditNumber(ctx context.Context, id int64, auditNumber string) (bool, error)
		// DeleteLicense 删除主体
		DeleteLicense(ctx context.Context, id int64, flag bool) (bool, error)
		// UpdateLicenseAuditLogId 设置主体资质关联的审核ID
		UpdateLicenseAuditLogId(ctx context.Context, id int64, latestAuditLogId int64) (bool, error)
		// Masker 资质信息脱敏
		Masker(license *co_entity.License) *co_entity.License
	}
)

var (
	localLicense ILicense
)

func License() ILicense {
	if localLicense == nil {
		panic("implement not found for interface ILicense, forgot register?")
	}
	return localLicense
}

func RegisterLicense(i ILicense) {
	localLicense = i
}
