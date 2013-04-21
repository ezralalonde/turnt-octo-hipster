package main

import (
	"fmt"
	"turnt-octo-hipster/max"
)

func main() {
	var sets int
	_, err := fmt.Scanf("%d", &sets)
	Check(err, "could not read sets")
	for ii := 0; ii < sets; ii++ {
		var col int
		_, err := fmt.Scanf("%d", &col)
		Check(err, "could not read number of columns")
		xx := ReadNextRow(col)
		yy := ReadNextRow(col)
		fmt.Printf("The maximum distance is %v\n\n", max.MaximumDistance(xx, yy))
	}
}

func Check(err error, str string) {
	if err != nil {
		panic(str)
	}
}

func ReadNextRow(col int) (xx []int) {
	for jj := 0; jj < col; jj++ {
		var val int
		_, err := fmt.Scanf("%d", &val)
		Check(err, "could not read row")
		xx = append(xx, val)
	}
	return
}
