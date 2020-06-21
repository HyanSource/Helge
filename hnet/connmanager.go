package hnet

import (
	"errors"
	"fmt"
	"sync"

	"github.com/HyanSource/Helge/hinterface"
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

	t.connMutex.Lock()
	defer t.connMutex.Unlock()

	t.connections[conn.GetConnID()] = conn

	fmt.Println("connection add success len:", len(t.connections))
}

/**/
func (t *ConnManager) Remove(connid uint32) {
	t.connMutex.Lock()
	defer t.connMutex.Unlock()

	delete(t.connections, connid)
}

/**/
func (t *ConnManager) Get(connid uint32) (hinterface.Iconnection, error) {

	t.connMutex.RLock()
	defer t.connMutex.RUnlock()

	if conn, ok := t.connections[connid]; ok {
		return conn, nil
	}

	return nil, errors.New("connection not found")
}

/**/
func (t *ConnManager) Length() int {
	t.connMutex.RLock()
	defer t.connMutex.RUnlock()

	return len(t.connections)
}

/**/
func (t *ConnManager) ClearCloseConn() {
	t.connMutex.Lock()
	defer t.connMutex.Unlock()
	/*停止並刪除連接*/
	for k, v := range t.connections {
		v.Stop()
		delete(t.connections, k)
	}

	fmt.Println("ClearClose success:")
}

/*判斷所有連接connid有無重複*/
func (t *ConnManager) ContainsConnID(connid uint32) bool {

	for k, _ := range t.connections {
		if k == connid {
			return true
		}
	}

	return false
}
