package hinterface

/*處理掉用hook的模塊*/
type IHook interface {
	// SetHook(f func(Iconnection))
	CallHook(conn Iconnection)
}
