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

func Path(path string) fn.Option {
	return fn.MakeOption("path", path)
}
