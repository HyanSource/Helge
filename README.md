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
func StartConn(conn hinterface.IConnection) {

}
```

#### 定義Global
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

---
- # hinterface.IServer
#### 1. Start()
啟動
#### 2. Stop()
停止
#### 3. Serve()
開啟業務
#### 4. AddRouter(magid uint32,router IRouter)
新增路由
#### 5. GetHook() IHook
取得hook模塊

---
- # hinterface.IHook
#### 1. SetHook(hookname string,f func(IConnection))
設置hook函數

---
- # hinterface.IConnection
#### 1. GetTCPConnection() *net.TCPConn
取得tcp連接
#### 2. GetConnID() uint32
取得connid
#### 3. RemoteAddr() net.Addr
取得客戶端地址
#### 4. SendMsg(msgid uint32,data []byte) error
傳送訊息(無緩衝)
#### 5. SendBuffMsg(msgid uint32,data []byte) error
傳送訊息(有緩衝)
#### 6. GetPropertys() IProperty
取得屬性模塊

---
- # hinterface.IProperty
#### 1. SetProperty(key string,value interface{})
設置屬性
#### 2. GetProperty(key string) (interface{},error)
取得屬性
#### 3. RemoveProperty(key string) bool
移除屬性

---
- # hinterface.IMessage
#### 1. GetDataLen() uint32
取得資料長度
#### 2. GetMsgId() uint32
取得訊息id
#### 3. GetData() []byte
取得訊息

---
- # hinterface.IRequest
#### 1. GetConnection() IConnection
取得連接
#### 2. GetMessage() IMessage
取得訊息

---
- # hinterface.IRouter
#### 1. PreHandle(request IRequest)
處理業務前方法
#### 2.Handle(request IRequest) 
處理業務方法
#### 3. PostHandle(request IRequest)
處理業務後方法

---