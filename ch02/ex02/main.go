package main

import (
	"fmt"
	"os"
	"strconv"
	"./tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		m := tempconv.Meter(t)
		ft := tempconv.Feet(t)
		fmt.Printf("%s = %s, %s = %s\n",
			m, tempconv.MToFt(m), ft, tempconv.FtToM(ft))

		p := tempconv.Pound(t)
		kg := tempconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n",
			p, tempconv.PToKg(p), kg, tempconv.KgToP(kg))
	}
}