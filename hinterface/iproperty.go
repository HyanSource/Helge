package hinterface

/*屬性接口模塊*/
type IProperty interface {
	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	RemoveProperty(key string) bool
}
