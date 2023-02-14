package company

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"math"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-community/utility/masker"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type sEmployee struct {
	modules co_interface.IModules
	dao     *co_dao.XDao
}

func NewEmployee(modules co_interface.IModules) *sEmployee {
	result := &sEmployee{
		modules: modules,
		dao:     modules.Dao(),
	}

	// 注入钩子函数
	result.injectHook()
	return result
}

// InjectHook 注入Audit的Hook
func (s *sEmployee) injectHook() {
	sys_service.Jwt().InstallHook(s.modules.GetConfig().UserType, s.jwtHookFunc)
	sys_service.SysAuth().InstallHook(sys_enum.Auth.ActionType.Login, s.modules.GetConfig().UserType, s.authHookFunc)
	sys_service.SysUser().InstallHook(sys_enum.User.Event.BeforeCreate, s.userHookFunc)
}

// AuthHookFunc 用户登录Hook函数
func (s *sEmployee) authHookFunc(ctx context.Context, _ sys_enum.AuthActionType, user *sys_model.SysUser) error {
	employee, _ := s.modules.Employee().GetEmployeeById(ctx, user.Id)
	if employee != nil {
		user.Detail.Realname = employee.Name
	}

	company, _ := s.modules.Company().GetCompanyById(ctx, employee.UnionMainId)
	if company != nil {
		user.Detail.UnionMainName = company.Name
	}

	return nil
}

// userHookFunc 新增用户Hook函数
func (s *sEmployee) userHookFunc(ctx context.Context, _ sys_enum.UserEvent, info sys_model.SysUser) (sys_model.SysUser, error) {
	employee, err := s.GetEmployeeById(ctx, info.Id)
	if err != nil {
		return info, nil
	}
	info.Detail.Realname = employee.Name

	company, err := s.modules.Company().GetCompanyById(ctx, employee.UnionMainId)
	if err != nil {
		return info, nil
	}
	info.Detail.UnionMainName = company.Name

	return info, nil
}

// JwtHookFunc Jwt钩子函数
func (s *sEmployee) jwtHookFunc(ctx context.Context, claims *sys_model.JwtCustomClaims) (*sys_model.JwtCustomClaims, error) {
	// 获取到当前user的主体id
	employee, err := s.GetEmployeeById(ctx, claims.Id)
	if employee == nil {
		return claims, err
	}

	// 这里还没登录成功不能使用 s.modules.Company().GetCompanyById，因为里面包含获取当前登录用户的 session 存在矛盾
	company, err := daoctl.GetByIdWithError[co_entity.Company](
		s.dao.Company.Ctx(ctx),
		employee.UnionMainId,
	)
	if company == nil || err != nil {
		return claims, sys_service.SysLogs().ErrorSimple(ctx, err, "主体id获取失败", s.dao.Company.Table())
	}

	claims.IsAdmin = claims.Type == -1 || claims.Id == company.UserId
	claims.UnionMainId = company.Id
	claims.ParentId = company.ParentId

	return claims, nil
}

// GetEmployeeById 根据ID获取员工信息
func (s *sEmployee) GetEmployeeById(ctx context.Context, id int64) (*co_model.EmployeeRes, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	data, err := daoctl.GetByIdWithError[co_model.EmployeeRes](
		s.dao.Employee.Ctx(ctx),
		id,
	)

	if err != nil {
		message := s.modules.T(ctx, "{#EmployeeName}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#EmployeeName}{#Data}")
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Employee.Table())
	}

	// 跨主体禁止查看员工信息，下级公司可查看上级公司员工信息
	if err == sql.ErrNoRows || data != nil && data.UnionMainId != sessionUser.UnionMainId && data.UnionMainId != sessionUser.ParentId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName} {#error_Data_NotFound}"), s.dao.Employee.Table())
	}

	return s.masker(s.makeMore(ctx, data)), nil
}

// GetEmployeeByName 根据Name获取员工信息
func (s *sEmployee) GetEmployeeByName(ctx context.Context, name string) (*co_model.EmployeeRes, error) {
	data, err := daoctl.ScanWithError[co_model.EmployeeRes](
		s.dao.Employee.Ctx(ctx).Where(co_do.CompanyEmployee{Name: name}),
	)

	if err != nil {
		message := s.modules.T(ctx, "{#EmployeeName}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#EmployeeName}{#Data}")
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Employee.Table())
	}

	return s.masker(s.makeMore(ctx, data)), nil
}

