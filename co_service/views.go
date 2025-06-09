// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package co_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type (
	ICompanyView interface {
		// GetCompanyById 根据公司ID获取公司信息。
		// 该方法首先尝试从数据库中获取公司信息。如果成功获取信息且makeResource参数为真，
		// 则进一步处理公司信息以生成更多的资源数据。
		// 参数:
		//
		//	ctx - 上下文，用于传递请求范围的信息。
		//	id - 公司的唯一标识符。
		//	makeResource - 指示是否需要进一步处理公司信息以生成额外的资源数据。
		//
		// 返回值:
		//
		//	*model.CompanyViewRes - 公司信息的视图资源对象，如果找不到则返回nil。
		//	error - 错误信息，如果执行过程中遇到任何问题则返回。
		GetCompanyById(ctx context.Context, id int64, makeResource bool) (*co_model.CompanyViewRes, error)
		// QueryCompanyList 查询公司列表
		// 该方法根据提供的搜索参数查询公司信息，并可选地处理额外的资源信息
		// 主要用于获取分页的公司列表，每个公司信息的详细程度取决于makeResource参数
		QueryCompanyList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.CompanyViewListRes, error)
	}
	IEmployeeView interface {
		// GetEmployeeById 根据员工ID获取员工详细信息。
		// 该方法首先尝试从数据库中获取员工信息。如果找到员工信息且makeResource参数为true，
		// 则进一步处理数据以生成更多的资源信息。
		// 参数:
		//
		//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
		//	id int64: 员工的唯一标识符。
		//	makeResource bool: 指示是否需要生成额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.EmployeeViewRes: 员工详细信息的视图模型，如果找不到则返回nil。
		//	error: 错误对象，如果操作成功则返回nil。
		GetEmployeeById(ctx context.Context, id int64, makeResource bool) (*co_model.EmployeeViewRes, error)
		// QueryEmployeeList 查询员工列表信息。
		// 该方法根据提供的搜索参数查询员工信息，并可选地构建额外的资源信息。
		// 参数:
		//
		//	ctx - 上下文，用于传递请求范围的上下文信息。
		//	params - 搜索参数，用于指定查询的条件。
		//	makeResource - 指示是否构建额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.EmployeeViewListRes - 包含员工列表的响应对象。
		//	error - 错误信息，如果执行过程中发生错误。
		QueryEmployeeList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.EmployeeViewListRes, error)
		QueryMyInviteEmployee(ctx context.Context, inviteUserId int64) (*[]co_model.EmployeeViewRes, error)
	}
	IFdAccountView interface {
		// GetFdAccountById 根据财务账户ID获取财务账户详细信息。
		// 该方法首先尝试从数据库中获取财务账户信息。如果找到财务账户信息且makeResource参数为true，
		// 则进一步处理数据以生成更多的资源信息。
		// 参数:
		//
		//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
		//	id int64: 财务账户的唯一标识符。
		//	makeResource bool: 指示是否需要生成额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.FdAccountViewRes: 财务账户详细信息的视图模型，如果找不到则返回nil。
		//	error: 错误对象，如果操作成功则返回nil。
		GetFdAccountById(ctx context.Context, id int64, makeResource bool) (*co_model.FdAccountViewRes, error)
		// QueryFdAccountList 查询财务账户列表信息。
		// 该方法根据提供的搜索参数查询财务账户信息，并可选地构建额外的资源信息。
		// 参数:
		//
		//	ctx - 上下文，用于传递请求范围的上下文信息。
		//	params - 搜索参数，用于指定查询的条件。
		//	makeResource - 指示是否构建额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.FdAccountViewListRes - 包含财务账户列表的响应对象。
		//	error - 错误信息，如果执行过程中发生错误。
		QueryFdAccountList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.FdAccountViewListRes, error)
		// GetUserDefaultFdAccountByUserId 根据用户ID和用户类型查询财务账户信息，并返回第一个匹配的财务账户
		// 该方法根据提供的用户ID和用户类型查询财务账户信息，并返回第一个匹配的财务账户
		// 参数:
		//
		//	ctx - 上下文，用于传递请求范围的上下文信息
		//	userId - 用户ID，用于指定查询的用户
		//	userType - 用户类型，用于指定查询的用户类型
		//
		// 返回值:
		//
		//	*co_model.FdAccountViewRes - 财务账户视图数据，如果查询失败则返回nil
		//	error - 错误信息，如果查询失败则返回错误
		GetUserDefaultFdAccountByUserId(ctx context.Context, userId int64, userType sys_enum.UserType) (*co_model.FdAccountViewRes, error)
	}
	IFdBillsView interface {
		GetBillsById(ctx context.Context, id int64, makeResource bool) (*co_model.CompanyBillsViewRes, error)
		// QueryBillsList 公司账单账单查询
		QueryBillsList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.CompanyBillsViewListRes, error)
	}
	IFdBankCardView interface {
		// GetFdBankCardById 根据银行卡ID获取银行卡详细信息。
		// 该方法首先尝试从数据库中获取银行卡信息。如果找到银行卡信息且makeResource参数为true，
		// 则进一步处理数据以生成更多的资源信息。
		// 参数:
		//
		//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
		//	id int64: 银行卡的唯一标识符。
		//	makeResource bool: 指示是否需要生成额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.FdBankCardViewRes: 银行卡详细信息的视图模型，如果找不到则返回nil。
		//	error: 错误对象，如果操作成功则返回nil。
		GetFdBankCardById(ctx context.Context, id int64, makeResource bool) (*co_model.FdBankCardViewRes, error)
		// QueryFdBankCardList 查询银行卡列表信息。
		// 该方法根据提供的搜索参数查询银行卡信息，并可选地构建额外的资源信息。
		// 参数:
		//
		//	ctx - 上下文，用于传递请求范围的上下文信息。
		//	params - 搜索参数，用于指定查询的条件。
		//	makeResource - 指示是否构建额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.FdBankCardViewListRes - 包含银行卡列表的响应对象。
		//	error - 错误信息，如果执行过程中发生错误。
		QueryFdBankCardList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.FdBankCardViewListRes, error)
	}
	IFdInvoiceView interface {
		// GetFdInvoiceById 根据财务发票ID获取财务发票详细信息。
		// 该方法首先尝试从数据库中获取财务发票信息。如果找到财务发票信息且makeResource参数为true，
		// 则进一步处理数据以生成更多的资源信息。
		// 参数:
		//
		//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
		//	id int64: 财务发票的唯一标识符。
		//	makeResource bool: 指示是否需要生成额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.FdInvoiceViewRes: 财务发票详细信息的视图模型，如果找不到则返回nil。
		//	error: 错误对象，如果操作成功则返回nil。
		GetFdInvoiceById(ctx context.Context, id int64, makeResource bool) (*co_model.FdInvoiceViewRes, error)
		// QueryFdInvoiceList 查询财务发票列表信息。
		// 该方法根据提供的搜索参数查询财务发票信息，并可选地构建额外的资源信息。
		// 参数:
		//
		//	ctx - 上下文，用于传递请求范围的上下文信息。
		//	params - 搜索参数，用于指定查询的条件。
		//	makeResource - 指示是否构建额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.FdInvoiceViewListRes - 包含财务发票列表的响应对象。
		//	error - 错误信息，如果执行过程中发生错误。
		QueryFdInvoiceList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.FdInvoiceViewListRes, error)
	}
	IFdRechargeView interface {
		// QueryAccountRecharge 查询充值记录列表
		QueryAccountRecharge(ctx context.Context, search *base_model.SearchParams) (*co_model.FdRechargeViewListRes, error)
		// GetAccountRechargeById 根据资金账户ID获取充值记录
		GetAccountRechargeById(ctx context.Context, id int64) (*co_model.FdRechargeViewRes, error)
		// GetRechargeByAccountId 根据资金账户ID获取充值记录
		GetRechargeByAccountId(ctx context.Context, id int64) (*co_model.FdRechargeViewListRes, error)
		// GetRechargeByUserId 根据用户ID获取充值记录
		GetRechargeByUserId(ctx context.Context, id int64) (*co_model.FdRechargeViewListRes, error)
		// GetRechargeByCompanyId 根据公司ID获取充值记录
		GetRechargeByCompanyId(ctx context.Context, id int64) (*co_model.FdRechargeViewListRes, error)
	}
	ITeamView interface {
		// GetTeamById 根据团队ID获取团队详细信息。
		// 该方法首先尝试从数据库中获取团队信息。如果找到团队信息且makeResource参数为true，
		// 则进一步处理数据以生成更多的资源信息。
		// 参数:
		//
		//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
		//	id int64: 团队的唯一标识符。
		//	makeResource bool: 指示是否需要生成额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.TeamViewRes: 团队详细信息的视图模型，如果找不到则返回nil。
		//	error: 错误对象，如果操作成功则返回nil。
		GetTeamById(ctx context.Context, id int64, makeResource bool) (*co_model.TeamViewRes, error)
		// QueryTeamList 查询团队列表信息。
		// 该方法根据提供的搜索参数查询团队信息，并可选地构建额外的资源信息。
		// 参数:
		//
		//	ctx - 上下文，用于传递请求范围的上下文信息。
		//	params - 搜索参数，用于指定查询的条件。
		//	makeResource - 指示是否构建额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.TeamViewListRes - 包含团队列表的响应对象。
		//	error - 错误信息，如果执行过程中发生错误。
		QueryTeamList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.TeamViewListRes, error)
	}
	ITeamMemberView interface {
		// GetTeamMemberById 根据团队成员ID获取团队成员详细信息。
		// 该方法首先尝试从数据库中获取团队成员信息。如果找到团队成员信息且makeResource参数为true，
		// 则进一步处理数据以生成更多的资源信息。
		// 参数:
		//
		//	ctx context.Context: 上下文对象，用于传递请求范围的信息。
		//	id int64: 团队成员的唯一标识符。
		//	makeResource bool: 指示是否需要生成额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.TeamMemberViewRes: 团队成员详细信息的视图模型，如果找不到则返回nil。
		//	error: 错误对象，如果操作成功则返回nil。
		GetTeamMemberById(ctx context.Context, id int64, makeResource bool) (*co_model.TeamMemberViewRes, error)
		// QueryTeamMemberList 查询团队成员列表信息。
		// 该方法根据提供的搜索参数查询团队成员信息，并可选地构建额外的资源信息。
		// 参数:
		//
		//	ctx - 上下文，用于传递请求范围的上下文信息。
		//	params - 搜索参数，用于指定查询的条件。
		//	makeResource - 指示是否构建额外的资源信息。
		//
		// 返回值:
		//
		//	*co_model.TeamMemberViewListRes - 包含团队成员列表的响应对象。
		//	error - 错误信息，如果执行过程中发生错误。
		QueryTeamMemberList(ctx context.Context, params *base_model.SearchParams, makeResource bool) (*co_model.TeamViewMemberListRes, error)
	}
)

