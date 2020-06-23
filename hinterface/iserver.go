package hinterface

/*服務器接口*/
type Iserver interface {
	Start() //啟動
	Stop()  //停止
	Serve() //開啟業務

	AddRouter(msgid uint32, router IRouter) //註冊路由的方法

	GetConnMgr() IConnManager //取得管理連接模塊
	GetMsgHandle() IMsgHandle //取得訊息處理模塊

	// SetHook(name string, f func(Iconnection)) //設置hook函數
	// CallHook(name string, conn Iconnection)   //呼叫hook函數
	GetHook() IHook //取得Hook模塊
}
