# go-clean-arch-template
golang 简洁架构模版，对应的前端项目为 [vue-web-template](https://github.com/niudaii/vue-web-template)。

### 技术架构

- RPC 通信：amqp-rpc
- HTTP Server：gin
- 分布式框架：asynq
- orm 框架：gorm
- 日志：zap
- 身份鉴权：jwt
- 权限管理：casbin

### 代码架构

- 简洁架构![Clean Coder Blog](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/CleanArchitecture.jpg)
- 面向包的设计
- 面向接口编程
- 依赖注入
- repo 层 Filter 思想 BuildWhere()
- 童子军军规

### API 文档

https://apifox.com/apidoc/project-2606572/

### 参考

https://github.com/amitshekhariitbhu/go-backend-clean-architecture

https://github.com/evrone/go-clean-template

https://www.jianshu.com/p/f3a025fb3053

《Clean Code》