var (
	localCompanyView    ICompanyView
	localEmployeeView   IEmployeeView
	localFdAccountView  IFdAccountView
	localFdBillsView    IFdBillsView
	localFdBankCardView IFdBankCardView
	localFdInvoiceView  IFdInvoiceView
	localFdRechargeView IFdRechargeView
	localTeamView       ITeamView
	localTeamMemberView ITeamMemberView
)

func CompanyView() ICompanyView {
	if localCompanyView == nil {
		panic("implement not found for interface ICompanyView, forgot register?")
	}
	return localCompanyView
}

func RegisterCompanyView(i ICompanyView) {
	localCompanyView = i
}

func EmployeeView() IEmployeeView {
	if localEmployeeView == nil {
		panic("implement not found for interface IEmployeeView, forgot register?")
	}
	return localEmployeeView
}

func RegisterEmployeeView(i IEmployeeView) {
	localEmployeeView = i
}

func FdAccountView() IFdAccountView {
	if localFdAccountView == nil {
		panic("implement not found for interface IFdAccountView, forgot register?")
	}
	return localFdAccountView
}

func RegisterFdAccountView(i IFdAccountView) {
	localFdAccountView = i
}

func FdBillsView() IFdBillsView {
	if localFdBillsView == nil {
		panic("implement not found for interface IFdBillsView, forgot register?")
	}
	return localFdBillsView
}

