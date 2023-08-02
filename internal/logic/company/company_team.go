package company

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/sys_rules"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_hook"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/yitter/idgenerator-go/idgen"
	"reflect"
)

type sTeam[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	TR co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	base_hook.ResponseFactoryHook[TR]

	// 邀约&加入团队Hook
	InviteJoinTeamHook base_hook.BaseHook[sys_enum.InviteType, co_hook.InviteJoinTeamHookFunc]

	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		TR,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
	dao co_dao.XDao
}

//
//func NewTeam[
//	ITCompanyRes co_model.ICompanyRes,
//	ITEmployeeRes co_model.IEmployeeRes,
//	TR co_model.ITeamRes,
//	ITFdAccountRes co_model.IFdAccountRes,
//	ITFdAccountBillRes co_model.IFdAccountBillRes,
//	ITFdBankCardRes co_model.IFdBankCardRes,
//	ITFdCurrencyRes co_model.IFdCurrencyRes,
//	ITFdInvoiceRes co_model.IFdInvoiceRes,
//	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
//](modules co_interface.IModules[
//	ITCompanyRes,
//	ITEmployeeRes,
//	TR,
//	ITFdAccountRes,
//	ITFdAccountBillRes,
//	ITFdBankCardRes,
//	ITFdCurrencyRes,
//	ITFdInvoiceRes,
//	ITFdInvoiceDetailRes,
//]) co_interface.ITeam[TR] {
//	result := &sTeam[
//		ITCompanyRes,
//		ITEmployeeRes,
//		TR,
//		ITFdAccountRes,
//		ITFdAccountBillRes,
//		ITFdBankCardRes,
//		ITFdCurrencyRes,
//		ITFdInvoiceRes,
//		ITFdInvoiceDetailRes,
//	]{
//		modules: modules,
//		dao:     *modules.Dao(),
//	}
//
//	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)
//
//	return result
//}

func NewTeam[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	TR co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) co_interface.ITeam[TR] {
	result := &sTeam[
		ITCompanyRes,
		ITEmployeeRes,
		TR,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		modules: modules,
		dao:     *modules.Dao(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetXDao(dao co_dao.XDao) {
	s.dao = dao
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.ITeamRes
	ret = &co_model.TeamRes{
		CompanyTeam: co_entity.CompanyTeam{},
		Owner:       nil,
		Captain:     nil,
		UnionMain:   nil,
		Parent:      nil,
	}
	return ret.(TR)
}

// GetTeamById 根据ID获取公司团队信息
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetTeamById(ctx context.Context, id int64) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	data, err := daoctl.GetByIdWithError[TR](
		s.dao.Team.Ctx(ctx), id,
	)

	if err != nil {
		message := s.modules.T(ctx, "{#teamOrGroup}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#teamOrGroup}{#error_Data_Get_Failed}")
		}
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Team.Table())
	}

	if !reflect.ValueOf(data).IsNil() {
		response = *data
	}

	// 需要进行跨主体判断
	if err == sql.ErrNoRows || !reflect.ValueOf(data).IsNil() &&
		response.Data().UnionMainId != sessionUser.UnionMainId &&
		!sessionUser.IsAdmin &&
		!sessionUser.IsSuperAdmin {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#TeamName} {#error_Data_NotFound}"), s.dao.Team.Table())
	}

	return s.makeMore(ctx, response), nil
}

// GetTeamByName 根据Name获取团队信息
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetTeamByName(ctx context.Context, name string) (response TR, err error) {
	data, err := daoctl.ScanWithError[TR](
		s.dao.Team.Ctx(ctx).
			Where(co_do.CompanyTeam{Name: name}),
	)

	if err != nil {
		message := s.modules.T(ctx, "{#teamOrGroup}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#teamOrGroup}{#error_Data_Get_Failed}")
		}
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Team.Table())
	}

	return s.makeMore(ctx, *data), nil
}

