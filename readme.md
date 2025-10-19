# Go Demo 教程与贡献指南

## 课程目录概览
- `01_variables`：讲解变量与常量的声明、零值、类型推断和短变量声明。运行 `go run ./01_variables/main.go`。
- `02_functions`：涵盖函数声明、多返回值、匿名函数、闭包及可变参数。运行 `go run ./02_functions/main.go`。
- `03_control_flow`：展示 if/switch、for 循环、range、defer 与 goto 的用法。运行 `go run ./03_control_flow/main.go`。
- `04_collections`：对数组、切片、map 以及 make/len/cap/append/copy 做系统示例。运行 `go run ./04_collections/main.go`。
- `05_structs_methods`：演示结构体、构造函数、方法集以及嵌入字段。运行 `go run ./05_structs_methods/main.go`。
- `06_interfaces`：说明接口定义、隐式实现、接口组合与类型断言。运行 `go run ./06_interfaces/main.go`。
- `07_error_handling`：演练 errors、fmt.Errorf、自定义错误、panic 与 recover。运行 `go run ./07_error_handling/main.go`。
- `08_pointers`：讲解指针基础、函数指针参数以及 new/make 区别。运行 `go run ./08_pointers/main.go`。
- `09_concurrency`：包含 goroutine、channel、缓冲通道、worker 池与 select。运行 `go run ./09_concurrency/main.go`。
- `10_generics`：介绍泛型函数、约束、泛型结构体与 map 构造。运行 `go run ./10_generics/main.go`。

## 仓库贡献指南

### 项目结构与模块组织
全部示例位于 `/` 目录，每个子目录采用两位数字前缀与蛇形命名，例如 `lessons/05_structs_methods`，并包含单独的 `main.go` 用于演示主题。新增课程时，请按序号追加目录，保证命名清晰、导入仅限标准库，若确需外部依赖需在示例中说明理由与使用方式。若课程需要额外数据文件，应与 `main.go` 保持同级并通过相对路径访问，避免跨目录引用。

### 构建、测试与开发命令
使用 `go run ./<lesson_name>/main.go` 独立运行课程示例，例如：
```bash
go run lessons/07_error_handling/main.go
```
提交前务必执行 `gofmt -w ./<lesson_name>/main.go`，若修改较多课程，可一并运行 `gofmt -w ./`。示例脚本需自包含，不依赖外部全局状态；若引入辅助脚本，请在对应课程目录添加简短说明文档。

### 代码风格与命名约定
遵循 Go 官方风格，由 `gofmt` 自动处理缩进（默认制表符）与空行，建议控制行长在 100 字符以内。作为教学范例的导出类型或函数使用 PascalCase，内部辅助函数及变量保持 camelCase。演示代码要体现最佳实践：始终检查错误、避免未使用变量、通过必要注释解释“为什么”，并保持现有中英文注释语气一致。

### 测试指南
当课程包含可验证逻辑时，在同级目录创建 `_test.go` 文件，使用标准 `testing` 框架编写单元测试。初始化模块后，可在仓库根目录执行 `go test ./lessons/...` 运行全部测试；若尚未初始化，请先执行一次 `go mod init go-demo`。测试函数命名应描述行为，例如 `TestUserGreeting`，同时覆盖主路径与课程中强调的边界情形。

### 提交与合并请求规范
提交信息请使用祈使句、聚焦单一主题，如“补充协程池示例”优于泛泛概述；相关课程改动尽量放在同一提交并保持 diff 精简，便于审阅。提交合并请求时请说明教学目标、列出影响的课程目录，并注明新增命令或所需数据；若示例产生交互式输出，可附上截图或终端片段辅助审核。
