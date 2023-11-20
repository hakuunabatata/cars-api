package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	result, err := sum(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func sum(a int, b int) (int, error) {
	if a+b > 10 {
		return 0, fmt.Errorf("sum > 10")
	}
	return a + b, nil
}
