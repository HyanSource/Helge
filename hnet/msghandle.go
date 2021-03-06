package hnet

import (
	"fmt"
	"strconv"

	"github.com/HyanSource/Helge/hinterface"
	"github.com/HyanSource/Helge/utils"
)

type MsgHandle struct {
	Apis         map[uint32]hinterface.IRouter
	WorkPoolSize uint32
	TaskQueue    []chan hinterface.IRequest
}

/*回傳一個新的處理訊息模塊*/
func NewMsgHandle() hinterface.IMsgHandle {
	return &MsgHandle{
		Apis:         make(map[uint32]hinterface.IRouter),
		WorkPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:    make([]chan hinterface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

//以非阻塞方式處理
func (t *MsgHandle) DoMsgHandler(request hinterface.IRequest) {

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
		t.TaskQueue[i] = make(chan hinterface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go t.StartOneWorker(i, t.TaskQueue[i]) //這裡應該不用傳chan
	}

}

/*啟動工作池*/
func (t *MsgHandle) StartOneWorker(workID int, taskQueue chan hinterface.IRequest) {
	fmt.Println("worker id:", workID)

	/*處理任務迴圈*/
	for {
		select {
		case Request := <-taskQueue:
			t.DoMsgHandler(Request)
		}
	}

}

/*消息給TaskQueue*/
func (t *MsgHandle) SendMsgToTaskQueue(request hinterface.IRequest) {
	//可以把這邊判斷workid 改成進來1次訊息加1 當到最大上限時歸0
	WorkerID := request.GetConnection().GetConnID() % t.WorkPoolSize

	t.TaskQueue[WorkerID] <- request
}
