package main

import (
	"fmt"
	"flag"
	"os"
	"runtime"
	"github.com/ccrook/paired_programming/katas/src/fizzbuzzplus/orig"
	greet "github.com/ccrook/paired_programming/katas/src/fizzbuzzplus/cmd/orig/internal"
)

var (
	GitCommit = "unknown"
	GitDescribe = "unknown"
	Build = "unknown"
)


type VersionInfo struct {
	Revision string
	Version string
	Build string
	GoVer string
}

func (v VersionInfo) String() string {
	return fmt.Sprintf("Version: %v\nRev: %v\nBuild: %v\nGoVer: %v\n", v.Version, v.Revision, v.Build, v.GoVer)
}


func main() {

	verFlag := flag.Bool("v", false, "print version")
	flag.Parse()

	ver := VersionInfo{
		Version: GitDescribe,
		Revision: GitCommit,
		Build: Build,
		GoVer: runtime.Version(),
	}

	if *verFlag {
		fmt.Print(ver)
		os.Exit(0)
	}
	for i := 0; i <= 10; i++ {
		if fizzbuzzplus.IsMultipleOrContains(i, 5) && fizzbuzzplus.IsMultipleOrContains(i, 3) {
			fmt.Println("FizzBuzz")
		} else if fizzbuzzplus.IsMultipleOrContains(i, 5) {
			fmt.Println("Buzz")
		} else if fizzbuzzplus.IsMultipleOrContains(i, 3) {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}

	greet.Hello()

	
}
