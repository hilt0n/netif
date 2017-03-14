package netif

import (
	"errors"
	"fmt"
	"net"
)

type AddrSource int
const (
	DHCP AddrSource = 1 + iota
	STATIC
	LOOPBACK
	MANUAL
)

type AddrFamily int
const (
	INET AddrFamily = 1 + iota
	INET6
)

// A representation of a network adapter
type NetworkAdapter struct {
	Name			string
	Hotplug			bool
	Auto			bool
	Address			net.IP
	Netmask			net.IP
	Network			net.IP
	Broadcast		net.IP
	Gateway			net.IP
	AddrSource		AddrSource
	AddrFamily		AddrFamily
}

type valueValidator struct {
	Type		string
	Required	bool
	In			[]string
}

var valueValidators = map[string]valueValidator {
	"hotplug": 		{Type: "bool"},
	"auto": 		{Type: "bool"},
	"name": 		{Required: true},
	"address": 		{Type: "IP"},
	"netmask": 		{Type: "IP"},
	"network": 		{Type: "IP"},
	"broadcast": 	{Type: "IP"},
	"gateway": 		{Type: "IP"},
	"addrFam":		{In: []string{"inet", "inet6"}},
	"source": 		{In: []string{"dhcp", "static", "loopback", "manual"}},
}

func (na *NetworkAdapter) validateAll() error {
	/*for k, v := range valueValidators {
		val := nil
		
	}*/
	return nil
}

func (na *NetworkAdapter) validateName() error {
	return nil
}

func (na *NetworkAdapter) validateAddress() error {
	return nil
}

func (na *NetworkAdapter) validateNetmask() error {
	return nil
}

func (na *NetworkAdapter) validateNetwork() error {
	return nil
}

func (na *NetworkAdapter) validateBroadcast() error {
	return nil
}

func (na *NetworkAdapter) validateGateway() error {
	return nil
}

func (na *NetworkAdapter) validateAddrFamily() error {
	return nil
}

func (na *NetworkAdapter) validateSource() error {
	return nil
}

func (na *NetworkAdapter) validateIP(strIP string) (error, net.IP){
	var ip net.IP
	if ip = net.ParseIP(strIP); ip == nil {
		return errors.New("invalide IP address"), nil
	}
	return nil, ip
}

func (na *NetworkAdapter) SetAddress(address string) error {
	err, addr := na.validateIP(address)
	if err == nil {
		na.Address = addr
	}
	return err
}

func (na *NetworkAdapter) SetNetmask(address string) error {
	err, addr := na.validateIP(address)
	if err == nil {
		na.Netmask = addr
	}
	return err
}

func (na *NetworkAdapter) SetGateway(address string) error {
	err, addr := na.validateIP(address)
	if err == nil {
		na.Gateway = addr
	}
	return err
}

func (na *NetworkAdapter) SetBroadcast(address string) error{
	err, addr := na.validateIP(address)
	if err == nil {
		na.Broadcast = addr
	}
	return err
}

func (na *NetworkAdapter) SetNetwork(address string) error{
	err, addr := na.validateIP(address)
	if err == nil {
		na.Network = addr
	}
	return err
}
	
func (na *NetworkAdapter) ParseAddressSource(AddressSource string) (error, AddrSource){
	// Parse the address source for an interface
	var src AddrSource
	switch AddressSource {
		case "dhcp":
			src = DHCP
		case "static":
			src = STATIC
		case "loopback":
			src = LOOPBACK
		case "manual":
			src = MANUAL
		default:
			return errors.New("invalid address source"), -1
	}
	return nil, src
}

func (na *NetworkAdapter) ParseAddressFamily(AddressFamily string) (error, AddrFamily){
	// Parse the address family for an interface
	var fam AddrFamily
	switch AddressFamily {
		case "inet":
			fam = INET
		case "inet6":
			fam = INET6
		default:
			return errors.New("invalid address family"), -1
			
	}
	return nil, fam
}

func (na *NetworkAdapter) Print() {
	fmt.Println("=== Interface ===")
	fmt.Println("=" + na.Name)
	fmt.Println("=================")
	if na.Hotplug {
		fmt.Println("  hotplug: yes")
	} else {
		fmt.Println("  hotplug: no")
	}
	if na.Auto {
		fmt.Println("     auto: yes")
	} else {
		fmt.Println("     auto: no")
	}
	fmt.Print("   source: ")
	fmt.Print(na.AddrSource)
	fmt.Print("\n")
	
	fmt.Print("   family: ")
	fmt.Print(na.AddrFamily)
	fmt.Print("\n")

	fmt.Print("  address: ")
	fmt.Print(na.Address)
	fmt.Print("\n")
	
	fmt.Print("  netmask: ")
	fmt.Print(na.Netmask)
	fmt.Print("\n")
	
	fmt.Print("  gateway: ")
	fmt.Print(na.Gateway)
	fmt.Print("\n")
	
	fmt.Print("broadcast: ")
	fmt.Print(na.Broadcast)
	fmt.Print("\n")
	
	fmt.Print("  network: ")
	fmt.Print(na.Network)
	fmt.Print("\n")
	
	fmt.Println("")
}