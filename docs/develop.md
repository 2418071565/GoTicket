
## 简介

在这个项目结构中，每个文件夹作用如下：

- `controllers`：用于处理 HTTP 请求和响应，通常是应用的入口。控制器是接收和解析客户端请求的第一步，并调用相关的服务层逻辑来处理这些请求，也负责处理各个模块出现的错误。

- `services`：包含业务逻辑的实现。服务层通常负责处理业务规则、复杂逻辑、数据的处理和转换等工作。

- `models`：定义数据模型（通常与数据库结构对应）和相关方法。模型通常用来表示应用的核心数据结构，可以包括字段定义、数据库表映射、方法和关系等。

- `pkg`：一些基础工具，如：日志输出，系统配置等。

- `dto`：用于定义 **Data Transfer Objects（数据传输对象）**，即在各个组件或服务间传递数据的对象格式。

- `storage`：数据存储，主要是数据库、cache 等定义及方法的实现。

- `docs`：项目文档。有 API信息、项目介绍、开发规范等。

- `test`：单元测试及 benchmark。


## 架构

![](image/architecture.png)

## 规范

### 命名

类名和函数名采用驼峰命名法，一些专有名词应全部大写，变量应采用蛇形命名法。如要控制包外的可见性，函数的首字母小写，其余按驼峰命名法命名，专有名词开头的单词应全部小写：

```go
// 包外可见
// 函数名
LogToFile()
POSTUser()
// 变量名
var Time_str
var DB

// 包内可见
// 函数名
logToFile()
postUser()
// 变量名
var time_str
var db
```

宏、全局常量、配置信息名字的单词间要用 `_` 分隔，且全部大写：

```go
DB_HOST
DB_MAX_IDLE_CONNS
```

目录及文件名要遵循蛇形命名法。

### 其他

字符串应使用双引号 `""` 括起来而非 \``。 \`` 只应出现在类的成员变量的 comments 中。

```go
// 字符串
"Hello GoTicket"

// comments
type Ticket struct {
  name string `This is a example`
}
```