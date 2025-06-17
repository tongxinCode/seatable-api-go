# SeaTable API Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/tongxinCode/seatable-api-go.svg)](https://pkg.go.dev/github.com/tongxinCode/seatable-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/tongxinCode/seatable-api-go)](https://goreportcard.com/report/github.com/tongxinCode/seatable-api-go)
[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org/dl/)

SeaTable API Go Client 是一个用于与 SeaTable 数据库进行交互的 Go 语言客户端库。它提供了完整的 CRUD 操作、文件管理、列管理以及实时 Socket.IO 连接功能。

## 安装

```bash
go get github.com/tongxinCode/seatable-api-go
```

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/tongxinCode/seatable-api-go"
)

func main() {
    // 初始化客户端
    client := seatable_api.Init("your-api-token", "https://cloud.seatable.io")
    
    // 认证
    err := client.Auth(false)
    if err != nil {
        log.Fatal("认证失败:", err)
    }
    
    // 获取元数据
    metadata, err := client.GetMetadata()
    if err != nil {
        log.Fatal("获取元数据失败:", err)
    }
    
    fmt.Printf("元数据: %+v\n", metadata)
}
```

### 行操作示例

```go
// 添加行
rowData := map[string]interface{}{
    "Name": "张三",
    "Age":  25,
    "Email": "zhangsan@example.com",
}
result, err := client.AppendRow("employees", rowData)
if err != nil {
    log.Fatal("添加行失败:", err)
}

// 批量添加行
rowsData := []interface{}{
    map[string]interface{}{"Name": "李四", "Age": 30},
    map[string]interface{}{"Name": "王五", "Age": 28},
}
result, err = client.BatchAppendRows("employees", rowsData)

// 更新行
updateData := map[string]interface{}{
    "Age": 26,
}
_, err = client.UpdateRow("employees", "row-id", updateData)

// 删除行
_, err = client.DeleteRow("employees", "row-id")
```

### 文件操作示例

```go
// 上传本地文件
result, err := client.UploadLocalFile(
    "local-file.txt",    // 本地文件路径
    "remote-file.txt",   // 远程文件名
    "documents/",        // 相对路径
    "file",              // 文件类型
    false,               // 是否替换
)

// 上传字节数据
data := []byte("Hello, SeaTable!")
reader := bytes.NewReader(data)
result, err = client.UploadBytesFile(
    "hello.txt",
    reader,
    "",
    "file",
    false,
)

// 下载文件
err = client.DownloadFile("file-url", "local-save-path.txt")
```

### 列管理示例

```go
// 插入新列
err := client.InsertColumn(
    "employees",           // 表名
    "Department",          // 列名
    seatable_api.ColumnTypes("text"), // 列类型
    "department_key",      // 列键
)

// 重命名列
err = client.RenameColumn("employees", "old_column_key", "NewColumnName")

// 调整列宽
err = client.ResizeColumn("employees", "column_key", 200)

// 冻结列
err = client.FreezeColumn("employees", "column_key", true)

// 删除列
err = client.DeleteColumn("employees", "column_key")
```

### 查询和过滤示例

```go
// 过滤行
filters := []map[string]interface{}{
    {
        "column_name": "Age",
        "filter_predicate": "Greater than",
        "filter_term": 25,
    },
}
result, err := client.FilterRows("employees", filters, "", "And")

// SQL 查询
sql := "SELECT * FROM employees WHERE age > 25"
result, err = client.Query(sql, true)

// 列出所有行
rows, err := client.ListRows("employees", "")
```

### Socket.IO 实时连接

```go
// 初始化带 Socket.IO 的客户端
client := seatable_api.Init("your-api-token", "https://cloud.seatable.io")
err := client.Auth(true) // 启用 Socket.IO

if err != nil {
    log.Fatal("认证失败:", err)
}

// 连接 Socket.IO
err = client.Client.Connect()
if err != nil {
    log.Fatal("Socket.IO 连接失败:", err)
}

