package main

import (
	"fmt"

	"github.com/HyanSource/hyannetserver/hnet"
	"github.com/HyanSource/hyannetserver/utils"
)

func main() {
	//server測試
	s := hnet.NewServer()
	s.Start()
	fmt.Println(utils.GlobalObject)
	s.Serve()

}
