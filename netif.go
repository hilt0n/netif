package netif

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (na NetworkAdapter) parseInterface(line string) {
	if !strings.HasPrefix(line, "iface") {
		return
	}
	
	sline := strings.Split(strings.TrimSpace(line), " ")
	for _, s := range sline {
		fmt.Println(s)
	}
} 

func ParseNetworkInterface(path string) []NetworkAdapter {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	var na NetworkAdapter
	
	for scanner.Scan() {
		// Remove the blank charactere at the beginning of the line
		line := strings.TrimSpace(scanner.Text())
		na.parseInterface(line)
		
		// If line begins with a comment, just pass
		if strings.HasPrefix(line, "#") {
			continue
		}
		
		// If line is empty
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		
		
		
		fmt.Println(line)
	}
	
	return nil
}
