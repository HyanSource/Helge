package hnet

import (
	"fmt"
	"strconv"

	"github.com/HyanSource/hyannetserver/hinterface"
	"github.com/HyanSource/hyannetserver/utils"
)

type MsgHandle struct {
	Apis         map[uint32]hinterface.IRouter
	WorkPoolSize uint32
	TaskQueue    []chan hinterface.Irequest
}

/*回傳一個新的處理訊息模塊*/
func NewMsgHandle() hinterface.IMsgHandle {
	return &MsgHandle{
		Apis:         make(map[uint32]hinterface.IRouter),
		WorkPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:    make([]chan hinterface.Irequest, utils.GlobalObject.WorkerPoolSize),
	}
}

//以非阻塞方式處理
func (t *MsgHandle) DoMsgHandler(request hinterface.Irequest) {

	handler, ok := t.Apis[request.GetMessage().GetMsgId()]

	if !ok {
		fmt.Println("not found router msgid:", request.GetMessage().GetMsgId())
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

/*添加路由*/
func (t *MsgHandle) AddRouter(msgid uint32, router hinterface.IRouter) {
	if _, ok := t.Apis[msgid]; ok {
		panic("repeat msgid:" + strconv.Itoa(int(msgid)))
	}

	t.Apis[msgid] = router
	fmt.Println("add api msgid", msgid)
}

/*啟動worker工作池*/
func (t *MsgHandle) StartWorkerPool() {
	//啟動worker
	for i := 0; i < int(t.WorkPoolSize); i++ {
		t.TaskQueue[i] = make(chan hinterface.Irequest, utils.GlobalObject.MaxWorkerTaskLen)
		go t.StartOneWorker(i, t.TaskQueue[i])
	}

}

/*啟動工作池*/
func (t *MsgHandle) StartOneWorker(workID int, taskQueue chan hinterface.Irequest) {
	fmt.Println("worker id:", workID)

	for {
		select {
		case Request := <-taskQueue:
			t.DoMsgHandler(Request)
		}
	}

}

/*消息給TaskQueue*/
func (t *MsgHandle) SendMsgToTaskQueue(request hinterface.Irequest) {
	WorkerID := request.GetConnection().GetConnID() % t.WorkPoolSize

	t.TaskQueue[WorkerID] <- request
}
