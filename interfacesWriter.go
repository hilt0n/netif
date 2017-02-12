package netif

import (
	"fmt"
)

type InterfacesWriter struct {
	filePath	string
	adapters[]	NetworkAdapter
}

func (i *InterfacesWriter) Write() {
	fmt.Println("Writer interface")
}