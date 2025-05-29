package company

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"reflect"
	"strings"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/base-library/utility/masker"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"

	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
)

type sCompany[
TR co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
ITFdRechargeRes co_model.IFdRechargeRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		TR,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]
	superAdminMainId int64
	dao              co_dao.XDao
	//makeMoreFunc func(ctx context.Context, data co_model.ICompanyRes, employeeModule co_interface.IEmployee[co_model.IEmployeeRes]) co_model.ICompanyRes
}

func NewCompany[
TR co_model.ICompanyRes,
ITEmployeeRes co_model.IEmployeeRes,
ITTeamRes co_model.ITeamRes,
ITFdAccountRes co_model.IFdAccountRes,
ITFdAccountBillRes co_model.IFdAccountBillsRes,
ITFdBankCardRes co_model.IFdBankCardRes,
ITFdInvoiceRes co_model.IFdInvoiceRes,
ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) co_interface.ICompany[TR] {
	result := &sCompany[
		TR,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]{
		modules: modules,
		dao:     *modules.Dao(),
	}

	userType := g.Cfg().MustGet(context.Background(), "service.superAdminMainId", 1)

	result.superAdminMainId = userType.Int64()
	//result.makeMoreFunc = MakeMore

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	// 订阅邀约用户注册Hook，然后将新用户设置到邀约userId中所属主体中
	sys_service.SysAuth().InstallInviteRegisterHook(sys_enum.Invite.Type.Register, result.SetNewUserJoinCompanyHook)

	return result
}

// SetNewUserJoinCompanyHook 将注册的新用户添加至邀约者的主体中
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetNewUserJoinCompanyHook(ctx context.Context, state sys_enum.InviteType, invite *sys_model.InviteRes, info *sys_entity.SysInvitePerson, registerInfo *sys_model.SysUser) (bool, error) {
	// 如下的逻辑：符合邀约的情况，为xx主体邀请新用户，下列逻辑只是将新用户设置为该主体的员工
	if state.Code() != sys_enum.Invite.Type.Register.Code() {
		return false, nil
	}

	// 找到userId对应的主体
	user, err := sys_service.SysUser().GetSysUserById(ctx, invite.UserId)
	if err != nil {
		return true, nil
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, user.Id)
	if err != nil || reflect.ValueOf(employee).IsNil() || employee.Data() == nil {
		return false, nil
	}

	// 将新用户设置至主体中  TODO 需要封装
	data := co_do.CompanyEmployee{
		Id:          registerInfo.Id,
		No:          nil, // 工号暂定
		Avatar:      nil, // 头像等后期用户登陆系统进行完善
		Name:        registerInfo.Username,
		Mobile:      registerInfo.Mobile,
		UnionMainId: employee.Data().UnionMainId,
		State:       0, // 状态：待确认
		CreatedBy:   invite.UserId,
		CreatedAt:   gtime.Now(),
	}

	affected, err := daoctl.InsertWithError(s.modules.Dao().Employee.Ctx(ctx).OmitNilData().Data(data))
	if affected == 0 || err != nil {
		return true, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Save_Failed"), s.modules.Dao().Employee.Table())
	}

	{
		SubInvitePersonData := sys_model.InvitePersonInfo{
			FormUserId: invite.UserId,
			ByUserId:   registerInfo.Id,
			InviteCode: invite.Code,
			InviteId:   invite.Id,
		}

		{
			// 邀请人
			InvitePerson, _ := sys_service.SysInvite().GetInvitePersonByUserId(ctx, invite.UserId)
			if InvitePerson != nil {
				SubInvitePersonData.CompanyIdentifierPrefix = InvitePerson.CompanyIdentifierPrefix
			}
		}

		// 创建邀约级别关系
		person, err := sys_service.SysInvite().CreateInvitePerson(ctx, &SubInvitePersonData)

		return person != nil, err
	}
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.ICompanyRes = &co_model.CompanyRes{
		Company:   &co_entity.Company{},
		AdminUser: nil,
	}
	return ret.(TR)
}

