package co_funs

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// CheckLicenseFiles 检查和处理许可相关的文件。
// 该函数接收一个上下文、审计许可信息和一个许可数据对象，
// 并返回处理后的许可数据对象和可能的错误。
// 泛型参数 T 可以是 co_entity.License 或 co_do.License 类型。
func CheckLicenseFiles[T co_entity.License | co_do.License](ctx context.Context, info co_model.AuditLicense, data *T) (response *T, err error) {
	// 创建一个新的许可对象以存储处理后的数据。
	newData := &co_entity.License{}
	// 将传入的许可数据转换为新的许可对象。
	gconv.Struct(data, newData)

	{
		// 获取用户ID。
		//userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id
		userId := info.UserId

		// 定义用户资源文件夹路径。
		userFolder := "resource/license/" + gconv.String(newData.Id)

		// 获取当前时间，用于文件名的生成。
		fileAt := gtime.Now().Format("YmdHis")

		// 检查并处理身份证头像面照片。
		if !gfile.Exists(info.IdCardFrontPath) {
			// 从缓存中获取上传的文件信息。
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.IdCardFrontPath), userId, "请上传身份证头像面")
			if err != nil {
				return nil, err
			}
			// 保存身份证头像面照片。
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/idCard/front_"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}

			// 存储持久化后的文件ID到许可对象中。
			newData.IdcardFrontPath = gconv.String(fileInfo.Id)
		}

		// 检查并处理身份证国徽面照片。
		if !gfile.Exists(info.IdCardBackPath) {
			// 从缓存中获取上传的文件信息。
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.IdCardBackPath), userId, "请上传身份证国徽面")
			if err != nil {
				return nil, err
			}
			// 保存身份证国徽面照片。
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/idCard/back_"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}

			// 存储持久化后的文件ID到许可对象中。
			newData.IdcardBackPath = gconv.String(fileInfo.Id)
		}

		// 检查并处理营业执照照片。
		if !gfile.Exists(info.BusinessLicensePath) {
			// 从缓存中获取上传的文件信息。
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.BusinessLicensePath), userId, "请上传营业执照图片")
			if err != nil {
				return nil, err
			}
			// 保存营业执照照片。
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/businessLicense/"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}

			// 存储持久化后的文件ID到许可对象中。
			newData.BusinessLicensePath = gconv.String(fileInfo.Id)
		}

		// 检查并处理门头照。
		if len(info.DoorPictures) > 0 {
			pictures := make([]co_model.AttachPictures, 0)
			for _, picture := range info.DoorPictures {
				if !gfile.Exists(picture.Id) {
					// 从缓存中获取上传的文件信息。
					fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(picture.Id), userId, "请上传门头照")
					if err != nil {
						return nil, err
					}
					// 保存门头照。
					fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/doorPictures/"+fileAt+fileInfoCache.Ext, fileInfoCache)
					if err != nil {
						return nil, err
					}
					picture.Id = gconv.String(fileInfo.Id)
					picture.Size = fileInfo.Size
					picture.Ext = fileInfo.Ext

					pictures = append(pictures, picture)
				}
			}
			// 将门头照信息编码为JSON字符串并存储到许可对象中。
			encodeString, _ := gjson.EncodeString(pictures)
			newData.DoorPicturesJson = encodeString
		}

		// 检查并处理其它照片。
		if len(info.OtherPictures) > 0 {
			pictures := make([]co_model.AttachPictures, 0)
			for _, picture := range info.OtherPictures {
				if !gfile.Exists(picture.Id) {
					// 从缓存中获取上传的文件信息。
					fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(picture.Id), userId, "请上传门头照")
					if err != nil {
						return nil, err
					}
					// 保存其它照片。
					fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/doorPictures/"+fileAt+fileInfoCache.Ext, fileInfoCache)
					if err != nil {
						return nil, err
					}
					picture.Id = gconv.String(fileInfo.Id)
					picture.Size = fileInfo.Size
					picture.Ext = fileInfo.Ext

					pictures = append(pictures, picture)
				}
			}
			// 将其它照片信息编码为JSON字符串并存储到许可对象中。
			encodeString, _ := gjson.EncodeString(pictures)
			newData.OtherPicturesJson = encodeString
		}
	}

	// 将处理后的许可数据转换回传入的数据对象。
	_ = gconv.Struct(newData, data)

	// 返回处理后的数据对象和错误信息。
	return data, err
}

// ApplyUserRestrictions 根据用户类型应用查询限制。
// 该函数检查当前会话的用户类型，如果不是平台用户，则限制查询结果仅为当前用户的数据。
// 这用于在查询数据库时，确保普通用户只能获取自己的信息。
func ApplyUserRestrictions(m *gdb.Model) *gdb.Model {
	// 获取当前会话中的用户信息。
	sessionContent := sys_service.SysSession().Get(m.GetCtx()).JwtClaimsUser

	// 如果是平台用户类型，不应用任何限制。
	if sessionContent.IsPlatformUserType {
		return m
	}

	// 对于非平台用户，限制查询结果仅为当前用户的数据。
	return m.Where(m.Builder().Where("UnionMainId", sessionContent.UnionMainId))
}
