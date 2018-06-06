package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input/hosts_access_log_00.txt")
	check(err)

	logs := strings.Split(string(f), "\n")

	type hostAttrs struct {
		name  string
		count int
	}
	hosts := make(map[string]*hostAttrs)
	regx := regexp.MustCompile(`^(\S+)\s.*\s(\d+)$`)

	for _, line := range logs {
		match := regx.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		name := match[1]
		byt, _ := strconv.Atoi(match[2])

		if _, ok := hosts[name]; !ok {
			hosts[name] = &hostAttrs{name: name, count: byt}
			continue
		}
		hosts[name].count++
	}

	var hostSlice []hostAttrs
	for _, v := range hosts {
		hostSlice = append(hostSlice, *v)
	}
	sort.Slice(hostSlice, func(i, j int) bool { return hostSlice.})
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
