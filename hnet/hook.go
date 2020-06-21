package hnet

import (
	"github.com/HyanSource/Helge/hinterface"
)

/*hook模塊*/
type Hook struct {
	hookfunc func(hinterface.Iconnection)
}

func NewHook(f func(hinterface.Iconnection)) hinterface.IHook {
	return &Hook{
		hookfunc: f,
	}
}

func (t *Hook) CallHook(conn hinterface.Iconnection) {
	if t.hookfunc != nil {
		t.hookfunc(conn)
	}
}
