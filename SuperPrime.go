package main

/**
* @Author: Quincy
* @Date: 2019/11/13 8:58
判断一个数是否是超级素数：
超级素数：从右往左，每去掉一位，便判断剩下的数是否为素数，如果是，则继续该流程；直到到达该数的最高位。
*/
import (
	"fmt"
)

//判断一个数是否为素数
func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	count := num / 2
	for i := 2; i < count; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

/*
生成规定长度范围内的素数
*/
func generatePrime(num int) []int {
	var isPrime = make([]bool, num) //判断是否为素数，初始值为false
	var storage = make([]int, 0)
	var pos = 0
	for i := 2; i < num; i++ {
		if !isPrime[pos] {
			storage = append(storage, i)
			pos++
		}
		for j := 0; j < pos && i*storage[j] < num; j++ {
			isPrime[i*storage[j]] = true
			if i%storage[j] == 0 {
				break
			}
		}
	}
	return storage
}

//func isSuper(values []int) []int {
//	nums := make([]int, 0)
//	for _, value := range values {
//
//		nums = append(nums, value)
//	}
//	return nums
//}
func main() {
	var primes = generatePrime(100)
	for _, prime := range primes {
		fmt.Println(prime)
	}
}
