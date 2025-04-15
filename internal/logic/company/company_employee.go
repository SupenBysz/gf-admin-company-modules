package company

import (
	"context"
	"database/sql"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/base-library/utility/masker"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type sEmployee[
	ITCompanyRes co_model.ICompanyRes,
	TR co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		ITCompanyRes,
		TR,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
	dao     co_dao.XDao
	hookArr *garray.Array
}

func NewEmployee[
	ITCompanyRes co_model.ICompanyRes,
	TR co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) co_interface.IEmployee[TR] {
	result := &sEmployee[
		ITCompanyRes,
		TR,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		modules: modules,
		dao:     *modules.Dao(),
		hookArr: garray.NewArray(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	// 注入钩子函数
	result.injectHook()
	//result.employee = (co_interface.IEmployee[TR])(result)

	//have GetModules() co_interface.IModules[ITCompanyRes, TR, ITTeamRes, ITFdAccountRes, ITFdAccountBillRes, ITFdBankCardRes, ITFdCurrencyRes, ITFdInvoiceRes, ITFdInvoiceDetailRes]
	//want GetModules() co_interface.IModules[co_model.ICompanyRes, TR, co_model.ITeamRes, co_model.IFdAccountRes, co_model.IFdAccountBillsRes, co_model.IFdBankCardRes, co_model.IFdCurrencyRes, co_model.IFdInvoiceRes, co_model.IFdInvoiceDetailRes]

	//*sEmployee[ITCompanyRes, TR, ITTeamRes, ITFdAccountRes, ITFdAccountBillRes, ITFdBankCardRes, ITFdCurrencyRes, ITFdInvoiceRes, ITFdInvoiceDetailRes]            as type co_interface.IEmployee[TR]
	//*sEmployee[ITCompanyRes, TR, ITTeamRes, ITFdAccountRes, ITFdAccountBillRes, ITFdBankCardRes, ITFdCurrencyRes, ITFdInvoiceRes, ITFdInvoiceDetailRes] does not implement co_interface.IEmployee[TR]
	//
	return result
}

//func (s *sEmployee[
//	ITCompanyRes,
//	TR,
//	ITTeamRes,
//	ITFdAccountRes,
//	ITFdAccountBillRes,
//	ITFdBankCardRes,
//	ITFdCurrencyRes,
//	ITFdInvoiceRes,
//	ITFdInvoiceDetailRes,
//]) GetModules() co_interface.IModules[
//	ITCompanyRes,
//	TR,
//	ITTeamRes,
//	ITFdAccountRes,
//	ITFdAccountBillRes,
//	ITFdBankCardRes,
//	ITFdCurrencyRes,
//	ITFdInvoiceRes,
//	ITFdInvoiceDetailRes,
//] {
//	result, _ := s.modules.(co_interface.IModules[
//		ITCompanyRes,
//		TR,
//		ITTeamRes,
//		ITFdAccountRes,
//		ITFdAccountBillRes,
//		ITFdBankCardRes,
//		ITFdCurrencyRes,
//		ITFdInvoiceRes,
//		ITFdInvoiceDetailRes,
//	])
//	return result
//}

func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
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
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.IEmployeeRes
	ret = &co_model.EmployeeRes{
		CompanyEmployee: co_entity.CompanyEmployee{},
		User:            &co_model.EmployeeUser{},
		Detail:          &sys_entity.SysUserDetail{},
		TeamList:        []co_model.Team{},
	}
	return ret.(TR)
}

// InjectHook 注入XXX的Hook
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) injectHook() {
	sys_service.Jwt().InstallHook(s.modules.GetConfig().UserType, s.jwtHookFunc)
	sys_service.SysAuth().InstallHook(sys_enum.Auth.ActionType.Login, s.modules.GetConfig().UserType, s.authHookFunc)
	sys_service.SysUser().InstallHook(sys_enum.User.Event.BeforeCreate, s.userHookFunc)
	sys_service.SysRole().InstallInviteRegisterHook(sys_enum.Role.Change.Remove, s.removeRoleMember)
}

