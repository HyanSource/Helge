package hnet

import (
	"bytes"
	"encoding/binary"

	"github.com/HyanSource/hyannetserver/hinterface"
)

type DataPack struct {
}

func NewDataPack() hinterface.IDataPack {
	return &DataPack{}
}

/*取得msgid以及datalen*/
func (t *DataPack) GetHeadLen() uint32 {
	return 4
}

/*把message物件轉換成[]byte*/
func (t *DataPack) Pack(msg hinterface.IMessage) ([]byte, error) {

	//byte緩衝
	databuff := bytes.NewBuffer([]byte{})

	//寫datalen
	if err := binary.Write(databuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//寫msgid
	if err := binary.Write(databuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//寫data數據
	if err := binary.Write(databuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return databuff.Bytes(), nil
}

/*把[]byte轉換成int 在解msgid以及datalen使用*/
func (t *DataPack) UnPack_Head(data []byte) (uint32, error) {

	databuff := bytes.NewBuffer(data)

	var getint uint32
	//把
	if err := binary.Read(databuff, binary.LittleEndian, &getint); err != nil {
		return 0, err
	}

	return getint, nil
}
