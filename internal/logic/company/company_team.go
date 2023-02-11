package company

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type sTeam struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewTeam(modules co_interface.IModules) co_interface.ITeam {
	return &sTeam{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetTeamById 根据ID获取公司团队信息
func (s *sTeam) GetTeamById(ctx context.Context, id int64) (*co_model.TeamRes, error) {
	data, err := daoctl.GetByIdWithError[co_model.TeamRes](
		s.dao.Team.Ctx(ctx), id,
	)

	if err != nil {
		message := s.modules.T(ctx, "{#teamOrGroup}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#teamOrGroup}{#error_Data_Get_Failed}")
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Team.Table())
	}

	return s.makeMore(ctx, data), nil
}

// GetTeamByName 根据Name获取团队信息
func (s *sTeam) GetTeamByName(ctx context.Context, name string) (*co_model.TeamRes, error) {
	data, err := daoctl.ScanWithError[co_model.TeamRes](
		s.dao.Team.Ctx(ctx).
			Where(co_do.CompanyTeam{Name: name}),
	)

	if err != nil {
		message := s.modules.T(ctx, "{#teamOrGroup}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#teamOrGroup}{#error_Data_Get_Failed}")
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Team.Table())
	}

	return s.makeMore(ctx, data), nil
}

// HasTeamByName 团队名称是否存在
func (s *sTeam) HasTeamByName(ctx context.Context, name string, excludeIds ...int64) bool {
	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	model := s.dao.Team.Ctx(ctx).Where(co_do.CompanyTeam{
		Name:        name,
		UnionMainId: unionMainId,
	})

	if len(excludeIds) > 0 {
		var ids []int64
		for _, id := range excludeIds {
			if id > 0 {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			model = model.WhereNotIn(s.dao.Team.Columns().Id, ids)
		}
	}

	count, _ := model.Count()
	return count > 0
}

// QueryTeamList 查询团队
func (s *sTeam) QueryTeamList(ctx context.Context, search *sys_model.SearchParams) (*co_model.TeamListRes, error) {
	// 跨主体查询条件过滤
	search = funs.FilterUnionMain(ctx, search, s.dao.Team.Columns().UnionMainId)

	data, err := daoctl.Query[*co_model.TeamRes](s.dao.Team.Ctx(ctx), search, false)

	items := make([]*co_model.TeamRes, 0)
	for _, item := range data.Records {
		items = append(items, s.makeMore(ctx, item))
	}
	data.Records = items

	return (*co_model.TeamListRes)(data), err
}

// QueryTeamMemberList 查询所有团队成员记录
func (s *sTeam) QueryTeamMemberList(ctx context.Context, search *sys_model.SearchParams) (*co_model.TeamMemberListRes, error) {
	model := s.dao.TeamMember.Ctx(ctx)

	data, err := daoctl.Query[*co_model.TeamMemberRes](model, search, false)

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	var UnionMain *co_model.CompanyRes
	if sessionUser.UnionMainId > 0 {
		UnionMain, _ = s.modules.Company().GetCompanyById(ctx, sessionUser.UnionMainId)
	}

	items := make([]*co_model.TeamMemberRes, 0)
	for _, item := range data.Records {
		if item.EmployeeId > 0 {
			item.Employee, _ = s.modules.Employee().GetEmployeeById(ctx, item.EmployeeId)
		}
		if item.InviteUserId > 0 {
			item.InviteUser, _ = s.modules.Employee().GetEmployeeById(ctx, item.InviteUserId)
		}
		if item.UnionMainId == sessionUser.UnionMainId {
			item.UnionMain = UnionMain
		} else if item.UnionMainId > 0 {
			UnionMain, _ = s.modules.Company().GetCompanyById(ctx, item.UnionMainId)
		}
		items = append(items, item)
	}
	data.Records = items

	return (*co_model.TeamMemberListRes)(data), err
}

// CreateTeam 创建团队或小组|信息
func (s *sTeam) CreateTeam(ctx context.Context, info *co_model.Team) (*co_model.TeamRes, error) {
	if info.ParentId > 0 {
		team, _ := s.GetTeamById(ctx, info.ParentId)
		if team == nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_ParentTeamNotFound"), s.dao.Team.Table())
		}
		if team.ParentId > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Group_ParentMustIsTeam"), s.dao.Team.Table())
		}
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断团队名称是否存在
	if s.HasTeamByName(ctx, info.Name) == true {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_TeamNameExist"), s.dao.Team.Table())
	}

	// 判断团队管理人信息是否存在
	if info.OwnerEmployeeId > 0 {
		_, err := s.modules.Employee().GetEmployeeById(ctx, info.OwnerEmployeeId)
		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamOwnerEmployee}{#error_Data_NotFound}"), s.dao.Team.Table())
		}
	}

	if info.CaptainEmployeeId > 0 {
		employee, err := s.modules.Employee().GetEmployeeById(ctx, info.CaptainEmployeeId)
		if err != nil || employee.UnionMainId != sessionUser.UnionMainId {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamOwnerEmployee}{#error_Data_NotFound}"), s.dao.Team.Table())
		}

		data, err := s.QueryTeamListByEmployee(ctx, employee.Id, employee.UnionMainId)
		if err != nil && err != sql.ErrNoRows {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamOwnerEmployee}{#error_Data_NotFound}"), s.dao.Team.Table())
		}

		if info.ParentId == 0 {
			for _, team := range data.Records {
				if team.ParentId == 0 {
					return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "TeamCaptainEmployee")+"不能是其它团队的队员", s.dao.Team.Table())
				}
			}
		}
	}

	data := co_do.CompanyTeam{
		Id:                idgen.NextId(),
		Name:              info.Name,
		Remark:            info.Remark,
		ParentId:          info.ParentId,
		OwnerEmployeeId:   info.OwnerEmployeeId,
		CaptainEmployeeId: info.CaptainEmployeeId,
		UnionMainId:       sessionUser.UnionMainId,
		CreatedAt:         gtime.Now(),
	}
	captain := co_do.CompanyTeamMember{
		Id:          idgen.NextId(),
		TeamId:      data.Id,
		EmployeeId:  info.CaptainEmployeeId,
		UnionMainId: sessionUser.UnionMainId,
		JoinAt:      gtime.Now(),
	}

	err := s.dao.Team.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 创建团队
		affected, err := daoctl.InsertWithError(
			s.dao.Team.Ctx(ctx).Data(data),
		)
		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_Save_Failed"), s.dao.Team.Table())
		}
		if info.CaptainEmployeeId > 0 {
			// 创建团队队长
			_, err = s.dao.TeamMember.Ctx(ctx).Data(captain).Insert()
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_Save_Failed")+"无法保存"+s.modules.T(ctx, "TeamCaptainEmployee")+"信息", s.dao.Team.Table())
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	result, err := s.GetTeamById(ctx, data.Id.(int64))
	return s.makeMore(ctx, result), err
}

