package main

import "fmt"

func fibnacci(n int) int {
	if n < 2 {
		return n
	}
	return fibnacci(n-2) + fibnacci(n-1)
}
func sqrt(x float64, i float64) (float64, float64) {
	remain := (i*i - x) / (2 * i)
	i = i - remain
	if remain > 0 {
		return sqrt(x, i)
	} else {
		return i, remain
	}
}
func get_sqrt(x float64) float64 {
	i, _ := sqrt(x, x)
	return i
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fibnacci(i))
	}
	fmt.Println(get_sqrt(2))
	fmt.Println(get_sqrt(3))
}
