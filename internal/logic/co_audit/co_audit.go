package co_audit

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_hook"
	"github.com/SupenBysz/gf-admin-company-modules/co_service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type hookInfo sys_model.KeyValueT[int64, co_hook.AuditHookInfo]

type sAudit struct {
	conf    gdb.CacheOption
	hookArr []hookInfo
}

func init() {
	co_service.RegisterAudit(NewAudit())
}

func NewAudit() *sAudit {
	return &sAudit{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
		hookArr: make([]hookInfo, 0),
	}
}

// InstallHook 安装Hook
func (s *sAudit) InstallHook(state co_enum.AuditEvent, category int, hookFunc co_hook.AuditHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: co_hook.AuditHookInfo{Key: state, Value: hookFunc, Category: category}}
	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sAudit) UnInstallHook(savedHookId int64) {
	newFuncArr := make([]hookInfo, 0)
	for _, item := range s.hookArr {
		if item.Key != savedHookId {
			newFuncArr = append(newFuncArr, item)
			continue
		}
	}
	s.hookArr = newFuncArr
}

// CleanAllHook 清除所有Hook
func (s *sAudit) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

// QueryAuditList 获取审核信息列表
func (s *sAudit) QueryAuditList(ctx context.Context, filter *base_model.SearchParams) (*co_model.AuditListRes, error) {
	if &filter.Pagination == nil {
		filter.Pagination = base_model.Pagination{
			PageNum:  1,
			PageSize: 20,
		}
	}

	filter.Filter = append(filter.Filter, base_model.FilterInfo{
		Field:       co_dao.Audit.Columns().Id,
		Where:       ">",
		IsOrWhere:   false,
		Value:       0,
		IsNullValue: false,
	})

	result, err := daoctl.Query[co_entity.Audit](co_dao.Audit.Ctx(ctx), filter, true)

	auditList := make([]co_entity.Audit, 0)
	for _, item := range result.Records {
		// 解析json字符串
		auditJsonData := item.AuditData
		auditData := co_model.AuditLicense{}
		gjson.DecodeTo(auditJsonData, &auditData)

		// 还未审核的图片从缓存中寻找  0 缓存  1 数据库

		// 将路径id换成可访问图片的url
		//if gstr.IsNumeric(auditData.IdcardFrontPath) {
		//	//auditData.IdcardFrontPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardFrontPath))
		//	auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardFrontPath)
		//	fmt.Println("身份证：", auditData.IdcardFrontPath)
		//}

		// 将路径src换成可访问图片的url
		{
			if gfile.IsFile(auditData.IdcardFrontPath) {
				//auditData.IdcardFrontPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardFrontPath))
				auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardFrontPath)
				fmt.Println("身份证：", auditData.IdcardFrontPath)

			}
			if gfile.IsFile(auditData.IdcardBackPath) {
				auditData.IdcardBackPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardBackPath)
				fmt.Println("身份证：", auditData.IdcardBackPath)
			}
			if gfile.IsFile(auditData.BusinessLicenseLegalPath) {
				auditData.BusinessLicenseLegalPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.BusinessLicenseLegalPath)
			}
			if gfile.IsFile(auditData.BusinessLicensePath) {
				auditData.BusinessLicensePath = sys_service.File().MakeFileUrlByPath(ctx, auditData.BusinessLicensePath)
			}
		}
		if err != nil {
			return nil, err
		}

		// 重新赋值
		rest := co_entity.Audit{}
		gconv.Struct(item, &rest)
		rest.AuditData = gjson.MustEncodeString(auditData)

		auditList = append(auditList, rest)
	}

	result.Records = auditList
	return (*co_model.AuditListRes)(result), err
}

