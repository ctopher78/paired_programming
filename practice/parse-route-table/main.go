package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input/show-ip-route.log")
	check(err)

	inlogs := strings.Split(string(f), "\n")

	protomap := make(map[string][]string)

	protoRgx := regexp.MustCompile(`^([A-Z])\s+(\S+)\s`)

	for _, line := range inlogs {
		match := protoRgx.FindStringSubmatch(line)
		if len(match) < 3 {
			continue
		}
		proto := match[1]
		prefix := match[2]
		// {
		protomap[proto] = append(protomap[proto], prefix)
		// 	continue
		// }
	}

	for k, v := range protomap {
		fmt.Println(k)
		for _, i := range v {
			fmt.Println(i)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
