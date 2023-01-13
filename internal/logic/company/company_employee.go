package company

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
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
	"github.com/SupenBysz/gf-admin-community/utility/kconv"
	"github.com/SupenBysz/gf-admin-community/utility/masker"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type sEmployee struct {
	modules co_interface.IModules
}

func NewEmployee(modules co_interface.IModules) co_interface.IEmployee {
	return &sEmployee{
		modules: modules,
	}
}

// GetEmployeeById 根据ID获取员工信息
func (s *sEmployee) GetEmployeeById(ctx context.Context, id int64) (*co_entity.CompanyEmployee, error) {
	data := co_entity.CompanyEmployee{}
	err := co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Scan(&data, co_do.CompanyEmployee{Id: id})

	if err != nil {
		message := s.modules.T(ctx, "{#employee_Name}{#error_Data_NotFound}")
		if err != sql.ErrNoRows {
			message = s.modules.T(ctx, "{#employee_Name}{#Data}")
		}
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, message, co_dao.CompanyEmployee(s.modules).Table())
	}
	return s.masker(&data), nil
}

// HasEmployeeByName 员工名称是否存在
func (s *sEmployee) HasEmployeeByName(ctx context.Context, name string, unionMainId int64, excludeId ...int64) bool {
	model := co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.CompanyEmployee{
		Name:        name,
		UnionMainId: unionMainId,
	})

	if len(excludeId) > 0 {
		model = model.WhereNotIn(co_dao.CompanyTeam(s.modules).Columns().Id, excludeId)
	}

	count, _ := model.Count()
	return count > 0
}

// HasEmployeeByNo 员工工号是否存在
func (s *sEmployee) HasEmployeeByNo(ctx context.Context, no string, unionMainId int64, excludeId ...int64) bool { // 如果工号为空则直接返回
	// 工号为空，且允许工号为空则不做校验
	if no == "" && s.modules.GetConfig().AllowEmptyNo == true {
		return false
	}

	model := co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(co_do.CompanyEmployee{
		No:          no,
		UnionMainId: unionMainId,
	})

	if len(excludeId) > 0 {
		model = model.WhereNotIn(co_dao.CompanyTeam(s.modules).Columns().Id, excludeId)
	}

	count, _ := model.Count()
	return count > 0
}

// GetEmployeeBySession 获取当前登录的员工信息
func (s *sEmployee) GetEmployeeBySession(ctx context.Context) (*co_entity.CompanyEmployee, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if user.Type != s.modules.GetConfig().UserType.Code() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_NotHasServerPermission"), co_dao.CompanyEmployee(s.modules).Table())
	}

	result, _ := s.GetEmployeeById(ctx, user.Id)
	if result == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_CheckLoginUser_Failed"), co_dao.CompanyEmployee(s.modules).Table())
	}
	return result, nil
}

// QueryEmployeeList 获取员工列表
func (s *sEmployee) QueryEmployeeList(ctx context.Context, search *sys_model.SearchParams) (*co_model.EmployeeListRes, error) { // 跨主体查询条件过滤
	search = funs.FilterUnionMain(ctx, search, co_dao.CompanyEmployee(s.modules).Columns().UnionMainId)

	// 查询符合过滤条件的员工信息
	result, err := daoctl.Query[*co_entity.CompanyEmployee](co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler), search, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "employee_Name")+"信息查询失败", co_dao.CompanyEmployee(s.modules).Table())
	}

	items := make([]*co_entity.CompanyEmployee, 0)
	for _, employeeInfo := range result.Records {
		items = append(items, s.masker(employeeInfo))
	}
	result.Records = items

	return (*co_model.EmployeeListRes)(result), nil
}

// CreateEmployee 创建员工信息
func (s *sEmployee) CreateEmployee(ctx context.Context, info *co_model.Employee) (*co_entity.CompanyEmployee, error) { //
	info.Id = 0
	return s.saveEmployee(ctx, info)
}

