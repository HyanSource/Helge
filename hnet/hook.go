package hnet

import (
	"github.com/HyanSource/Helge/hinterface"
)

/*hook模塊*/
type Hook struct {
	names     []string                                //存放可用名稱
	hookfuncs map[string]func(hinterface.Iconnection) //存放方法
}

func NewHook() hinterface.IHook {
	return &Hook{
		names:     []string{"start", "stop"},
		hookfuncs: make(map[string]func(hinterface.Iconnection)),
	}
}

func (t *Hook) SetHook(hookname string, f func(hinterface.Iconnection)) {
	if t.contains(hookname) {
		if _, ok := t.hookfuncs[hookname]; !ok {
			t.hookfuncs[hookname] = f
		}
	}
}

func (t *Hook) CallHook(hookname string, conn hinterface.Iconnection) {
	if _, ok := t.hookfuncs[hookname]; ok {
		t.hookfuncs[hookname](conn)
	}
}

/*判斷有無名稱*/
func (t *Hook) contains(hookname string) bool {
	for _, v := range t.names {
		if v == hookname {
			return true
		}
	}
	return false
}
