package main

func Fibonacci(iterations int) []int {
	var numbers []int
	x, y := 0, 1
	for i := 0; i < iterations; i++ {
		x, y = y, x+y
		numbers = append(numbers, x)
	}
	return numbers
}
