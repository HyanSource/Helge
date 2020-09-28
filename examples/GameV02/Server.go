package main

import (
	_ "github.com/HyanSource/Helge/examples/GameV02/slot"
	"github.com/HyanSource/Helge/hnet"
)

func main() {
	s := hnet.NewServer()

	s.Serve()
}