// UpdateTeam 更新团队或小组|信息
func (s *sTeam) UpdateTeam(ctx context.Context, id int64, name string, remark string) (*co_model.TeamRes, error) {

	if s.HasTeamByName(ctx, name, id) == true {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_TeamNameExist"), s.dao.Team.Table())
	}

	data := co_do.CompanyTeam{
		Name:      name,
		Remark:    remark,
		UpdatedAt: gtime.Now(),
	}

	rowsAffected, err := daoctl.UpdateWithError(
		s.dao.Team.Ctx(ctx).
			Data(data).
			Where(co_do.CompanyTeam{Id: id}),
	)

	if rowsAffected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_Save_Failed"), s.dao.Team.Table())
	}

	result, err := s.GetTeamById(ctx, id)
	return s.makeMore(ctx, result), err
}

// GetTeamMemberList 获取团队成员|列表
func (s *sTeam) GetTeamMemberList(ctx context.Context, id int64) (*co_model.EmployeeListRes, error) {
	team, err := s.GetTeamById(ctx, id)
	if err != nil {
		return nil, err
	}

	// 团队成员信息
	items, err := daoctl.ScanWithError[[]*co_entity.CompanyTeamMember](
		s.dao.TeamMember.Ctx(ctx).Where(co_do.CompanyTeamMember{
			TeamId:      team.Id,
			UnionMainId: team.UnionMainId,
		}),
	)

	ids := make([]int64, 0)
	for _, item := range *items {
		ids = append(ids, item.EmployeeId)
	}

	return s.modules.Employee().QueryEmployeeList(ctx, &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0),
			sys_model.FilterInfo{
				Field: s.dao.Employee.Columns().Id,
				Where: "in",
				Value: ids,
			},
			sys_model.FilterInfo{
				Field: s.dao.Employee.Columns().UnionMainId,
				Where: "=",
				Value: team.UnionMainId,
			},
		),
	})
}

