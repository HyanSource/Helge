package hinterface

type IRequest interface {
	GetConnection() IConnection //請求連接
	GetMessage() IMessage       //訊息
}
