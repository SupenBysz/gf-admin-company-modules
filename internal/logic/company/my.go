package company

import (
	"context"
	"reflect"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/base_model/base_enum"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

type sMy[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
] struct {
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
		ITFdRechargeRes,
	]
	dao co_dao.XDao
}

func NewMy[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillsRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
	ITFdRechargeRes co_model.IFdRechargeRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) co_interface.IMy {
	return &sMy[
		ITCompanyRes,
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
}

// GetProfile 获取当前员工及用户信息
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetProfile(ctx context.Context) (*co_model.MyProfileRes, error) {
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser

	user, err := sys_service.SysUser().GetUserDetail(ctx, session.Id)
	//user, err := sys_service.SysUser().GetSysUserById(ctx, session.Id)
	if err != nil {
		return nil, err
	}

	// 超级管理员直接返回用户信息
	if session.Type == sys_enum.User.Type.SuperAdmin.Code() {
		return &co_model.MyProfileRes{
			IsSuperAdmin: true,
			User:         user,
		}, nil
	}

	//employee, err := s.modules.Employee().GetEmployeeById(ctx, session.Id)
	employee, err := s.modules.Employee().GetEmployeeDetailById(ctx, session.Id)
	if err != nil || reflect.ValueOf(employee).IsNil() {
		return &co_model.MyProfileRes{
			User:     user,
			Employee: nil,
		}, nil
	}

	res := co_model.MyProfileRes{
		User:     user,
		Employee: employee.Data(),
	}

	response, err := s.modules.Company().GetCompanyById(ctx, employee.Data().UnionMainId)
	if response.Data().UserId == employee.Data().Id {
		res.IsAdmin = true
	}

	return &res, nil
}

// GetCompany 获取当前公司信息
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetCompany(ctx context.Context) (*co_model.MyCompanyRes, error) {
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if session.Type == sys_enum.User.Type.SuperAdmin.Code() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_SuperAdminNotServer"), "my")
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, session.SysUser.Id)
	if err != nil || reflect.ValueOf(employee).IsNil() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_employee_not_found"), co_dao.CompanyEmployee.Table())
	}

	// 公司信息
	company, err := s.modules.Company().GetCompanyById(ctx, employee.Data().UnionMainId)
	if err != nil {
		return nil, err
	}

	result := kconv.Struct(company, &co_model.MyCompanyRes{})

	return result, nil
}

// GetTeams 获取当前员工团队信息
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetTeams(ctx context.Context) (res co_model.MyTeamListRes, err error) {
	res = co_model.MyTeamListRes{}
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 判断身份类型（超级管理员不支持此操作）
	if session.Type == sys_enum.User.Type.SuperAdmin.Code() {
		return co_model.MyTeamListRes{}, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_SuperAdminNotServer"), "my")
	}

	employee, err := s.modules.Employee().GetEmployeeById(ctx, session.Id)
	if err != nil {
		return co_model.MyTeamListRes{}, err
	}

	// 团队列表
	teamList, err := s.modules.Team().QueryTeamListByEmployee(ctx, employee.Data().Id, employee.Data().UnionMainId)
	if err != nil {
		return co_model.MyTeamListRes{}, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#TeamList}{#error_Data_Get_Failed}"), s.dao.Team.Table())
	}

	// 团队成员列表
	for _, team := range teamList.Records {
		var teamInfo co_model.MyTeamRes

		// 团队
		teamInfo.TeamRes = *team.Data()

		// 团队成员列表
		memberList, err := s.modules.Team().GetEmployeeListByTeamId(ctx, team.Data().Id)
		if err != nil {
			return co_model.MyTeamListRes{}, err
		}

		teamInfo.MemberItems = []*co_model.EmployeeRes{}

		for _, record := range memberList.Records {

			teamInfo.MemberItems = append(teamInfo.MemberItems, record.Data())
		}

		// 赋值
		res = append(res, teamInfo)
	}

	return res, nil
}

