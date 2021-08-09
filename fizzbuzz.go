package main

import "strconv"

func FizzBuzz(iterations int) []string {
	var result []string
	for i := 1; i < iterations; i++ {
		isFizzBuzz := [2]bool{i%3 == 0, i%5 == 0}
		switch isFizzBuzz {
		case [2]bool{true, true}:
			result = append(result, "FizzBuzz")
		case [2]bool{true, false}:
			result = append(result, "Fizz")
		case [2]bool{false, true}:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
