package co_dao

type XDao struct {
	Company    Company
	Employee   CompanyEmployee
	Team       CompanyTeam
	TeamMember CompanyTeamMember

	FdAccount       FdAccount
	FdAccountBill   FdAccountBill
	FdInvoice       FdInvoice
	FdInvoiceDetail FdInvoiceDetail
	FdCurrency      FdCurrency
	FdBankCard      FdBankCard
}
