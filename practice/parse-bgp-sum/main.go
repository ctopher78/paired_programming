package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type neighbor struct {
	nei   string
	ver   string
	as    string
	estab bool
}

type neighbors []neighbor

func main() {
	f, err := ioutil.ReadFile("input/cisco-show-ip-bgp-sum.log")
	check(err)

	loglines := strings.Split(string(f), "\n")

	var n neighbors

	nbrRgx := regexp.MustCompile(`\s+`)
	neiRgx := regexp.MustCompile(`^\S+[.:]`)

	for _, line := range loglines {
		parts := nbrRgx.Split(line, -1)
		if !neiRgx.MatchString(line) {
			continue
		}
		estab := true
		if _, err := strconv.Atoi(parts[len(parts)-1]); err != nil {
			estab = false
		}
		nei := neighbor{
			nei:   parts[0],
			ver:   parts[1],
			as:    parts[2],
			estab: estab,
		}
		n = append(n, nei)
	}

	for _, i := range n {
		if !i.estab {
			fmt.Println(i)
		}
	}

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
