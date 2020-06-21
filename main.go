package main

import (
	"github.com/HyanSource/Helge/hnet"
)

func main() {
	server := hnet.NewServer()
	server.Start()
	server.Serve()
}
