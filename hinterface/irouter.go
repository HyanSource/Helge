package hinterface

/*路由用*/
type IRouter interface {
	PreHandle(request IRequest)  //
	Handle(request IRequest)     //
	PostHandle(request IRequest) //
}
