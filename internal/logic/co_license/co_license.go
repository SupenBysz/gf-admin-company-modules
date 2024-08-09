package co_license

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/SupenBysz/gf-admin-company-modules/utility/co_funs"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/masker"
	"time"
)

type sLicense struct {
	conf gdb.CacheOption
}

func init() {
	co_service.RegisterLicense(NewLicense())
}

func NewLicense() *sLicense {
	result := &sLicense{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}

	// 订阅审核Hook,审核通过添加主体资质主体信息
	sys_service.SysAudit().InstallHook(sys_enum.Audit.Action.Approve, sys_enum.Audit.Category.CompanyLicenseAudit.Code(), result.AuditChange)

	// 订阅审核数据获取Hook, 将审核数据渲染成主体资质然后进行输出
	sys_service.SysAudit().InstallHook(sys_enum.Audit.Action.Approve, sys_enum.Audit.Category.CompanyLicenseAudit.Code(), result.GetAuditData)

	return result
}

// GetAuditData 订阅审核数据获取Hook, 将审核数据渲染成主体资质然后进行输出  (数据处理渲染)
func (s *sLicense) GetAuditData(ctx context.Context, auditEvent sys_enum.AuditEvent, info *sys_entity.SysAudit) error {
	//  处理审核
	if info == nil {
		return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核数据为空", "Audit")
	}

	if (auditEvent.Code() & sys_enum.Audit.Event.GetAuditData.Code()) == sys_enum.Audit.Event.GetAuditData.Code() {
		if (info.Category & sys_enum.Audit.Category.CompanyLicenseAudit.Code()) == sys_enum.Audit.Category.CompanyLicenseAudit.Code() {

			auditData := co_model.AuditLicense{}

			//解析json字符串
			//decode, _ := gjson.Decode(info.AuditData)
			gjson.DecodeTo(info.AuditData, &auditData)
			if auditData.IdcardNo == "" { // 说明不是默认结构，业务层自己封装了结构
				return nil
			}

			// 将路径id换成可访问图片的url
			{
				tempIdcardFrontPath := ""
				if gstr.IsNumeric(auditData.IdcardFrontPath) {
					// TODO 没审核通过的文件是没有保存的， 如何获取到，（从缓存获取）
					if uploadFile, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(auditData.IdcardFrontPath), auditData.UserId); err == nil && uploadFile != nil {
						tempIdcardFrontPath = uploadFile.Src
					}
				}
				if tempIdcardFrontPath != "" {
					auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, tempIdcardFrontPath)
				}

			}
			{
				tempIdcardBackPath := ""
				if gstr.IsNumeric(auditData.IdcardBackPath) {
					// TODO 没审核通过的文件是没有保存的， 如何获取到，（从缓存获取）
					if uploadFile, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(auditData.IdcardBackPath), auditData.UserId); err == nil && uploadFile != nil {
						tempIdcardBackPath = uploadFile.Src
					}
				}
				if tempIdcardBackPath != "" {
					auditData.IdcardBackPath = sys_service.File().MakeFileUrlByPath(ctx, tempIdcardBackPath)
				}

			}
			{
				tempBusinessLicensePath := ""
				if gstr.IsNumeric(auditData.BusinessLicensePath) {
					// TODO 没审核通过的文件是没有保存的， 如何获取到，（从缓存获取）
					if uploadFile, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(auditData.BusinessLicensePath), auditData.UserId); err == nil && uploadFile != nil {
						tempBusinessLicensePath = uploadFile.Src
					}
				}
				if tempBusinessLicensePath != "" {
					auditData.BusinessLicensePath = sys_service.File().MakeFileUrlByPath(ctx, tempBusinessLicensePath)
				}

			}

			if auditData.IdcardNo != "" { // 说明是默认结构
				// 重新赋值  将id转为可访问路径
				info.AuditData = gjson.MustEncodeString(auditData)
			} else { // 业务层自己自定义的审核机构，业务层自己解析即可

			}

		}
	}
	return nil
}

// AuditChange 审核成功的处理逻辑 Hook （数据存储）
func (s *sLicense) AuditChange(ctx context.Context, auditEvent sys_enum.AuditEvent, info *sys_entity.SysAudit) error {
	//  处理审核
	if (auditEvent.Code() & sys_enum.Audit.Event.ExecAudit.Code()) == sys_enum.Audit.Event.ExecAudit.Code() {
		// 审核通过
		if (info.State & sys_enum.Audit.Action.Approve.Code()) == sys_enum.Audit.Action.Approve.Code() {
			// 创建主体资质
			//license := co_model.License{}
			license := co_model.AuditLicense{}
			gjson.DecodeTo(info.AuditData, &license)
			if license.IdcardNo == "" { // 业务层自己处理审核通过的逻辑
				return nil
			}
			licenseRes, err := co_service.License().CreateLicense(ctx, license)
			//licenseRes, err := sys_service.SysPersonLicense().CreateLicense(ctx, license)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核通过后主体资质创建失败", co_dao.License.Table())
			}

			// 设置主体资质的审核编号
			ret, err := co_service.License().SetLicenseAuditNumber(ctx, licenseRes.Id, gconv.String(info.Id))
			if err != nil || ret == false {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "", co_dao.License.Table())
			}
		}
	}

	return nil
}

