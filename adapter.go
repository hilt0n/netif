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
	Name       string
	Hotplug    bool
	Auto       bool
	Address    net.IP
	Netmask    net.IP
	Network    net.IP
	Broadcast  net.IP
	Gateway    net.IP
	AddrSource AddrSource
	AddrFamily AddrFamily
}

type valueValidator struct {
	Type     string
	Required bool
	In       []string
}

var valueValidators = map[string]valueValidator{
	"hotplug":   {Type: "bool"},
	"auto":      {Type: "bool"},
	"name":      {Required: true},
	"address":   {Type: "IP"},
	"netmask":   {Type: "IP"},
	"network":   {Type: "IP"},
	"broadcast": {Type: "IP"},
	"gateway":   {Type: "IP"},
	"addrFam":   {In: []string{"inet", "inet6"}},
	"source":    {In: []string{"dhcp", "static", "loopback", "manual"}},
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

func (na *NetworkAdapter) validateIP(strIP string) (net.IP, error) {
	var ip net.IP
	if ip = net.ParseIP(strIP); ip == nil {
		return nil, errors.New("invalid IP address")
	}
	return ip, nil
}

func (na *NetworkAdapter) SetAddress(address string) error {
	addr, err := na.validateIP(address)
	if err != nil {
		return err
	}
	na.Address = addr
	return nil
}

func (na *NetworkAdapter) SetNetmask(address string) error {
	addr, err := na.validateIP(address)
	if err == nil {
		na.Netmask = addr
	}
	return err
}

func (na *NetworkAdapter) SetGateway(address string) error {
	addr, err := na.validateIP(address)
	if err == nil {
		na.Gateway = addr
	}
	return err
}

func (na *NetworkAdapter) SetBroadcast(address string) error {
	addr, err := na.validateIP(address)
	if err == nil {
		na.Broadcast = addr
	}
	return err
}

func (na *NetworkAdapter) SetNetwork(address string) error {
	addr, err := na.validateIP(address)
	if err == nil {
		na.Network = addr
	}
	return err
}

func (na *NetworkAdapter) SetConfigType(configType string) error {
	switch configType {
	case "DHCP":
		na.AddrSource = DHCP
	case "STATIC":
		na.AddrSource = STATIC
	default:
		return fmt.Errorf("unexpected configType: %s", configType)
	}
	return nil
}

func (na *NetworkAdapter) ParseAddressSource(AddressSource string) (AddrSource, error) {
	// Parse the address source for an interface
	var src AddrSource
	switch AddressSource {
	case "static":
		src = STATIC
	case "dhcp":
		src = DHCP
	case "loopback":
		src = LOOPBACK
	case "manual":
		src = MANUAL
	default:
		return -1, errors.New("invalid address source")
	}
	return src, nil
}

func (na *NetworkAdapter) ParseAddressFamily(AddressFamily string) (AddrFamily, error) {
	// Parse the address family for an interface
	var fam AddrFamily
	switch AddressFamily {
	case "inet":
		fam = INET
	case "inet6":
		fam = INET6
	default:
		return -1, errors.New("invalid address family")

	}
	return fam, nil
}