// HasTeamByName 团队名称是否存在
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) HasTeamByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool {
	if unionMainId == 0 {
		unionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId
	}

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
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryTeamList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error) {
	// 过滤UnionMainId字段查询条件
	search = s.modules.Company().FilterUnionMainId(ctx, search)

	data, err := daoctl.Query[TR](s.dao.Team.Ctx(ctx), search, false)

	items := make([]TR, 0)
	for _, item := range data.Records {
		items = append(items, s.makeMore(ctx, item))
	}
	data.Records = items

	return data, err
}

// QueryTeamMemberList 查询所有团队成员记录
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryTeamMemberList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[*co_model.TeamMemberRes], error) {
	// 过滤UnionMainId字段查询条件
	search = s.modules.Company().FilterUnionMainId(ctx, search)
	model := s.dao.TeamMember.Ctx(ctx)

	data, err := daoctl.Query[*co_model.TeamMemberRes](model, search, false)

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	var UnionMain co_model.ICompanyRes
	if sessionUser.UnionMainId > 0 {
		UnionMain, _ = s.modules.Company().GetCompanyById(ctx, sessionUser.UnionMainId)
	}

	items := make([]*co_model.TeamMemberRes, 0)
	for _, item := range data.Records {
		if item.EmployeeId > 0 {
			v, _ := s.modules.Employee().GetEmployeeById(ctx, item.EmployeeId)
			if !reflect.ValueOf(v).IsNil() {
				item.Employee = v.Data()
			}
		}
		if item.InviteUserId > 0 {
			v, _ := s.modules.Employee().GetEmployeeById(ctx, item.InviteUserId)
			if !reflect.ValueOf(v).IsNil() {
				item.InviteUser = v.Data()
			}
		}
		if item.UnionMainId == sessionUser.UnionMainId {
			item.UnionMain = UnionMain.Data()
		} else if item.UnionMainId > 0 {
			UnionMain, _ = s.modules.Company().GetCompanyById(ctx, item.UnionMainId)
		}
		items = append(items, item)
	}
	data.Records = items

	return data, err
}

