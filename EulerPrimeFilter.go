package main

import "fmt"

/**
欧拉素数筛选法;
欧拉老爷子太强了！！！
*/
func main() {
	var num int //比num小的素数,即num为筛选的上限
	fmt.Scanf("%d", &num)
	var primes = ELPrimeFilter(num)
	for _, prime := range primes {
		fmt.Println(prime)
	}

}

func ELPrimeFilter(num int) []int {
	var visit = make([]bool, num) //用来标记是否是素数；经过验证，visit的初始值为false
	var prime = make([]int, 0)    //用来存储筛选的素数
	var pos = 0
	for i := 2; i < num; i++ {
		if !visit[i] { //如果i是素数，则把i放入prime切片中，prime数组的索引加一
			prime = append(prime, i)
			pos++
		}
		for j := 0; j < pos && i*prime[j] < num; j++ {
			visit[i*prime[j]] = true
			if i%prime[j] == 0 {
				break
			}
		}
	}
	return prime
}
