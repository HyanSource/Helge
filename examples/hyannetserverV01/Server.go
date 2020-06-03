package main

import (
	"github.com/HyanSource/hyannetserver/hnet"
)

func main() {
	//server測試
	s := hnet.NewServer()
	s.Start()
	s.Serve()
}