// HasEmployeeByName 员工名称是否存在
func (s *sEmployee) HasEmployeeByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool {
	if unionMainId <= 0 {
		unionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId
	}

	model := s.dao.Employee.Ctx(ctx).Where(co_do.CompanyEmployee{
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
			model = model.WhereNotIn(s.dao.Employee.Columns().Id, ids)
		}
	}

	count, _ := model.Count()
	return count > 0
}

// HasEmployeeByNo 员工工号是否存在
func (s *sEmployee) HasEmployeeByNo(ctx context.Context, no string, unionMainId int64, excludeIds ...int64) bool { // 如果工号为空则直接返回
	// 工号为空，且允许工号为空则不做校验
	if no == "" && s.modules.GetConfig().AllowEmptyNo == true {
		return false
	}

	if unionMainId <= 0 {
		unionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId
	}

	model := s.dao.Employee.Ctx(ctx).Where(co_do.CompanyEmployee{
		No:          no,
		UnionMainId: unionMainId,
	})

	if len(excludeIds) > 0 && excludeIds[0] > 0 {
		var ids []int64
		for _, id := range excludeIds {
			if id > 0 {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			model = model.WhereNotIn(s.dao.Employee.Columns().Id, ids)
		}
	}

	count, _ := model.Count()
	return count > 0
}

// GetEmployeeBySession 获取当前登录的员工信息
func (s *sEmployee) GetEmployeeBySession(ctx context.Context) (*co_model.EmployeeRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if user.Type != s.modules.GetConfig().UserType.Code() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_NotHasServerPermission"), s.dao.Employee.Table())
	}

	result, _ := s.GetEmployeeById(ctx, user.Id)
	if result == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_CheckLoginUser_Failed"), s.dao.Employee.Table())
	}
	return result, nil
}

// QueryEmployeeList 获取员工列表
func (s *sEmployee) QueryEmployeeList(ctx context.Context, search *sys_model.SearchParams) (*co_model.EmployeeListRes, error) { // 跨主体查询条件过滤
	// 过滤UnionMainId字段查询条件
	search = s.modules.Company().FilterUnionMainId(ctx, search)

	model := s.dao.Employee.Ctx(ctx)

	excludeIds := make([]int64, 0)

	// 处理扩展条件，扩展支持 teamId，employeeId, inviteUserId, unionMainId 字段过滤支持
	{
		teamSearch := funs.SearchFilterEx(*search, "teamId", "employeeId", "inviteUserId", "unionMainId")

		if len(teamSearch.Filter) > 0 {
			items, _ := s.modules.Team().QueryTeamMemberList(ctx, teamSearch)

			if len(items.Records) > 0 {
				for _, item := range items.Records {
					excludeIds = append(excludeIds, item.EmployeeId)
				}
			}

			if len(excludeIds) > 0 {
				// 过滤掉重复的id
				excludeIds = gconv.Int64s(garray.NewSortedStrArrayFrom(gconv.Strings(excludeIds)).Unique().Slice())
				model = model.WhereIn(s.dao.Employee.Columns().Id, excludeIds)
			}
		}
	}

	// 查询符合过滤条件的员工信息
	result, err := daoctl.Query[*co_model.EmployeeRes](model.
		With(co_model.EmployeeRes{}.Detail, co_model.EmployeeRes{}.User), search, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName}{#error_Data_Get_Failed}"), s.dao.Employee.Table())
	}

	items := make([]*co_model.EmployeeRes, 0)
	for _, record := range result.Records {
		items = append(items, s.masker(s.makeMore(ctx, record)))
	}
	result.Records = items

	return (*co_model.EmployeeListRes)(result), nil
}

// CreateEmployee 创建员工信息
func (s *sEmployee) CreateEmployee(ctx context.Context, info *co_model.Employee) (*co_model.EmployeeRes, error) {
	info.Id = 0

	return s.saveEmployee(ctx, info)
}

// UpdateEmployee 更新员工信息
func (s *sEmployee) UpdateEmployee(ctx context.Context, info *co_model.Employee) (*co_model.EmployeeRes, error) {
	return s.saveEmployee(ctx, info)
}

