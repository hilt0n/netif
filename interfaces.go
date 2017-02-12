package netif

import (
	"fmt"
)

func NewInterfaces() *Interfaces{
	return &Interfaces{
		InterfacesPath: "/etc/network/interfaces",
	}
}

type Interfaces struct {
	InterfacesReader
	InterfacesWriter
	
	InterfacesPath		string
	Adapters[]			NetworkAdapter
}

func (i *Interfaces) Init() {
	fmt.Println("Init")
}

func (i *Interfaces) UpdateAdapters() {
	// (re)read interfaces file and save adapters
	i.Adapters = NewInterfacesReader(i.InterfacesPath).ParseInterfaces()

	for _, adapter := range i.Adapters {
		adapter.Print()
	}
}