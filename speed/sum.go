package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	total   float64 = 0
	upto, _         = strconv.ParseFloat(os.Args[1], 64)
)

func main() {
	for curr := 0.0; curr < upto; curr++ {
		total += curr
	}

	fmt.Println(total)
}