// CreateTeam 创建团队或小组|信息
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) CreateTeam(ctx context.Context, info *co_model.Team) (response TR, err error) {
	if info.ParentId > 0 {
		team, _ := s.GetTeamById(ctx, info.ParentId)
		if reflect.ValueOf(team).IsNil() {
			return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_ParentTeamNotFound"), s.dao.Team.Table())
		}
		if team.Data().ParentId > 0 {
			return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Group_ParentMustIsTeam"), s.dao.Team.Table())
		}
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断团队名称是否存在
	if s.HasTeamByName(ctx, info.Name, sessionUser.UnionMainId) == true {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_TeamNameExist"), s.dao.Team.Table())
	}

	// 判断团队管理人信息是否存在
	if info.OwnerEmployeeId > 0 {
		_, err := s.modules.Employee().GetEmployeeById(ctx, info.OwnerEmployeeId)
		if err != nil {
			return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamOwnerEmployee}{#error_Data_NotFound}"), s.dao.Team.Table())
		}
	}

	if info.CaptainEmployeeId > 0 {
		employee, err := s.modules.Employee().GetEmployeeById(ctx, info.CaptainEmployeeId)
		if err != nil || employee.Data().UnionMainId != sessionUser.UnionMainId {
			return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamOwnerEmployee}{#error_Data_NotFound}"), s.dao.Team.Table())
		}

		data, err := s.QueryTeamListByEmployee(ctx, employee.Data().Id, employee.Data().UnionMainId)
		if err != nil && err != sql.ErrNoRows {
			return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamOwnerEmployee}{#error_Data_NotFound}"), s.dao.Team.Table())
		}

		if info.ParentId == 0 {
			for _, team := range data.Records {
				if team.Data().ParentId == 0 {
					return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamCaptainEmployee}{#error_Team_NotInOtherTeam}"), s.dao.Team.Table())
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
	member := co_do.CompanyTeamMember{
		Id:          idgen.NextId(),
		TeamId:      data.Id,
		EmployeeId:  info.CaptainEmployeeId,
		UnionMainId: sessionUser.UnionMainId,
		JoinAt:      gtime.Now(),
	}
	err = s.dao.Team.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 重载Do模型
		doData, err := info.OverrideDo.DoFactory(data)
		if err != nil {
			return err
		}

		// 创建团队
		affected, err := daoctl.InsertWithError(
			s.dao.Team.Ctx(ctx).Data(doData),
		)
		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_Save_Failed"), s.dao.Team.Table())
		}
		if info.CaptainEmployeeId > 0 {

			// 构建待写入数据库的Do数据对象
			captain, err := info.TeamMemberDo.DoFactory(member)

			if err != nil {
				return err
			}

			// 创建团队队长
			_, err = s.dao.TeamMember.Ctx(ctx).Data(captain).Insert()
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#error_Team_Save_Failed}{#error_Team_TeamCaptainEmployee_NotSave}"), s.dao.Team.Table())
			}
		}
		return nil
	})
	if err != nil {
		return response, err
	}

	result, err := s.GetTeamById(ctx, data.Id.(int64))
	return s.makeMore(ctx, result), err
}

// UpdateTeam 更新团队或小组|信息
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateTeam(ctx context.Context, id int64, name string, remark string) (response TR, err error) {

	if s.HasTeamByName(ctx, name, id) == true {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_TeamNameExist"), s.dao.Team.Table())
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
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_Save_Failed"), s.dao.Team.Table())
	}

	result, err := s.GetTeamById(ctx, id)
	return s.makeMore(ctx, result), err
}

// QueryTeamListByEmployee 根据员工查询团队
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryTeamListByEmployee(ctx context.Context, employeeId int64, unionMainId int64) (*base_model.CollectRes[TR], error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

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

	// 跨主体判断
	if err == sql.ErrNoRows || data != nil && unionMainId != sessionUser.UnionMainId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#TeamMember} {#error_Data_NotFound}"), s.dao.TeamMember.Table())
	}

	var teamIds []int64
	for _, member := range *data {
		teamIds = append(teamIds, member.TeamId)
	}

	return s.QueryTeamList(ctx, &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0),
			base_model.FilterInfo{
				Field: s.dao.Team.Columns().UnionMainId,
				Where: "=",
				Value: unionMainId,
			},
			base_model.FilterInfo{
				Field: s.dao.Team.Columns().Id,
				Where: "in",
				Value: teamIds,
			},
		),
	})
}

