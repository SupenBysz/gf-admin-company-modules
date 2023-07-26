## 一、技术选型

### 开发工具

![推荐](https://img.shields.io/badge/goland-2022.2.3+-blue) ![img](https://img.shields.io/badge/vscode-1.71.2+-blue)

- 推荐 goland 2022.2.3+

#### 语言/框架/类库

<img src="https://img.shields.io/badge/golang-1.19+-green" alt="" />
<img src="https://img.shields.io/badge/goframe-2.3.2+-green" alt="" />
<img src="https://img.shields.io/badge/casbin-2.60.0+-green" alt="" />
<img src="https://img.shields.io/badge/golang jwt-4.4.3+-green" alt="" />
<img src="https://img.shields.io/badge/idgenerator go-1.3.2+-green" alt="" />

#### 数据库

<img src="https://img.shields.io/badge/PostgreSQL-14.5+-red" alt="" />

#### 脚本工具

<img src="https://img.shields.io/badge/goframe cli-2.4.1+-red" alt="" />

## 二、项目结构
```
/
├── api               对外接口：接口定义
│   ├── xx_api            接口所需请求参数封装。
│   └── xx_v1             对外提供服务的输入/输出数据结构定义，包含多主体接口定义。
├── database          迁移脚本：数据库脚本及迁移或升级脚本
├── co_consts         常量定义：组件专用常量或公共类型及变量定义。
├── co_controller     接口处理：接收/解析用户输入参数的入口/接口层。
├── co_interface      接口封装：封装controller控制层和service层接口定义。
│   ├── i_controller          不同模块的控制器层接口封装。
│   └── service_interface.go  各模块服务接口service定义。  
├── co_model          结构模型：数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义。
│   ├── co_dao           数据访问：数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 CURD 方法
│   ├── co_do            领域对象：用于dao数据操作中业务模型与实例模型转换，由工具维护，用户不能修改。
│   ├── co_entity        数据模型：数据模型是模型与数据集合的一对一关系，由工具维护，用户不能修改。
│   ├── co_enum          枚举定义：枚举类型及变量定义。
│   └── co_hook          Hook定义：钩子函数定义。
├── co_module         操作模型：全局操作模型定义及初始化方法封装。
├── co_permission     权限定义：定义接口权限树结构。
├── co_router         接口路由：用于接口的path路由定义层，利用具体的路由绑定控制器。
├── docs              说明文档：业务流程图 + 接口逻辑实现图。
├── example           运行例子：当前组件模块用于测试调试启动入口文件，启动所需文件示例。
│   ├── ...           示例文件：适配company所需示例文件。
│   └── main.go       入口指令：启动文件，当前作为业务项目引用应用的参考。
├── hack              工具脚本：存放组件开发工具、脚本等内容。例如，CLI工具的配置，各种shell/bat脚本等文件
│   └──   tpls           模板文件：包含dao，do，entity生成的模板，一般配合Makefile使用
├── i18n              国际化输出：国际化输出文件
├── internal          组件逻辑：组件逻辑存放目录。通过internal特性对外部隐藏可见性。
│   ├── boot          权限初始化：权限树等初始化。
│   └── logic         模块封装：组件逻辑封装管理。
├── manifest          交付清单：包含程序编译、部署、运行、配置的文件。
│   └── config        配置管理：配置文件存放目录。
├── resources         静态资源：静态资源文件。这些文件往往可以通过 资源打包/镜像编译 的形式注入到发布文件中。
├── utility           实用类库：一般以独立文件形式的通用类库
├── .env              环境配置：权限最高的配置文件。
├── .gitignore        Git忽略：git提交忽略文件定义。
├── go.mod            依赖管理：使用Go Module包管理的依赖描述文件。
├── Makefile          Make工具：使用自定义模版等实现gf工具的封装&替换配置文件。
└── main.go           依赖库入口
```

### gf-admin-company-modules 属于筷满客项目的多主体组件库，具备：【公司、团队、员工、财务】、我的
特性：
> 此项目封装了多主体的底层基础，其使用者可以根据实际所需延伸出多主体或者单主体 <br>
> 使用者可以根据自身需求直接引用或者二次封装 <br>
> ...

> ### company 特性：
> 1、支持多实例（即【公司、团队、员工、财务】支持全局引用，或按需引用） <br>
> 2、支持实例模型表绑定到不同或相同的数据表 <br>
> 3、支持不同实例拥有各自独立的权限控制 <br>
> 4、支持路由全局引入，或按需引入 <br>
> 5、支持多语言国际化，并支持手动或自动加载不同实例的国际化语言文件 <br>
> 6、支持自动缓存，源数据有变动时自动更新缓存；本地内存模式缓存（不支持分布式及微服务），redis模式（支持分布式及微服务） <br>
> 7、支持自动生成OpenApi接口以及API文档 <br>
> 8、支持分布式微服务应用 <br>
> 9、支持一键引用，零逻辑实现公司、团队、员工、财务等业务接口服务 <br>
> 10、支持自定义重写或二次封装【公司、团队、员工、财务】等相关控制器方法或接口描述信息 <br>
> 11、支持通过Hook方法实现【公司、团队、员工、财务】来附加业务层需要的自定义字段数据 <br>

描述：
> 接口包含两部分：接口定义（api）+ 接口实现（controller）<br>

> 服务包含两部分：模型结构（model）+ 服务接口（sys_service）<br>
> 负责接收并响应调用的输入与输出，包括对输入参数的过滤、转换、校验，对输出数据结构的维护，并调用 service 实现业务逻辑处理。

## 三、项目编译/运行

### 前言：

编译后执行必须文件如下

- i18n目录需要，里头包含项目所需国际化输出文件
- manifest/config/config.yaml 配置文件
- hack/config.yaml 配置文件，用于配置gf dao生成数据模型表   （可选）
- resources目录需要，是serverRoot目录，里头包含ip资源库等资源
- .env 配置文件，优先级最高  (可选)

### 部署&编译流程：

1. 创建并导入数据库文件 latest.sql

2. 修改数据库连接配置，在.env文件或者mainifeast/config/config.yaml 修改,

3. 下载项目文件

4. go  mod tidy 下载并清理依赖

5. `go  build main.go 或者 gf build main.go -n main -a amd64 -s linux -p .`

6. 运行
> 由于是以工具库方式提供服务，需要进行引用运行。<br>
> 根据example目录下的示例文件进行适配引用。
> 项目被引用运行时，是可以实现调试目的。<br>
> 相关引用参考 example/main.go internal/boot/boot.go


### 项目脚手架make命令:

- `make dao`: 根据“hack”文件夹中的配置文件为“entity/dao/do”生成go文件，具体生成的模型目标文件夹需要在Makefile文件标明。
- `make service`: 解析“逻辑”文件夹生成service服务接口。具体生成的service服务接口目标文件夹需要在Makefile文件标明。
> 不推荐直接在本项目使用 gf gen 等命令，将会导致文件生成的位置不在预期位置，如果一定要用先了解相关参数后再用。


### 项目脚手架gf命令：

- `gf version`： 查看框架版本

- `gf init projectName` : 用于创建gofream框架结构项目
- `gf build ... `：编译项目
- `gf run main.go `：运行
- `gf gen dao`：根据hack中的配置文件，生成模型文件，如果配置了make脚手架，不建议直接使用该命令生成。
- `gf gen service`：生成服务接口，如果配置了make脚手架，不建议直接使用该命令生成。
- ``gf [COMMAND] -h``: 查看命令帮助说明



## 四、代码风格规范
&emsp;&emsp;代码风格并不影响程序的运行，也不会给你的程序带来潜在的危险。但一段好的代码风格，能让人阅读起来特别舒服，特别是阅读别人代码的时候。
## 代码风格规范
### 1、包名
&emsp;&emsp;统一使用小写，用下划线作为分隔符，以自然语义缩写或全称作为分隔
### 2、类名
&emsp;&emsp;由于golang属于面向过程编程，没有真正意义上的类。<br>&emsp;&emsp;本项目将逻辑封装成服务对象、控制器对象等，通过封装后的对象进行关业务调用，以实现仿类操作效果，这里称作“伪类”。那么伪类的定义规范如下：
> 伪类采用大驼峰的命名形式，所谓大驼峰就是首字母大写，例如UpperCameCase
### 3、常量
- 全局常量：遵行类名的命名一样的方式。
- 局部常量：局部常量则采用小驼峰的形式。所谓局部常量指的是方法/内/包内的常量。
### 4、大括号代码域
&emsp;&emsp;遵循JAVA风格，缩进直接紧接后面括号，例如：
```
int Foo(bool isBar) {
    if (isBar) {
        bar();
        return 1;
    }
    return 0
}
```
### 5、tab缩进
&emsp;&emsp;每个tab缩进，用4个空格代替。
## 接口风格规范
&emsp;&emsp;为了能更好的同步接口相关约束，规范接口文档的自动生成、并支持接口的相关参数自动校验项目接口做如下规范定义。
### 接口规范
1、所有逻辑接口需要封装成控制器对象，统一保存在controller文件夹下<br>
2、一个控制器接口方法首字母大写的驼峰命名方式<br>
3、每个接口方法只能包含2个参数，分别为接口上下文对象 Context 和 请求参数定义（以Req结尾的类型参数定义）。接口方法返回值只能有2个固定返回值，第一个是要返回的数据，第二个是错误异常信息。<br>例如：
```
func (a *cAuth) Login(ctx context.Context, req *sysapi.LoginReq) (res *sysapi.LoginRes, err error) {
    return nil, nil
}
```
- 方法参数：ctx context.Context 为请求上下文，不需要用到上下文时，变量 ctx 可用下划线代替
- 方法参数：req *sys_api.LoginReq 为请求参数定义，必须以Req结尾，且必须是指针对象，不需要用到前端请求参数时，变量 req 可用下划线代替
- 方法返回：*sys_api.LoginRes 为方法第一个返回参数，且必须以Res结尾，返回值的变量名称为可选，如：(*sys_api.LoginRes,error)
- 方法返回：err error 为方法第二个返回参数，必填项，该参数类型固定且不可修改
### 请求参数规范
&emsp;&emsp;请求参数的定义统一在对象成员属性中实现约束、描述、验证规则等相关定义，例如：
```
type LoginReq struct {
    g.Meta `path:"/login" method:"post" summary:"登录" tags:"鉴权"`
	Username string `json:"username" v:"required#请输入用户名" dc:"登录账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"登录密码"`
    Captcha  string `json:"captcha" v:"required#请输入验证吗" dc:"验证码"`
}
```
- LoginReq 作为请求对象的类型名称，必须以首字母大写，且以Res结尾
- g.Meta path:定义请求路径，method:定义请求类型,summary:定义接口概要说明，tags:定义接口在文档生成上的分类
- 请求参数属性：Username，Password，Captcha，定义了请求参数的细则，包括但不限于json:说明参数格式，v:参数校验规则和校验失败后的提示信息,dc:参数描述



## 五、Git Commit 语义化规范
### 作用：
- 能加快 Code Review 的流程
- 根据 Git Commit 的元数据生成 Changelog
- 后续维护者可以知道 Feature 被修改的原因
- 统一团队 Git Commit 日志标准，便于后续代码 Review 和版本发布
- main、dev两个分支提交类型限定为：feat,fix,docs,style,refactor,perf,test,chore,revert等
- 提交格式参考：
```
<type>(<scope>):<subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```
- Type 代表了某次提交的类型，比如是修复一个 bug 还是增加了一个新的 feature。常用类型如下：
    - feat：新增 feature
    - fix：修复 bug
    - docs：仅仅修改了文档，比如 README、CHANGELOG 等等
    - style：仅仅修改了空格、格式缩进、多行等等，不改变逻辑
    - refactor：代码重构，没有加新的功能或者修复 bug
    - pref：优化相关，比如提升性能、体验等
    - test：测试用例，包括单元测试、集成测试等
    - chore：改变构建流程，或者增加依赖库、工具等
    - revert：回滚到上一个版本
- Scope 字段用于说明本次 commit 所影响的范围，比如视图层、数据模型或者路由模块等，是一个可选参数。
- Subject 字段是本次 commit 的一个概要，需要用最简洁的语言来说明本次修改的内容。
- Body 可以使用多行文本详细地说明本次提交所改动的一些细节，从而帮助后续的使用者们更好地了解代码的内容。
- Footer 大部分只用于两种情况
    - 不兼容变动
        - 如果本次的 commit 与前一个版本的代码无法兼容，那么 Footer 部分需要以 BREAKING CHANGE 开头，后面描述本次变动的详细情况以及迁移到新版本代码的方法。
    - 关闭 Issue
        - 如果当前 commit 针对某个 issue，那么可以在 Footer 部分关闭这个 issue 。以 Close 开头，后面用 # 号标识对应的 Issue 号码。

### Commitizen
&emsp;&emsp;如果觉得自己写规范的 comment 信息太麻烦，那么可以使用Commitizen ，可以辅助撰写一个合格的 Commit message 工具。


## 注意事项
### 项目规范
```
项目规范：
- 所有项目包名统一规范用小写，下划线隔开

- 包名不支持横杠,项目名支持横杠，包名不支持

项目引入问题：
- 项目引入是有先后顺序区分的。
```
###  方法命名规范
```
方法命名规范: 
Get 前缀代表获取单一记录数据
Query 代表返回列表数据，需要包含分页信息
Update 代表更新数据信息
Delete 代表删除数据
Set 代表设置相关状态或类型

注意：Query 前缀返回的数据结构必须是 CollectRes
```

### 接口规范
```
不要缩写  例如： addr  u_id 

使用驼峰命名，不要使用短下划线_   例如：order_by --> orderBy    user_id  --> userId 
```

###  参数使用规范
```
  字段名不要使用硬编码，如果数据库改动了，那么重新生成的文件不会自动适配。

  OperatorEmployee.Columns().Id 这种方式
```