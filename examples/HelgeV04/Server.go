package main

import (
	"fmt"

	"github.com/HyanSource/Helge/hinterface"
	"github.com/HyanSource/Helge/hnet"
)

type PingRouter struct {
	hnet.BaseRouter
}

func (t *PingRouter) Handle(request hinterface.Irequest) {
	fmt.Println("recv:", string(request.GetMessage().GetData()))

	err := request.GetConnection().SendBuffMsg(0, []byte("Ping"))

	if err != nil {
		fmt.Println(err)
	}
}

type PingRouter2 struct {
	hnet.BaseRouter
}

func (t *PingRouter2) Handle(request hinterface.Irequest) {
	fmt.Println("recv:", string(request.GetMessage().GetData()))

	err := request.GetConnection().SendBuffMsg(1, []byte("Ping2"))

	if err != nil {
		fmt.Println(err)
	}
}

func StartConn(conn hinterface.Iconnection) {
	fmt.Println(conn.GetConnID(), " startconn")

	conn.GetPropertys().SetProperty("name", "Helge test")
}

func StopConn(conn hinterface.Iconnection) {
	fmt.Println(conn.GetConnID(), " stopconn")

	name, err := conn.GetPropertys().GetProperty("name")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("property name:", name)
}

func main() {
	s := hnet.NewServer()

	//添加路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &PingRouter2{})

	//設置hook
	s.GetHook().SetHook("start", StartConn)
	s.GetHook().SetHook("stop", StopConn)

	s.Serve()
}
