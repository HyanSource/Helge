package main

import (
	"github.com/HyanSource/hyannetserver/hnet"
)

func main() {
	server := hnet.NewServer()
	server.Start()
	server.Serve()
}
