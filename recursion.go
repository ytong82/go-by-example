package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return fact(n-1) * n
	}
}

func main() {
	fmt.Printf("The fact of %v is %v \n", 8, fact(8))
}