// QueryTeamListByEmployee 根据员工查询团队
func (s *sTeam) QueryTeamListByEmployee(ctx context.Context, employeeId int64, unionMainId int64) (*co_model.TeamListRes, error) {

	if unionMainId == 0 {
		unionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId
	}

	data, err := daoctl.ScanWithError[[]*co_entity.CompanyTeamMember](
		s.dao.TeamMember.Ctx(ctx).
			Where(co_do.CompanyTeamMember{EmployeeId: employeeId, UnionMainId: unionMainId}),
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_NotFound"), s.dao.Team.Table())
	}

	var teamIds []int64
	for _, member := range *data {
		teamIds = append(teamIds, member.TeamId)
	}

	return s.QueryTeamList(ctx, &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0),
			sys_model.FilterInfo{
				Field: s.dao.Team.Columns().UnionMainId,
				Where: "=",
				Value: unionMainId,
			},
			sys_model.FilterInfo{
				Field: s.dao.Team.Columns().Id,
				Where: "in",
				Value: teamIds,
			},
		),
	})
}

// SetTeamMember 设置团队队员或小组组员
func (s *sTeam) SetTeamMember(ctx context.Context, teamId int64, employeeIds []int64) (api_v1.BoolRes, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 获取团队所有旧成员
	teamMemberArr, err := daoctl.ScanWithError[[]*co_entity.CompanyTeamMember](
		s.dao.TeamMember.Ctx(ctx).
			Where(co_do.CompanyTeamMember{
				TeamId:      teamId,
				UnionMainId: sessionUser.UnionMainId,
			}),
	)

	// 待移除的团队成员
	waitIds := make([]int64, 0)
	// 已存在的团队成员
	existIds := make([]int64, 0)

	// 遍历所有旧成员
	for _, member := range *teamMemberArr {
		if len(employeeIds) == 0 {
			waitIds = append(existIds, member.EmployeeId)
			continue
		}
		// 遍历待加入团队的员工
		for _, employeeId := range employeeIds {
			if member.EmployeeId != employeeId {
				// 追加已移除的团队成员ID到待移除数组
				waitIds = append(waitIds, employeeId)
			} else {
				existIds = append(existIds, employeeId)
			}
		}
	}

	// 新团队成员Ids
	newTeamMemberIds := make([]int64, 0)
	for _, employeeId := range employeeIds {
		has := false
		for _, id := range existIds {
			if employeeId == id {
				has = true
			}
		}
		if has == false {
			newTeamMemberIds = append(newTeamMemberIds, employeeId)
		}
	}

	// 如果新团队成员为空，则直接移除所有团队成员
	if len(newTeamMemberIds) <= 0 {
		model := s.dao.TeamMember.Ctx(ctx).
			Where(co_do.CompanyTeamMember{TeamId: teamId, UnionMainId: sessionUser.UnionMainId})

		if len(existIds) > 0 {
			model = model.WhereNotIn(s.dao.TeamMember.Columns().EmployeeId, existIds)
		}

		if _, err = model.Delete(); err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_DeleteMember_Failed"), s.dao.Team.Table())
		}
		return true, nil
	}

	// 校验新团队成员是否存在
	res, err := s.modules.Employee().QueryEmployeeList(ctx, &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0),
			sys_model.FilterInfo{
				Field: s.dao.Employee.Columns().Id,
				Where: "in",
				Value: newTeamMemberIds,
			},
			sys_model.FilterInfo{
				Field: s.dao.Employee.Columns().UnionMainId,
				Where: "=",
				Value: sessionUser.UnionMainId,
			},
		),
		Pagination: sys_model.Pagination{
			PageNum:  1,
			PageSize: 1000,
		},
	})

	if res.Total < gconv.Int64(len(newTeamMemberIds)) {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_NewTeam_NotFoundMembers"), s.dao.Team.Table())
	}

	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	//
	if team.ParentId == 0 {
		count, _ := s.dao.TeamMember.Ctx(ctx).
			WhereIn(s.dao.TeamMember.Columns().EmployeeId, newTeamMemberIds).
			Where(s.dao.TeamMember.Columns().UnionMainId, sessionUser.UnionMainId).Count()
		if count > 0 {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_MemberIsHasTeam"), s.dao.Team.Table())
		}
	}

	err = s.dao.TeamMember.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 清理团队成员
		_, err = s.dao.TeamMember.Ctx(ctx).
			WhereIn(s.dao.TeamMember.Columns().Id, existIds).
			Delete()

		if err != nil {
			return err
		}

		for _, employeeId := range newTeamMemberIds {
			affected, err := daoctl.InsertWithError(
				s.dao.TeamMember.Ctx(ctx).Data(
					co_do.CompanyTeamMember{
						Id:          idgen.NextId(),
						TeamId:      team.Id,
						EmployeeId:  employeeId,
						UnionMainId: sessionUser.UnionMainId,
						JoinAt:      gtime.Now(),
					},
				),
			)
			if affected == 0 || err != nil {
				return err
			}
		}
		return nil
	})

	return err == nil, err
}