// saveEmployee 保存员工信息
func (s *sEmployee) saveEmployee(ctx context.Context, info *co_model.Employee) (*co_model.EmployeeRes, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 除匿名用户外，其它用户在有权限的情况下均可以创建或更新员工信息，001 代表默认管理员工号
	// info.Id == 0 仅单纯新建员工时需要初始化用户归属为当前操作员所在 UnionMainId
	if sessionUser.Type > 0 && info.Id == 0 && info.No != "001" {
		info.UnionMainId = sessionUser.UnionMainId
	}

	// 校验员工名称是否已存在
	if true == s.HasEmployeeByName(ctx, info.Name, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NameAlreadyExists}"), s.dao.Employee.Table())
	}

	// 校验工号是否允许为空
	if info.No == "" && s.modules.GetConfig().AllowEmptyNo == false {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NoNotNull}"), s.dao.Employee.Table())
	}

	// 校验工号是否已存在
	if true == s.HasEmployeeByNo(ctx, info.No, info.UnionMainId, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NoAlreadyExists}"), s.dao.Employee.Table())
	}

	data := &co_do.CompanyEmployee{}
	gconv.Struct(info, data)

	err := s.dao.Employee.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var avatarFile *sys_model.FileInfo
		if info.Avatar != "" {
			// 校验员工头像并保存
			fileInfo, err := sys_service.File().GetFileById(ctx, gconv.Int64(info.Avatar), s.modules.T(ctx, "{#Avatar}{#error_File_FileVoid}"))

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "", s.dao.Employee.Table())
			}
			avatarFile = fileInfo
			avatarFile.Src = s.modules.GetConfig().StoragePath + "/employee/" + gconv.String(data.Id) + "/avatar." + avatarFile.Ext

			info.Avatar = gconv.String(fileInfo.Id)
		}

		if info.Id == 0 {
			// 创建员工信息
			data.Id = idgen.NextId()
			data.CreatedBy = sessionUser.Id
			data.CreatedAt = gtime.Now()
			data.UnionMainId = info.UnionMainId

			{
				// 创建登录信息
				passwordLen := len(gconv.String(data.Id))
				password := gstr.SubStr(gconv.String(data.Id), passwordLen-6, 6)

				newUser, err := sys_service.SysUser().CreateUser(ctx, sys_model.UserInnerRegister{
					Username:        strconv.FormatInt(gconv.Int64(data.Id), 36),
					Password:        password,
					ConfirmPassword: password,
					Mobile:          gconv.String(data.Mobile),
				},
					sys_enum.User.State.Normal,
					s.modules.GetConfig().UserType,
					gconv.Int64(data.Id),
				)
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_User_Save_Failed"), s.dao.Employee.Table())
				}

				data.Id = newUser.Id
			}

			affected, err := daoctl.InsertWithError(s.dao.Employee.Ctx(ctx).Data(data))

			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Save_Failed"), s.dao.Employee.Table())
			}
		} else {
			// 更新员工信息
			data.UpdatedBy = sessionUser.Id
			data.UpdatedAt = gtime.Now()
			// unionMainId不能修改，强制为nil
			data.UnionMainId = nil
			data.Mobile = nil

			_, err := daoctl.UpdateWithError(s.dao.Employee.Ctx(ctx).Data(data).Where(co_do.CompanyEmployee{Id: data.Id}))
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Save_Failed"), s.dao.Employee.Table())
			}
		}

		// 保存文件
		if avatarFile != nil {
			avatarFile, err := sys_service.File().SaveFile(ctx, avatarFile.Src, avatarFile)
			_, err = sys_dao.SysFile.Ctx(ctx).Insert(avatarFile)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Avatar}{#error_File_Save_Failed}"), s.dao.Employee.Table())
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return s.GetEmployeeById(ctx, gconv.Int64(data.Id))
}

