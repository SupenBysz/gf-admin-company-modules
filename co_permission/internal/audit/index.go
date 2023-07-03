package audit

type audit struct {
	PermissionType permissionType
}

var Audit = audit{
	PermissionType: PermissionType,
}