// SetTeamOwner 设置团队或小组的负责人
func (s *sTeam) SetTeamOwner(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error) {
	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	if team.OwnerEmployeeId == employeeId {
		return true, nil
	}

	// 需要删除团队负责人的情况
	if team.Id != 0 && employeeId == 0 {
		affected, err := daoctl.UpdateWithError(s.dao.Team.Ctx(ctx).
			Where(co_do.CompanyTeam{Id: team.Id}).
			Data(co_do.CompanyTeam{OwnerEmployeeId: 0}),
		)
		return affected == 1, err
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, employeeId)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 校验数据主体是否一致
	if sessionUser.UnionMainId != team.UnionMainId || sessionUser.UnionMainId != employee.UnionMainId {
		if team.ParentId <= 0 {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamOrEmployee_Check_Failed"), s.dao.Team.Table())
		} else {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_GroupOrEmployee_Check_Failed"), s.dao.Team.Table())
		}
	}

	affected, err := daoctl.UpdateWithError(
		s.dao.Team.Ctx(ctx).
			Data(co_do.CompanyTeam{OwnerEmployeeId: employee.Id}).
			Where(co_do.CompanyTeam{Id: team.Id}),
	)

	return affected == 1, err
}

// SetTeamCaptain 设置团队队长或小组组长
func (s *sTeam) SetTeamCaptain(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error) {
	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	if team.CaptainEmployeeId == employeeId {
		return true, nil
	}

	// 需要删除团队队长或者组长的情况
	if employeeId == 0 && team.Id != 0 {
		affected, err := daoctl.UpdateWithError(
			s.dao.Team.Ctx(ctx).
				Data(co_do.CompanyTeam{CaptainEmployeeId: 0}).
				Where(co_do.CompanyTeam{Id: team.Id}),
		)
		return affected == 1, err
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, employeeId)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 校验数据主体是否一致
	if sessionUser.UnionMainId != team.UnionMainId || sessionUser.UnionMainId != employee.UnionMainId {
		if team.ParentId <= 0 {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamOrEmployee_Check_Failed"), s.dao.Team.Table())
		} else {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_GroupOrEmployee_Check_Failed"), s.dao.Team.Table())
		}
	}

	// 员工能否设置为队长
	canCaptain := false
	{
		// 查询员工所在的所有团队信息
		data, err := s.QueryTeamListByEmployee(ctx, employee.Id, employee.UnionMainId)
		if err != nil && err != sql.ErrNoRows {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamCaptainEmployee}{#error_Data_NotFound}"), s.dao.Team.Table())
		}

		for _, item := range data.Records {
			// 判断要设置的是团队还是小组 ParentId == 0团队，ParentId > 0小组
			if team.ParentId == 0 && item.ParentId == 0 {
				// 如果员工是其它团队成员则返回
				if item.Id != team.Id {
					return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_MemberIsHasTeam"), s.dao.Team.Table())
				} else {
					canCaptain = true
				}
			}
		}
	}

	if team.ParentId == 0 && !canCaptain {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamCaptainEmployee_MustInTeam"), s.dao.Team.Table())
	}

	affected, err := daoctl.UpdateWithError(
		s.dao.Team.Ctx(ctx).
			Where(co_do.CompanyTeam{Id: team.Id}).
			Data(co_do.CompanyTeam{CaptainEmployeeId: employee.Id}),
	)

	return affected == 1, err
}