// UpdateEmployee 更新员工信息
func (s *sEmployee) UpdateEmployee(ctx context.Context, info *co_model.Employee) (*co_entity.CompanyEmployee, error) { //
	return s.saveEmployee(ctx, info)
}

// saveEmployee 保存员工信息
func (s *sEmployee) saveEmployee(ctx context.Context, info *co_model.Employee) (*co_entity.CompanyEmployee, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if sessionUser.Type > 0 && info.UnionMainId == 0 {
		info.UnionMainId = sessionUser.UnionMainId
	}

	if info.UnionMainId == 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "employee_Name")+"所属主体不能为空，请选择后提交", co_dao.CompanyEmployee(s.modules).Table())
	}

	if sessionUser.Type > 0 && sessionUser.UnionMainId != info.UnionMainId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "employee_Name")+"所属主体校验失败，请确认后提交", co_dao.CompanyEmployee(s.modules).Table())
	}

	// 校验员工名称是否已存在
	if true == s.HasEmployeeByName(ctx, info.Name, sessionUser.UnionMainId, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "employee_Name")+"名称已存在，请修改后提交", co_dao.CompanyEmployee(s.modules).Table())
	}

	// 校验工号是否允许为空
	if info.No == "" && s.modules.GetConfig().AllowEmptyNo == false {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "employee_Name")+"工号不能为空，请修改后提交", co_dao.CompanyEmployee(s.modules).Table())
	}

	// 校验工号是否已存在
	if true == s.HasEmployeeByNo(ctx, info.No, sessionUser.UnionMainId, info.Id) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "employee_Name")+"工号已存在，请修改后提交", co_dao.CompanyEmployee(s.modules).Table())
	}

	data := &co_do.CompanyEmployee{}
	gconv.Struct(info, data)

	err := co_dao.CompanyEmployee(s.modules).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var avatarFile *sys_model.FileInfo
		if info.Avatar != "" {
			// 校验员工头像并保存
			fileInfo, err := sys_service.File().GetFileById(ctx, gconv.Int64(info.Avatar), "头像"+s.modules.T(ctx, "error_File_FileVoid"))

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "", co_dao.CompanyEmployee(s.modules).Table())
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
					return err
				}

				data.Id = newUser.UserInfo.Id
			}

			_, err := co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(data).Insert()
			if err != nil {
				return err
			}
		} else {
			// 更新员工信息
			data.UpdatedBy = sessionUser.Id
			data.UpdatedAt = gtime.Now()

			_, err := co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(data).Where(co_do.CompanyEmployee{Id: data.Id}).Update()
			if err != nil {
				return err
			}
		}

		// 保存文件
		if avatarFile != nil {
			avatarFile, err := sys_service.File().SaveFile(ctx, avatarFile.Src, avatarFile)
			_, err = sys_dao.SysFile.Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(avatarFile)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "头像"+s.modules.T(ctx, "error_File_Save_Failed"), co_dao.CompanyEmployee(s.modules).Table())
			}
		}
		return nil
	})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Save_Failed"), co_dao.CompanyEmployee(s.modules).Table())
	}

	return s.GetEmployeeById(ctx, gconv.Int64(data.Id))
}

