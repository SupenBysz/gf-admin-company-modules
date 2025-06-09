package co_model

type ClientConfig struct {
	XClientToken string `json:"identifier"     dc:"客户端token"`
	// 是否允许跨级创建子公司单位
	AllowSkipLevelCreateCompany bool `json:"allow_skip_level_create_company" dc:"是否允许跨级创建子公司单位"`
	// 佣金模式，仅作用于公司超管财务账户，即公司财务账户，一般场景为机构、代理、子公司、加盟等关系业务场景）
	// 仅作用于公司主体，0:不启用佣金机制，1：相对上级佣金百分比，2：相对交易金额百分比,3:相较于交易佣金百分比。
	// 注意：一旦选定模式后，在一个统计周期内不能修改，否则会引起统计错误，在下一个统计周期才能正常
	// co_enum.CompanyCommissionMode
	CompanyCommissionModel int `json:"company_commission_model"     dc:"公司佣金模式"`
	// 佣金分配级别，仅作用于公司超管财务账户，即公司财务账户），1：一级佣金，2：二级佣金，3：三级佣金，……以此类推
	CompanyCommissionAllocationLevel int `json:"company_commission_allocation_level" dc:"公司佣金分配级别"`
	// 员工提成模式，仅作用于员工，0不启用提成机制，1:相较于上级，2:相较于相较于交易金额百分比，3:相较于交易佣金百分比，
	// 注意：一旦选定模式后，在一个统计周期内不能修改，否则会引起统计错误，在下一个统计周期才能正常
	// co_enum.EmployeeCommissionMode
	EmployeeCommissionModel int `json:"employee_commission_model"     dc:"员工提成模式"`
	// 员工提成分配等级模式，仅作用于员工：1师徒/邀请模式、2部门/团队/小组模式、3角色模式，
	// 注意：一旦选定模式后，在一个统计周期内不能修改，否则会引起统计错误，在下一个统计周期才能正常
	// co_enum.EmployeeCommissionLevelMode
	EmployeeCommissionLevelMode int `json:"employee_commission_level_mode"     dc:"员工提成分配等级模式"`
	// 员工提成分配级别（仅作用于员工）
	EmployeeCommissionAllocationLevel int `json:"employee_commission_allocation_level" dc:"员工提成分配级别"`
}
