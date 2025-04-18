package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_controller"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_controller/lincense"
	_ "github.com/SupenBysz/gf-admin-company-modules/example/internal/boot/internal"

	"github.com/SupenBysz/gf-admin-company-modules/example/internal/consts"
	"github.com/SupenBysz/gf-admin-company-modules/example/router"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gmode"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			{
				// OpenApi自定义信息
				oai.Info.Title = `API Reference`
				oai.Config.CommonResponse = api_v1.JsonRes{}
				oai.Config.CommonResponseDataField = `Data`
			}

			{
				// 静态目录设置
				uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
				if uploadPath == "" {
					g.Log().Fatal(ctx, "文件上传配置路径不能为空!")
				}
				if !gfile.Exists(uploadPath) {
					_ = gfile.Mkdir(uploadPath)
				}
				// 上传目录添加至静态资源
				s.AddStaticPath("/upload", uploadPath)
			}

			{
				// HOOK, 开发阶段禁止浏览器缓存,方便调试
				if gmode.IsDevelop() {
					s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
						r.Response.Header().Set("Cache-Control", "no-store")
					})
				}
			}

			{
				// ImportPermissionTree 导入权限结构
				_ = sys_service.SysPermission().ImportPermissionTree(ctx, consts.Global.PermissionTree, nil)

				// 导入财务服务权限结构 (可选)
				_ = sys_service.SysPermission().ImportPermissionTree(ctx, consts.Global.FinancePermissionTree, nil)

				// CASBIN 初始化
				sys_service.Casbin().Enforcer()
			}

			// 初始化路由
			apiPrefix := g.Cfg().MustGet(ctx, "service.apiPrefix").String()
			s.Group(apiPrefix, func(group *ghttp.RouterGroup) {
				// 注册中间件
				group.Middleware(
					// sys_service.Middleware().Casbin,
					sys_service.Middleware().CTX,
					sys_service.Middleware().ResponseHandler,
				)

				// 匿名路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 鉴权：登录，注册，找回密码等
					group.Group("/auth", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.Auth) })
					// 图型验证码、短信验证码、地区
					group.Group("/common", func(group *ghttp.RouterGroup) {
						group.Bind(
							// 图型验证码
							sys_controller.Captcha,
							// 短信验证码
							//sys_controller.SysSms,
							// 地区
							sys_controller.SysArea,
						)
					})
				})

				// 权限路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 注册中间件
					group.Middleware(
						sys_service.Middleware().Auth,
					)

					// 注册公共模块路由
					group.Group("/common", func(group *ghttp.RouterGroup) {
						group.Bind(co_controller.ModuleConf)
					})

					// 注册公司模块路由 （包含：公司、团队、员工）
					router.ModulesGroup(consts.Global.IModules, group)

					// 注册财务模块路由 (可选)
					router.FinanceGroup(consts.Global.IModules, group)

					// 审核管理
					group.Group("/audit", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysAudit) })

					// 个人资质管理
					group.Group("/person_license", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysLicense) })

					// 主体资质管理
					group.Group("/license", func(group *ghttp.RouterGroup) { group.Bind(lincense.License) })

				})
			})
			s.Run()
			return nil
		},
	}
)
