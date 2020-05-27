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

func NewConntion() {

}
