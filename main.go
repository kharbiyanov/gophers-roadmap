package main

import "fmt"

func main() {
	fmt.Println("Fibonacci result:", Fibonacci(10))
	fmt.Println("FizzBuzz result:", FizzBuzz(100))

	palindrome := Palindrome{}
	for _, word := range []string{"заказ", "Кабак", "Вадим"} {
		if isPalindrome := palindrome.Check(word); isPalindrome {
			fmt.Printf("Palindrome check: %s is a palindrome.\n", word)
		} else {
			fmt.Printf("Palindrome check: %s is not a palindrome.\n", word)
		}
	}

	oddEvenSumIterations := 10
	oddSum, evenSum := OddEvenSum(oddEvenSumIterations)
	fmt.Printf("OddEvenSum result for %d iterations - Sum of odds: %d; Sum of evens: %d;\n", oddEvenSumIterations, oddSum, evenSum)

	toCheckDuplicates := []interface{}{1,2,3,4,2,"Hello", "World", "Hello"}
	if hasDuplicate, duplicatedItems := HasDuplicate(toCheckDuplicates); hasDuplicate{
		fmt.Printf("Array of \"%v\" has duplicated items: %v", toCheckDuplicates, duplicatedItems)
	} else {
		fmt.Printf("Array of \"%v\" has not duplicated items", toCheckDuplicates)
	}


}