// removeRoleMember 判断是否能移除角色成员逻辑
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) removeRoleMember(ctx context.Context, event sys_enum.RoleMemberChange, role sys_entity.SysRole, sysUser *sys_model.SysUser) (bool, error) {
	if s.dao.Employee == nil {
		return false, nil
	}

	// 不能将系统管理员移除出默认的管理员角色
	if (event.Code() & sys_enum.Role.Change.Remove.Code()) == sys_enum.Role.Change.Remove.Code() {
		if role.IsSystem && role.Name == s.modules.T(ctx, "admin_role_name") {
			employee, err := s.modules.Employee().GetEmployeeById(ctx, sysUser.Id)
			if err != nil {
				return false, err
			}

			company, err := s.modules.Company().GetCompanyById(ctx, employee.Data().UnionMainId)
			if err != nil {
				return false, err
			}

			if company.Data().UserId == sysUser.Id {
				return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_cannot_remove_admin_role"), s.dao.Company.Table())
			}

			return true, nil
		}
	}

	return true, nil
}

// AuthHookFunc 用户登录Hook函数
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) authHookFunc(ctx context.Context, _ sys_enum.AuthActionType, user *sys_model.SysUser) error {
	if s.dao.Employee == nil {
		return nil
	}

	data, err := daoctl.GetByIdWithError[TR](
		s.dao.Employee.Ctx(ctx),
		user.Id,
	)
	if err != nil && err == sql.ErrNoRows {
		return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_user_not_exist"), s.dao.Employee.Table())
	}

	var employee TR
	if data != nil && !reflect.ValueOf(data).IsNil() {
		employee = *data
		if !reflect.ValueOf(employee).IsNil() {
			user.Detail.Realname = employee.Data().Name
		} else {
			return nil
		}
	}

	if !reflect.ValueOf(data).IsNil() && employee.Data().UnionMainId != 0 {
		company, _ := s.modules.Company().GetCompanyById(ctx, employee.Data().UnionMainId)
		if !reflect.ValueOf(company).IsNil() {
			user.Detail.UnionMainName = company.Data().Name
		}
	}

	return nil
}

// userHookFunc 新增用户Hook函数
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) userHookFunc(ctx context.Context, _ sys_enum.UserEvent, info sys_model.SysUser) (sys_model.SysUser, error) {

	data, err := daoctl.GetByIdWithError[TR](
		s.dao.Employee.Ctx(ctx),
		info.Id,
	)
	if err != nil || reflect.ValueOf(*data).IsNil() {
		return info, nil
	}
	var employee TR
	if !reflect.ValueOf(data).IsNil() {
		employee = *data
	}

	info.Detail.Realname = employee.Data().Name

	company, err := s.modules.Company().GetCompanyById(ctx, employee.Data().UnionMainId)
	if err != nil {
		return info, nil
	}
	info.Detail.UnionMainName = company.Data().Name

	return info, nil
}

// JwtHookFunc Jwt钩子函数
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) jwtHookFunc(ctx context.Context, claims *sys_model.JwtCustomClaims) (*sys_model.JwtCustomClaims, error) {
	if s.dao.Employee == nil {
		return claims, nil
	}
	// 获取到当前user的主体id，由于JWT钩子函数大多是登录成功前调用，所以这里不能使用 s.GetEmployeeById 方法调用获取员工数据
	employee, err := daoctl.GetByIdWithError[co_model.EmployeeRes](
		s.dao.Employee.Ctx(ctx),
		claims.Id,
	)
	if employee == nil {
		return claims, err
	}

	// 这里还没登录成功不能使用 s.modules.Company().GetCompanyById，因为里面包含获取当前登录用户的 session 存在矛盾
	company, err := daoctl.GetByIdWithError[co_entity.Company](
		s.dao.Company.Ctx(ctx),
		employee.UnionMainId,
	)
	if company == nil || err != nil {
		return claims, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_company_id_fetch_failed"), s.dao.Company.Table())
	}

	// 是否是管理员：用户类型=admin 或者 登录用户id = 主体的管理员ID
	claims.IsAdmin = claims.Type == sys_enum.User.Type.Admin.Code() || claims.Id == company.UserId
	claims.UnionMainId = company.Id
	claims.ParentId = company.ParentId

	return claims, nil
}

