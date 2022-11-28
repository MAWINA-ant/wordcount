package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	count := count(src)
	if count < 0 {
		os.Exit(0)
	}
	fmt.Println(count)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}

func count(src string) int {
	if len(src) == 0 {
		return 0
	}
	lstStr := strings.Split(src, " ")
	return len(lstStr)
}
