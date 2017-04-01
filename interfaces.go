package netif

import "github.com/n-marshall/fn"

type InterfaceSet struct {
	InterfacesReader

	InterfacesPath string
	Adapters       []NetworkAdapter
}

func NewInterfaceSet(opts ...fn.Option) *InterfaceSet {
	fnConfig := fn.MakeConfig(
		fn.Defaults{"path": "/etc/network/interfaces"},
		opts,
	)
	path := fnConfig.GetString("path")

	return &InterfaceSet{
		InterfacesPath: path,
	}
}

func (i *InterfaceSet) UpdateAdapters() {
	// (re)read interfaces file and save adapters
	i.Adapters = NewInterfacesReader(i.InterfacesPath).ParseInterfaces()

	// for _, adapter := range i.Adapters {
	// 	adapter.Print()
	// }
}