// GetCompanyById 根据ID获取获取公司信息
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetCompanyById(ctx context.Context, id int64) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
	if id == 0 {
		return response, sys_service.SysLogs().WarnSimple(ctx, nil, s.modules.T(ctx, "error_Id_NotNull"), s.dao.Company.Table())
	}
	m := s.dao.Company.Ctx(ctx)

	if !sessionUser.IsSuperAdmin && sessionUser.UnionMainId != s.superAdminMainId && sessionUser.UnionMainId != id {
		m = m.Where(co_do.Company{ParentId: sessionUser.UnionMainId}).WhereOr(co_do.Company{Id: sessionUser.UnionMainId})
	}

	data, err := daoctl.GetByIdWithError[TR](m, id)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Get_Failed}"), s.dao.Company.Table())
		}
	}
	// 为什么data为空，还是会进去if
	if !reflect.ValueOf(data).IsNil() {
		response = *data
	}

	if errors.Is(err, sql.ErrNoRows) {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.masker(s.MakeMore(ctx, response)), nil
}

// GetCompanyByName 根据Name获取获取公司信息
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetCompanyByName(ctx context.Context, name string) (response TR, err error) {
	data, err := daoctl.ScanWithError[TR](
		s.dao.Company.Ctx(ctx).
			Where(co_do.Company{Name: name}),
	)

	if err != nil || data == nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	if !reflect.ValueOf(data).IsNil() {
		response = *data
	}

	return s.masker(s.MakeMore(ctx, response)), nil
}

// HasCompanyByName 判断名称是否存在
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) HasCompanyByName(ctx context.Context, name string, excludeIds ...int64) bool {
	model := s.dao.Company.Ctx(ctx)

	if len(excludeIds) > 0 {
		var ids []int64
		for _, id := range excludeIds {
			if id > 0 {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			model = model.WhereNotIn(s.dao.Company.Columns().Id, ids)
		}
	}

	count, _ := model.Where(co_do.Company{Name: name}).Count()
	return count > 0
}

// QueryCompanyList 查询公司列表
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) QueryCompanyList(ctx context.Context, filter *base_model.SearchParams, isExport ...bool) (*base_model.CollectRes[TR], error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
	export := false

	if len(isExport) > 0 {
		export = isExport[0]
	}

	m := s.dao.Company.Ctx(ctx)

	if !sessionUser.IsSuperAdmin && sessionUser.UnionMainId != s.superAdminMainId {
		m = m.Where(co_do.Company{ParentId: sessionUser.UnionMainId}).WhereOr(co_do.Company{Id: sessionUser.UnionMainId})
	}

	data, err := daoctl.Query[TR](m, filter, export)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	if data.Total > 0 {
		items := make([]TR, 0)
		// 脱敏处理
		for _, item := range data.Records {
			items = append(items, s.masker(s.MakeMore(ctx, item)))
		}
		data.Records = items
	}

	return data, nil
}

// CreateCompany 创建公司信息
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) CreateCompany(ctx context.Context, info *co_model.Company, bindUser *sys_model.SysUser) (response TR, err error) {
	info.Id = 0
	return s.saveCompany(ctx, info, bindUser)
}