// GetAuditById 根据ID获取审核信息
func (s *sAudit) GetAuditById(ctx context.Context, id int64) *co_entity.Audit {

	result, err := daoctl.GetByIdWithError[co_entity.Audit](co_dao.Audit.Ctx(ctx), id)

	if err != nil {
		return nil
	}

	// 解析json字符串
	auditData := co_model.AuditLicense{}
	gjson.DecodeTo(result.AuditData, &auditData)

	// 还未审核的图片从缓存中寻找  0 缓存  1 数据库

	// 将路径id换成可访问图片的url
	{
		if gstr.IsNumeric(auditData.IdcardFrontPath) {
			auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardFrontPath)
		}
		if gstr.IsNumeric(auditData.IdcardBackPath) {
			auditData.IdcardBackPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardBackPath)
		}
		if gstr.IsNumeric(auditData.BusinessLicenseLegalPath) {
			auditData.BusinessLicenseLegalPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.BusinessLicenseLegalPath)
		}
		if gstr.IsNumeric(auditData.BusinessLicensePath) {
			auditData.BusinessLicensePath = sys_service.File().MakeFileUrlByPath(ctx, auditData.BusinessLicensePath)
		}
	}
	// fmt.Println(auditData.IdcardFrontPath + " --- " + auditData.IdcardBackPath + " --- " + auditData.BusinessLicensePath + " --- " + auditData.BusinessLicenseLegalPath)

	// 重新赋值  将id转为可访问路径
	result.AuditData = gjson.MustEncodeString(auditData)

	return result

}

// Audit存，将userId 和 上传id从缓存中读取出，然后将file.Src作为身份证、营业执照字段的值，  idCardPath：文件id  idCardPath：/tmp/upload/20230413/20230413/6504378708918341/crvld008yix5scyuio.jpeg

// Audit取，拿出路劲转成带签名的url，

// GetAuditByLatestUnionMainId 获取最新的业务主体审核信息
func (s *sAudit) GetAuditByLatestUnionMainId(ctx context.Context, unionMainId int64) *co_entity.Audit {
	result := co_entity.Audit{}
	err := co_dao.Audit.Ctx(ctx).Where(co_do.Audit{UnionMainId: unionMainId}).OrderDesc(co_dao.Audit.Columns().CreatedAt).Limit(1).Scan(&result)
	if err != nil {
		return nil
	}

	// 将路径src换成可访问图片的url
	auditData := co_model.AuditLicense{}
	gjson.DecodeTo(result.AuditData, &auditData)

	{
		if gfile.IsFile(auditData.IdcardFrontPath) {
			//auditData.IdcardFrontPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardFrontPath))
			auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardFrontPath)

		}
		if gfile.IsFile(auditData.IdcardBackPath) {
			auditData.IdcardBackPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardBackPath)
		}
		if gfile.IsFile(auditData.BusinessLicenseLegalPath) {
			auditData.BusinessLicenseLegalPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.BusinessLicenseLegalPath)
		}
		if gfile.IsFile(auditData.BusinessLicensePath) {
			auditData.BusinessLicensePath = sys_service.File().MakeFileUrlByPath(ctx, auditData.BusinessLicensePath)
		}
	}

	result.AuditData = gjson.MustEncodeString(auditData)

	return &result
}

