package hinterface

/*裝包和解包模塊*/
type IDataPack interface {
	GetHeadLen() uint32                   //包頭長度
	Pack(msg IMessage) ([]byte, error)    //裝包
	UnPack(data []byte) (IMessage, error) //解包
}