// SetTeamMember 设置团队队员或小组组员
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetTeamMember(ctx context.Context, teamId int64, employeeIds []int64) (api_v1.BoolRes, error) {
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
	res, err := s.modules.Employee().QueryEmployeeList(ctx, &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0),
			base_model.FilterInfo{
				Field: s.dao.Employee.Columns().Id,
				Where: "in",
				Value: newTeamMemberIds,
			},
			base_model.FilterInfo{
				Field: s.dao.Employee.Columns().UnionMainId,
				Where: "=",
				Value: sessionUser.UnionMainId,
			},
		),
		Pagination: base_model.Pagination{
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
	if team.Data().ParentId == 0 {
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
						TeamId:      team.Data().Id,
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
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetTeamOwner(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error) {
	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	if team.Data().OwnerEmployeeId == employeeId {
		return true, nil
	}

	// 需要删除团队负责人的情况
	if team.Data().Id != 0 && employeeId == 0 {
		affected, err := daoctl.UpdateWithError(s.dao.Team.Ctx(ctx).
			Where(co_do.CompanyTeam{Id: team.Data().Id}).
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
	if sessionUser.UnionMainId != team.Data().UnionMainId || sessionUser.UnionMainId != employee.Data().UnionMainId {
		if team.Data().ParentId <= 0 {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamOrEmployee_Check_Failed"), s.dao.Team.Table())
		} else {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_GroupOrEmployee_Check_Failed"), s.dao.Team.Table())
		}
	}

	affected, err := daoctl.UpdateWithError(
		s.dao.Team.Ctx(ctx).
			Data(co_do.CompanyTeam{OwnerEmployeeId: employee.Data().Id}).
			Where(co_do.CompanyTeam{Id: team.Data().Id}),
	)

	return affected == 1, err
}

// SetTeamCaptain 设置团队队长或小组组长
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetTeamCaptain(ctx context.Context, teamId int64, employeeId int64) (api_v1.BoolRes, error) {
	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	if team.Data().CaptainEmployeeId == employeeId {
		return true, nil
	}

	// 需要删除团队队长或者组长的情况
	if employeeId == 0 && team.Data().Id != 0 {
		affected, err := daoctl.UpdateWithError(
			s.dao.Team.Ctx(ctx).
				Data(co_do.CompanyTeam{CaptainEmployeeId: 0}).
				Where(co_do.CompanyTeam{Id: team.Data().Id}),
		)
		return affected == 1, err
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, employeeId)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 校验数据主体是否一致
	if sessionUser.UnionMainId != team.Data().UnionMainId || sessionUser.UnionMainId != employee.Data().UnionMainId {
		if team.Data().ParentId <= 0 {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamOrEmployee_Check_Failed"), s.dao.Team.Table())
		} else {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_GroupOrEmployee_Check_Failed"), s.dao.Team.Table())
		}
	}

	// 员工能否设置为队长
	canCaptain := false
	{
		// 查询员工所在的所有团队信息
		data, err := s.QueryTeamListByEmployee(ctx, employee.Data().Id, employee.Data().UnionMainId)
		if err != nil && err != sql.ErrNoRows {
			return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamCaptainEmployee}{#error_Data_NotFound}"), s.dao.Team.Table())
		}

		for _, item := range data.Records {
			// 判断要设置的是团队还是小组 ParentId == 0团队，ParentId > 0小组
			if team.Data().ParentId == 0 && item.Data().ParentId == 0 {
				// 如果员工是其它团队成员则返回
				if item.Data().Id != team.Data().Id {
					return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Team_MemberIsHasTeam"), s.dao.Team.Table())
				} else {
					canCaptain = true
				}
			}
		}
	}

	if team.Data().ParentId == 0 && !canCaptain {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_TeamCaptainEmployee_MustInTeam"), s.dao.Team.Table())
	}

	affected, err := daoctl.UpdateWithError(
		s.dao.Team.Ctx(ctx).
			Where(co_do.CompanyTeam{Id: team.Data().Id}).
			Data(co_do.CompanyTeam{CaptainEmployeeId: employee.Data().Id}),
	)

	return affected == 1, err
}

// DeleteTeam 删除团队
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) DeleteTeam(ctx context.Context, teamId int64) (api_v1.BoolRes, error) {
	team, err := s.GetTeamById(ctx, teamId)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 查询团队成员数量
	count, err := s.dao.TeamMember.Ctx(ctx).
		Where(co_do.CompanyTeamMember{
			TeamId:      team.Data().Id,
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
			Where(co_do.CompanyTeam{Id: team.Data().Id}),
	)

	return affected == 1, err
}

// DeleteTeamMemberByEmployee 删除某个员工的所有团队成员记录
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) DeleteTeamMemberByEmployee(ctx context.Context, employeeId int64) (bool, error) {
	affected, err := daoctl.DeleteWithError(s.dao.TeamMember.Ctx(ctx).Where(co_do.CompanyTeamMember{EmployeeId: employeeId}))

	return affected > 0, err
}

// GetEmployeeListByTeamId 获取团队成员|列表
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeListByTeamId(ctx context.Context, teamId int64) (*base_model.CollectRes[co_model.IEmployeeRes], error) {
	team, err := s.modules.Team().GetTeamById(ctx, teamId)
	if err != nil {
		return nil, err
	}

	// 团队成员信息
	items, err := daoctl.ScanWithError[[]*co_entity.CompanyTeamMember](
		s.dao.TeamMember.Ctx(ctx).Where(co_do.CompanyTeamMember{
			TeamId:      team.Data().Id,
			UnionMainId: team.Data().UnionMainId,
		}),
	)

	ids := make([]int64, 0)
	for _, item := range *items {
		ids = append(ids, item.EmployeeId)
	}
	ret, err := s.modules.Employee().QueryEmployeeList(ctx, &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0),
			base_model.FilterInfo{
				Field: s.dao.Employee.Columns().Id,
				Where: "in",
				Value: ids,
			},
			base_model.FilterInfo{
				Field: s.dao.Employee.Columns().UnionMainId,
				Where: "=",
				Value: team.Data().UnionMainId,
			},
		),
	})
	//kconv.Struct(ret, &base_model.CollectRes[co_model.IEmployeeRes]{})

	result := base_model.CollectRes[co_model.IEmployeeRes]{}
	for _, record := range ret.Records {
		i := new(ITEmployeeRes)
		res := kconv.Struct(record, i)
		result.Records = append(result.Records, *res)
	}

	return &result, err
}

