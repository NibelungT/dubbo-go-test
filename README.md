# Dubbo-Go 调用 Java Dubbo 服务示例

## 项目介绍

本项目是一个跨语言RPC调用的示例，展示了如何使用Go语言通过dubbo-go框架调用Java实现的Dubbo服务。项目包含两个主要部分：

- `go-service`: Go语言客户端，使用dubbo-go框架发起RPC调用
- `java-client`: Java语言服务端，提供Dubbo服务接口

通过本项目，您可以了解如何在微服务架构中实现Go和Java服务之间的无缝通信。

## 技术栈

### Go服务端

- Go 1.24
- dubbo-go v3 (Apache Dubbo的Go语言实现)
- dubbo-go-hessian2 (序列化库)
- Nacos (服务注册与发现)

### Java服务端

- Java 8
- Spring Boot 2.3.12.RELEASE
- Apache Dubbo 2.7.15
- Nacos 2.1.0 (服务注册与发现)

## 项目结构

```
├── go-service/            # Go客户端代码
│   ├── main.go            # 主程序入口
│   ├── rpc/               # RPC相关代码
│   │   ├── dubbo.go       # Dubbo客户端配置
│   │   └── GreetingService.go  # 服务调用实现
│   ├── go.mod             # Go模块依赖
│   └── go.sum             # 依赖版本锁定文件
│
└── java-client/           # Java服务端代码
    ├── pom.xml            # Maven项目配置
    └── src/               # Java源代码
```

## 功能特性

本项目实现了以下功能：

1. Go服务通过dubbo-go调用Java Dubbo服务的三种方式：
   - 简单字符串参数调用 (`SayHi`)
   - 复杂对象参数调用 (`Greet`)
   - 泛化调用 (`SayHiGeneric`)

2. 使用Nacos作为服务注册中心，实现服务的自动发现

3. 使用Hessian2序列化协议，确保跨语言数据传输的兼容性

## 快速开始

### 前置条件

- 安装Go 1.24或以上版本
- 安装Java JDK 8或以上版本
- 安装Maven
- 安装并启动Nacos服务（默认地址：localhost:8848）

### 启动Java服务端

1. 进入java-client目录
```bash
cd java-client
```

2. 编译并启动服务
```bash
mvn clean package
mvn spring-boot:run
```

### 启动Go客户端

1. 进入go-service目录
```bash
cd go-service
```

2. 运行Go客户端
```bash
go run main.go
```

## 核心代码说明

### Go客户端配置 (dubbo.go)

```go
service, err := dubbo.NewInstance(
    dubbo.WithName("go-service"),
    dubbo.WithRegistry(
        registry.WithNacos(),
        registry.WithAddress("localhost:8848"),
        registry.WithRegisterInterface(),
    ),
)

cli, err := service.NewClient(
    client.WithClientProtocolDubbo(),
    client.WithClientSerialization(constant.Hessian2Serialization),
)
```

### Go调用Java服务示例 (GreetingService.go)

```go
// 简单参数调用
func SayHi(param string) (string, error) {
    // ...
    err = req.CallUnary(context.Background(), []interface{}{param}, &response, "sayHi")
    // ...
}

// 复杂对象参数调用
func Greet(param *GreetRequest) (*GreetResponse, error) {
    // ...
    err = req.CallUnary(context.Background(), []interface{}{param}, &response, "greet")
    // ...
}
```

## 注意事项

1. 确保Java服务和Go客户端使用相同的接口名称、方法名称和参数类型

2. 在Go客户端中，需要为Java对象定义对应的结构体，并实现`JavaClassName()`方法

3. 使用`dubbo_go_hessian2.RegisterPOJO()`注册POJO类型，确保序列化/反序列化正确

4. Nacos服务地址需要在两端保持一致

## 常见问题

1. **连接Nacos失败**
   - 检查Nacos服务是否正常运行
   - 确认配置的Nacos地址是否正确

2. **调用服务失败**
   - 检查Java服务是否正确注册到Nacos
   - 确认接口名称、版本号和分组是否匹配
   - 检查参数类型是否兼容

## 参考资料

- [Apache Dubbo官方文档](https://dubbo.apache.org/zh/docs/)
- [Dubbo-go GitHub仓库](https://github.com/apache/dubbo-go)
- [Nacos官方文档](https://nacos.io/zh-cn/docs/what-is-nacos.html)

## 许可证

[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)
