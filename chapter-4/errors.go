package main 

import ( "errors"
		 "fmt"
)

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("cannot divide by zero")
	}
	return dividend / divisor, nil
}

func main() {
	result, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result of division is", result)
}