// GetEmployeeById 根据ID获取员工信息
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeById(ctx context.Context, id int64) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	data, err := daoctl.GetByIdWithError[TR](
		s.dao.Employee.Ctx(ctx),
		id,
	)

	if err != nil {
		message := s.modules.T(ctx, "{#EmployeeName}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#EmployeeName}{#Data}")
		}
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Employee.Table())
	}

	// 有Bug，这里data为空，也进来了
	if !reflect.ValueOf(data).IsNil() {
		response = *data
	}

	// 跨主体禁止查看员工信息，下级公司可查看上级公司员工信息
	if err == sql.ErrNoRows ||
		!reflect.ValueOf(data).IsNil() && sessionUser != nil &&
			sessionUser.Id != 0 && !reflect.ValueOf(response).IsNil() &&
			response.Data().UnionMainId != sessionUser.UnionMainId &&
			response.Data().UnionMainId != sessionUser.ParentId &&
			!sessionUser.IsAdmin &&
			!sessionUser.IsSuperAdmin {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName} {#error_Data_NotFound}"), s.dao.Employee.Table())
	}

	return s.masker(s.makeMore(ctx, response)), nil
}

// GetEmployeeByName 根据Name获取员工信息
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeByName(ctx context.Context, name string) (response TR, err error) {
	data, err := daoctl.ScanWithError[TR](
		s.dao.Employee.Ctx(ctx).Where(co_do.CompanyEmployee{Name: name}),
	)

	if err != nil {
		message := s.modules.T(ctx, "{#EmployeeName}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#EmployeeName}{#Data}")
		}
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Employee.Table())
	}

	return s.masker(s.makeMore(ctx, *data)), nil
}

// HasEmployeeByName 员工名称是否存在
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) HasEmployeeByName(ctx context.Context, name string, unionMainId int64, excludeIds ...int64) bool {
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
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) HasEmployeeByNo(ctx context.Context, no string, unionMainId int64, excludeIds ...int64) bool { // 如果工号为空则直接返回
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
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeBySession(ctx context.Context) (response TR, err error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if user.Type != s.modules.GetConfig().UserType.Code() {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_NotHasServerPermission"), s.dao.Employee.Table())
	}

	result, _ := s.GetEmployeeById(ctx, user.Id)
	if reflect.ValueOf(result).IsNil() {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_CheckLoginUser_Failed"), s.dao.Employee.Table())
	}
	return result, nil
}

// QueryEmployeeList 获取员工列表
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryEmployeeList(ctx context.Context, search *base_model.SearchParams) (*base_model.CollectRes[TR], error) { // 跨主体查询条件过滤
	// 过滤UnionMainId字段查询条件
	search = s.modules.Company().FilterUnionMainId(ctx, search)

	model := s.dao.Employee.Ctx(ctx)

	includeIds := make([]int64, 0)
	var teamId int64

	// 处理扩展条件，扩展支持 teamId，employeeId, inviteUserId, unionMainId 字段过滤支持
	{
		teamSearch := base_funs.SearchFilterEx(search, "teamId.remove", "employeeId", "inviteUserId.remove", "unionMainId")
		for _, filter := range teamSearch.Filter {
			if filter.Field == "team_id" {
				teamId = gconv.Int64(filter.Value)
				break
			}
		}

		if len(teamSearch.Filter) > 0 && teamId != 0 {

			items, _ := s.modules.Team().QueryTeamMemberList(ctx, teamSearch, true)

			if len(items.Records) > 0 {
				for _, item := range items.Records {
					includeIds = append(includeIds, item.EmployeeId)
				}
			}

			if len(includeIds) > 0 {
				// 过滤掉重复的id
				includeIds = gconv.Int64s(garray.NewSortedStrArrayFrom(gconv.Strings(includeIds)).Unique().Slice())
				model = model.WhereIn(s.dao.Employee.Columns().Id, includeIds)
			}
		}
	}
	isExport := false
	if ctx.Value("isExport") == nil {
		r := g.RequestFromCtx(ctx)
		isExport = r.GetForm("isExport", false).Bool()
	} else {
		isExport = gconv.Bool(ctx.Value("isExport"))
	}

	// 团队里面没有成员的情况，但是限定了团队ID，返回的列表就不应该是全部学生，应该是[]
	if teamId != 0 && len(includeIds) <= 0 {
		return &base_model.CollectRes[TR]{
			Records: make([]TR, 0),
		}, nil
	}

	//result, err := daoctl.Query[TR](model.With(co_model.EmployeeRes{}.Detail, co_model.EmployeeRes{}.User), search, isExport)
	result, err := daoctl.Query[TR](model, search, isExport)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName}{#error_Data_Get_Failed}"), s.dao.Employee.Table())
	}

	items := make([]TR, 0)
	for _, record := range result.Records {
		items = append(items, s.masker(s.makeMore(ctx, record)))
	}
	result.Records = items

	return result, nil
}

