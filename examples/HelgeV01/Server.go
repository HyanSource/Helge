package main

import (
	"fmt"

	"github.com/HyanSource/Helge/hnet"
	"github.com/HyanSource/Helge/utils"
)

func main() {
	//server測試
	s := hnet.NewServer()
	s.Start()
	fmt.Println(utils.GlobalObject)
	s.Serve()

}
