package hnet

import (
	"github.com/HyanSource/hyannetserver/hinterface"
)

type MsgHandle struct {
	Apis         map[uint32]hinterface.IRouter
	WorkPoolSize uint32
	TaskQueue    []chan hinterface.Irequest
}

func NewMsgHandle() hinterface.IMsgHandle {
	return &MsgHandle{}
}

func (t *MsgHandle) DoMsgHandler(request hinterface.Irequest) {

}

func (t *MsgHandle) AddRouter(msgid uint32, router hinterface.IRouter) {

}

func (t *MsgHandle) StartWorkerPool() {

}

func (t *MsgHandle) SendMsgToTaskQueue(Request hinterface.Irequest) {

}
