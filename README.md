#### Helge 海格爾 golang 的TCPServer框架

#### 使用方式

1. 初始化server
2. 定義路由
3. 設置hook
4. 定義Global
5. 開始服務
```go
func main() {
    //初始化
    s :=hnet.NewServer()

    //路由
    s.AddRouter(0,&PingRouter{})

    //設置hook
    s.GetHook().SetHook("start")

    //開始
    s.Serve()
}
```

#### 定義路由
```go
type PingRouter struct {
    hnet.BaseRouter
}

func (t *PingRouter) Handle(request hinterface.Irequest) {
    //do soming
}

```
#### 定義hook
```go
func StartConn(conn hinterface.Iconnection) {

}
```

### 定義Global
1. 創建一個conf的資料夾
2. 在資料夾下新增Helge.json
```json
{
    "Name":"HelgeV01",
    "Host":"127.0.0.1",
    "TCPPort":8124,
    "MaxConn":10,
    "WorkerPoolSize":10
}
```
- Name
Server名稱
- Host
主機IP
- TCPPort
主機Port號
- MaxConn
連線數上限
- WorkerPoolSize
啟動工作池的數量