// DeleteTeam 删除团队
func (s *sTeam) DeleteTeam(ctx context.Context, teamId int64) (api_v1.BoolRes, error) {
	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 查询团队成员数量
	count, err := s.dao.TeamMember.Ctx(ctx).
		Where(co_do.CompanyTeamMember{
			TeamId:      team.Id,
			UnionMainId: sessionUser.UnionMainId,
		}).Count()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamMember}{#error_Data_Get_Failed}"), s.dao.Team.Table())
	}

	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_NeedRemoveTeamMember"), s.dao.Team.Table())
	}

	affected, err := daoctl.DeleteWithError(
		s.dao.Team.Ctx(ctx).Unscoped().
			Where(co_do.CompanyTeam{Id: team.Id}),
	)

	return affected == 1, err
}

// DeleteTeamMemberByEmployee 删除某个员工的所有团队成员记录
func (s *sTeam) DeleteTeamMemberByEmployee(ctx context.Context, employeeId int64) (bool, error) {
	affected, err := daoctl.DeleteWithError(s.dao.TeamMember.Ctx(ctx).Where(co_do.CompanyTeamMember{EmployeeId: employeeId}))

	return affected > 0, err
}

// makeMore 按需加载附加数据
func (s *sTeam) makeMore(ctx context.Context, data *co_model.TeamRes) *co_model.TeamRes {
	if data.OwnerEmployeeId > 0 {
		funs.AttrMake[co_model.TeamRes](ctx,
			s.dao.Team.Columns().OwnerEmployeeId,
			func() *co_model.EmployeeRes {
				data.Owner, _ = s.modules.Employee().GetEmployeeById(ctx, data.OwnerEmployeeId)
				user, _ := sys_service.SysUser().GetSysUserById(ctx, data.OwnerEmployeeId)
				gconv.Struct(user.SysUser, &data.Owner.User)
				gconv.Struct(user.Detail, &data.Owner.Detail)
				return data.Owner
			},
		)
	}
	if data.CaptainEmployeeId > 0 {
		funs.AttrMake[co_model.TeamRes](ctx,
			s.dao.Team.Columns().CaptainEmployeeId,
			func() *co_model.EmployeeRes {
				data.Captain, _ = s.modules.Employee().GetEmployeeById(ctx, data.CaptainEmployeeId)
				user, _ := sys_service.SysUser().GetSysUserById(ctx, data.CaptainEmployeeId)
				gconv.Struct(user.SysUser, &data.Captain.User)
				gconv.Struct(user.Detail, &data.Captain.Detail)
				return data.Captain
			},
		)
	}
	if data.UnionMainId > 0 {
		funs.AttrMake[co_model.TeamRes](ctx,
			s.dao.Team.Columns().UnionMainId,
			func() *co_model.CompanyRes {
				data.UnionMain, _ = s.modules.Company().GetCompanyById(ctx, data.UnionMainId)
				return data.UnionMain
			},
		)
	}
	if data.ParentId > 0 {
		funs.AttrMake[co_model.TeamRes](ctx,
			s.dao.Team.Columns().ParentId,
			func() *co_model.TeamRes {
				data.Parent, _ = s.modules.Team().GetTeamById(ctx, data.ParentId)
				return data.Parent
			},
		)
	}
	return data
}
