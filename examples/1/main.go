package main

import ni "github.com/Wifx/netif"

func main() {
	i := ni.NewInterfaceSet(ni.Path("input"))
	i.UpdateAdapters()

	i.Write(
		ni.Path("output"),
	)
}
