package company

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/masker"
	"github.com/yitter/idgenerator-go/idgen"
	"reflect"

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
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		TR,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
	dao *co_dao.XDao
}

func NewCompany[
	TR co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) co_interface.ICompany[TR] {
	result := &sCompany[
		TR,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes]{
		modules: modules,
		dao:     modules.Dao(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.ICompanyRes
	ret = &co_model.CompanyRes{
		Company:   co_entity.Company{},
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
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetCompanyById(ctx context.Context, id int64) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	data, err := daoctl.GetByIdWithError[TR](
		s.dao.Company.Ctx(ctx),
		id,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Get_Failed}"), s.dao.Company.Table())
		}
	}

	if !reflect.ValueOf(data).IsNil() {
		response = *data
	}

	if err == sql.ErrNoRows || !reflect.ValueOf(data).IsNil() && response.Data().Id != sessionUser.UnionMainId && response.Data().ParentId != sessionUser.UnionMainId {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.masker(s.makeMore(ctx, response)), nil
}

// GetCompanyByName 根据Name获取获取公司信息
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
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

	return s.masker(s.makeMore(ctx, response)), nil
}

// HasCompanyByName 判断名称是否存在
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
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
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryCompanyList(ctx context.Context, filter *base_model.SearchParams) (*base_model.CollectRes[TR], error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
	data, err := daoctl.Query[TR](
		s.dao.Company.Ctx(ctx).
			Where(co_do.Company{ParentId: sessionUser.UnionMainId}),
		filter,
		false,
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	if data.Total > 0 {
		items := make([]TR, 0)
		// 脱敏处理
		for _, item := range data.Records {
			items = append(items, s.masker(s.makeMore(ctx, item)))
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
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) CreateCompany(ctx context.Context, info *co_model.Company) (response TR, err error) {
	info.Id = 0
	return s.saveCompany(ctx, info)
}

// UpdateCompany 更新公司信息
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateCompany(ctx context.Context, info *co_model.Company) (response TR, err error) {
	if info.Id <= 0 {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}
	return s.saveCompany(ctx, info)
}

// SaveCompany 保存公司信息
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) saveCompany(ctx context.Context, info *co_model.Company) (response TR, err error) {
	// 名称重名检测
	if s.HasCompanyByName(ctx, info.Name, info.Id) {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "{#CompanyName} {#error_NameAlreadyExists}"), s.dao.Company.Table())
	}

	// 获取登录用户
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 构建公司ID
	UnionMainId := idgen.NextId()

	data := co_do.Company{
		Id:            info.Id,
		Name:          info.Name,
		ContactName:   info.ContactName,
		ContactMobile: info.ContactMobile,
		Remark:        info.Remark,
	}

	// 启用事务
	err = s.dao.Company.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var employee co_model.IEmployeeRes
		if info.Id == 0 {
			// 是否创建默认员工和角色
			if s.modules.GetConfig().IsCreateDefaultEmployeeAndRole {
				// 构建员工信息
				employee, err = s.modules.Employee().CreateEmployee(ctx, &co_model.Employee{
					No:          "001",
					Name:        info.ContactName,
					Mobile:      info.ContactMobile,
					UnionMainId: UnionMainId,
					State:       co_enum.Employee.State.Normal.Code(),
					HiredAt:     gtime.Now(),
				})
				if err != nil {
					return err
				}

				// 构建角色信息
				roleData := sys_model.SysRole{
					Name:        "管理员",
					UnionMainId: UnionMainId,
					IsSystem:    true,
				}
				roleInfo, err := sys_service.SysRole().Create(ctx, roleData)
				if err != nil {
					return err
				}
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

			data.Id = UnionMainId
			data.ParentId = sessionUser.UnionMainId
			data.CreatedBy = sessionUser.Id
			data.CreatedAt = gtime.Now()

			affected, err := daoctl.InsertWithError(
				s.dao.Company.Ctx(ctx),
				data,
			)
			if affected == 0 || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
			}
		} else {
			if gstr.Contains(info.ContactMobile, "***") || info.ContactMobile == "" {
				data.ContactMobile = nil
			}

			data.UpdatedBy = sessionUser.Id
			data.UpdatedAt = gtime.Now()
			_, err = daoctl.UpdateWithError(
				s.dao.Company.Ctx(ctx).
					Where(co_do.Company{Id: info.Id}),
				data,
			)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
			}
		}
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Save_Failed}"), s.dao.Company.Table())
		}
		return nil
	})

	if err != nil {
		return response, err
	}

	return s.GetCompanyById(ctx, data.Id.(int64))
}

// GetCompanyDetail 获取公司详情，包含完整商务联系人电话
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetCompanyDetail(ctx context.Context, id int64) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	data, err := daoctl.GetByIdWithError[TR](
		s.dao.Company.Ctx(ctx),
		id,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_Get_Failed}"), s.dao.Company.Table())
		}
	}

	if !reflect.ValueOf(data).IsNil() {
		response = *data
	}

	if err == sql.ErrNoRows || !reflect.ValueOf(data).IsNil() && response.Data().Id != sessionUser.UnionMainId && response.Data().ParentId != sessionUser.UnionMainId {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#CompanyName} {#error_Data_NotFound}"), s.dao.Company.Table())
	}

	return s.makeMore(ctx, response), nil
}

// FilterUnionMainId 跨主体查询条件过滤
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
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

		if len(search.Filter) == 0 {
			search.Filter = append(search.Filter, base_model.FilterInfo{
				Field:     "union_main_id",
				Where:     "=",
				IsOrWhere: false,
				Value:     sessionUser.UnionMainId,
			})
		}
	}

	// 遍历所有过滤条件：
	for _, field := range search.Filter {
		// 过滤所有自定义主体ID条件
		if gstr.ToLower(field.Field) == gstr.ToLower("unionMainId") {
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

// makeMore 按需加载附加数据
func (s *sCompany[
	TR,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) makeMore(ctx context.Context, data TR) TR {
	if reflect.ValueOf(data).IsNil() || data.Data() == nil {
		return data
	}

	if data.Data().UserId > 0 {
		// 附加数据 employee
		funs.AttrMake[co_model.CompanyRes](ctx, s.dao.Company.Columns().UserId,
			func() *co_model.EmployeeRes {
				employee, _ := s.modules.Employee().GetEmployeeById(ctx, data.Data().UserId)
				if employee.Data() == nil {
					return nil
				}
				data.Data().AdminUser = employee.Data()

				user, _ := sys_service.SysUser().GetSysUserById(ctx, data.Data().UserId)
				if user != nil {
					gconv.Struct(user.SysUser, &data.Data().AdminUser.User)
					gconv.Struct(user.Detail, &data.Data().AdminUser.Detail)
				}

				return data.Data().AdminUser
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
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) masker(data TR) TR {
	if reflect.ValueOf(data).IsNil() {
		return data
	}

	data.Data().ContactMobile = masker.MaskString(data.Data().ContactMobile, masker.MaskPhone)

	return data
}
