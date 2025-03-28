package co_v1

import "github.com/gogf/gf/v2/frame/g"

type ModuleConfInfoReq struct {
	g.Meta `method:"post" path:"/moduleTypeInfo" summary:"获取模块配置信息" tags:"模块配置信息"`
}
