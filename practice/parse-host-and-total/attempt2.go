/*
count the number of unique hosts and sort hosts by number of occurances

host1 10
host2 6
host6 2
...
*/

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var log = `unicomp6.unicomp.net - - [01/Jul/1995:00:00:06 -0400] "GET /shuttle/countdown/ HTTP/1.0" 200 3985
burger.letters.com - - [01/Jul/1995:00:00:11 -0400] "GET /shuttle/countdown/liftoff.html HTTP/1.0" 304 0
burger.letters.com - - [01/Jul/1995:00:00:12 -0400] "GET /images/NASA-logosmall.gif HTTP/1.0" 304 0
burger.letters.com - - [01/Jul/1995:00:00:12 -0400] "GET /shuttle/countdown/video/livevideo.gif HTTP/1.0" 200 0
d104.aa.net - - [01/Jul/1995:00:00:13 -0400] "GET /shuttle/countdown/ HTTP/1.0" 200 3985
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /shuttle/countdown/count.gif HTTP/1.0" 200 40310
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /images/NASA-logosmall.gif HTTP/1.0" 200 786
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /images/KSC-logosmall.gif HTTP/1.0" 200 1204
d104.aa.net - - [01/Jul/1995:00:00:15 -0400] "GET /shuttle/countdown/count.gif HTTP/1.0" 200 40310
d104.aa.net - - [01/Jul/1995:00:00:15 -0400] "GET /images/NASA-logosmall.gif HTTP/1.0" 200 786`

func main() {

	type host struct {
		name  string
		count int
	}
	hosts := make(map[string]*host)
	var sorted []host

	loglines := strings.Split(log, "\n")

	regx := regexp.MustCompile(`^(\S+) -`)

	for _, line := range loglines {
		// if !strings.Contains(line, "HTTP") {
		// 	continue
		// }

		match := regx.FindStringSubmatch(line)

		// if len(match) < 2 {
		// 	log.Println("warning")
		// 	continue
		// }
		name := match[1]

		if _, ok := hosts[name]; !ok {
			hosts[name] = &host{
				name:  name,
				count: 1,
			}
			continue
		}
		hosts[name].count++
	}

	for _, v := range hosts {
		sorted = append(sorted, *v)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].count > sorted[j].count })

	fmt.Println(sorted)
}
