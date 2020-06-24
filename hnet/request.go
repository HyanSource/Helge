package hnet

import (
	"github.com/HyanSource/Helge/hinterface"
)

type Request struct {
	conn hinterface.IConnection
	msg  hinterface.IMessage
}

func (t *Request) GetConnection() hinterface.IConnection {
	return t.conn
}

func (t *Request) GetMessage() hinterface.IMessage {
	return t.msg
}
