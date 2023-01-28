package main

import (
	"flag"
	"fmt"
)

func main() {
	var levels int
	flag.IntVar(&levels, "levels", 5, "number of levels in tree")
	flag.Parse()

	drawChristmasTree(levels)
}

func drawChristmasTree(levels int) {
	for l := 1; l <= levels; l++ {
		for space := 1; space <= levels-l; space++ {
			fmt.Print(" ")
		}
		for star := 1; star <= l*2-1; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
