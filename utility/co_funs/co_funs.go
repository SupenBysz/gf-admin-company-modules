package co_funs

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func CheckLicenseFiles[T co_entity.License | co_do.License](ctx context.Context, info co_model.AuditLicense, data *T) (response *T, err error) {
	newData := &co_entity.License{}
	gconv.Struct(data, newData)

	{
		//userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id
		userId := info.UserId

		// 用户资源文件夹
		userFolder := "./resources/license/" + gconv.String(newData.Id)
		fileAt := gtime.Now().Format("YmdHis")
		if !gfile.Exists(info.IdcardFrontPath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.IdcardFrontPath), userId, "请上传身份证头像面")
			if err != nil {
				return nil, err
			}
			// 保存员工身份证头像面
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/idCard/front_"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}

			//  注意：实际存储的License 需要存储持久化后的文件ID，而不是路径
			newData.IdcardFrontPath = gconv.String(fileInfo.Id)
		}

		if !gfile.Exists(info.IdcardBackPath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.IdcardBackPath), userId, "请上传身份证国徽面")
			if err != nil {
				return nil, err
			}
			// 保存员工身份证国徽面
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/idCard/back_"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}

			//  注意：实际存储的License 需要存储持久化后的文件ID，而不是路径
			newData.IdcardBackPath = gconv.String(fileInfo.Id)
		}

		if !gfile.Exists(info.BusinessLicensePath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.BusinessLicensePath), userId, "请上传营业执照图片")
			if err != nil {
				return nil, err
			}
			// 保存营业执照图片
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/businessLicense/"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}

			//  注意：实际存储的License 需要存储持久化后的文件ID，而不是路径
			newData.BusinessLicensePath = gconv.String(fileInfo.Id)
		}

		//if info.BusinessLicenseLegalPath != "" && !gfile.Exists(info.BusinessLicenseLegalPath) {
		//	// 检测缓存文件
		//	fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.BusinessLicenseLegalPath), userId, "请上传法人证件照图片")
		//	if err != nil {
		//		return nil, err
		//	}
		//	// 保存法人证件照图片
		//	fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/businessLicense/"+fileAt+fileInfoCache.Ext, fileInfoCache)
		//	if err != nil {
		//		return nil, err
		//	}
		//	newData.BusinessLicenseLegalPath = fileInfo.Src
		//}
	}

	gconv.Struct(newData, data)
	return data, err
}
