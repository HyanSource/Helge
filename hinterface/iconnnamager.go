package hinterface

/*管理連接接口*/
type IConnManager interface {
	Add(conn Iconnection)                   //新增
	Remove(conn Iconnection)                //移除(應該還有一個用id去移除的方法)
	Get(connid uint32) (Iconnection, error) //用connid取得連接
	Length() int                            //目前總連接數量
	ClearCloseConn()                        //清除以及關閉所有連接
}
