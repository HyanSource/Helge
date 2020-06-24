package hinterface

type Irequest interface {
	GetConnection() IConnection //請求連接
	GetMessage() IMessage       //訊息
}
