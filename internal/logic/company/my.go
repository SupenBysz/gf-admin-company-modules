package company

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/kconv"
	"github.com/SupenBysz/gf-admin-community/utility/masker"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type sMy struct {
	modules co_interface.IModules
}

func NewMy(modules co_interface.IModules) co_interface.IMy {
	return &sMy{
		modules: modules,
	}
}

// GetProfile 获取当前员工及用户信息
func (s *sMy) GetProfile(ctx context.Context) (*co_model.MyProfileRes, error) {
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser

	user, err := sys_service.SysUser().GetSysUserById(ctx, session.Id)
	if err != nil {
		return nil, err
	}
	// 信息脱敏
	user.Password = masker.MaskString(user.Password, masker.Password)

	// 超级管理员不脱敏
	if session.Type == sys_enum.User.Type.SuperAdmin.Code() {
		return &co_model.MyProfileRes{
			SysUser: *user,
		}, nil
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, session.Id)
	if err != nil && employee == nil {
		return &co_model.MyProfileRes{
			SysUser:         *user,
			CompanyEmployee: nil,
		}, nil
	}

	// 信息脱敏
	employee = s.modules.Employee().Masker(employee)

	return &co_model.MyProfileRes{
		SysUser:         *user,
		CompanyEmployee: employee,
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

	// 如果当前登录员工不是管理员公司数据需要脱敏
	if employee.Id != company.UserId {
		company = s.modules.Company().Masker(company)
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
		return nil, err
	}

	// 团队成员列表
	for _, team := range teamList.Records {
		var teamInfo co_model.MyTeamRes

		// 团队
		teamInfo.CompanyTeam = *team

		// 团队成员列表
		memberList, err := s.modules.Team().GetTeamMemberList(ctx, team.Id)
		if err != nil {
			return nil, err
		}

		teamInfo.EmployeeListRes = *memberList

		// 员工信息脱敏,保留手机号
		employeeList := make([]*co_entity.CompanyEmployee, 0)
		for _, member := range teamInfo.EmployeeListRes.Records {
			mobile := member.Mobile
			member = s.modules.Employee().Masker(member)

			member.Mobile = mobile
			employeeList = append(employeeList, member)
		}

		// 赋值
		kconv.Struct(employeeList, &teamInfo.EmployeeListRes)
		res = append(res, teamInfo)
	}

	return res, nil
}

// Masker 员工信息脱敏
func (s *sMy) masker(employee *co_entity.CompanyEmployee) *co_entity.CompanyEmployee {
	if employee == nil {
		return nil
	}
	employee.Mobile = masker.MaskString(employee.Mobile, masker.MaskPhone)
	employee.LastActiveIp = masker.MaskString(employee.LastActiveIp, masker.MaskIPv4)
	return employee
}