// UpdateCompany 更新公司信息
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateCompany(ctx context.Context, info *co_model.Company) (response TR, err error) {
	if info.Id <= 0 {
		return response, sys_service.SysLogs().WarnSimple(ctx, nil, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}
	return s.saveCompany(ctx, info, nil)
}

func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetCompanyState(ctx context.Context, companyId int64, companyState co_enum.CompanyState) (bool, error) {
	if companyId <= 0 {
		return false, sys_service.SysLogs().WarnSimple(ctx, nil, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	affected, err := daoctl.UpdateWithError(s.dao.Company.Ctx(ctx).Where(s.dao.Company.Columns().Id, companyId), co_do.Company{State: companyState.Code()})

	return affected > 0, err
}

// SaveCompany 保存公司信息
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) saveCompany(ctx context.Context, info *co_model.Company, bindUser *sys_model.SysUser) (response TR, err error) {
	// 名称重名检测
	if info.Name != nil {
		if s.HasCompanyByName(ctx, *info.Name, info.Id) {
			return response, sys_service.SysLogs().WarnSimple(ctx, nil, s.modules.T(ctx, "{#CompanyName} {#error_NameAlreadyExists}"), s.dao.Company.Table())
		}
	}

	// 构建公司ID
	unionMainId := idgen.NextId()

	data := kconv.Struct(info, &co_do.Company{})

	//data := co_do.Company{
	//	Id:            info.Id,
	//	Name:          info.Name,
	//	ContactName:   info.ContactName,
	//	ContactMobile: info.ContactMobile,
	//	Remark:        info.Remark,
	//	Address:       info.Address,
	//	LicenseId:     info.LicenseId,
	//	LicenseState:  info.LicenseState,
	//}

	// 获取登录用户
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if bindUser == nil && !sessionUser.IsSuperAdmin && sessionUser.UnionMainId != s.superAdminMainId && sessionUser.Type < s.modules.GetConfig().UserType.Code() {
		return response, sys_service.SysLogs().WarnSimple(ctx, nil, s.modules.T(ctx, "error_insufficient_permissions"), s.dao.Company.Table())
	}

	// 启用事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var employee co_model.IEmployeeRes
		if info.Id == 0 {
			if bindUser != nil && bindUser.Type != s.modules.GetConfig().UserType.Code() {
				ok, err := sys_service.SysUser().SetUserType(ctx, bindUser.Id, s.modules.GetConfig().UserType)
				if !ok || err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_save_failed_cannot_update_related_user"), s.dao.Company.Table())
				}
			}

			// 是否创建默认员工和角色
			if s.modules.GetConfig().IsCreateDefaultEmployeeAndRole {
				employeeDoData, err := info.Employee.DoFactory(&co_model.Employee{
					No:          "001",
					Name:        *info.ContactName,
					Mobile:      *info.ContactMobile,
					UnionMainId: unionMainId,
					State:       co_enum.Employee.State.Normal.Code(),
					HiredAt:     gtime.Now(),
				})
				if err != nil {
					return err
				}
				employeeData := employeeDoData.(*co_model.Employee)

				// 1.构建员工信息 + user登录信息
				employee, err = s.modules.Employee().CreateEmployee(ctx, employeeData, bindUser)
				if err != nil {
					return err
				}

				// 2.构建角色信息
				roleData := sys_model.SysRole{
					Name:        s.modules.T(ctx, "admin_role_name"),
					UnionMainId: unionMainId,
					IsSystem:    true,
				}
				roleInfo, err := sys_service.SysRole().Create(ctx, roleData)
				if err != nil {
					return err
				}
				// 设置首个员工为：自己内部管理员
				_, err = sys_service.SysUser().SetUserRoleIds(ctx, []int64{roleInfo.Id}, employee.Data().Id)
				if err != nil {
					return err
				}
			}

			if employee != nil {
				// 如果需要创建默认的用户和角色的时候才会有employee，所以进行非空判断，不然会有空指针错误
				data.UserId = employee.Data().Id
			} else {
				data.UserId = 0
			}

			// 3.构建公司信息
			data.Id = unionMainId
			info.Id = unionMainId
			if sessionUser.Type > 4 {
				data.ParentId = sessionUser.UnionMainId
			}
			data.CreatedBy = sessionUser.Id
			data.CreatedAt = gtime.Now()
			//data.LicenseId = 0 // 首次创建没有主体id

			// 重载Do模型
			doData, err := info.OverrideDo.DoFactory(*data)
			if err != nil {
				return err
			}

			affected, err := daoctl.InsertWithError(
				s.dao.Company.Ctx(ctx),
				doData,
			)
			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
			}

			// 4.创建主财务账号  通用账户
			accountData := co_do.FdAccount{}
			_ = gconv.Struct(info, &accountData)

			account := &co_model.FdAccountRegister{
				Name: *info.Name,
				//UnionLicenseId:     0, // 刚注册的公司暂时还没有主体资质

				UnionUserId:        gconv.Int64(data.UserId),
				UnionMainId:        unionMainId,
				CurrencyCode:       "CNY",
				PrecisionOfBalance: 100,
				SceneType:          0,                                         // 不限
				AccountType:        co_enum.Finance.AccountType.System.Code(), // 一个主体只会有一个系统财务账号，并且编号为空
				AccountNumber:      "",                                        // 账户编号
				AllowExceed:        1,                                         // 系统账号默认是可以存在负余额
			}

			createAccount, err := s.modules.Account().CreateAccount(ctx, *account, sessionUser.Id)
			if err != nil || reflect.ValueOf(createAccount).IsNil() {
				return err
			}

		} else {
			//if gstr.Contains(*info.ContactMobile, "***") || *info.ContactMobile == "" {
			//	data.ContactMobile = nil
			//}

			data.UpdatedBy = sessionUser.Id
			data.UpdatedAt = gtime.Now()
			data.Id = nil

			// 重载Do模型
			doData, err := info.OverrideDo.DoFactory(*data)
			if err != nil {
				return err
			}

			_, err = daoctl.UpdateWithError(
				s.dao.Company.Ctx(ctx).
					Where(co_do.Company{Id: info.Id}).OmitNilData(),
				doData,
			)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
			}
		}

		// 保存LOGO
		if info.LogoId != nil && *info.LogoId > 0 {
			logoInfo, err := sys_service.File().GetFileById(ctx, *info.LogoId, "")
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
			}

			uploadPath := s.modules.GetConfig().StoragePath
			tempPath := g.Cfg().MustGet(ctx, "upload.tempPath").String()

			if logoInfo != nil && !strings.HasPrefix(logoInfo.Src, uploadPath) {
				if strings.HasPrefix(logoInfo.Src, tempPath) {
					targetFilePath := uploadPath + "/" + gconv.String(info.Id) + "/logo" + logoInfo.Ext
					_, err := sys_service.File().SaveFile(ctx, targetFilePath, logoInfo, true)
					if err != nil {
						return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		return response, err
	}

	return s.GetCompanyById(ctx, info.Id)
}

// GetCompanyDetail 获取公司详情，包含完整商务联系人电话
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetCompanyDetail(ctx context.Context, id int64) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	m := s.dao.Company.Ctx(ctx)

	if !sessionUser.IsSuperAdmin && sessionUser.UnionMainId != s.superAdminMainId {
		m = m.Where(co_do.Company{ParentId: sessionUser.UnionMainId}).WhereOr(co_do.Company{Id: sessionUser.UnionMainId})
	}

	data, err := daoctl.GetByIdWithError[TR](m, id)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Get_Failed}"), s.dao.Company.Table())
		}
	}

	if !reflect.ValueOf(data).IsNil() {
		response = *data
	}

	if errors.Is(err, sql.ErrNoRows) {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.MakeMore(ctx, response), nil
}