// CreateEmployee 创建员工信息
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) CreateEmployee(ctx context.Context, info *co_model.Employee, bindUser *sys_model.SysUser) (response TR, err error) {
	info.Id = 0

	return s.saveEmployee(ctx, info, bindUser)
}

// UpdateEmployee 更新员工信息
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateEmployee(ctx context.Context, info *co_model.UpdateEmployee) (response TR, err error) {
	//data := kconv.Struct(info, &co_model.Employee{})
	//
	//return s.saveEmployee(ctx, data)
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
	employee, err := s.GetEmployeeById(ctx, info.Id)
	if err != nil {
		return response, err
	}
	id := employee.Data().Id

	// 校验员工名称是否已存在
	if info.Name != nil {
		// 重名 & 系统不允许员工重名
		if true == s.HasEmployeeByName(ctx, *info.Name, info.Id) && !co_consts.Global.EmployeeNameCanRepeated {
			return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NameAlreadyExists}"), s.dao.Employee.Table())
		}
	}
	// 校验工号是否允许为空
	if info.No != nil {
		if *info.No == "" && s.modules.GetConfig().AllowEmptyNo == false {
			return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NoNotNull}"), s.dao.Employee.Table())
		}
	}
	if info.No != nil {
		// 校验工号是否已存在
		if true == s.HasEmployeeByNo(ctx, *info.No, employee.Data().UnionMainId, info.Id) {
			return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NoAlreadyExists}"), s.dao.Employee.Table())
		}
	}

	data := kconv.Struct(info, &co_do.CompanyEmployee{})
	if err != nil {
		return response, err
	}

	err = s.dao.Employee.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if employee.Data() != nil {

			var avatarFile *sys_model.FileInfo
			storageSrc := ""
			if info.Avatar != "" {
				// 校验员工头像并保存
				fileInfo, err := sys_service.File().GetFileById(ctx, gconv.Int64(data.Avatar), s.modules.T(ctx, "{#Avatar}{#error_File_FileVoid}"))

				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "", s.dao.Employee.Table())
				}
				avatarFile = fileInfo

				storageSrc = s.modules.GetConfig().StoragePath + "/employee/" + gconv.String(sessionUser.Id) + "/avatar." + fileInfo.Ext

				//avatarFile.Src = s.modules.GetConfig().StoragePath + "/employee/" + gconv.String(data.Id) + "/avatar." + avatarFile.Ext

				info.Avatar = gconv.String(fileInfo.Id)
			}

			// 更新员工信息
			data.UpdatedBy = sessionUser.Id
			data.UpdatedAt = gtime.Now()
			// unionMainId不能修改，强制为nil
			data.UnionMainId = nil
			//data.Mobile = nil
			data.Id = nil

			// 重载Do模型
			doData, err := info.OverrideDo.DoFactory(*data)
			if err != nil {
				return err
			}

			_, err = daoctl.UpdateWithError(s.dao.Employee.Ctx(ctx).Data(doData).OmitNilData().Where(co_do.CompanyEmployee{Id: id}))
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Save_Failed"), s.dao.Employee.Table())
			}

			// 保存文件
			if avatarFile != nil {
				avatarFile, err = sys_service.File().SaveFile(ctx, storageSrc, avatarFile)
				//_, err = sys_dao.SysFile.Ctx(ctx).Insert(avatarFile)
				if err != nil || avatarFile == nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Avatar}{#error_File_Save_Failed}"), s.dao.Employee.Table())
				}
			}
		}

		return nil
	})
	if err != nil {
		return response, err
	}

	return s.GetEmployeeById(ctx, id)
}

// UpdateEmployeeAvatar 更新员工头像
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateEmployeeAvatar(ctx context.Context, id int64, avatar string) bool {
	result, err := daoctl.UpdateWithError(s.dao.Employee.Ctx(ctx).Where(s.dao.Employee.Columns().Id, id).Data(co_do.CompanyEmployee{
		Avatar: avatar,
	}))

	if err != nil && result == 0 {
		return false
	}
	return result > 0
}

