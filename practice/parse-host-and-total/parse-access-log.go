package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var logFile = "hosts_access_log_00.txt"

func main() {
	in, err := ioutil.ReadFile(logFile) // []byte
	check(err)

	logs := strings.Split(string(in), "\n") // string

	out, err := os.Create("records_" + logFile) // takes []byte
	check(err)
	defer out.Close()

	hostRgx := regexp.MustCompile(`^(\S+) `)
	hosts := make(map[string]int)

	for _, line := range logs {
		if !strings.Contains(line, "HTTP") {
			continue
		}
		host := hostRgx.FindStringSubmatch(line) // []string
		if len(host) < 2 {
			log.Println("no match found")
			continue
		}
		h := host[1]
		if _, ok := hosts[h]; !ok {
			hosts[h] = 1
			continue
		}
		hosts[h]++
	}

	for host, count := range hosts {
		fmt.Println(host, count)
		out.WriteString(fmt.Sprintf("%v %v\n", host, count))
	}

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
