package hnet

import (
	"fmt"
	"net"

	"github.com/HyanSource/hyannetserver/hinterface"
)

type Server struct {
	Name      string //名稱
	IPVersion string //ipv4或其他
	IP        string //ip
	Port      int    //port號
}

func NewServer() hinterface.Iserver {

	return &Server{} //以後讀取json檔或是全域
}

func (t *Server) Start() {
	fmt.Println("")

	go func() {

		//1.獲取tcp的addr
		addr, err := net.ResolveTCPAddr(t.IPVersion, fmt.Sprintf("%s:%d", t.IP, t.Port))

		if err != nil {
			panic("resolve tcp addr err:" + err.Error())
		}

		//2.監聽服務器地址
		listener, err := net.ListenTCP(t.IPVersion, addr)

		if err != nil {
			panic("listen err:" + err.Error())
		}

		fmt.Println("listen OK")

		//原本的作法是用cid去當成connection以及connmanager的id 可能要想其他的方式
		//比如 隨機生成之後去判斷connmanager的map有無key在新增

		//3.啟動server
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP err:", err)
				continue
			}
			//顯示客戶端的ip port
			fmt.Println("Get conn remote addr:" + conn.RemoteAddr().String())

			//超過最大連線時 關閉此連線

			//處理新連接的請求 應該有conn和handler

			//啟動此連接的業務處理
		}

	}()
}

func (t *Server) Stop() {

}

func (t *Server) Serve() {

	select {}
}
