package audit

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_audit_v1"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
)

// Audit 审核记录
var Audit = cAudit{}

type cAudit struct{}

// GetAuditLogList 获取审核信息|列表
func (c *cAudit) GetAuditLogList(ctx context.Context, req *co_audit_v1.QueryAuditListReq) (*co_audit_v1.AuditListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, co_permission.Audit.PermissionType.List); has != true {
		return nil, err
	}

	result, err := co_service.Audit().QueryAuditList(ctx, &req.SearchParams)

	return (*co_audit_v1.AuditListRes)(result), err
}

// SetAuditApprove 审批通过
func (c *cAudit) SetAuditApprove(ctx context.Context, req *co_audit_v1.SetAuditApproveReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, co_permission.Audit.PermissionType.Update); has != true {
		return false, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	result, err := co_service.Audit().UpdateAudit(ctx, req.Id, co_enum.Audit.Action.Approve.Code(), "", user.Id)

	return result == true, err
}

// SetAuditReject 审批不通过
func (c *cAudit) SetAuditReject(ctx context.Context, req *co_audit_v1.SetAuditRejectReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, co_permission.Audit.PermissionType.Update); has != true {
		return false, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	result, err := co_service.Audit().UpdateAudit(ctx, req.Id, co_enum.Audit.Action.Reject.Code(), req.Reply, user.Id)
	return result == true, err
}

// GetAuditById 根据ID获取资质审核信息
func (c *cAudit) GetAuditById(ctx context.Context, req *co_audit_v1.GetAuditByIdReq) (*co_audit_v1.AuditRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, co_permission.Audit.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result := co_service.Audit().GetAuditById(ctx, req.Id)
	return (*co_audit_v1.AuditRes)(result), nil
}