// SetMyMobile 设置我的手机号
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetMyMobile(ctx context.Context, newMobile string, captcha string, password string) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	_, err := sys_service.SysSms().Verify(ctx, newMobile, captcha, base_enum.Captcha.Type.SetMobile)
	if err != nil {
		return false, err
	}

	// 如果原手机号码和新号码一致，直接返回true
	userInfo, err := sys_service.SysUser().GetUserDetail(ctx, sessionUser.Id)
	if err != nil {
		return false, err
	}
	if newMobile == userInfo.Mobile {
		return true, nil
	}

	// 判断密码是否正确
	checkPassword, _ := sys_service.SysUser().CheckPassword(ctx, sessionUser.Id, password)
	if checkPassword != true {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_SetMobile_Failed"), s.dao.Employee.Table())
	}

	// 更新数据
	affected, err := daoctl.UpdateWithError(s.dao.Employee.Ctx(ctx).
		Data(co_do.CompanyEmployee{Mobile: newMobile, UpdatedBy: sessionUser.Id, UpdatedAt: gtime.Now()}).
		Where(co_do.CompanyEmployee{Id: sessionUser.Id}))

	if err != nil || affected == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_SetMobile_Failed"), s.dao.Employee.Table())
	}

	return true, nil
}

// SetMyMail 设置我的邮箱
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetMyMail(ctx context.Context, oldMail string, newMail string, captcha string, password string) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	_, err := sys_service.SysMails().Verify(ctx, newMail, captcha, base_enum.Captcha.Type.SetMail)
	if err != nil {
		return false, err
	}

	// 如果原手机号码和新号码一致，直接返回true
	userInfo, err := sys_service.SysUser().GetUserDetail(ctx, sessionUser.Id)
	if err != nil {
		return false, err
	}
	if newMail == userInfo.Email {
		return true, nil
	}

	// 判断密码是否正确
	checkPassword, _ := sys_service.SysUser().CheckPassword(ctx, sessionUser.Id, password)
	if checkPassword != true {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_SetMail_Failed"), s.dao.Employee.Table())
	}

	// 更新数据
	affected, err := daoctl.UpdateWithError(s.dao.Employee.Ctx(ctx).
		Data(co_do.CompanyEmployee{Email: newMail, UpdatedBy: sessionUser.Id, UpdatedAt: gtime.Now()}).
		Where(co_do.CompanyEmployee{Id: sessionUser.Id}))

	if err != nil || affected == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Employee_SetMail_Failed"), s.dao.Employee.Table())
	}

	return true, nil
}

// SetMyAvatar 设置我的头像
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) SetMyAvatar(ctx context.Context, imageId int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 上传 --> 保存

	// 校验员工头像并保存
	fileInfo, err := sys_service.File().GetFileById(ctx, imageId, s.modules.T(ctx, "avatar")+s.modules.T(ctx, "error_File_FileVoid"))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "", s.dao.Employee.Table())
	}

	err = s.dao.Employee.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		storageAddr := s.modules.GetConfig().StoragePath + "/employee/" + gconv.String(sessionUser.Id) + "/avatar" + fileInfo.Ext

		_, err = sys_service.File().SaveFile(ctx, storageAddr, fileInfo)

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "avatar")+s.modules.T(ctx, "error_File_Save_Failed"), s.dao.Employee.Table())
		}

		//avatar := s.modules.Employee().UpdateEmployeeAvatar(ctx, sessionUser.Id, fileInfo.Url)
		updateAvatar := s.modules.Employee().UpdateEmployeeAvatar(ctx, sessionUser.Id, gconv.String(fileInfo.Id))

		if updateAvatar == false {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "avatar")+s.modules.T(ctx, "error_Data_Save_Failed"), s.dao.Employee.Table())
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// GetAccountBills 我的账单|列表
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccountBills(ctx context.Context, searchParams *base_model.SearchParams) (*co_model.MyAccountBillRes, error) {
	// 1、获取到当前登录用户
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 2、根据userId查询该用户的所有财务账号  我可能有两个不同货币类型的财务账号
	accounts, err := s.modules.Account().QueryAccountListByUserId(ctx, user.SysUser.Id)
	if err != nil {
		return &co_model.MyAccountBillRes{}, err
	}

	accountBillDetailList := co_model.MyAccountBillRes{}

	// 3、遍历每一个账号，把账单统计出来
	for _, account := range accounts.Records {
		bills, err := s.modules.AccountBills().GetAccountBillsByAccountId(ctx, account.Data().Id, searchParams)
		// base_model.call[co_model.IFdAccountBillsRes]
		if err != nil {
			return nil, err
		}

		// 账号信息
		var accountInfo co_entity.FdAccount
		gconv.Struct(account, &accountInfo)

		// 账单信息
		accountBillDetailList = append(accountBillDetailList, co_model.AccountBillRes{
			Account: accountInfo,
			Bill:    kconv.Struct(bills, &co_model.FdAccountBillsListRes{}),
		})
	}

	return &accountBillDetailList, nil
}

