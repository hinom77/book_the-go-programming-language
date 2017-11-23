package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	algo1()
	fmt.Printf("Pattern1: %.20fs\n", time.Since(start1).Seconds())

	start2 := time.Now()
	algo2()
	fmt.Printf("Pattern2: %.20fs\n", time.Since(start2).Seconds())
}

func algo1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func algo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}


