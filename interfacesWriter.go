package netif

import (
	//"errors"
	"fmt"
	"os"
)

type InterfacesWriter struct {
	filePath	string
	backupPath	string
	adapters[]	NetworkAdapter
}

func (iw *InterfacesWriter) WriteInterfaces() error {
	fmt.Println("Writer interface to " + iw.filePath)
	
	// Backup interface file here
	//===========================
	
	// try to open the interface file for writing
	f, err := os.OpenFile(iw.filePath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
	if err != nil {
		// Restore interface file here
		//===========================
		return err
	}
	defer f.Close()
	
	err = iw.WriteInterfacesToFile(f)
	if err != nil {
		// Restore interface file here
		//===========================
	}
	return err
}

func (iw *InterfacesWriter) WriteInterfacesToFile(file *os.File) error {
	
	for naIdx, _ := range iw.adapters {
		iw.writeAdapter(file, &iw.adapters[naIdx])
	}
	
	return nil
}

func (iw *InterfacesWriter) writeAdapter(f *os.File, na *NetworkAdapter) {
	iw.writeAuto(f, na)
	iw.writeHotplug(f, na)
	iw.writeAddrFamily(f, na)
	
	fmt.Fprintln(f, "")
}

func (iw *InterfacesWriter) writeAuto(f *os.File, na *NetworkAdapter) {
	if na.Auto {
		fmt.Fprintf(f, "auto %s\n", na.Name)
	}
}

func (iw *InterfacesWriter) writeHotplug(f *os.File, na *NetworkAdapter) {
	if na.Hotplug {
		fmt.Fprintf(f, "allow-hotplug %s\n", na.Name) 
	}
}

func (iw *InterfacesWriter) writeAddrFamily(f *os.File, na *NetworkAdapter) {
	var familyStr = "inet"
	var sourceStr = "dhcp"
	switch na.AddrFamily {
		default:
		case INET:
			familyStr = "inet"
		case INET6:
			familyStr = "inet6"
	}
	switch na.AddrSource {
		default:
		case DHCP:
			sourceStr = "dhcp"
		case STATIC:
			sourceStr = "static"
		case LOOPBACK:
			sourceStr = "loopback"
		case MANUAL:
			sourceStr = "manual"
	}
	fmt.Fprintf(f, "iface %s %s %s\n", na.Name, familyStr, sourceStr)
}



