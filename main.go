package main

import (
	"fmt"
	"turnt-octo-hipster/max"
)

func main() {
	var sets int
	_, err := fmt.Scanf("%d", &sets)
	if err != nil {
		panic("could not read sets")
	}
	for ii := 0; ii < sets; ii++ {
		var col int
		_, err := fmt.Scanf("%d", &col)
		if err != nil {
			panic("could not read number of columns")
		}
		var xx, yy []int
		for jj := 0; jj < col; jj++ {
			var val int
			_, err := fmt.Scanf("%d", &val)
			if err != nil {
				panic("could not read row")
			}
			xx = append(xx, val)
		}
		for jj := 0; jj < col; jj++ {
			var val int
			_, err := fmt.Scanf("%d", &val)
			if err != nil {
				panic("could not read row")
			}
			yy = append(yy, val)
		}
		fmt.Printf("The maximum distance is %v\n\n", max.MaximumDistance(xx, yy))
	}
}
