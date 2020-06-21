package hnet

import (
	"github.com/HyanSource/Helge/hinterface"
)

//先嵌入Base之類再根據需求重寫
type BaseRouter struct{}

func (t *BaseRouter) PreHandle(request hinterface.Irequest)  {}
func (t *BaseRouter) Handle(request hinterface.Irequest)     {}
func (t *BaseRouter) PostHandle(Request hinterface.Irequest) {}