// saveEmployee 保存员工信息
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) saveEmployee(ctx context.Context, info *co_model.Employee, bindUser *sys_model.SysUser) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 除匿名用户外，其它用户在有权限的情况下均可以创建或更新员工信息，001 代表默认管理员工号
	// info.Id == 0 仅单纯新建员工时需要初始化用户归属为当前操作员所在 UnionMainId
	if sessionUser.Type > 0 && info.Id == 0 && info.No != "001" {
		info.UnionMainId = sessionUser.UnionMainId
	}

	// 校验员工名称是否已存在
	if true == s.HasEmployeeByName(ctx, info.Name, info.Id) && !co_consts.Global.EmployeeNameCanRepeated {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NameAlreadyExists}"), s.dao.Employee.Table())
	}

	// 校验工号是否允许为空
	if info.No == "" && s.modules.GetConfig().AllowEmptyNo == false {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NoNotNull}"), s.dao.Employee.Table())
	}

	// 校验工号是否已存在
	if true == s.HasEmployeeByNo(ctx, info.No, info.UnionMainId, info.Id) {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#EmployeeName}{#error_NoAlreadyExists}"), s.dao.Employee.Table())
	}

	data := &co_do.CompanyEmployee{}
	gconv.Struct(info, data)

	err = s.dao.Employee.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var avatarFile *sys_model.FileInfo
		storageSrc := ""
		if info.Avatar != "" {
			// 校验员工头像并保存
			fileInfo, err := sys_service.File().GetFileById(ctx, gconv.Int64(info.Avatar), s.modules.T(ctx, "{#Avatar}{#error_File_FileVoid}"))

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "", s.dao.Employee.Table())
			}
			avatarFile = fileInfo

			storageSrc = s.modules.GetConfig().StoragePath + "/employee/" + gconv.String(sessionUser.Id) + "/avatar." + fileInfo.Ext

			//avatarFile.Src = s.modules.GetConfig().StoragePath + "/employee/" + gconv.String(data.Id) + "/avatar." + avatarFile.Ext

			info.Avatar = gconv.String(fileInfo.Id)
		}

		if info.Id == 0 {
			if bindUser == nil {
				data.Id = idgen.NextId()
				// 创建登录信息
				passwordLen := len(gconv.String(data.Id))
				password := gstr.SubStr(gconv.String(data.Id), passwordLen-6, 6)

				bindUser, err = sys_service.SysUser().CreateUser(ctx, sys_model.UserInnerRegister{
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
			}

			// 创建员工信息
			data.Id = bindUser.Id
			data.CreatedBy = sessionUser.Id
			data.CreatedAt = gtime.Now()
			data.UnionMainId = info.UnionMainId

			// 重载Do模型
			doData, err := info.OverrideDo.DoFactory(*data)
			if err != nil {
				return err
			}

			affected, err := daoctl.InsertWithError(s.dao.Employee.Ctx(ctx).Data(doData))

			if sessionUser.Id == data.Id {
				sessionUser.UnionMainId = info.UnionMainId
			}

			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Save_Failed"), s.dao.Employee.Table())
			}
		}

		// 更新逻辑剥离至UpdateEmployee方法
		//else {
		//	// 更新员工信息
		//	data.UpdatedBy = sessionUser.Id
		//	data.UpdatedAt = gtime.Now()
		//	// unionMainId不能修改，强制为nil
		//	data.UnionMainId = nil
		//	data.Mobile = nil
		//
		//	// 重载Do模型
		//	doData, err := info.OverrideDo.DoFactory(*data)
		//	if err != nil {
		//		return err
		//	}
		//
		//	_, err = daoctl.UpdateWithError(s.dao.Employee.Ctx(ctx).Data(doData).Where(co_do.CompanyEmployee{Id: data.Id}))
		//	if err != nil {
		//		return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Save_Failed"), s.dao.Employee.Table())
		//	}
		//}

		// 保存文件
		if avatarFile != nil {
			avatarFile, err = sys_service.File().SaveFile(ctx, storageSrc, avatarFile)
			//_, err = sys_dao.SysFile.Ctx(ctx).Insert(avatarFile)
			if err != nil || avatarFile == nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Avatar}{#error_File_Save_Failed}"), s.dao.Employee.Table())
			}
		}
		return nil
	})
	if err != nil {
		return response, err
	}

	return s.GetEmployeeById(ctx, gconv.Int64(data.Id))
}

