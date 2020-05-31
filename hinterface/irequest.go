package hinterface

type Irequest interface {
	GetConnection() Iconnection //請求連接
	GetMessage() IMessage       //訊息
}
