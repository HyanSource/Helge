package hinterface

/*玩家連接接口模塊*/
type Iconnection interface {
	Start() //啟動連接
	Stop()  //停止連接

	GetTCPConnection()  //獲取原始socket tcpconn
	GetConnID()         //獲取連線id
	RemoteAddr()        //獲取客戶端地址
	SendMsg() error     //將數據給TCP客戶端 (無緩衝)
	SendBuffMsg() error //將數據給TCP客戶端 (有緩衝)
	SetProperty()       //設置屬性
	getProperty()       //獲取屬性
	RemoveProperty()    //移除屬性
}