// DeleteEmployee 删除员工信息
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) DeleteEmployee(ctx context.Context, id int64) (bool, error) {
	// 这个下面两行查询会过滤掉DeletedAt不为空的
	//employee, err := s.GetEmployeeById(ctx, id)
	//employee, err := s.GetEmployeeDetailById(ctx, id)

	//  Unscoped 找到被软删除的记录
	var employee TR
	err := s.dao.Employee.Ctx(ctx).Where(s.dao.Employee.Columns().Id, id).Unscoped().Scan(&employee)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName}{#error_Disabled_Delete}"), s.dao.Employee.Table())
	}

	if s.modules.GetConfig().HardDeleteWaitAt == -1 || s.modules.GetConfig().HardDeleteWaitAt == math.MaxInt64 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName}{#error_Nonsupport_Delete}"), s.dao.Employee.Table())
	}

	err = s.dao.Employee.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		if s.modules.GetConfig().HardDeleteWaitAt > 0 && !reflect.ValueOf(employee).IsNil() && employee.Data() != nil && employee.Data().DeletedAt == nil {
			// 设置账户状态为已注销
			_, err = sys_service.SysUser().SetUserState(ctx, employee.Data().Id, sys_enum.User.State.Canceled)
			if err != nil {
				return err
			}
			// 设置员工状态为已注销
			_, err = s.dao.Employee.Ctx(ctx).
				Data(co_do.CompanyEmployee{State: co_enum.Employee.State.Canceled.Code()}).
				Where(co_do.CompanyEmployee{Id: employee.Data().Id}).
				Update()
			if err != nil {
				return err
			}
			// 软删除
			_, err = s.dao.Employee.Ctx(ctx).Delete(co_do.CompanyEmployee{Id: employee.Data().Id})
			if err != nil {
				return err
			}
		} else {
			if employee.Data().DeletedAt != nil {
				HardDeleteWaitAt := time.Hour * (time.Duration)(s.modules.GetConfig().HardDeleteWaitAt)

				if gtime.Now().Before(employee.Data().DeletedAt.Add(HardDeleteWaitAt)) {
					hours := gtime.Now().Sub(employee.Data().DeletedAt.Add(HardDeleteWaitAt)).Hours()
					message := s.modules.T(ctx, "error_Employee_Delete_Failed") + s.modules.Tf(ctx, "data_protection_wait_hours", math.Abs(hours))
					return sys_service.SysLogs().ErrorSimple(ctx, err, message, s.dao.Employee.Table())
				}
			}

			// 员工移出团队|小组
			_, err := s.deleteEmployeeTeam(ctx, employee.Data().Id)
			if err != nil {
				return err
			}

			// 删除员工
			_, err = s.dao.Employee.Ctx(ctx).Unscoped().Delete(co_do.CompanyEmployee{Id: employee.Data().Id})
			if err != nil {
				return err
			}

			// 删除用户
			_, err = sys_service.SysUser().DeleteUser(ctx, employee.Data().Id)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Delete_Failed")+": "+err.Error(), s.dao.Employee.Table())
	}
	return true, nil
}

// deleteEmployeeTeam 员工移出小组 | 团队
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) deleteEmployeeTeam(ctx context.Context, employeeId int64) (bool, error) {
	// 直接删除属于员工的团队成员记录
	isSuccess, err := s.modules.Team().DeleteTeamMemberByEmployee(ctx, employeeId)
	if err != nil && isSuccess == false {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Team_DeleteMember_Failed"), s.dao.Employee.Table())
	}

	// 查找到员工是管理员或者队长的团队
	teamList, err := s.modules.Team().QueryTeamList(ctx, &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
			Field:     s.dao.Team.Columns().CaptainEmployeeId,
			Where:     "=",
			Value:     employeeId,
			IsOrWhere: true,
		}, base_model.FilterInfo{
			Field:     s.dao.Team.Columns().OwnerEmployeeId,
			Where:     "=",
			Value:     employeeId,
			IsOrWhere: true,
		}),
	})

	// 假如是队长或者组长，需要将团队表的队长或者组长设置为0
	if len(teamList.Records) > 0 {
		for _, item := range teamList.Records {
			if item.Data().CaptainEmployeeId == employeeId { // 队长或者组长
				ret, err := s.modules.Team().SetTeamCaptain(ctx, item.Data().Id, 0)
				if err != nil || ret == false {
					return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Delete_Failed"), s.dao.Employee.Table())
				}
			}

			if item.Data().OwnerEmployeeId == employeeId { // 团队负责人
				ret, err := s.modules.Team().SetTeamOwner(ctx, item.Data().Id, 0)
				if err != nil || ret == false {
					return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Delete_Failed"), s.dao.Employee.Table())
				}
			}
		}
	}
	return true, nil
}

