package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type counts struct {
	packets int
	bytes   int
}

type logs struct {
	flows map[string]*counts
}

func main() {
	in, err := ioutil.ReadFile("input/ip-access-list.log")
	check(err)

	logstring := string(in)

	f := make(map[string]*counts)
	l := logs{f}
	l.parse(logstring)
	for k, v := range l.flows {
		fmt.Println(k, v.packets, v.bytes)
	}
}

func (l logs) parse(logs string) error {
	loglines := strings.Split(logs, "\n")

	ttrgx := regexp.MustCompile(`.*denied \w+ (.*), (\d+) \w+ (\d+):$`)

	for _, line := range loglines {
		if !strings.Contains(line, "SEC-6-IPACCESSLOGP") || !strings.Contains(line, "denied") {
			fmt.Println("Here")
			continue
		}
		match := ttrgx.FindStringSubmatch(line)
		if len(match) < 3 {
			log.Println("error: problem with regex")
			continue
		}
		stream := match[1]
		pkt, err := strconv.Atoi(match[2])
		if err != nil {
			return err
		}
		byt, err := strconv.Atoi(match[3])
		if err != nil {
			return err
		}
		if _, ok := l.flows[stream]; !ok {
			l.flows[stream] = &counts{pkt, byt}
			continue
		}
		l.flows[stream].packets += pkt
		l.flows[stream].bytes += byt
	}

	return nil

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
