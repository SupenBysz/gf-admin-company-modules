package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/gogf/gf/v2/database/gdb"
)

type TableName struct {
	Company         string `p:"company" dc:"公司表名"`
	Employee        string `p:"employee" dc:"员工表名"`
	Team            string `p:"team" dc:"团队表名"`
	TeamMember      string `p:"teamMember" dc:"团队成员表名"`
	FdAccount       string `p:"fdAccount" dc:"财务账户表名"`
	FdAccountBill   string `p:"fdAccountBill" dc:"财务账单表名"`
	FdBankCard      string `p:"fdBankCard" dc:"财务银行卡表名"`
	FdCurrency      string `p:"fdCurrency" dc:"财务结算货币单位表名"`
	FdInvoice       string `p:"fdInvoice" dc:"财务发票抬头表名"`
	FdInvoiceDetail string `p:"fdInvoiceDetail" dc:"财务发票明细表名"`
}

type Identifier struct {
	Company         string `p:"company" dc:"公司标识符"`
	Employee        string `p:"employee" dc:"员工标识符"`
	Team            string `p:"team" dc:"团队标识符"`
	TeamMember      string `p:"teamMember" dc:"团队成员标识符"`
	FdAccount       string `p:"fdAccount" dc:"财务账户标识符"`
	FdAccountBill   string `p:"fdAccountBill" dc:"财务账单标识符"`
	FdBankCard      string `p:"fdBankCard" dc:"财务银行卡标识符"`
	FdCurrency      string `p:"fdCurrency" dc:"财务结算货币单位标识符"`
	FdInvoice       string `p:"fdInvoice" dc:"财务发票抬头标识符"`
	FdInvoiceDetail string `p:"fdInvoiceDetail" dc:"财务发票明细标识符"`
}

type Config struct {
	DB                             gdb.DB            `p:"-" dc:"数据库连接"`
	AllowEmptyNo                   bool              `p:"allowEmptyNo" dc:"允许员工工号为空" default:"false"`
	IsCreateDefaultEmployeeAndRole bool              `p:"isCreateDefaultEmployeeAndRole" dc:"是否创建默认员工和角色"`
	HardDeleteWaitAt               int64             `p:"hardDeleteWaitAt" dc:"硬删除等待时限,单位/小时" default:"12"`
	KeyIndex                       string            `p:"keyIndex" dc:"配置索引"`
	I18nName                       string            `p:"i18NName" dc:"i18n文件名称"`
	RoutePrefix                    string            `p:"routePrefix" dc:"路由前缀"`
	StoragePath                    string            `p:"storagePath" dc:"资源存储路径"`
	UserType                       sys_enum.UserType `p:"userType" dc:"用户类型"`
	Identifier                     Identifier        `p:"identifier" dc:"标识符"`
	TableName                      TableName         `p:"tableName" dc:"模块表名"`
}