// GetEmployeeDetailById 根据ID获取员工详细信息
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeDetailById(ctx context.Context, id int64) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	model := s.dao.Employee.Ctx(ctx)

	if sessionUser.IsAdmin == false {
		// 判断用户是否有权限
		can, _ := sys_service.SysPermission().CheckPermission(ctx, co_permission.Employee.PermissionType(s.modules).MoreDetail)
		if can == false {
			model = model.Where(sys_do.SysFile{UnionMainId: sessionUser.UnionMainId})
		}
	}

	data, err := daoctl.GetByIdWithError[TR](
		model,
		id,
	)

	//data, err := daoctl.ScanWithError[TR](model.Where(co_do.CompanyEmployee{Id: id}))

	if err != nil || data == nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetEmployeeDetailById_Failed"), s.dao.Employee.Table())
	}

	if !reflect.ValueOf(data).IsNil() {
		response = *data
	}

	//if err == sql.ErrNoRows ||
	//	!reflect.ValueOf(data).IsNil() && sessionUser != nil &&
	//		sessionUser.Id != 0 &&
	//		response.Data().UnionMainId != sessionUser.UnionMainId &&
	//		response.Data().UnionMainId != sessionUser.ParentId &&
	//		!sessionUser.IsAdmin {
	//}

	if err == sql.ErrNoRows ||
		!reflect.ValueOf(data).IsNil() && sessionUser != nil &&
			sessionUser.Id != 0 && !reflect.ValueOf(response).IsNil() &&
			response.Data().UnionMainId != sessionUser.UnionMainId &&
			//response.Data().UnionMainId != sessionUser.ParentId &&
			!sessionUser.IsAdmin &&
			!sessionUser.IsSuperAdmin {
		// 下级公司也不可查看上级公司员工详细信息
		if response.Data().UnionMainId == sessionUser.ParentId {
			return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_NotHasServerPermission"), s.dao.Employee.Table())
		}
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName} {#error_Data_NotFound}"), s.dao.Employee.Table())
	}

	// 跨主体禁止查看员工信息，
	//if err == sql.ErrNoRows || !reflect.ValueOf(data).IsNil() && response.Data().UnionMainId != sessionUser.UnionMainId && !sessionUser.IsAdmin {
	//	// 下级公司也不可查看上级公司员工详细信息
	//	if response.Data().UnionMainId == sessionUser.ParentId {
	//		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_NotHasServerPermission"), s.dao.Employee.Table())
	//	}
	//	return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#EmployeeName} {#error_Data_NotFound}"), s.dao.Employee.Table())
	//}

	return s.makeMore(ctx, response), err
}

// GetEmployeeListByRoleId 根据角色ID获取所有所属员工
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeListByRoleId(ctx context.Context, roleId int64) (*base_model.CollectRes[TR], error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	userIds, err := sys_service.SysRole().GetRoleMemberIds(ctx, roleId, sessionUser.UnionMainId)
	if err != nil {
		return &base_model.CollectRes[TR]{
			PaginationRes: base_model.PaginationRes{
				Pagination: base_model.Pagination{
					PageNum:  1,
					PageSize: 20,
				},
				PageTotal: 0,
				Total:     0,
			},
		}, nil
	}

	result, err := daoctl.Query[TR](
		s.dao.Employee.Ctx(ctx),
		&base_model.SearchParams{
			Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
				Field: s.dao.Employee.Columns().Id,
				Where: "in",
				Value: userIds,
			}),
			OrderBy:    nil,
			Pagination: base_model.Pagination{},
		},
		true,
	)

	items := make([]TR, 0)
	for _, record := range result.Records {
		items = append(items, s.masker(s.makeMore(ctx, record)))
	}
	result.Records = items

	return result, err
}

// SetEmployeeState 设置员工状态
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) SetEmployeeState(ctx context.Context, id int64, state int) (bool, error) {
	_, err := s.modules.Employee().GetEmployeeById(ctx, id)
	if err != nil {
		return false, err
	}

	result, err := s.dao.Employee.Ctx(ctx).Where(s.dao.Employee.Columns().Id, id).Update(co_do.CompanyEmployee{State: state})
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_update_employee_state_failed"), s.dao.Employee.Table())
	}

	return true, nil

}

