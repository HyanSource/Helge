package hnet

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/HyanSource/hyannetserver/hinterface"
	"github.com/HyanSource/hyannetserver/utils"
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

func NewConntion(server hinterface.Iserver, conn *net.TCPConn, connID uint32) hinterface.Iconnection {

	c := &Connection{
		TCPServer:    server,
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
		msgBuffChan:  make(chan []byte, utils.GlobalObject.MaxMsgChanLen),
	}

	c.TCPServer.GetConnMgr().Add(c)

	return c
}

func (t *Connection) Start() {
	go t.StartWrite()
	go t.StartReader()

	//hook
}

func (t *Connection) Stop() {
	if t.isClosed {
		return
	}

	t.isClosed = true

	//hook

	t.Conn.Close()

	t.TCPServer.GetConnMgr().Remove(t)

	close(t.ExitBuffChan)
	close(t.msgBuffChan)
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

	if t.isClosed {
		return errors.New("sendmsg connection is closed")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMsg(msgid, data))

	if err != nil {
		fmt.Println("pack error msg id:", msgid)
		return errors.New("pack error")
	}

	t.msgChan <- msg

	return nil
}

func (t *Connection) SendBuffMsg(msgid uint32, data []byte) error {

	if t.isClosed {
		return errors.New("sendbuffmsg connection is closed")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMsg(msgid, data))

	if err != nil {
		fmt.Println("pack error msg id:", msgid)
		return errors.New("Pack error")
	}

	t.msgBuffChan <- msg

	return nil
}

//goroutine
func (t *Connection) StartWrite() {
	defer fmt.Println(t.RemoteAddr().String(), " conn write exit")
	for {
		select {
		case data := <-t.msgChan:
			if _, err := t.Conn.Write(data); err != nil {
				fmt.Println("msgchan data conn write err:", err)
				return
			}
		case data, ok := <-t.msgBuffChan:
			if ok {
				if _, err := t.Conn.Write(data); err != nil {
					fmt.Println("msgbuffchan data conn write err:", err)
					return
				}
			} else {
				fmt.Println("msgbuffchan is closed")
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
		headdatalen := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(t.Conn, headdatalen); err != nil {
			fmt.Println("headdatalen read err:", err)
			break
		}
		datalen, err := dp.UnPack_Head(headdatalen)
		if err != nil {
			fmt.Println("datalen unpack err:", err)
			return
		}

		/*判斷長度*/
		if utils.GlobalObject.MaxPacketSize > 0 && datalen > utils.GlobalObject.MaxPacketSize {
			fmt.Println("too large msg data received")
			return
		}

		//讀取msgid
		headmsgid := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(t.Conn, headmsgid); err != nil {
			fmt.Println("headmsgid read err:", err)
			break
		}

		msgid, err := dp.UnPack_Head(headmsgid)
		if err != nil {
			fmt.Println("msgid unpack err:", err)
			return
		}

		//加上一個判斷有無超過最大byte長度(全域變數)

		//讀取data
		bodydata := make([]byte, datalen)
		if _, err := io.ReadFull(t.Conn, bodydata); err != nil {
			fmt.Println("bodydata read err:", err)
			break
		}

		msg := NewMsg(msgid, bodydata)
		//未來在Request新增一個New的方法
		req := Request{
			conn: t,
			msg:  msg,
		}

		// fmt.Println(req)

		if utils.GlobalObject.MaxWorkerTaskLen > 0 {
			t.TCPServer.GetMsgHandle().SendMsgToTaskQueue(&req)
		} else {
			go t.TCPServer.GetMsgHandle().DoMsgHandler(&req)
		}

	}

}

//屬性應該要獨立寫成1個模塊
