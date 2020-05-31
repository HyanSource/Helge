package hnet

import (
	"fmt"
	"io"
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
	go t.StartWrite()
	go t.StartReader()
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
	defer fmt.Println(t.RemoteAddr().String(), " conn write exit")
	for {
		select {
		case data := <-t.msgChan:
			if _, err := t.Conn.Write(data); err != nil {
				return
			}
		case data, ok := <-t.msgBuffChan:
			if ok {
				if _, err := t.Conn.Write(data); err != nil {
					return
				}
			} else {
				break
			}

		case <-t.ExitBuffChan:
			return
		}
	}
}

//goroutine
func (t *Connection) StartReader() {
	defer fmt.Println(t.RemoteAddr().String(), " conn read exit")

	for {
		//照自己修改的解包流程
		dp := NewDataPack()

		//讀取datalen
		head_datalen := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(t.Conn, head_datalen); err != nil {
			break
		}
		datalen, err := dp.UnPack_Head(head_datalen)
		if err != nil {
			break
		}
		//讀取msgid
		head_msgid := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(t.Conn, head_msgid); err != nil {
			break
		}

		msgid, err := dp.UnPack_Head(head_msgid)
		if err != nil {
			break
		}

		//加上一個判斷有無超過最大byte長度(全域變數)

		//讀取data
		body_data := make([]byte, datalen)
		if _, err := io.ReadFull(t.Conn, body_data); err != nil {
			break
		}

		msg := NewMsg(msgid, body_data)
		//未來在Request新增一個New的方法
		req := Request{
			conn: t,
			msg:  msg,
		}
	}

}
