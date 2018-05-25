#zhuzi-bootstrap

一是封装功能方便使用, 二是规范项目结构模板

### 基础功能 

- log: 简单封装go-kit/kit/log, 支持Error, Info等等级
- xorm: 简单封装, 支持读写分离
- config: 读取yml, 一个项目只有一个配置文件
- nsq: 简单封装
- grpc: 客户端和服务端
- http: 封装gin

### 项目结构

|目录|说明
|-|
|/api|也就是mvc的控制层, 仅仅用于处理输入输出, 不做任何业务逻辑, http或者rpc都在里面|
|/model|模型层, 处理数据来源, 数据来源可以是其他服务(grpc)或者mysql. 因遵守单一职责原则|
|/service|一般来说一个控制器只需要调用一个model, 这时就不需要service层, 但如果一个controller需要调用多个model的时候, 一般会有一些逻辑判断, 这就违反了controller不做逻辑的原则, 也因为controller不能重用, 这时候就需要service层来组装model以供controller使用|
|/lib|无业务无关的方法/库|
|/proto|用于存放grpc所需要的proto文件|

### QA

为什么lib下的包不新建仓库 在其他项目引用即可, 而是每次都复制过来?
> 1. 将项目最小耦合, 不依赖其他库
> 2. 为了让开发人员更自由. 如果觉得lib库不合理, 可自行修改而不会影响到其他项目.

vendor目录是用来干嘛的?
> 应该提交到git里, 方便CI go build. 有点影响git clone的效率, 但除此之外没想到更好的办法. 
> vendor通过`govendor`生成: 
>  - `go get -u github.com/kardianos/govendor`
>  - `govendor init`
>  - `govendor add +e`

config.test.yml / config.prod.yml / config.yml 三者的区别
> - config.test.yml 用于测试环境, 由开发人员管理. 当push到develop分支后程序会使用此配置文件运行.
> - config.prod.yml 用于生产环境, 由管理服务器的人管理, 开发人员不应该轻易修改. 当push到master分支后程序会使用此配置文件运行
> - config.yml 用于开发人员做本地测试与开发, 由于每个人本地环境不一致, 所以不应该上传到GIT.
>
> ps: 程序会优先使用config.yml文件, 当config.yml文件不存在时会使用config.test.yml
