package co_dao

type XDao struct {
	Company    CompanyDao
	Employee   CompanyEmployeeDao
	Team       CompanyTeamDao
	TeamMember CompanyTeamMemberDao

	FdAccount       FdAccountDao
	FdAccountBills  FdAccountBillsDao
	FdInvoice       FdInvoiceDao
	FdInvoiceDetail FdInvoiceDetailDao
	FdBankCard      FdBankCardDao
	FdAccountDetail FdAccountDetailDao
	FdRecharge      FdRechargeDao
}
