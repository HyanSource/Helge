#### Helge 海格爾 golang 的TCPServer框架

#### 使用方式

1. 初始化server
2. 定義路由
3. 設置hook
4. 開始服務
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