/*
Challenge 1:

Parse the "show ip int brie" output and print all interfaces that aren't up.

Input:
Router# show ip interface brief
Interface     IP-Address     OK?  Method  Status                  Protocol
Ethernet0     10.108.00.5    YES  NVRAM   up                      up
Ethernet1     unassigned     YES  unset   administratively down   down
Loopback0     10.108.200.5   YES  NVRAM   up                      up
Serial0       10.108.100.5   YES  NVRAM   up                      up
Serial1       10.108.40.5    YES  NVRAM   up                      up
Serial2       10.108.100.5   YES  manual  proto                   down
Serial3       unassigned     YES  unset   administratively down   down


Output:
Ethernet1 administratively down
Serial3 administratively down
*/
package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	input := strings.Split(showIntBrief, "\n")

	var down [][]string // ["eth0", "status"]
	regx := regexp.MustCompile(`^(\S+)\s+\S+\s+\S+\s+\S+\s+(\S+(?: \S+)?)\s+down`)

	for _, line := range input {
		if !strings.Contains(line, "down") {
			continue
		}

		match := regx.FindStringSubmatch(line)

		if len(match) != 3 {
			log.Println("assumption warning")
			continue
		}

		down = append(down, match[1:])
	}

	for _, iface := range down {
		fmt.Println(iface[0], iface[1])
	}

}

var showIntBrief = `Interface     IP-Address     OK?  Method  Status                  Protocol
Ethernet0     10.108.00.5    YES  NVRAM   up                      up      
Ethernet1     unassigned     YES  unset   administratively down   down    
Loopback0     10.108.200.5   YES  NVRAM   up                      up      
Serial0       10.108.100.5   YES  NVRAM   up                      up      
Serial1       10.108.40.5    YES  NVRAM   up                      up      
Serial2       10.108.100.5   YES  manual  proto                   down     
Serial3       unassigned     YES  unset   administratively down   down `
