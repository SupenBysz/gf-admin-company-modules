// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao/internal"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
)

type FdInvoiceDetailDao = dao_interface.TIDao[internal.FdInvoiceDetailColumns]

func NewFdInvoiceDetail(dao ...dao_interface.IDao) FdInvoiceDetailDao {
	return (FdInvoiceDetailDao)(internal.NewFdInvoiceDetailDao(dao...))
}

var (
	FdInvoiceDetail = NewFdInvoiceDetail()
)
