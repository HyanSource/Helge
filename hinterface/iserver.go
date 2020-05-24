package hinterface

/*服務器接口*/
type Iserver interface {
	Start() //啟動
	Stop()  //停止
	Serve() //開啟業務
}
