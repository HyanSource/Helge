package hnet

import (
	"net"

	"github.com/HyanSource/hyannetserver/hinterface"
)

type Connection struct {
	TCPServer    hinterface.Iserver //iserver hook函數用(尚未使用)
	Conn         *net.TCPConn       //當前玩家socket
	ConnID       uint32             //id
	isClosed     bool               //關閉channel用
	ExitBuffChan chan bool          //退出停止用
	msgChan      chan []byte        //無緩衝chan
	msgBuffChan  chan []byte        //有緩衝chan
}

func NewConntion() hinterface.Iconnection {
	return &Connection{}
}

func (t *Connection) Start() {

}

func (t *Connection) Stop() {

}

func (t *Connection) GetTCPConnection() *net.TCPConn {
	return t.Conn
}

func (t *Connection) GetConnID() uint32 {
	return t.ConnID
}

func (t *Connection) RemoteAddr() net.Addr {
	return t.Conn.RemoteAddr()
}

func (t *Connection) SendMsg(msgid uint32, data []byte) error {
	return nil
}

func (t *Connection) SendBuffMsg(msgid uint32, data []byte) error {
	return nil
}

//goroutine
func (t *Connection) StartWrite() {

}

//goroutine
func (t *Connection) StartReader() {

}