// GetTeamInviteCode 获取团队邀约码
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetTeamInviteCode(ctx context.Context, teamId, userId int64) (*co_model.TeamInviteCodeRes, error) {
	// 1.获取团队信息
	team, err := s.modules.Team().GetTeamById(ctx, teamId)
	if err != nil {
		return nil, err
	}

	// 2.生成团队邀约码
	encodeStr, _ := gjson.EncodeString(g.Map{
		"teamId": teamId, // 团队邀约码信息存储团队ID即可
	})
	data := &sys_model.Invite{
		UserId:         userId,
		Value:          encodeStr,
		ExpireAt:       gtime.Now().AddDate(0, 0, sys_consts.Global.InviteCodeExpireDay), // 过期时间
		ActivateNumber: sys_consts.Global.InviteCodeMaxActivateNumber,                    //
		State:          1,                                                                //  默认正常
		Type:           sys_enum.Invite.Type.JoinTeam.Code(),
	}
	if sys_consts.Global.InviteCodeExpireDay == 0 {
		data.ExpireAt = nil
	}

	invite, err := sys_service.SysInvite().CreateInvite(ctx, data)

	// 3.返回响应
	res := co_model.TeamInviteCodeRes{}
	res.InviteRes = invite
	res.Team = team.Data().CompanyTeam

	return &res, nil
}

// JoinTeamByInviteCode 扫码邀约码进入团队
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) JoinTeamByInviteCode(ctx context.Context, inviteCode string, userId int64) (bool, error) {
	// 1.解析邀约码，获取团队信息
	//id := invite_id.CodeToInviteId(inviteCode)
	inviteInfo, err := sys_rules.CheckInviteCode(ctx, inviteCode)
	info := g.Map{
		"teamId": 0,
	}
	gjson.DecodeTo(inviteInfo.Value, &info)
	teamId := gconv.Int64(info["teamId"])

	err = s.dao.Team.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1.获取团队信息
		team, err := s.modules.Team().GetTeamById(ctx, teamId)
		if err != nil {
			return err
		}

		// 2.将扫码人员加入团队
		_, err = s.SetTeamMember(ctx, teamId, []int64{userId})
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_JoinTeamByInviteCode_Failed"), s.dao.TeamMember.Table())
		}

		// 3.需要处理邀约信息的：减少次数、改变状态
		needToSettleInvite := true

		// 广播团队邀约处理Hook
		if inviteCode != "" {
			s.InviteJoinTeamHook.Iterator(func(key sys_enum.InviteType, value co_hook.InviteJoinTeamHookFunc) {
				// 判断订阅的Hook类型是否一致
				if key.Code()&inviteInfo.Type == inviteInfo.Type {
					// 业务类型一致则调用注入的Hook函数
					g.Try(ctx, func(ctx context.Context) {
						needToSettleInvite, err = value(ctx, sys_enum.Invite.Type.Register, inviteInfo, team)
						if err != nil {
							return
						}
					})
				}
			})
		}

		// 业务层没有处理邀约
		if needToSettleInvite {
			if sys_consts.Global.InviteCodeMaxActivateNumber != 0 { // 非无上限
				// 修改邀约次数（里面包含了判断邀约次数从而修改邀约状态的逻辑）
				_, err = sys_service.SysInvite().SetInviteNumber(ctx, inviteInfo.Id, 1, false)
				if err != nil {
					return err
				}
			}
		}

		return nil

	})

	return err == nil, nil
}