func RegisterFdBillsView(i IFdBillsView) {
	localFdBillsView = i
}

func FdBankCardView() IFdBankCardView {
	if localFdBankCardView == nil {
		panic("implement not found for interface IFdBankCardView, forgot register?")
	}
	return localFdBankCardView
}

func RegisterFdBankCardView(i IFdBankCardView) {
	localFdBankCardView = i
}

func FdInvoiceView() IFdInvoiceView {
	if localFdInvoiceView == nil {
		panic("implement not found for interface IFdInvoiceView, forgot register?")
	}
	return localFdInvoiceView
}

func RegisterFdInvoiceView(i IFdInvoiceView) {
	localFdInvoiceView = i
}

func FdRechargeView() IFdRechargeView {
	if localFdRechargeView == nil {
		panic("implement not found for interface IFdRechargeView, forgot register?")
	}
	return localFdRechargeView
}

func RegisterFdRechargeView(i IFdRechargeView) {
	localFdRechargeView = i
}

func TeamView() ITeamView {
	if localTeamView == nil {
		panic("implement not found for interface ITeamView, forgot register?")
	}
	return localTeamView
}

func RegisterTeamView(i ITeamView) {
	localTeamView = i
}

func TeamMemberView() ITeamMemberView {
	if localTeamMemberView == nil {
		panic("implement not found for interface ITeamMemberView, forgot register?")
	}
	return localTeamMemberView
}

func RegisterTeamMemberView(i ITeamMemberView) {
	localTeamMemberView = i
}