// SetCommissionRate 设置佣金比例
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetCommissionRate(ctx context.Context, companyId int64, commissionRate int, actionUserId int64) (bool, error) {
	response, err := s.GetCompanyById(ctx, companyId)
	if err != nil {
		return false, err
	}

	if response.Data().UserId == actionUserId {
		return false, errors.New("error_only_your_superior_has_the_authority_to_set_it._please_contact_your_superior")
	}

	actionUser, err := sys_service.SysUser().GetSysUserById(ctx, actionUserId)
	if err != nil {
		return false, err
	}

	// 限制同级用户或单位进行佣金的配置
	if actionUser.Type <= s.modules.GetConfig().UserType.Code() {
		return false, errors.New("error_only_your_superior_has_the_authority_to_set_it._please_contact_your_superior")
	}

	canMaxCommissionRate := 100

	// 佣金模式如果相对成交金额则不超过上级佣金百分比，否则将相对于上级佣金收益的百分比
	if co_consts.Global.CommissionModel.Code() == co_enum.Common.CommissionMode.TradeAmount.Code() && response.Data().ParentId > 0 {
		parentCompany, err := s.GetCompanyById(ctx, response.Data().ParentId)
		if err != nil {
			return false, err
		}
		if parentCompany.Data().ParentId > 0 {
			canMaxCommissionRate = parentCompany.Data().CommissionRate
		}
	}

	if commissionRate < 0 || commissionRate > canMaxCommissionRate {
		return false, errors.New("error_commission_rate_must_be_between_0_and_" + gconv.String(canMaxCommissionRate))
	}

	affected, err := daoctl.UpdateWithError(s.modules.Dao().Company.Ctx(ctx).Where(s.modules.Dao().Company.Columns().Id, companyId).Data(co_do.Company{
		CommissionRate: commissionRate,
	}))

	if err != nil && affected == 0 {
		return false, err
	}
	
	return affected == 1, err
}

