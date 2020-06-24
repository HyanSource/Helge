package hinterface

/*處理掉用hook的模塊*/
type IHook interface {
	SetHook(hookname string, f func(IConnection))
	CallHook(hookname string, conn IConnection)
}
