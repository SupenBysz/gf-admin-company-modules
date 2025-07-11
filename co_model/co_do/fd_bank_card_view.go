// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FdBankCardView is the golang structure of table co_fd_bank_card_view for DAO operations like Where/Data.
type FdBankCardView struct {
	g.Meta        `orm:"table:co_fd_bank_card_view, do:true"`
	Id            interface{} //
	BankName      interface{} //
	CardType      interface{} //
	CardNumber    interface{} //
	ExpiredAt     *gtime.Time //
	HolderName    interface{} //
	BankOfAccount interface{} //
	State         interface{} //
	Remark        interface{} //
	UserId        interface{} //
	CreatedAt     *gtime.Time //
	CreatedBy     interface{} //
	UpdatedAt     *gtime.Time //
	UpdatedBy     interface{} //
	DeletedAt     *gtime.Time //
	DeletedBy     interface{} //
	CompanyType   interface{} //
}