// GetLicenseById 根据ID获取主体认证|信息
func (s *sLicense) GetLicenseById(ctx context.Context, id int64) (*co_entity.License, error) {
	data := co_entity.License{}
	err := co_dao.License.Ctx(ctx).Scan(&data, co_do.License{Id: id})

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "主体信息不存在", co_dao.License.Table())
	}

	// 需要将持久化的文件ID替换成可访问的接口URL
	s.buildURL(ctx, &data)

	return &data, nil
}

// QueryLicenseList 查询主体认证|列表
func (s *sLicense) QueryLicenseList(ctx context.Context, search base_model.SearchParams) (*co_model.LicenseListRes, error) {
	result, err := daoctl.Query[co_entity.License](co_dao.License.Ctx(ctx), &search, false)
	if err != nil {
		return &co_model.LicenseListRes{}, err
	}

	response := co_model.LicenseListRes{}
	for _, record := range result.Records {
		// 需要将持久化的文件ID替换成可访问的接口URL
		s.buildURL(ctx, &record)

		response.Records = append(response.Records, record)
	}

	response.PaginationRes = result.PaginationRes

	return &response, err
	//return (*co_model.LicenseListRes)(result), err
}

// CreateLicense 新增主体资质|信息
func (s *sLicense) CreateLicense(ctx context.Context, info co_model.AuditLicense) (*co_entity.License, error) {
	//result := co_model.AuditLicense{}
	result := co_entity.License{}
	gconv.Struct(info.License, &result)

	if info.LicenseId == 0 {
		result.Id = idgen.NextId()
	} else {
		result.Id = info.LicenseId
	}

	//result.State = 0
	//result.AuthType = 0
	result.CreatedAt = gtime.Now()

	{
		_, err := co_funs.CheckLicenseFiles(ctx, info, &result)
		if err != nil {
			return nil, err
		}
	}

	{
		// 创建主体信息
		_, err := co_dao.License.Ctx(ctx).Insert(result)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "新增主体信息失败", co_dao.License.Table())
		}

		if err != nil {
			return nil, err
		}
	}

	// 需要将持久化的文件ID替换成可访问的接口URL
	s.buildURL(ctx, &result)

	return &result, nil
}

// UpdateLicense 更新主体认证，如果是已经通过的认证，需要重新认证通过后才生效|信息
func (s *sLicense) UpdateLicense(ctx context.Context, info co_model.AuditLicense, id int64) (*co_entity.License, error) {
	//data, err := s.GetLicenseById(ctx, id)
	data := co_entity.License{}
	err := co_dao.License.Ctx(ctx).Scan(&data, co_do.License{Id: id})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，主体信息不存在", co_dao.License.Table())
	}

	if data.State == -1 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "操作是不，主体信息被冻结，禁止修改"), "", co_dao.License.Table())
	}

	newData := co_do.License{}

	gconv.Struct(info.License, &newData)

	{
		_, err := co_funs.CheckLicenseFiles(ctx, info, &newData)
		if err != nil {
			return nil, err
		}
	}

	err = co_dao.License.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		newAudit := sys_model.Audit{
			Id:          idgen.NextId(),
			State:       0,
			UnionMainId: data.Id,
			Category:    1,
			AuditData:   gjson.MustEncodeString(data),
			ExpireAt:    gtime.Now().Add(time.Hour * 24 * 7),
		}

		{
			audit := sys_service.SysAudit().GetAuditById(ctx, data.LatestAuditLogId)
			// 未审核通过的主体资质，直接更改待审核的资质信息
			if audit != nil && audit.State == 0 {
				_, err := tx.Ctx(ctx).Model(co_dao.License.Table()).Where(co_do.License{Id: id}).OmitNil().Save(&newData)
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，更新主体信息失败", co_dao.License.Table())
				}

				// 更新待审核的审核信息
				newAudit.Id = audit.Id
				// TODO
				//co_service.Audit().UpdateAudit()
				_, err = sys_dao.SysAudit.Ctx(ctx).Data(newAudit).Where(sys_do.SysAudit{Id: audit.Id}).Update()
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "更新审核信息失败", co_dao.License.Table())
				}
				return nil
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return s.GetLicenseById(ctx, id)
}

