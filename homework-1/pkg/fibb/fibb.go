package fibb

import "fmt"

/*
Calculate - calculates sum of previous two numbers in fibonacci sequence fro argument n
*/
func Fib(n int) int {
	if n > 20 {
		fmt.Println("We don't support numbers higher than 20")
		return -1
	}

	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}
