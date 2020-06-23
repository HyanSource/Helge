package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type GlobalObj struct {
	/*Server*/
	// TcpServer hinterface.Iserver
	Host    string
	TCPPort int
	Name    string

	/*hnet*/
	Version string

	MaxPacketSize    uint32
	MaxConn          int
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32
	MaxMsgChanLen    uint32

	/*config file path*/
	ConfFilePath string
}

/*全域變數*/
var GlobalObject *GlobalObj

/*判斷文件存在*/
func (t *GlobalObj) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**/
func (t *GlobalObj) Reload() {
	if confFileExists, _ := t.PathExists(t.ConfFilePath); !confFileExists {
		fmt.Println(t.ConfFilePath, " is not exist")
		return
	}

	data, err := ioutil.ReadFile(t.ConfFilePath)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, t)

	if err != nil {
		panic(err)
	}

}

func init() {
	/*初始化設置默認值*/
	GlobalObject = &GlobalObj{
		/**/
		Name:             "",
		Version:          "V0.1",
		TCPPort:          8123,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPacketSize:    4096,
		ConfFilePath:     "conf/Helge.json",
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
	}
	GlobalObject.Reload()
}
