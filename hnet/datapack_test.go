package hnet

import (
	"fmt"
	"testing"

	"github.com/HyanSource/Helge/utils"
)

/*封包拆包的測試*/
func TestDataPack(t *testing.T) {
	dp := NewDataPack()

	data, err := dp.Pack(NewMsg(100, []byte("hello world")))
	if err != nil {
		fmt.Println("data err:", err)
		return
	}
	fmt.Println(data)

	datalen, err := dp.UnPack_Head(data)
	if err != nil {
		fmt.Println("datalen err:", err)
		return
	}

	if utils.GlobalObject.MaxPacketSize > 0 && utils.GlobalObject.MaxPacketSize > datalen {
		fmt.Println("too large msg data received")
		return
	}

	data = data[4:]

	fmt.Println(datalen)
	datamsgid, err := dp.UnPack_Head(data)
	if err != nil {
		fmt.Println("datamsgid err:", err)
		return
	}
	data = data[4:]

	fmt.Println(datamsgid)

	fmt.Println(string(data))
}