// Masker 员工信息脱敏，并加载附加数据
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) masker(employee TR) TR {
	if reflect.ValueOf(employee).IsNil() {
		return employee
	}

	employee.Data().Mobile = masker.MaskString(employee.Data().Mobile, masker.MaskPhone)
	employee.Data().LastActiveIp = masker.MaskString(employee.Data().LastActiveIp, masker.MaskIPv4)

	if employee.Data().Detail != nil {
		employee.Data().Detail.LastLoginIp = masker.MaskString(employee.Data().Detail.LastLoginIp, masker.MaskIPv4)
	}

	// 将头像换成可访问url
	if gconv.Int64(employee.Data().Avatar) > 0 {
		employee.Data().Avatar = sys_service.File().MakeFileUrl(context.Background(), gconv.Int64(employee.Data().Avatar))
	}

	return employee
}

// makeMore 按需加载附加数据
func (s *sEmployee[
	ITCompanyRes,
	TR,
	ITTeamRes,
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

	//exclude := &garray.StrArray{}
	//arrayform := r.GetForm("exclude")
	//if arrayform != nil {
	//	array := arrayform.Array()
	//	arr :=
	//	exclude = garray.NewStrArrayFrom(*arr)
	//
	//}

	//if ctx.Value("exclude") == nil {
	//	r := g.RequestFromCtx(ctx)
	//	array := r.GetForm("exclude").Array()
	//	arr := kconv.Struct(array, &[]string{})
	//	exclude = garray.NewStrArrayFrom(*arr)
	//} else {
	//	array := ctx.Value("isExport")
	//	arr := kconv.Struct(array, &[]string{})
	//	exclude = garray.NewStrArrayFrom(*arr)
	//}

	// team附加数据
	if data.Data().UnionMainId > 0 {
		//if data.Data().UnionMainId > 0 && !exclude.Contains("teamList") {
		//if data.Data().UnionMainId > 0 && exclude.Contains("teamList") {
		base_funs.AttrMake[TR](ctx,
			s.dao.Employee.Columns().UnionMainId,
			func() []ITTeamRes {
				g.Try(ctx, func(ctx context.Context) {
					// 获取到该员工的所有团队成员信息记录ids
					ids, err := s.dao.TeamMember.Ctx(ctx).
						Where(co_do.CompanyTeamMember{EmployeeId: data.Data().CompanyEmployee.Id}).Fields([]string{co_dao.CompanyTeamMember.Columns().TeamId}).All()

					temIds := ids.Array()

					if err != nil {
						return
					}

					if len(ids) == 0 {
						data.Data().TeamList = nil
						return
					}

					// 记录该员工所在所有团队
					if len(temIds) > 0 {
						s.dao.Team.Ctx(ctx).
							WhereIn(s.dao.Team.Columns().Id, temIds).Scan(&data.Data().TeamList)
					}
					TeamList := data.Data().TeamList

					data.Data().TeamList = nil
					// 添加附加数据
					data.Data().SetTeamList(TeamList)

					// 业务层添加附加数据
					data.SetTeamList(data.Data().TeamList)
				})

				return kconv.Struct(data.Data().TeamList, []ITTeamRes{})
			},
		)
	}

	// user相关附加数据
	if data.Data().CompanyEmployee.Id > 0 {
		//if exclude.Contains("user") {
		//	data.Data().User = nil
		//} else {
		base_funs.AttrMake[TR](ctx,
			s.dao.Employee.Columns().Id,
			func() (res TR) {
				// 为什么要在内部订阅
				//ctx = base_funs.AttrBuilder[TR, []ITTeamRes](ctx, s.modules.Dao().Employee.Columns().UnionMainId)
				//ctx = base_funs.AttrBuilder[TR, TR](ctx, s.modules.Dao().Employee.Columns().Id)
				//ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)

				if data.Data().CompanyEmployee.Id == 0 {
					return res
				}

				user, _ := sys_service.SysUser().GetSysUserById(ctx, data.Data().CompanyEmployee.Id)

				if user != nil {
					kconv.Struct(user, &data.Data().User)
					data.Data().SetUser(data.Data().User)
					data.SetUser(data.Data().User)

					gconv.Struct(user.SysUser, &data.Data().User)
					gconv.Struct(user.Detail, &data.Data().Detail)
				}

				return data
			},
		)
		//}
	}

	return data
}
