package hnet

import (
	"github.com/HyanSource/Helge/hinterface"
)

type Message struct {
	dataLen uint32 //資料長度
	id      uint32 //資料id
	data    []byte //資料byte 可以搭配protobuf
}

func NewMsg(id uint32, data []byte) hinterface.IMessage {

	return &Message{
		dataLen: uint32(len(data)),
		id:      id,
		data:    data,
	}
}

func (t *Message) GetDataLen() uint32 {
	return t.dataLen
}

func (t *Message) GetMsgId() uint32 {
	return t.id
}

func (t *Message) GetData() []byte {
	return t.data
}
