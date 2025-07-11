// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyView is the golang structure for table company_view.
type CompanyView struct {
	Id             int64       `json:"id"             orm:"id"              description:""`
	Name           string      `json:"name"           orm:"name"            description:""`
	ContactName    string      `json:"contactName"    orm:"contact_name"    description:""`
	ContactMobile  string      `json:"contactMobile"  orm:"contact_mobile"  description:""`
	UserId         int64       `json:"userId"         orm:"user_id"         description:""`
	State          int         `json:"state"          orm:"state"           description:""`
	Remark         string      `json:"remark"         orm:"remark"          description:""`
	CreatedBy      int64       `json:"createdBy"      orm:"created_by"      description:""`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
	UpdatedBy      int64       `json:"updatedBy"      orm:"updated_by"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""`
	DeletedBy      int64       `json:"deletedBy"      orm:"deleted_by"      description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""`
	ParentId       int64       `json:"parentId"       orm:"parent_id"       description:""`
	Address        string      `json:"address"        orm:"address"         description:""`
	LicenseId      int64       `json:"licenseId"      orm:"license_id"      description:""`
	LicenseState   int         `json:"licenseState"   orm:"license_state"   description:""`
	LogoId         int64       `json:"logoId"         orm:"logo_id"         description:""`
	StartLevel     int         `json:"startLevel"     orm:"start_level"     description:""`
	CountryCode    string      `json:"countryCode"    orm:"country_code"    description:""`
	Region         string      `json:"region"         orm:"region"          description:""`
	Score          int         `json:"score"          orm:"score"           description:""`
	CommissionRate int         `json:"commissionRate" orm:"commission_rate" description:""`
	CompanyType    int         `json:"companyType"    orm:"company_type"    description:""`
}