// DeleteEmployee 删除员工信息
func (s *sEmployee) DeleteEmployee(ctx context.Context, id int64) (bool, error) {
	employee, err := s.GetEmployeeById(ctx, id)
	if err != nil {
		return false, err
	}

	err = co_dao.CompanyEmployee(s.modules).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if s.modules.GetConfig().HardDeleteWaitAt > 0 && employee.DeletedAt == nil {
			// 设置账户状态为已注销
			_, err = sys_service.SysUser().SetUserState(ctx, employee.Id, sys_enum.User.State.Canceled)
			if err != nil {
				return err
			}
			// 设置员工状态为已注销
			_, err = co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).
				Data(co_do.CompanyEmployee{State: co_enum.Employee.State.Canceled.Code()}).
				Where(co_do.CompanyEmployee{Id: employee.Id}).
				Update()
			if err != nil {
				return err
			}
			// 软删除
			_, err = co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Delete(co_do.CompanyEmployee{Id: employee.Id})
			if err != nil {
				return err
			}
		} else {
			if employee.DeletedAt != nil {
				HardDeleteWaitAt := time.Hour * (time.Duration)(s.modules.GetConfig().HardDeleteWaitAt)

				if gtime.Now().Before(employee.DeletedAt.Add(HardDeleteWaitAt)) {
					hours := gtime.Now().Sub(employee.DeletedAt.Add(HardDeleteWaitAt)).Hours()
					message := s.modules.T(ctx, "error_Employee_Delete_Failed") + "数据延期保护中\r请于 " + gconv.String(math.Abs(hours)) + " 小时后操作"
					return sys_service.SysLogs().ErrorSimple(ctx, err, message, co_dao.CompanyEmployee(s.modules).Table())
				}
			}

			// 员工移出团队

			// 员工移出小组

			// 删除员工
			_, err = co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler).Unscoped().Delete(co_do.CompanyEmployee{Id: employee.Id})
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
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Delete_Failed"), co_dao.CompanyEmployee(s.modules).Table())
	}
	return true, nil
}

// SetEmployeeMobile 设置手机号
func (s *sEmployee) SetEmployeeMobile(ctx context.Context, newMobile int64, captcha string) (bool, error) {
	_, err := sys_service.SysSms().Verify(ctx, newMobile, captcha)
	if err != nil {
		return false, err
	}

	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	_, err = co_dao.CompanyEmployee(s.modules).Ctx(ctx).
		Data(co_do.CompanyEmployee{Mobile: newMobile, UpdatedBy: sessionUser.Id, UpdatedAt: gtime.Now()}).
		Where(co_do.CompanyEmployee{Id: sessionUser.Id}).
		Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_SetMobile_Failed"), co_dao.CompanyEmployee(s.modules).Table())
	}

	return true, nil
}

// SetEmployeeAvatar 设置员工头像
func (s *sEmployee) SetEmployeeAvatar(ctx context.Context, imageId int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 校验员工头像并保存
	fileInfo, err := sys_service.File().GetFileById(ctx, imageId, "头像"+s.modules.T(ctx, "error_File_FileVoid"))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "", co_dao.CompanyEmployee(s.modules).Table())
	}

	storageAddr := s.modules.GetConfig().StoragePath + "/employee/" + gconv.String(sessionUser.Id) + "/avatar." + fileInfo.Ext

	_, err = sys_service.File().SaveFile(ctx, storageAddr, fileInfo)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "头像"+s.modules.T(ctx, "error_File_Save_Failed"), co_dao.CompanyEmployee(s.modules).Table())
	}
	return true, nil
}

// GetEmployeeDetailById 根据ID获取员工详细信息
func (s *sEmployee) GetEmployeeDetailById(ctx context.Context, id int64) (*co_entity.CompanyEmployee, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	data := co_entity.CompanyEmployee{}
	model := co_dao.CompanyEmployee(s.modules).Ctx(ctx).Hook(daoctl.CacheHookHandler)

	if sessionUser.IsAdmin == false {
		// 判断用户是否有权限
		can, _ := sys_service.SysPermission().CheckPermission(ctx, co_enum.Employee.PermissionType(s.modules).MoreDetail)
		if can == false {
			model = model.Where(sys_do.SysFile{UnionMainId: sessionUser.UnionMainId})
		}
	}

	err := model.Where(co_do.CompanyEmployee{Id: id}).Scan(&data)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "employee_Name")+"详情信息查询失败", co_dao.CompanyEmployee(s.modules).Table())
	}

	return kconv.StructWithError(data, &co_entity.CompanyEmployee{})
}

// Masker 员工信息脱敏
func (s *sEmployee) masker(employee *co_entity.CompanyEmployee) *co_entity.CompanyEmployee {
	if employee == nil {
		return nil
	}
	employee.Mobile = masker.MaskString(employee.Mobile, masker.MaskPhone)
	employee.LastActiveIp = masker.MaskString(employee.LastActiveIp, masker.MaskIPv4)
	return employee
}
