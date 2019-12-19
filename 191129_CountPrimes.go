package main

//统计所有小于非负整数 n 的质数的数量。
func main() {
	println(countPrimes(10))
}
func countPrimes(n int) int {
	var primes = PrimeFilter(n)
	return len(primes)
}
func PrimeFilter(n int) []int {
	var isPrime = make([]bool, n) //记录是否为素数,默认值为false
	var primes = make([]int, 0)   //用来存储素数
	var pos = 0

	for i := 2; i < n; i++ {
		if !isPrime[i] { //如果是素数
			primes = append(primes, i)
			pos++
		}
		for j := 0; j < pos && i*primes[j] < n; j++ {
			isPrime[i*primes[j]] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}
	return primes
}
