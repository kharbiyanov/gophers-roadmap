package main

func OddEvenSum(iterations int) (oddSum, evenSum int) {
	for i := 0; i <= iterations; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}

	return
}
