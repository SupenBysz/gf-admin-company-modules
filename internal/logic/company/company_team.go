package company

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/funs"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type sTeam struct {
	modules co_interface.IModules
}

func NewTeam(modules co_interface.IModules) co_interface.ITeam {
	return &sTeam{
		modules: modules,
	}
}

// GetTeamById 根据ID获取公司团队信息
func (s *sTeam) GetTeamById(ctx context.Context, id int64) (*co_entity.CompanyTeam, error) {
	data := co_entity.CompanyTeam{}
	err := co_dao.CompanyTeam(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Scan(&data, co_do.CompanyTeam{Id: id})

	if err != nil {
		message := s.modules.T(ctx, "teamOrGroup") + s.modules.T(ctx, "error_Data_NotFound")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "teamOrGroup") + s.modules.T(ctx, "error_Data_Get_Failed")
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, message, co_dao.CompanyEmployee(s.modules).Table())
	}
	return &data, nil
}

// HasTeamByName 团队名称是否存在
func (s *sTeam) HasTeamByName(ctx context.Context, name string, unionMainId int64, excludeId ...int64) bool {
	model := co_dao.CompanyTeam(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.CompanyTeam{
		Name:        name,
		UnionMainId: unionMainId,
	})

	if len(excludeId) > 0 {
		model = model.WhereNotIn(co_dao.CompanyTeam(s.modules).Columns().Id, excludeId)
	}

	count, _ := model.Count()
	return count > 0
}

// QueryTeamList 查询团队
func (s *sTeam) QueryTeamList(ctx context.Context, search *sys_model.SearchParams) (*co_model.TeamListRes, error) {
	// 跨主体查询条件过滤
	search = funs.FilterUnionMain(ctx, search, co_dao.CompanyTeam(s.modules).Columns().UnionMainId)

	result, err := daoctl.Query[*co_entity.CompanyTeam](co_dao.CompanyTeam(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler), search, false)

	return (*co_model.TeamListRes)(result), err
}

// CreateTeam 创建团队或小组|信息
func (s *sTeam) CreateTeam(ctx context.Context, info *co_model.Team) (*co_entity.CompanyTeam, error) {
	if info.ParentId > 0 {
		team, _ := s.GetTeamById(ctx, info.ParentId)
		if team == nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_ParentTeamNotFound"), co_dao.CompanyTeam(s.modules).Table())
		}
		if team.ParentId > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Group_ParentMustIsTeam"), co_dao.CompanyTeam(s.modules).Table())
		}
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断团队名称是否存在
	if s.HasTeamByName(ctx, info.Name, sessionUser.UnionMainId) == true {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_TeamNameExist"), co_dao.CompanyTeam(s.modules).Table())
	}

	// 判断团队管理人信息是否存在
	if info.OwnerEmployeeId > 0 {
		_, err := s.modules.Employee().GetEmployeeById(ctx, info.OwnerEmployeeId)
		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "TeamOwnerEmployee")+s.modules.T(ctx, "error_Data_NotFound"), co_dao.CompanyTeam(s.modules).Table())
		}
	}

	if info.CaptainEmployeeId > 0 {
		employee, err := s.modules.Employee().GetEmployeeById(ctx, info.CaptainEmployeeId)
		if err != nil || employee.UnionMainId != sessionUser.UnionMainId {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "TeamCaptainEmployee")+s.modules.T(ctx, "error_Data_NotFound"), co_dao.CompanyTeam(s.modules).Table())
		}

		data, err := s.QueryTeamListByEmployee(ctx, employee.Id, employee.UnionMainId)
		if err != nil && err != sql.ErrNoRows {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "TeamCaptainEmployee")+s.modules.T(ctx, "error_Data_NotFound"), co_dao.CompanyTeam(s.modules).Table())
		}

		if info.ParentId == 0 {
			for _, team := range *data.List {
				if team.ParentId == 0 {
					return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "TeamCaptainEmployee")+"不能是其它团队的队员", co_dao.CompanyTeam(s.modules).Table())
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

	err := co_dao.CompanyTeam(s.modules).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 创建团队
		_, err := co_dao.CompanyTeam(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(data).Insert()
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_Save_Failed"), co_dao.CompanyTeam(s.modules).Table())
		}
		// 创建团队队长
		_, err = co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(captain).Insert()
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_Save_Failed")+"无法保存"+s.modules.T(ctx, "TeamCaptainEmployee")+"信息", co_dao.CompanyTeam(s.modules).Table())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return s.GetTeamById(ctx, data.Id.(int64))
}