// GetAccounts 获取我的财务账号|列表
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetAccounts(ctx context.Context) (*co_model.FdAccountListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	accountList, err := s.modules.Account().QueryAccountListByUserId(ctx, user.Id)
	if err != nil {
		return &co_model.FdAccountListRes{}, nil
	}

	return kconv.Struct(accountList, &co_model.FdAccountListRes{}), nil
}

// GetBankCards 获取我的银行卡｜列表
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetBankCards(ctx context.Context) (*co_model.FdBankCardListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	bankCardList, err := s.modules.BankCard().QueryBankCardListByUserId(ctx, user.Id)
	if err != nil {
		return &co_model.FdBankCardListRes{}, nil
	}

	return kconv.Struct(bankCardList, &co_model.FdBankCardListRes{}), nil
}

// GetInvoices 获取我的发票抬头|列表
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetInvoices(ctx context.Context) (*co_model.FdInvoiceListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 获取到所有的发票抬头+ 纳税识别号
	invoiceList, err := s.modules.Invoice().QueryInvoiceList(ctx, nil, user.Id)
	if err != nil {
		return &co_model.FdInvoiceListRes{}, nil
	}

	return kconv.Struct(invoiceList, &co_model.FdInvoiceListRes{}), nil
}

// UpdateAccount  修改我的财务账号
func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) UpdateAccount(ctx context.Context, accountId int64, info *co_model.UpdateAccount) (api_v1.BoolRes, error) {
	//user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := s.modules.Account().UpdateAccount(ctx, accountId, info)
	if err != nil {
		return false, err
	}

	return ret == true, nil
}

func (s *sMy[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
	ITFdRechargeRes,
]) GetMyCompanyPermissionList(ctx context.Context, permissionType *int) (*sys_model.MyPermissionListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	unionMainAdminId := gconv.Int64(0)

	if user.IsAdmin {
		unionMainAdminId = user.Id
	} else if user.UnionMainId > 0 {
		response, err := s.modules.Company().GetCompanyById(ctx, user.UnionMainId)
		if err != nil {
			return nil, err
		}
		unionMainAdminId = response.Data().UserId
	}

	// 获取用户的权限ids
	ids, err := sys_service.SysPermission().GetPermissionsByResource(ctx, gconv.String(unionMainAdminId))

	if err != nil {
		return nil, err
	}

	Filter := append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
		Field: sys_dao.SysPermission.Columns().Id,
		Where: "in",
		Value: ids,
	})

	if permissionType != nil && *permissionType > 0 {
		Filter = append(Filter, base_model.FilterInfo{
			Field: sys_dao.SysPermission.Columns().Type,
			Where: "=",
			Value: *permissionType,
		})
	}

	// 获取用户的权限list
	result, err := sys_service.SysPermission().QueryPermissionList(ctx, base_model.SearchParams{
		Filter: Filter,
	})

	if err != nil {
		return nil, err
	}

	ret := sys_model.MyPermissionListRes{}
	ret = result.Records

	return &ret, err
}
