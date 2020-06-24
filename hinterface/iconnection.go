package hinterface

import "net"

/*玩家連接接口模塊*/
type IConnection interface {
	Start() //啟動連接
	Stop()  //停止連接

	GetTCPConnection() *net.TCPConn              //獲取原始socket tcpconn
	GetConnID() uint32                           //獲取連線id
	RemoteAddr() net.Addr                        //獲取客戶端地址
	SendMsg(msgid uint32, data []byte) error     //將數據給TCP客戶端 (無緩衝)
	SendBuffMsg(msgid uint32, data []byte) error //將數據給TCP客戶端 (有緩衝)

	GetPropertys() IProperty //取得屬性模塊
}
