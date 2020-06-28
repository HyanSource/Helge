package main

import (
	"github.com/HyanSource/Helge/hinterface"
	"github.com/HyanSource/Helge/hnet"
)

/*登入模塊*/
type SignIn struct {
	hnet.BaseRouter
}

func (t *SignIn) Handle(request hinterface.IRequest) {

}

/*遊玩模塊*/
type Spin struct {
	hnet.BaseRouter
}

func (t *Spin) Handle(request hinterface.IRequest) {

}

/*連接時的hook方法*/
func StartConnection(conn hinterface.IConnection) {

}

/*離線時的hook方法*/
func stopConnection(conn hinterface.IConnection) {

}

func main() {
	s := hnet.NewServer()

	s.AddRouter(0, &SignIn{}) //登入業務
	s.AddRouter(100, &Spin{}) //遊玩業務

	s.GetHook().SetHook("start", StartConnection)
	s.GetHook().SetHook("stop", stopConnection)

	s.Serve()
}
