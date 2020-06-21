package hnet

import (
	"github.com/HyanSource/Helge/hinterface"
)

type Message struct {
	DataLen uint32
	Id      uint32
	Data    []byte
}

func NewMsg(id uint32, data []byte) hinterface.IMessage {

	return &Message{
		DataLen: uint32(len(data)),
		Id:      id,
		Data:    data,
	}
}

func (t *Message) GetDataLen() uint32 {
	return t.DataLen
}

func (t *Message) GetMsgId() uint32 {
	return t.Id
}

func (t *Message) GetData() []byte {
	return t.Data
}
