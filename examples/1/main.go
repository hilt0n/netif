package main

import ni "github.com/Wifx/netif"

func main() {
	is := ni.Parse(
		ni.Path("input"),
	)

	is.Write(
		ni.Path("output"),
	)
}
