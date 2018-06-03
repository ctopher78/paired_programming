package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input/ip-access-list.log")
	check(err)

	logs := strings.Split(string(f), "\n")

	type data struct {
		srcIP string
		pks   int
		byts  int
	}

	sources := make(map[string]*data)

	regx := regexp.MustCompile(`.*denied \w+ ([0-9.]+)\(.*, (\d+) \w+ (\d+):`)

	for _, line := range logs {
		if !strings.Contains(line, "denied") {
			continue
		}

		match := regx.FindStringSubmatch(line)

		if len(match) < 4 {
			log.Printf("interesting line, but no match found")
			continue
		}

		srcIP := match[1]
		check(err)

		pkts, err := strconv.Atoi(match[2])
		check(err)

		byts, err := strconv.Atoi(match[3])
		check(err)

		d := data{
			srcIP,
			pkts,
			byts,
		}
		if _, ok := sources[srcIP]; !ok {
			sources[srcIP] = &d
			continue
		}
		sources[srcIP].srcIP = d.srcIP
		sources[srcIP].pks += d.pks
		sources[srcIP].byts += d.byts
	}

	var sorted []data
	for _, v := range sources {
		sorted = append(sorted, *v)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].byts > sorted[j].byts })
	for _, e := range sorted {
		fmt.Println(e.srcIP, e.pks, e.byts)
	}

	// sort.Slice(sources, func(i, j int) bool { return sources[i].byts > sources[j].byts })
	// for _, s := range sources {
	// 	fmt.Println(s.srcIP, s.pks, s.byts)
	// }
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