// GetLicenseByLatestAuditId 获取最新的审核记录Id获取资质信息
func (s *sLicense) GetLicenseByLatestAuditId(ctx context.Context, auditId int64) *co_entity.License {
	result := co_entity.License{}
	err := co_dao.License.Ctx(ctx).Where(co_do.License{LatestAuditLogId: auditId}).OrderDesc(co_dao.License.Columns().CreatedAt).Limit(1).Scan(&result)
	if err != nil {
		return nil
	}

	// 需要将持久化的文件ID替换成可访问的接口URL
	s.buildURL(ctx, &result)

	return &result
}

// SetLicenseState 设置主体信息状态
func (s *sLicense) SetLicenseState(ctx context.Context, id int64, state int) (bool, error) {
	//_, err := s.GetLicenseById(ctx, id)
	data := co_entity.License{}
	err := co_dao.License.Ctx(ctx).Scan(&data, co_do.License{Id: id})
	if err != nil {
		return false, err
	}

	_, err = co_dao.License.Ctx(ctx).Data(co_do.License{State: state}).OmitNilData().Where(co_do.License{Id: id}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "更新主体状态信息失败", co_dao.License.Table())
	}

	return true, nil
}

// SetLicenseAuditNumber 设置主体神审核编号
func (s *sLicense) SetLicenseAuditNumber(ctx context.Context, id int64, auditNumber string) (bool, error) {
	//_, err := s.GetLicenseById(ctx, id)
	data := co_entity.License{}
	err := co_dao.License.Ctx(ctx).Scan(&data, co_do.License{Id: id})
	if err != nil {
		return false, err
	}

	_, err = co_dao.License.Ctx(ctx).Data(co_do.License{LatestAuditLogId: auditNumber}).OmitNilData().Where(co_do.License{Id: id}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "更新主体证照审核编号失败", co_dao.License.Table())
	}
	return true, nil
}

// DeleteLicense 删除主体
func (s *sLicense) DeleteLicense(ctx context.Context, id int64, flag bool) (bool, error) {
	return false, nil
}

// UpdateLicenseAuditLogId 设置主体资质关联的审核ID
func (s *sLicense) UpdateLicenseAuditLogId(ctx context.Context, id int64, latestAuditLogId int64) (bool, error) {
	auditLog := sys_service.SysAudit().GetAuditById(ctx, latestAuditLogId)
	if nil == auditLog {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "资质信息校验失败", co_dao.License.Table())
	}

	audit := co_model.AuditLicense{}

	err := gjson.DecodeTo(auditLog.AuditData, &audit)

	if err != nil || audit.LicenseId != id {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "资质校验失败", co_dao.License.Table())
	}

	// 构建资质对象
	license := co_entity.License{}
	// 加载资质信息
	err = co_dao.License.Ctx(ctx).Scan(&license, co_do.License{Id: id})
	// 如果资质不存在则无需更新，直接返回
	if err == sql.ErrNoRows {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	if license.BusinessLicenseCreditCode != audit.BusinessLicenseCreditCode {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "禁止修改组织机构代码", co_dao.License.Table())
	}

	// 将新创建的主体认证信息关联至主体
	affected, err := daoctl.UpdateWithError(co_dao.License.Ctx(ctx).
		Data(co_do.License{LatestAuditLogId: latestAuditLogId}).
		Where(co_do.License{Id: id}))

	return affected > 0, err
}

// Masker 资质信息脱敏
func (s *sLicense) Masker(license *co_entity.License) *co_entity.License {
	license.PersonContactMobile = masker.MaskString(license.PersonContactMobile, masker.MaskPhone)
	license.IdcardNo = masker.MaskString(license.IdcardNo, masker.IDCard)
	license.BusinessLicensePath = ""
	license.BusinessLicenseLegalPath = ""
	license.IdcardFrontPath = ""
	license.IdcardBackPath = ""

	return license
}

// buildURL 将文件id替换成可访问的URL
func (s *sLicense) buildURL(ctx context.Context, data *co_entity.License) {

	{
		if gstr.IsNumeric(data.BusinessLicensePath) {
			data.BusinessLicensePath = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.BusinessLicensePath))
		}

	}
	{
		if gstr.IsNumeric(data.BusinessLicenseLegalPath) {
			data.BusinessLicenseLegalPath = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.BusinessLicenseLegalPath))
		}

	}
	{
		if gstr.IsNumeric(data.IdcardFrontPath) {
			data.IdcardFrontPath = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.IdcardFrontPath))
		}

	}
	{
		if gstr.IsNumeric(data.IdcardBackPath) {
			data.IdcardBackPath = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.IdcardBackPath))
		}

	}

	//return data
}