// UpdateTeam 更新团队或小组|信息
func (s *sTeam) UpdateTeam(ctx context.Context, id int64, name string, remark string) (*co_entity.CompanyTeam, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if s.HasTeamByName(ctx, name, sessionUser.UnionMainId) == true {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_TeamNameExist"), co_dao.CompanyTeam(s.modules).Table())
	}

	data := co_do.CompanyTeam{
		Name:      name,
		Remark:    remark,
		UpdatedAt: gtime.Now(),
	}
	result, _ := co_dao.CompanyTeam(s.modules).Ctx(ctx).
		Hook(daoctl.CacheHookHandler).Data(data).
		Where(co_do.CompanyTeam{Id: id}).
		Update()

	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_Save_Failed"), co_dao.CompanyTeam(s.modules).Table())
	}

	return s.GetTeamById(ctx, id)
}

// GetTeamMemberList 获取团队成员|列表
func (s *sTeam) GetTeamMemberList(ctx context.Context, id int64) (*co_model.EmployeeListRes, error) {
	team, err := s.GetTeamById(ctx, id)
	if err != nil {
		return nil, err
	}

	// 团队成员信息
	var items []co_entity.CompanyTeamMember
	err = co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.CompanyTeamMember{
		TeamId:      team.Id,
		UnionMainId: team.UnionMainId,
	}).Scan(&items)

	ids := make([]int64, 0)
	for _, item := range items {
		ids = append(ids, item.EmployeeId)
	}

	return s.modules.Employee().QueryEmployeeList(ctx, &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0),
			sys_model.FilterInfo{
				Field: co_dao.CompanyEmployee(s.modules).Columns().Id,
				Where: "in",
				Value: ids,
			},
			sys_model.FilterInfo{
				Field: co_dao.CompanyEmployee(s.modules).Columns().UnionMainId,
				Where: "=",
				Value: team.UnionMainId,
			},
		),
	})
}

// QueryTeamListByEmployee 根据员工查询团队
func (s *sTeam) QueryTeamListByEmployee(ctx context.Context, employeeId int64, unionMainId int64) (*co_model.TeamListRes, error) {
	data := &[]*co_entity.CompanyTeamMember{}

	if unionMainId == 0 {
		unionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId
	}

	err := co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
		Where(co_do.CompanyTeamMember{EmployeeId: employeeId, UnionMainId: unionMainId}).
		Scan(data)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_NotFound"), co_dao.CompanyTeam(s.modules).Table())
	}

	var teamIds []int64
	for _, member := range *data {
		teamIds = append(teamIds, member.TeamId)
	}

	return s.QueryTeamList(ctx, &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0),
			sys_model.FilterInfo{
				Field: co_dao.CompanyTeam(s.modules).Columns().UnionMainId,
				Where: "=",
				Value: unionMainId,
			},
			sys_model.FilterInfo{
				Field: co_dao.CompanyTeam(s.modules).Columns().Id,
				Where: "in",
				Value: teamIds,
			},
		),
	})
}