// SetCompanyAdminUser 设置主体的管理员用户
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetCompanyAdminUser(ctx context.Context, sysUserId, unionMainId int64) (bool, error) {
	// 用户是否存在
	sysUser, err := sys_service.SysUser().GetSysUserById(ctx, sysUserId)
	if err != nil {
		return false, err
	}

	// 主体是否存在
	company, err := daoctl.GetByIdWithError[co_model.Company](
		s.dao.Company.Ctx(ctx),
		unionMainId,
	)

	if company == nil || err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_company_not_exist"), s.dao.Company.Table())
	}

	// 是否是本主体员工
	isCompanyEmployee := false

	// 1、 判断用户是够存在别的主体中了
	employee, _ := s.modules.Employee().GetEmployeeById(ctx, sysUser.Id)
	if !reflect.ValueOf(employee).IsNil() && employee.Data().UnionMainId != 0 {
		if employee.Data().UnionMainId != unionMainId {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_user_belongs_to_other_company"), s.dao.Company.Table())
		} else {
			isCompanyEmployee = true
		}
	}

	// 2、将用户添加为主体的员工
	if !isCompanyEmployee {
		// 不能：如下方法会将当前登陆用户作为本主体的员工操作添加员工这一行为
		// s.modules.Employee().CreateEmployee()

		data := co_do.CompanyEmployee{
			Id:          sysUser.Id,
			No:          nil, // 工号暂定
			Avatar:      nil, // 头像等后期用户登陆系统进行完善
			Name:        sysUser.Username,
			Mobile:      sysUser.Mobile,
			UnionMainId: company.Id,
			State:       0,   // 状态：待确认
			HiredAt:     nil, // 入职时间：nil
			CreatedBy:   0,   // 系统创建：0
			CreatedAt:   gtime.Now(),
		}
		affected, err := daoctl.InsertWithError(s.modules.Dao().Employee.Ctx(ctx).OmitNilData().Data(data))
		if affected == 0 || err != nil {
			return true, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_Save_Failed"), s.modules.Dao().Employee.Table())
		}
	}

	// 3、修改主体的UserId
	affected, err := daoctl.UpdateWithError(s.dao.Company.Ctx(ctx).Where(s.dao.Company.Columns().Id, company.Id).Data(s.dao.Company.Columns().UserId, sysUser.Id))
	if affected == 0 || err != nil {
		return true, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_company_admin_user_set_failed"), s.dao.Company.Table())
	}

	// TODO 4、是否需要创建管理员角色、并设置为该用户...

	return true, nil
}