// 监听事件
err = client.Client.On("custom-event", func(channel *gosocketio.Channel, data interface{}) {
    fmt.Printf("收到事件: %+v\n", data)
})
```

## API 参考

### 类型

#### Base
主要的客户端结构体，包含所有 SeaTable API 操作。

```go
type Base struct {
    Token           string            // API Token
    ServerURL       string            // SeaTable 服务器 URL
    DtableServerURL string            // DTable 服务器 URL
    DtableDbUrl     string            // DTable 数据库 URL
    JwtToken        string            // JWT Token
    JwtExp          int64             // JWT 过期时间
    Headers         map[string]string // HTTP 请求头
    WorkspaceID     int               // 工作空间 ID
    DtableUUID      string            // DTable UUID
    DtableName      string            // DTable 名称
    Timeout         int               // 请求超时时间（秒）
    Client          *SocketIO         // Socket.IO 客户端
}
```

#### ColumnTypes
列类型定义。

```go
type ColumnTypes string
```

### 函数

#### Init
初始化 SeaTable API 客户端。

```go
func Init(token string, serverURL string) *Base
```

#### Auth
进行身份认证。

```go
func (s *Base) Auth(withSocketIO bool) error
```

### 行操作方法

#### AppendRow
添加单行数据。

```go
func (s *Base) AppendRow(tableName string, rowData interface{}) (map[string]interface{}, error)
```

#### BatchAppendRows
批量添加行数据。

```go
func (s *Base) BatchAppendRows(tableName string, rowsData []interface{}) (map[string]interface{}, error)
```

#### InsertRow
在指定行后插入新行。

```go
func (s *Base) InsertRow(tableName string, rowData interface{}, anchorRowID string) (map[string]interface{}, error)
```

#### UpdateRow
更新指定行数据。

```go
func (s *Base) UpdateRow(tableName string, rowID string, rowData interface{}) (map[string]interface{}, error)
```

#### DeleteRow
删除指定行。

```go
func (s *Base) DeleteRow(tableName, rowID string) (map[string]interface{}, error)
```

#### BatchDeleteRows
批量删除行。

```go
func (s *Base) BatchDeleteRows(tableName string, rowIDs interface{}) (map[string]interface{}, error)
```

#### FilterRows
根据条件过滤行。

```go
func (s *Base) FilterRows(tableName string, filters []map[string]interface{}, viewName string, filterConjunction string) (interface{}, error)
```

#### ListRows
列出表中的所有行。

```go
func (s *Base) ListRows(tableName, viewName string) (interface{}, error)
```

### 列操作方法

#### InsertColumn
插入新列。

```go
func (s *Base) InsertColumn(tableName, columnName string, columnType ColumnTypes, columnKey string) (map[string]interface{}, error)
```

#### RenameColumn
重命名列。

```go
func (s *Base) RenameColumn(tableName, columnKey, newColumnName string) (map[string]interface{}, error)
```

#### ResizeColumn
调整列宽。

```go
func (s *Base) ResizeColumn(tableName, columnKey string, newColumnWidth int) (map[string]interface{}, error)
```

#### FreezeColumn
冻结/解冻列。

```go
func (s *Base) FreezeColumn(tableName, columnKey string, frozen bool) (map[string]interface{}, error)
```

#### ModifyColumnType
修改列类型。

```go
func (s *Base) ModifyColumnType(tableName, columnKey string, newColumnType ColumnTypes) (map[string]interface{}, error)
```

#### DeleteColumn
删除列。

```go
func (s *Base) DeleteColumn(tableName, columnKey string) (map[string]interface{}, error)
```

### 文件操作方法

#### UploadLocalFile
上传本地文件。

```go
func (s *Base) UploadLocalFile(filePath, name, relativePath, fileType string, replace bool) (map[string]interface{}, error)
```

#### UploadBytesFile
上传字节数据作为文件。

```go
func (s *Base) UploadBytesFile(name string, r io.Reader, relativePath, fileType string, replace bool) (map[string]interface{}, error)
```

#### DownloadFile
下载文件。

```go
func (s *Base) DownloadFile(url, savePath string) error
```

#### GetFileDownloadLink
获取文件下载链接。

```go
func (s *Base) GetFileDownloadLink(path string) (interface{}, error)
```

#### GetFileUploadLink
获取文件上传链接。

```go
func (s *Base) GetFileUploadLink() (map[string]interface{}, error)
```

### 链接操作方法

#### AddLink
添加行链接。

```go
func (s *Base) AddLink(linkID, tableName, otherTableName, rowID, otherRowID string) (map[string]interface{}, error)
```

#### RemoveLink
移除行链接。

```go
func (s *Base) RemoveLink(linkID, tableName, otherTableName, rowID, otherRowID string) (map[string]interface{}, error)
```

### 查询方法

#### Query
执行 SQL 查询。

```go
func (s *Base) Query(sql string, convert bool) (interface{}, error)
```

#### GetMetadata
获取数据库元数据。

```go
func (s *Base) GetMetadata() (interface{}, error)
```

## 常量

### 列操作常量
```go
const (
    RENAME_COLUMN      = "rename_column"
    RESIZE_COLUMN      = "resize_column"
    FREEZE_COLUMN      = "freeze_column"
    MOVE_COLUMN        = "move_column"
    MODIFY_COLUMN_TYPE = "modify_column_type"
    DELETE_COLUMN      = "delete_column"
)
```

### Socket.IO 事件常量
```go
const (
    JOIN_ROOM        = "join-room"
    UPDATE_DTABLE    = "update-dtable"
    NEW_NOTIFICATION = "new-notification"
)
```

## 错误处理

所有方法都返回 `error` 类型，建议始终检查错误：

```go
result, err := client.AppendRow("table", data)
if err != nil {
    log.Printf("操作失败: %v", err)
    return
}
```

## 依赖

- Go 1.24+
- [github.com/graarh/golang-socketio](https://github.com/graarh/golang-socketio) - Socket.IO 客户端
- [github.com/mattn/go-isatty](https://github.com/mattn/go-isatty) - 终端检测

## 贡献

欢迎提交 Issue 和 Pull Request！

## 相关链接

- [SeaTable 官网](https://seatable.io/)
- [SeaTable API 文档](https://seatable.io/docs/developers/web-api/)
- [Go 官方文档](https://golang.org/doc/)
