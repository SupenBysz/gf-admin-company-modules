package company

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/en_crypto"
	"github.com/kysion/base-library/utility/kconv"
)

type sMy struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewMy(modules co_interface.IModules) co_interface.IMy {
	return &sMy{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetProfile 获取当前员工及用户信息
func (s *sMy) GetProfile(ctx context.Context) (*co_model.MyProfileRes, error) {
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser

	user, err := sys_service.SysUser().GetSysUserById(ctx, session.Id)
	if err != nil {
		return nil, err
	}

	// 超级管理员直接返回用户信息
	if session.Type == sys_enum.User.Type.SuperAdmin.Code() {
		return &co_model.MyProfileRes{
			User: user,
		}, nil
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, session.Id)
	if err != nil && employee == nil {
		return &co_model.MyProfileRes{
			User:     user,
			Employee: nil,
		}, nil
	}

	return &co_model.MyProfileRes{
		User:     user,
		Employee: employee,
	}, nil
}

// GetCompany 获取当前公司信息
func (s *sMy) GetCompany(ctx context.Context) (*co_model.MyCompanyRes, error) {
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if session.Type == sys_enum.User.Type.SuperAdmin.Code() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_SuperAdminNotServer"), "my")
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, session.Id)
	if err != nil {
		return nil, err
	}

	// 公司信息
	company, err := s.modules.Company().GetCompanyById(ctx, employee.UnionMainId)
	if err != nil {
		return nil, err
	}

	result := kconv.Struct(company, &co_model.MyCompanyRes{})

	return result, nil
}

// GetTeams 获取当前员工团队信息
func (s *sMy) GetTeams(ctx context.Context) (res co_model.MyTeamListRes, err error) {
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断身份类型（超级管理员不支持此操作）
	if session.Type == sys_enum.User.Type.SuperAdmin.Code() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_SuperAdminNotServer"), "my")
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, session.Id)
	if err != nil {
		return nil, err
	}

	// 团队列表
	teamList, err := s.modules.Team().QueryTeamListByEmployee(ctx, employee.Id, employee.UnionMainId)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamList}{#error_Data_Get_Failed}"), s.dao.Team.Table())
	}

	// 团队成员列表
	for _, team := range teamList.Records {
		var teamInfo co_model.MyTeamRes

		// 团队
		teamInfo.TeamRes = *team

		// 团队成员列表
		memberList, err := s.modules.Team().GetTeamMemberList(ctx, team.Id)
		if err != nil {
			return nil, err
		}

		teamInfo.EmployeeListRes = *memberList

		// 赋值
		res = append(res, teamInfo)
	}

	return res, nil
}

// SetMyMobile 设置我的手机号
func (s *sMy) SetMyMobile(ctx context.Context, newMobile int64, captcha string, password string) (bool, error) {
	_, err := sys_service.SysSms().Verify(ctx, newMobile, captcha)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 如果原手机号码和新号码一致，直接返回true
	userInfo, err := sys_service.SysUser().GetUserDetail(ctx, sessionUser.Id)
	if err != nil {
		return false, err
	}

	if newMobile == gconv.Int64(userInfo.Mobile) {
		return true, nil
	}

	pwdHash, err := en_crypto.PwdHash(password, gconv.String(sessionUser.Id))

	checkPassword, _ := sys_service.SysUser().CheckPassword(ctx, sessionUser.Id, pwdHash)
	if checkPassword != true {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_SetMobile_Failed"), s.dao.Employee.Table())
	}

	_, err = s.dao.Employee.Ctx(ctx).
		Data(co_do.CompanyEmployee{Mobile: newMobile, UpdatedBy: sessionUser.Id, UpdatedAt: gtime.Now()}).
		Where(co_do.CompanyEmployee{Id: sessionUser.Id}).
		Update()

	affected, err := daoctl.UpdateWithError(s.dao.Employee.Ctx(ctx).
		Data(co_do.CompanyEmployee{Mobile: newMobile, UpdatedBy: sessionUser.Id, UpdatedAt: gtime.Now()}).
		Where(co_do.CompanyEmployee{Id: sessionUser.Id}))

	if err != nil || affected == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_SetMobile_Failed"), s.dao.Employee.Table())
	}

	return true, nil
}

// SetMyAvatar 设置我的头像
func (s *sMy) SetMyAvatar(ctx context.Context, imageId int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 校验员工头像并保存
	fileInfo, err := sys_service.File().GetFileById(ctx, imageId, "头像"+s.modules.T(ctx, "error_File_FileVoid"))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "", s.dao.Employee.Table())
	}

	storageAddr := s.modules.GetConfig().StoragePath + "/employee/" + gconv.String(sessionUser.Id) + "/avatar." + fileInfo.Ext

	_, err = sys_service.File().SaveFile(ctx, storageAddr, fileInfo)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "头像"+s.modules.T(ctx, "error_File_Save_Failed"), s.dao.Employee.Table())
	}
	return true, nil
}