// FilterUnionMainId 跨主体查询条件过滤
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) FilterUnionMainId(ctx context.Context, search *base_model.SearchParams) *base_model.SearchParams {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	filter := make([]base_model.FilterInfo, 0)

	if search == nil || len(search.Filter) == 0 {
		if search == nil {
			search = &base_model.SearchParams{
				Pagination: base_model.Pagination{
					PageNum:  1,
					PageSize: 20,
				},
			}
		}
	}

	hasUnionMainId := false
	for _, field := range search.Filter {
		if gstr.CaseSnake(field.Field) == "union_main_id" {
			hasUnionMainId = true
			break
		}
	}

	if !hasUnionMainId && sessionUser.UnionMainId != s.superAdminMainId {
		search.Filter = append(search.Filter, base_model.FilterInfo{
			Field:     "union_main_id",
			Where:     "=",
			IsOrWhere: false,
			Value:     sessionUser.UnionMainId,
		})
	}

	// 遍历所有过滤条件：
	for _, field := range search.Filter {
		// 过滤所有自定义主体ID条件
		if gstr.ToLower(field.Field) == gstr.ToLower("unionMainId") || gstr.CaseSnake(field.Field) == "union_main_id" {
			unionMainId := gconv.Int64(field.Value)
			if unionMainId == sessionUser.UnionMainId || unionMainId <= 0 {
				filter = append(filter, field)
				continue
			}
			company, err := s.modules.Company().GetCompanyById(ctx, unionMainId)
			if err != nil || (company.Data().ParentId != unionMainId && company.Data().Id != unionMainId) {
				field.Value = sessionUser.UnionMainId
				filter = append(filter, field)
				continue
			}
		}
		filter = append(filter, field)
	}

	search.Filter = filter

	return search
}

// MakeMore 按需加载附加数据
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,

// ]) MakeMore(ctx context.Context, data TR, employeeModule co_interface.IEmployee[ITEmployeeRes]) TR {
]) MakeMore(ctx context.Context, data TR) TR {
	if reflect.ValueOf(data).IsNil() || data.Data() == nil {
		return data
	}

	if data.Data().UserId > 0 {
		// 附加数据 employee
		base_funs.AttrMake[TR](ctx, co_dao.Company.Columns().UserId,
			func() ITEmployeeRes {
				// 订阅自定义类型的员工数据信息
				ctx = base_funs.AttrBuilder[ITEmployeeRes, ITEmployeeRes](ctx, s.modules.Dao().Employee.Columns().Id)

				// 追加订阅自定义类型的员工扩展数据
				ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_model.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)

				employee, err := s.modules.Employee().GetEmployeeById(ctx, data.Data().UserId)
				//if err != nil || reflect.ValueOf(employee.Data()).IsNil() {
				if err != nil || reflect.ValueOf(employee).IsNil() || employee.Data() == nil {
					return employee
				}
				//// 将头像中的文件id换成可访问地址
				//employee.Data().Avatar = sys_service.File().MakeFileUrl(ctx, gconv.Int64(employee.Data().Avatar))
				//var dd TR = *employee

				// 给Company中对象的AdminUser成员赋值
				data.Data().SetAdminUser(employee.Data())
				// 给自定义类型的AdminUser成员赋值
				data.SetAdminUser(employee)

				return employee
			},
		)
	}

	return data
}

// Masker 信息脱敏
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) masker(data TR) TR {
	if reflect.ValueOf(data).IsNil() {
		return data
	}

	data.Data().ContactMobile = masker.MaskString(data.Data().ContactMobile, masker.MaskPhone)

	ref := reflect.ValueOf(data).Elem()

	if ref.FieldByName("ContactMobile").CanSet() {
		ref.FieldByName("ContactMobile").SetString(data.Data().ContactMobile)
	}
	return data
}