// DeleteEmployee 删除员工信息
func (s *sEmployee) DeleteEmployee(ctx context.Context, id int64) (bool, error) {
	employee, err := s.GetEmployeeById(ctx, id)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName}{#error_Data_Get_Failed}"), s.dao.Employee.Table())
	}

	err = s.dao.Employee.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if s.modules.GetConfig().HardDeleteWaitAt > 0 && employee.DeletedAt == nil {
			// 设置账户状态为已注销
			_, err = sys_service.SysUser().SetUserState(ctx, employee.Id, sys_enum.User.State.Canceled)
			if err != nil {
				return err
			}
			// 设置员工状态为已注销
			_, err = s.dao.Employee.Ctx(ctx).
				Data(co_do.CompanyEmployee{State: co_enum.Employee.State.Canceled.Code()}).
				Where(co_do.CompanyEmployee{Id: employee.Id}).
				Update()
			if err != nil {
				return err
			}
			// 软删除
			_, err = s.dao.Employee.Ctx(ctx).Delete(co_do.CompanyEmployee{Id: employee.Id})
			if err != nil {
				return err
			}
		} else {
			if employee.DeletedAt != nil {
				HardDeleteWaitAt := time.Hour * (time.Duration)(s.modules.GetConfig().HardDeleteWaitAt)

				if gtime.Now().Before(employee.DeletedAt.Add(HardDeleteWaitAt)) {
					hours := gtime.Now().Sub(employee.DeletedAt.Add(HardDeleteWaitAt)).Hours()
					message := s.modules.T(ctx, "error_Employee_Delete_Failed") + "数据延期保护中\r请于 " + gconv.String(math.Abs(hours)) + " 小时后操作"
					return sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Employee.Table())
				}
			}

			// 员工移出团队|小组
			_, err := s.setEmployeeTeam(ctx, employee.Id)
			if err != nil {
				return err
			}

			// 删除员工
			_, err = s.dao.Employee.Ctx(ctx).Unscoped().Delete(co_do.CompanyEmployee{Id: employee.Id})
			if err != nil {
				return err
			}
			// 删除用户
			_, err = sys_service.SysUser().DeleteUser(ctx, employee.Id)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Delete_Failed"), s.dao.Employee.Table())
	}
	return true, nil
}

// setEmployeeTeam 员工移出小组 | 团队
func (s *sEmployee) setEmployeeTeam(ctx context.Context, employeeId int64) (bool, error) {
	// 直接删除属于员工的团队成员记录
	isSuccess, err := s.modules.Team().DeleteTeamMemberByEmployee(ctx, employeeId)
	if err != nil && isSuccess == false {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_DeleteMember_Failed"), s.dao.Employee.Table())
	}

	// 查找到员工是管理员或者队长的团队
	teamList, err := s.modules.Team().QueryTeamList(ctx, &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0), sys_model.FilterInfo{
			Field:     s.dao.Team.Columns().CaptainEmployeeId,
			Where:     "=",
			Value:     employeeId,
			IsOrWhere: true,
		}, sys_model.FilterInfo{
			Field:     s.dao.Team.Columns().OwnerEmployeeId,
			Where:     "=",
			Value:     employeeId,
			IsOrWhere: true,
		}),
	})

	// 假如是队长或者组长，需要将团队表的队长或者组长设置为0
	if len(teamList.Records) > 0 {
		for _, item := range teamList.Records {
			if item.CaptainEmployeeId == employeeId { // 队长或者组长
				ret, err := s.modules.Team().SetTeamCaptain(ctx, item.Id, 0)
				if err != nil || ret == false {
					return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Delete_Failed"), s.dao.Employee.Table())
				}
			}

			if item.OwnerEmployeeId == employeeId { // 团队负责人
				ret, err := s.modules.Team().SetTeamOwner(ctx, item.Id, 0)
				if err != nil || ret == false {
					return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Delete_Failed"), s.dao.Employee.Table())
				}
			}
		}
	}
	return true, nil
}

// GetEmployeeDetailById 根据ID获取员工详细信息
func (s *sEmployee) GetEmployeeDetailById(ctx context.Context, id int64) (*co_model.EmployeeRes, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	model := s.dao.Employee.Ctx(ctx)

	if sessionUser.IsAdmin == false {
		// 判断用户是否有权限
		can, _ := sys_service.SysPermission().CheckPermission(ctx, co_enum.Employee.PermissionType(s.modules).MoreDetail)
		if can == false {
			model = model.Where(sys_do.SysFile{UnionMainId: sessionUser.UnionMainId})
		}
	}

	data, err := daoctl.ScanWithError[co_model.EmployeeRes](model.Where(co_do.CompanyEmployee{Id: id}))

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetEmployeeDetailById_Failed"), s.dao.Employee.Table())
	}

	// 跨主体禁止查看员工信息，
	if err == sql.ErrNoRows || data != nil && data.UnionMainId != sessionUser.UnionMainId {
		// 下级公司也不可查看上级公司员工详细信息
		if data.UnionMainId == sessionUser.ParentId {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_NotHasServerPermission"), s.dao.Employee.Table())
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName} {#error_Data_NotFound}"), s.dao.Employee.Table())
	}

	return s.makeMore(ctx, data), err
}

