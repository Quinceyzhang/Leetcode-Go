package main

import (
	"fmt"
)

func main() {
	var num = []int{9, 9, 9}

	var num1 = plusOne(num)
	for _, item := range num1 {
		fmt.Println(item)

	}

}

func plusOne(digits []int) []int {
	var length = len(digits)
	var carry int
	if length == 1 {
		if digits[length-1] == 9 {
			var num = [] int{1, 0}
			return num
		} else {
			digits[length-1]++
			return digits
		}
	} else {
		digits[length-1]++
		if digits[length-1] >= 10 {
			digits[length-1] = digits[length-1] - 10
			carry = 1
		}
		for i := length - 2; i >= 0; i-- {

			digits[i] = digits[i] + carry
			carry = 0
			if digits[i] >= 10 {
				digits[i] = digits[i] - 10
				carry = 1
			}
			if i == 0 && carry == 1 {
				var num = []int{1}
				num = append(num, digits...)
				return num
			}

		}
		return digits
	}
}
