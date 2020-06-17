package hnet

import (
	"fmt"
	"net"

	"github.com/HyanSource/hyannetserver/hinterface"
	"github.com/HyanSource/hyannetserver/utils"
)

type Server struct {
	Name      string                      //名稱
	IPVersion string                      //ipv4或其他
	IP        string                      //ip
	Port      int                         //port號
	MsgHandle hinterface.IMsgHandle       //消息模塊
	ConnMgr   hinterface.IConnManager     //連接管理
	Hooks     map[string]hinterface.IHook //管理hook的map
}

func NewServer() hinterface.Iserver {

	return &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TCPPort,
		MsgHandle: NewMsgHandle(),
		ConnMgr:   NewConnManager(),
		Hooks:     make(map[string]hinterface.IHook),
	} //以後讀取json檔或是全域
}

func (t *Server) Start() {
	fmt.Println(t.IPVersion, " ", t.IP, ":", t.Port)

	go func() {

		t.MsgHandle.StartWorkerPool()

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
			if t.ConnMgr.Length() >= utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			var cid uint32
			cid = 0

			for t.ConnMgr.ContainsConnID(cid) {
				cid++
			}

			fmt.Println("cid:", cid)

			//處理新連接的請求 應該有conn和handler
			dealConn := NewConntion(t, conn, cid)

			//啟動此連接的業務處理
			go dealConn.Start()
		}

	}()
}

/*停止服務*/
func (t *Server) Stop() {
	t.GetConnMgr().ClearCloseConn()
}

/*運行*/
func (t *Server) Serve() {
	t.Start()

	select {}
}

func (t *Server) AddRouter(msgid uint32, router hinterface.IRouter) {
	t.MsgHandle.AddRouter(msgid, router)
}

func (t *Server) GetConnMgr() hinterface.IConnManager {
	return t.ConnMgr
}

func (t *Server) GetMsgHandle() hinterface.IMsgHandle {
	return t.MsgHandle
}

//設置hook函數
func (t *Server) SetHook(name string, f func(hinterface.Iconnection)) {

	if _, ok := t.Hooks[name]; ok {
		panic("SetHook Repeat:" + name)
	}
	t.Hooks[name] = NewHook(f)
}

//呼叫hook函數
func (t *Server) CallHook(name string, conn hinterface.Iconnection) {

	if h, ok := t.Hooks[name]; ok {
		h.CallHook(conn)
	}

}