// makeMore 按需加载附加数据
func (s *sTeam[
	ITCompanyRes,
	ITEmployeeRes,
	TR,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) makeMore(ctx context.Context, data TR) TR {
	if reflect.ValueOf(data).IsNil() {
		return data
	}

	// 附加数据1：团队负责人Owner
	if data.Data().OwnerEmployeeId > 0 {
		base_funs.AttrMake[TR](ctx,
			s.dao.Team.Columns().OwnerEmployeeId,
			func() ITEmployeeRes {
				var returnRes ITEmployeeRes
				if data.Data().OwnerEmployeeId == 0 {
					return returnRes
				}

				employee, _ := s.modules.Employee().GetEmployeeById(ctx, data.Data().OwnerEmployeeId)
				if !reflect.ValueOf(employee).IsNil() {
					data.Data().Owner = employee.Data()

					// 附加数据填充
					data.Data().SetOwner(employee.Data())
					// 业务层附加数据填充
					data.SetOwner(employee)
				}

				user, _ := sys_service.SysUser().GetSysUserById(ctx, data.Data().OwnerEmployeeId)
				if user != nil && data.Data().Owner != nil {
					gconv.Struct(user.SysUser, &data.Data().Owner.User)
					gconv.Struct(user.Detail, &data.Data().Owner.Detail)
				}

				return employee
			},
		)
	}

	// 附加数据2：团队队长Captain
	if data.Data().CaptainEmployeeId > 0 {
		base_funs.AttrMake[TR](ctx,
			s.dao.Team.Columns().CaptainEmployeeId,
			func() ITEmployeeRes {
				var returnRes ITEmployeeRes
				if data.Data().CaptainEmployeeId == 0 {
					return returnRes
				}

				employee, _ := s.modules.Employee().GetEmployeeById(ctx, data.Data().CaptainEmployeeId)
				if !reflect.ValueOf(employee).IsNil() {
					data.Data().Captain = employee.Data()

					// 附加数据填充
					data.Data().SetCaptain(employee.Data())
					// 业务层附加数据填充
					data.SetCaptain(employee)
				}

				user, _ := sys_service.SysUser().GetSysUserById(ctx, data.Data().CaptainEmployeeId)
				if user != nil && data.Data().Captain != nil {
					gconv.Struct(user.SysUser, &data.Data().Captain.User)
					gconv.Struct(user.Detail, &data.Data().Captain.Detail)
				}

				return employee
			},
		)
	}

	// 附加数据3：团队主体UnionMain
	if data.Data().UnionMainId > 0 {
		base_funs.AttrMake[TR](ctx,
			s.dao.Team.Columns().UnionMainId,
			func() ITCompanyRes {
				var returnRes ITCompanyRes
				if data.Data().UnionMainId == 0 {
					return returnRes
				}

				unionMain, _ := s.modules.Company().GetCompanyById(ctx, data.Data().UnionMainId)
				if !reflect.ValueOf(unionMain).IsNil() {
					data.Data().UnionMain = unionMain.Data()

					data.Data().SetUnionMain(unionMain)
					data.SetUnionMain(unionMain)
				}

				return unionMain
			},
		)
	}

	// 附加数据4：团队或小组父级
	if data.Data().ParentId > 0 {
		base_funs.AttrMake[TR](ctx,
			s.dao.Team.Columns().ParentId,
			func() TR {
				var returnRes TR
				if data.Data().ParentId == 0 {
					return returnRes
				}

				team, _ := s.modules.Team().GetTeamById(ctx, data.Data().ParentId)
				if !reflect.ValueOf(team).IsNil() {
					data.Data().Parent = team.Data()

					data.Data().SetParentTeam(team)
					data.SetParentTeam(team)
				}

				return team
			},
		)
	}
	return data
}
