package main

import (
	"fmt"

	"github.com/HyanSource/Helge/examples/GameV01/pb"
	"github.com/HyanSource/Helge/hinterface"
	"github.com/HyanSource/Helge/hnet"
	"github.com/golang/protobuf/proto"
)

/*登入模塊*/
type SignIn struct {
	hnet.BaseRouter
}

func (t *SignIn) Handle(request hinterface.IRequest) {

	playerdata := &pb.PlayerData{
		Id:    request.GetConnection().GetConnID(),
		Money: 100000,
	}

	b, err := proto.Marshal(playerdata)
	if err != nil {
		fmt.Println("playerdata marshal err:", err)
		return
	}

	request.GetConnection().SendMsg(100, b)
}

/*遊玩模塊*/
type Spin struct {
	hnet.BaseRouter
}

func (t *Spin) Handle(request hinterface.IRequest) {

	tabledata := &pb.TableData{
		Table:    "",
		Getmoney: 0,
	}

	b, err := proto.Marshal(tabledata)
	if err != nil {
		fmt.Println("tabledata marshal err:", err)
		return
	}

	request.GetConnection().SendMsg(200, b)
}

/*連接時的hook方法*/
func StartConnection(conn hinterface.IConnection) {

	// msg := &pb.StartConn{Connid: 1}
	// bd, err := proto.Marshal(msg)

	// if err != nil {
	// 	fmt.Println("startconn err:", err)
	// 	return
	// }

}

/*離線時的hook方法*/
func stopConnection(conn hinterface.IConnection) {
	// msg := &pb.StopConn{Connid: conn.GetConnID()}
	// bd, err := proto.Marshal(msg)

	// if err != nil {
	// 	fmt.Println("stopconn err:", err)
	// 	return
	// }

}

func main() {
	s := hnet.NewServer()

	s.AddRouter(1, &SignIn{}) //登入業務
	s.AddRouter(150, &Spin{}) //遊玩業務

	s.GetHook().SetHook("start", StartConnection)
	s.GetHook().SetHook("stop", stopConnection)

	s.Serve()
}
