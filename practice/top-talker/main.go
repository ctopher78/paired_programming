package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input/hosts_access_log_00.txt")
	check(err)

	logs := strings.Split(string(f), "\n")

	hosts := make(map[string]int)
	regx := regexp.MustCompile(`^(\S+)\s.*\s(\d+)$`)

	for _, line := range logs {
		if !strings.Contains(line, "GET") {
			continue
		}
		match := regx.FindStringSubmatch(line)
		if len(match) < 3 {
			log.Println("regex problem")
			continue
		}
		name := match[1]
		byt, err := strconv.Atoi(match[2])
		check(err)

		if _, ok := hosts[name]; !ok {
			hosts[name] = byt
			continue
		}
		hosts[name] += byt
	}

	var top = struct {
		host string
		byt  int
	}{
		byt: 0,
	}

	for h, b := range hosts {
		if b > top.byt {
			top.host = h
			top.byt = b
		}
	}
	fmt.Println(top)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
