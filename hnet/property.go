package hnet

import (
	"errors"

	"github.com/HyanSource/Helge/hinterface"
)

type Property struct {
	propertys map[string]interface{}
}

func NewProperty() hinterface.IProperty {

	return &Property{
		propertys: make(map[string]interface{}),
	}
}

func (t *Property) SetProperty(key string, value interface{}) {
	if _, ok := t.propertys[key]; ok {
		return
	}
	t.propertys[key] = value
}

func (t *Property) GetProperty(key string) (interface{}, error) {

	if _, ok := t.propertys[key]; ok {
		return t.propertys[key], nil
	}

	return nil, errors.New("property not found key:" + key)
}

func (t *Property) RemoveProperty(key string) bool {

	if _, ok := t.propertys[key]; ok {
		delete(t.propertys, key)
		return true
	}

	return false
}
