package hnet

//先嵌入Base之類再根據需求重寫
type BaseRouter struct{}

func (t *BaseRouter) PreHandle()  {}
func (t *BaseRouter) Handle()     {}
func (t *BaseRouter) PostHandle() {}
