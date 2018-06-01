package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	f, err := ioutil.ReadFile("input/show-interface.txt")
	check(err)
	showint := string(f)
	fmt.Println(showint)
	regx := regexp.MustCompile(`(?s).*BW (\d+).*input rate (\d+).*output rate (\d+)`)
	match := regx.FindStringSubmatch(showint)
	fmt.Println(match[1:])
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
