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

func (i *Interfaces) WriteInterfaces() {
	// Write adapters to interfaces file
	i.writerFactory().WriteInterfaces()
}

func (i *Interfaces) writerFactory() *InterfacesWriter{
	// Create a writer object
	iw := InterfacesWriter {
		filePath: i.InterfacesPath,
		backupPath : "",
		adapters: i.Adapters,
	}
	
	return &iw
}