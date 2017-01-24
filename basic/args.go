package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func printArgs() {
	l := len(os.Args)
	for i := 0; i < l; i++ {
		fmt.Printf("os.Args[%d]:%s\n", i, os.Args[i])
	}
}

func initStrings(n int) []string {
	strs := make([]string, n, n)
	for i := 0; i < n; i++ {
		strs[i] = strconv.Itoa(i)
	}
	return strs
}

func concatString(n int) float64 {
	strs := initStrings(n)
	start := time.Now()
	str := ""
	for _, s := range strs {
		str += s
	}
	return time.Now().Sub(start).Seconds()
}

func joinString(n int) float64 {
	strs := initStrings(n)
	start := time.Now()
	strings.Join(strs, "")
	return time.Now().Sub(start).Seconds()
}

var n = flag.Int("n", 10000, "loop")

func main() {
	flag.Parse()
	printArgs()

	fmt.Printf("concat: %.6fs\n", concatString(*n))
	fmt.Printf("join  : %.6fs\n", joinString(*n))
}
