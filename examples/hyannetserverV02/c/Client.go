package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/HyanSource/hyannetserver/utils"

	"github.com/hyansource/hyannetserver/hnet"
)

func main() {
	fmt.Println("client start")

	time.Sleep(5 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8124")

	if err != nil {
		fmt.Println("client err:", err)
		return
	}

	for {
		dp := hnet.NewDataPack()
		msg, _ := dp.Pack(hnet.NewMsg(0, []byte("hnet hello")))

		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write err:", err)
			return
		}

		headDatalen := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headDatalen)
		if err != nil {
			fmt.Println("read head error:", err)
			break
		}

		datalen, err := dp.UnPack_Head(headDatalen)
		if err != nil {
			fmt.Println("datalen unpack err:", err)
			return
		}

		/*判斷長度*/

		if utils.GlobalObject.MaxPacketSize > 0 && datalen > utils.GlobalObject.MaxPacketSize {
			fmt.Println("too large msg data")
			return
		}

		headmsgid := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headmsgid); err != nil {
			fmt.Println("headmsgid err:", err)
			break
		}

		msgid, err := dp.UnPack_Head(headmsgid)
		if err != nil {
			fmt.Println("msgid unpack err:", err)
			return
		}

		bodydata := make([]byte, datalen)
		if _, err := io.ReadFull(conn, bodydata); err != nil {
			fmt.Println("body read err:", err)
			break
		}

		if datalen > 0 {
			fmt.Println("msgid:", msgid, " datalen:", datalen, " bodydata:", string(bodydata))
		}

		time.Sleep(1 * time.Second)
	}

}
