package main

import (
	"fmt"

	"github.com/HyanSource/hyannetserver/hnet"
)

func main() {
	server := hnet.NewServer()
	server.Start()

	server.Serve()

	fmt.Println("OK")
}
