package hinterface

/*路由用*/
type IRouter interface {
	PreHandle(request Irequest)  //
	Handle(request Irequest)     //
	PostHandle(request Irequest) //
}
