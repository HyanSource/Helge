package hinterface

/* 裝包和解包模塊
 * 目前此解包採用的是先解出包(長度和msgid) 再來是(內容)
 */
type IDataPack interface {
	GetHeadLen() uint32                //包頭長度
	Pack(msg IMessage) ([]byte, error) //裝包
	// UnPack(data []byte) (IMessage, error) //解包
	UnPack_Head(data []byte) (uint32, error) //解包取得int 回傳
}