// SetTeamMember 设置团队队员或小组组员
func (s *sTeam) SetTeamMember(ctx context.Context, teamId int64, employeeIds []int64) (api_v1.BoolRes, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	var teamMemberArr []*co_entity.CompanyTeamMember
	// 获取团队所有旧成员
	err := co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
		Where(co_do.CompanyTeamMember{
			TeamId:      teamId,
			UnionMainId: sessionUser.UnionMainId,
		}).Scan(&teamMemberArr)

	// 待移除的团队成员
	waitIds := make([]int64, 0)
	// 已存在的团队成员
	existIds := make([]int64, 0)

	// 遍历所有旧成员
	for _, member := range teamMemberArr {
		if len(employeeIds) == 0 {
			existIds = append(existIds, member.EmployeeId)
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
		_, err = co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
			Where(
				co_do.CompanyTeamMember{
					TeamId:      teamId,
					UnionMainId: sessionUser.UnionMainId,
				},
			).Delete()
		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_DeleteMember_Failed"), co_dao.CompanyTeam(s.modules).Table())
		}
		return true, nil
	}

	// 校验新团队成员是否存在
	res, err := s.modules.Employee().QueryEmployeeList(ctx, &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0),
			sys_model.FilterInfo{
				Field: co_dao.CompanyEmployee(s.modules).Columns().Id,
				Where: "in",
				Value: newTeamMemberIds,
			},
			sys_model.FilterInfo{
				Field: co_dao.CompanyEmployee(s.modules).Columns().UnionMainId,
				Where: "=",
				Value: sessionUser.UnionMainId,
			},
		),
		Pagination: sys_model.Pagination{
			Page:     1,
			PageSize: 1000,
		},
	})

	if res.Total < gconv.Int64(len(newTeamMemberIds)) {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_NewTeam_NotFoundMembers"), co_dao.CompanyTeam(s.modules).Table())
	}

	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	//
	if team.ParentId == 0 {
		count, _ := co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
			WhereIn(co_dao.CompanyTeamMember(s.modules).Columns().EmployeeId, newTeamMemberIds).
			Where(co_dao.CompanyTeamMember(s.modules).Columns().UnionMainId, sessionUser.UnionMainId).Count()
		if count > 0 {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_MemberIsHasTeam"), co_dao.CompanyTeam(s.modules).Table())
		}
	}

	err = co_dao.CompanyTeamMember(s.modules).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 清理团队成员
		_, err = co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
			WhereIn(co_dao.CompanyTeamMember(s.modules).Columns().Id, existIds).
			Delete()

		if err != nil {
			return err
		}

		for _, employeeId := range newTeamMemberIds {
			_, err = co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(
				co_do.CompanyTeamMember{
					Id:          idgen.NextId(),
					TeamId:      team.Id,
					EmployeeId:  employeeId,
					UnionMainId: sessionUser.UnionMainId,
					JoinAt:      gtime.Now(),
				},
			)
			if err != nil {
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

	employee, err := s.modules.Employee().GetEmployeeById(ctx, employeeId)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 校验数据主体是否一致
	if sessionUser.UnionMainId != team.UnionMainId || sessionUser.UnionMainId != employee.UnionMainId {
		if team.ParentId <= 0 {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamOrEmployee_Check_Failed"), co_dao.CompanyTeam(s.modules).Table())
		} else {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_GroupOrEmployee_Check_Failed"), co_dao.CompanyTeam(s.modules).Table())
		}
	}

	result, _ := co_dao.CompanyTeam(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
		Data(co_do.CompanyTeam{OwnerEmployeeId: employee.Id}).
		Where(co_do.CompanyTeam{Id: team.Id}).
		Update()

	rowsAffected, err := result.RowsAffected()

	return rowsAffected == 1, err
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

	employee, err := s.modules.Employee().GetEmployeeById(ctx, employeeId)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 校验数据主体是否一致
	if sessionUser.UnionMainId != team.UnionMainId || sessionUser.UnionMainId != employee.UnionMainId {
		if team.ParentId <= 0 {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamOrEmployee_Check_Failed"), co_dao.CompanyTeam(s.modules).Table())
		} else {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_GroupOrEmployee_Check_Failed"), co_dao.CompanyTeam(s.modules).Table())
		}
	}

	// 员工能否设置为队长
	canCaptain := false
	{
		// 查询员工所在的所有团队信息
		data, err := s.QueryTeamListByEmployee(ctx, employee.Id, employee.UnionMainId)
		if err != nil && err != sql.ErrNoRows {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "TeamCaptainEmployee")+s.modules.T(ctx, "error_Data_NotFound"), co_dao.CompanyTeam(s.modules).Table())
		}

		for _, item := range *data.List {
			// 判断要设置的是团队还是小组 ParentId == 0团队，ParentId > 0小组
			if team.ParentId == 0 && item.ParentId == 0 {
				// 如果员工是其它团队成员则返回
				if item.Id != team.Id {
					return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_MemberIsHasTeam"), co_dao.CompanyTeam(s.modules).Table())
				} else {
					canCaptain = true
				}
			}
		}
	}

	if team.ParentId == 0 && !canCaptain {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamCaptainEmployee_MustInTeam"), co_dao.CompanyTeam(s.modules).Table())
	}

	result, _ := co_dao.CompanyTeam(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
		Data(co_do.CompanyTeam{CaptainEmployeeId: employee.Id}).
		Where(co_do.CompanyTeam{Id: team.Id}).
		Update()

	rowsAffected, err := result.RowsAffected()

	return rowsAffected == 1, err
}

// DeleteTeam 删除团队
func (s *sTeam) DeleteTeam(ctx context.Context, teamId int64) (api_v1.BoolRes, error) {
	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 查询团队成员数量
	count, err := co_dao.CompanyTeamMember(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
		Where(co_do.CompanyTeamMember{
			TeamId:      team.Id,
			UnionMainId: sessionUser.UnionMainId,
		}).Count()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "TeamMember")+s.modules.T(ctx, "error_Data_Get_Failed"), co_dao.CompanyTeam(s.modules).Table())
	}

	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_NeedRemoveTeamMember"), co_dao.CompanyTeam(s.modules).Table())
	}

	result, _ := co_dao.CompanyTeam(s.modules).Ctx(ctx).Unscoped().Hook(daoctl.CacheHookHandler).Where(co_do.CompanyTeam{Id: team.Id}).Delete()

	rowsAffected, err := result.RowsAffected()

	return rowsAffected == 1, err
}