// GetEmployeeListByRoleId 根据角色ID获取所有所属员工
func (s *sEmployee) GetEmployeeListByRoleId(ctx context.Context, roleId int64) (*co_model.EmployeeListRes, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	userIds, err := sys_service.SysRole().GetRoleMemberIds(ctx, roleId, sessionUser.UnionMainId)
	if err != nil {
		return &co_model.EmployeeListRes{
			PaginationRes: sys_model.PaginationRes{
				Pagination: sys_model.Pagination{
					PageNum:  1,
					PageSize: 20,
				},
				PageTotal: 0,
				Total:     0,
			},
		}, nil
	}

	result, err := daoctl.Query[*co_model.EmployeeRes](
		s.dao.Employee.Ctx(ctx),
		&sys_model.SearchParams{
			Filter: append(make([]sys_model.FilterInfo, 0), sys_model.FilterInfo{
				Field: s.dao.Employee.Columns().Id,
				Where: "in",
				Value: userIds,
			}),
			OrderBy:    nil,
			Pagination: sys_model.Pagination{},
		},
		true,
	)

	items := make([]*co_model.EmployeeRes, 0)
	for _, record := range result.Records {
		items = append(items, s.masker(s.makeMore(ctx, record)))
	}
	result.Records = items

	return (*co_model.EmployeeListRes)(result), err
}

// Masker 员工信息脱敏，并加载附加数据
func (s *sEmployee) masker(employee *co_model.EmployeeRes) *co_model.EmployeeRes {
	if employee == nil {
		return nil
	}
	employee.Mobile = masker.MaskString(employee.Mobile, masker.MaskPhone)
	employee.LastActiveIp = masker.MaskString(employee.LastActiveIp, masker.MaskIPv4)
	employee.Detail.LastLoginIp = masker.MaskString(employee.Detail.LastLoginIp, masker.MaskIPv4)
	return employee
}

// makeMore 按需加载附加数据
func (s *sEmployee) makeMore(ctx context.Context, data *co_model.EmployeeRes) *co_model.EmployeeRes {
	if data == nil {
		return nil
	}

	// team附加数据
	if data.UnionMainId > 0 {
		funs.AttrMake[co_model.EmployeeRes](ctx,
			s.dao.Employee.Columns().UnionMainId,
			func() []co_model.Team {
				g.Try(ctx, func(ctx context.Context) {
					// 获取到该员工的所有团队成员信息记录
					teamMemberItems := make([]*co_entity.CompanyTeamMember, 0)
					s.dao.TeamMember.Ctx(ctx).
						Where(co_do.CompanyTeamMember{EmployeeId: data.CompanyEmployee.Id}).Scan(&teamMemberItems)
					if len(teamMemberItems) == 0 {
						data.TeamList = nil
						return
					}

					// 记录该员工所在所有团队
					teamIds := make([]int64, 0)
					for _, memberItem := range teamMemberItems {
						teamIds = append(teamIds, memberItem.TeamId)
					}
					s.dao.Team.Ctx(ctx).
						WhereIn(s.dao.Team.Columns().Id, teamIds).Scan(&data.TeamList)
				})
				return data.TeamList
			},
		)
	}

	// user相关附加数据
	if data.CompanyEmployee.Id > 0 {
		funs.AttrMake[co_model.EmployeeRes](ctx,
			s.dao.Employee.Columns().Id,
			func() *co_model.EmployeeRes {
				if data.CompanyEmployee.Id == 0 {
					return nil
				}

				user, _ := sys_service.SysUser().GetSysUserById(ctx, data.CompanyEmployee.Id)
				if user != nil {
					gconv.Struct(user.SysUser, &data.User)
					gconv.Struct(user.Detail, &data.Detail)
				}

				return data
			},
		)
	}

	return data
}
