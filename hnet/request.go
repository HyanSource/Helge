package hnet

import (
	"github.com/HyanSource/Helge/hinterface"
)

type Request struct {
	conn hinterface.Iconnection
	msg  hinterface.IMessage
}

func (t *Request) GetConnection() hinterface.Iconnection {
	return t.conn
}

func (t *Request) GetMessage() hinterface.IMessage {
	return t.msg
}
