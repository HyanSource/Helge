package hnet

import (
	"github.com/HyanSource/hyannetserver/hinterface"
)

type DataPack struct {
}

func NewDataPack() hinterface.IDataPack {
	return &DataPack{}
}

func (t *DataPack) GetHeadLen() uint32 {
	return 0
}

func (t *DataPack) Pack(msg hinterface.IMessage) ([]byte, error) {
	return nil, nil
}

func (t *DataPack) UnPack(data []byte) (hinterface.IMessage, error) {
	return nil, nil
}
