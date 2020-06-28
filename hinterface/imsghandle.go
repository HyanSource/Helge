package hinterface

/*管理消息接口*/
type IMsgHandle interface {
	DoMsgHandler(request IRequest)          //以非阻塞處理消息
	AddRouter(msgid uint32, router IRouter) //增加路由
	StartWorkerPool()                       //啟動工作池
	SendMsgToTaskQueue(request IRequest)    //將消息給工作池處理
}
