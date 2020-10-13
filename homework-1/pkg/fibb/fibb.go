package fibb

import (
	"fmt"
	"os"
)

func Fib(n int) int {
	if n > 20 {
		fmt.Println("We don't support numbers higher than 20")
		os.Exit(1)
	}

	var a, b int = 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}