// CreateAudit 创建审核信息
func (s *sAudit) CreateAudit(ctx context.Context, info co_model.CreateAudit) (*co_entity.Audit, error) {
	// 校验参数
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return nil, err
	}

	// 如果业务没有设置审核服务时限则加载默认设置
	if info.ExpireAt == nil {
		day := g.Cfg().MustGet(ctx, "service.auditExpireDay.default", 7).Float64()
		info.ExpireAt = gtime.Now().Add(time.Duration(time.Hour.Seconds() * 24 * day))
	}

	data := co_entity.Audit{}
	audit := co_entity.Audit{}
	gconv.Struct(info, &data)

	err := co_dao.Audit.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		{
			// 查询当前关联业务ID是否有审核记录
			err := co_dao.Audit.Ctx(ctx).Where(co_do.Audit{
				UnionMainId: info.UnionMainId,
				Category:    info.Category,
			}).Scan(&audit)
			if err != nil && err != sql.ErrNoRows {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "查询校验信息失败", co_dao.Audit.Table())
			}
			// 如果当前有审核记录，则转存入历史记录中，并删除当前申请记录，避免后续步骤创建记录时重复导致的失败
			if audit.Id > 0 {
				historyItems := make([]co_entity.Audit, 0)
				g.Try(ctx, func(ctx context.Context) {
					// 判断历史记录是否为空
					if audit.HistoryItems != "" {
						// 解码json字符串为列表为切片对象
						gjson.DecodeTo(audit.HistoryItems, &historyItems)
						// 清空记录中的历史记录，便于后面压入记录中导致冗余的历史记录
						audit.HistoryItems = ""
					}
					// 判断当前审核状态是否审核中，只对已审核的记录压入历史记录中
					if audit.State != 0 {
						// 将记录压入列表
						historyItems = append(historyItems, audit)
					}
					// 编码切片列表为JSON字符串
					data.HistoryItems = gjson.MustEncodeString(historyItems)
				})

				_, err = co_dao.Audit.Ctx(ctx).Delete(co_do.Audit{Id: audit.Id})
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "保存审核前置信息失败", co_dao.Audit.Table())
				}
			}
		}

		data.Id = idgen.NextId()
		data.CreatedAt = gtime.Now()

		_, err := co_dao.Audit.Ctx(ctx).Data(data).Insert()

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "保存审核信息失败", co_dao.Audit.Table())
		}

		stateType := co_enum.Audit.Event.Created
		if info.Id > 0 {
			stateType = co_enum.Audit.Event.ReSubmit
		}

		for _, hook := range s.hookArr {
			// 判断注入的Hook业务类型是否一致
			if hook.Value.Category&info.Category == info.Category {
				// 业务类型一致则调用注入的Hook函数
				err = hook.Value.Value(ctx, stateType, data)
			}
			gerror.NewCode(gcode.CodeInvalidConfiguration, "")
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "创建审核信息失败", co_dao.Audit.Table())
	}
	return s.GetAuditById(ctx, data.Id), nil
}

// UpdateAudit 处理审核信息
func (s *sAudit) UpdateAudit(ctx context.Context, id int64, state int, reply string, auditUserId int64) (bool, error) {
	if state == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核行为类型错误", co_dao.Audit.Table())
	}

	if state == -1 && reply == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", co_dao.Audit.Table())
	}

	info := s.GetAuditById(ctx, id)
	if info == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "ID参数错误", co_dao.Audit.Table())
	}

	if info.State != 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "禁止单次申请重复审核业务", co_dao.Audit.Table())
	}

	err := co_dao.Audit.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := co_dao.Audit.Ctx(ctx).OmitNilData().Data(co_do.Audit{
			State:        state,
			Reply:        reply,
			AuditReplyAt: gtime.Now(),
			AuditUserId:  auditUserId,
		}).Where(co_do.Audit{
			Id:          info.Id,
			UnionMainId: info.UnionMainId,
			Category:    info.Category,
		}).Update()

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核信息保存失败", co_dao.Audit.Table())
		}

		data := s.GetAuditById(ctx, info.Id)
		if data == nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "获取审核信息失败", co_dao.Audit.Table())
		}

		// 审核通过
		if (data.State & co_enum.Audit.Action.Approve.Code()) == co_enum.Audit.Action.Approve.Code() {
			// 创建主体资质
			license := co_model.AuditLicense{}
			gjson.DecodeTo(data.AuditData, &license)

			licenseRes, err := co_service.License().CreateLicense(ctx, license.License)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核通过后主体资质创建失败", co_dao.License.Table())
			}

			// 设置主体资质的审核编号
			ret, err := co_service.License().SetLicenseAuditNumber(ctx, licenseRes.Id, gconv.String(data.Id))
			if err != nil || ret == false {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "", co_dao.License.Table())
			}
		}

		for _, hook := range s.hookArr {
			// 判断注入的Hook业务类型是否一致
			if hook.Value.Category&info.Category == info.Category {
				// 业务类型一致则调用注入的Hook函数
				err = hook.Value.Value(ctx, co_enum.Audit.Event.ExecAudit, *data)
			}
			gerror.NewCode(gcode.CodeInvalidConfiguration, "")
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err == nil, err
}
