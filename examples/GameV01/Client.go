package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/HyanSource/Helge/examples/GameV01/pb"

	"github.com/golang/protobuf/proto"
)

func main() {
	//設置CPU核心數 默認已經設置了
	runtime.GOMAXPROCS(runtime.NumCPU())

	count := 10

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()

			c := NewTcpClient("127.0.0.1", 8124)
			c.Start()
		}()
		time.Sleep(3 * time.Second)
	}

	fmt.Println("client wait")
	wg.Wait()
	fmt.Println("client over")
	select {}
}

type Message struct {
	Len   uint32
	MsgId uint32
	Data  []byte
}

/**/

type ITcpClient interface {
	Start()
	DoMsg(msg *Message)
	Unpack(headdata []byte) (data uint32, err error)
	Pack(msgid uint32, dataBytes []byte) (out []byte, err error)
	SendMsg(msgid uint32, data proto.Message)
}

type TcpClient struct {
	conn       net.Conn //tcp
	Id         int32
	isOnline   chan bool
	PlayerData Data
}

type Data struct {
	Money int32
}

func (t *TcpClient) Start() {

	//goroutine reading

	go func() {
		defer func() { t.isOnline <- false }()
		for {
			//讀取datalen
			headdata := make([]byte, 4)
			if _, err := io.ReadFull(t.conn, headdata); err != nil {
				fmt.Println(err)
				return
			}

			datalen, err := t.Unpack(headdata)

			if err != nil {
				fmt.Println("datalen unpack err:", err)
				return
			}

			headmsgid := make([]byte, 4)
			if _, err := io.ReadFull(t.conn, headmsgid); err != nil {
				fmt.Println(err)
				return
			}

			msgid, err := t.Unpack(headmsgid)

			if err != nil {
				fmt.Println("msgid unpack err:", err)
				return
			}

			bodydata := make([]byte, datalen)

			if _, err := io.ReadFull(t.conn, bodydata); err != nil {
				fmt.Println("bodydata read err:", err)
				return
			}

			m := &Message{
				Len:   uint32(len(bodydata)),
				MsgId: msgid,
				Data:  bodydata,
			}

			fmt.Println(m)

			t.DoMsg(m)
		}
	}()

	//goroutine writing
	go func() {

		//登入
		signid := &pb.SignIn{
			Name:     "Sccot",
			Password: "123456",
		}
		t.SendMsg(1, signid)

		select {
		case isonline := <-t.isOnline: //online api

			if isonline {
				//online
				go func() {
					for {
						//seconds
						time.Sleep(time.Second * 5)

						spin := &pb.Spin{
							Bet: 1,
						}

						t.SendMsg(150, spin)
					}
				}()
			} else {
				//offlline
				return
			}
			break
		}

	}()
}

//處理收到Message的業務
func (t *TcpClient) DoMsg(msg *Message) {

	switch msg.MsgId {
	case 100: //玩家訊息數據
		playerdata := &pb.PlayerData{}
		proto.Unmarshal(msg.Data, playerdata)
		fmt.Println(playerdata.Id, " ", playerdata.Money)
		t.isOnline <- true
		break
	case 200: //玩家獲得盤面
		tabledata := &pb.TableData{}
		proto.Unmarshal(msg.Data, tabledata)
		fmt.Println(tabledata.Table, " ", tabledata.Getmoney)
		break
	}

}

//解出datalen and msgid
func (t *TcpClient) Unpack(headdata []byte) (data uint32, err error) {

	databuff := bytes.NewBuffer(headdata)

	if err = binary.Read(databuff, binary.LittleEndian, &data); err != nil {
		return 0, err
	}

	return data, nil
}

func (t *TcpClient) Pack(msgid uint32, dataBytes []byte) (out []byte, err error) {

	outbuff := bytes.NewBuffer([]byte{})

	//寫Len
	if err = binary.Write(outbuff, binary.LittleEndian, uint32(len(dataBytes))); err != nil {
		return nil, err
	}

	//寫MsgId
	if err = binary.Write(outbuff, binary.LittleEndian, msgid); err != nil {
		return nil, err
	}

	if err = binary.Write(outbuff, binary.LittleEndian, dataBytes); err != nil {
		return nil, err
	}

	out = outbuff.Bytes()

	return out, nil
}

func (t *TcpClient) SendMsg(msgid uint32, data proto.Message) {
	//Marshal
	binarydata, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal err:", err)
		return
	}

	senddata, err := t.Pack(msgid, binarydata)
	if err != nil {
		fmt.Println("pack err:", err)
		return
	}

	if _, err := t.conn.Write(senddata); err != nil {
		fmt.Println("conn write err:", err)
		return
	}

}

//業務

//初始化
func NewTcpClient(ip string, port int) ITcpClient {

	addstr := ip + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", addstr)

	if err != nil {
		panic("net dial err:" + err.Error())
		return nil
	}

	client := &TcpClient{
		conn:     conn,
		isOnline: make(chan bool),
	}

	return client
}
