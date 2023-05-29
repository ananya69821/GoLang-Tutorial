package main 

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(square(5))
	fmt.Println(square(-5))
}

func square(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("can't take square root of negative number")
	}
	return math.Sqrt(x), nil
}