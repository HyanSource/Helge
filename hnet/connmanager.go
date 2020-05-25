package hnet

import (
	"sync"

	"github.com/HyanSource/hyannetserver/hinterface"
)

/*連接管理模塊*/
type ConnManager struct {
	connections map[uint32]hinterface.Iconnection //
	connMutex   sync.RWMutex                      //單寫多讀鎖     占用時阻止寫入 但不會阻止讀取            //
}

/**/
func NewConnManager() hinterface.IConnManager {

	//尚未實作 所以沒有辦法去New模塊
	return &ConnManager{
		connections: make(map[uint32]hinterface.Iconnection),
	}
}

/**/
func (t *ConnManager) Add(conn hinterface.Iconnection) {

}

/**/
func (t *ConnManager) Remove(conn hinterface.Iconnection) {

}

/**/
func (t *ConnManager) Get(connid uint32) (hinterface.Iconnection, error) {

	return nil, nil
}

/**/
func (t *ConnManager) Length() int {
	return 0
}

/**/
func (t *ConnManager) ClearCloseConn() {

}
