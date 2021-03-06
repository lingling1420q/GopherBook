# Web 开发

> Go-web

由浅入深：

**简单的 web 服务**

- 启动的方式
- 返回的响应
    - json
    - string
    - xml
- 写入的方式
    - GET
    - POST
    - PostForm
        - 登入操作
- 模版的使用
    - FuncMap
    - 继承
    - 遍历


**正规的设计**

- 需求文档: PRD
- 模型设计: 
    - model
    - serializer
- 项目组织
    - layout
        - 日志处理
        - 错误信息处理
        - 命令行工具
        - 配置文件
        - ...
- 路由设计
    - 规范
        - 带版本信息
        - 操作资源：名称
        - 正确的设计请求方法
        - 正确的设计状态码
- 请求
    - 参数验证
    - 整型、字符串、时间等
- 响应
    - json
- 文档
    - swagger api
- 测试
    - postman/vscode
    - 接口测试
- 总结


### 说明

- 如何启动 BeeQuick.v1 服务

```text
- 本地安装 Docker
- 本地安装 docker-compose
- 启动 MySQL： cd deployments && docker-compose -f docker-compose.yml up -d
- go build -o BeeQuick -v -ldflags "-X main.Env=dev" -tags=jsoniter 
- ./BeeQuick
- 调用相应的接口
```

- 如何联系模型设计的能力

```text
- 研究 APP / 小程序
- 自己建立模型，看是否能够满足要求
```

