package co_dao

type XDao struct {
	Company    CompanyDao
	Employee   CompanyEmployeeDao
	Team       CompanyTeamDao
	TeamMember CompanyTeamMemberDao

	FdAccount       FdAccountDao
	FdAccountBill   FdAccountBillDao
	FdInvoice       FdInvoiceDao
	FdInvoiceDetail FdInvoiceDetailDao
	FdCurrency      FdCurrencyDao
	FdBankCard      FdBankCardDao
	FdAccountDetail FdAccountDetailDao
}
