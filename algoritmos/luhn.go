package main

import (
	"fmt"
	"strconv"
)

func LuhnCheck(cardNumber int64) bool {

	cardStr := strconv.FormatInt(cardNumber, 10)
	digits := make([]int, len(cardStr))

	for i, r := range cardStr {
		digits[i] = int(r - '0')
	}

	sum := 0
	alt := false
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i]
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}
	return sum%10 == 0
}

func main() {

	fmt.Println("Welcome to Luhn Algorithm to validate card numbers")
	number := 4940190000370787
	fmt.Printf("Number %v is %v \n", number, LuhnCheck(int64(number)))
}
