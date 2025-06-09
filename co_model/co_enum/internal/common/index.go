package common

type common struct {
	AppealState                 appealState
	CompanyCommissionMode       companyCommissionMode
	EmployeeCommissionMode      employeeCommissionMode
	EmployeeCommissionLevelMode employeeCommissionLevelMode
}

var Common = common{
	AppealState:                 AppealState,
	CompanyCommissionMode:       CompanyCommissionMode,
	EmployeeCommissionMode:      EmployeeCommissionMode,
	EmployeeCommissionLevelMode: EmployeeCommissionLevelMode,
}
