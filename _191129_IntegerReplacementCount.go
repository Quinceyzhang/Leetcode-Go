package main

//给定一个正整数 n，你可以做如下操作：
//
//1. 如果 n 是偶数，则用 n / 2替换 n。
//2. 如果 n 是奇数，则可以用 n + 1或n - 1替换 n。
//n 变为 1 所需的最小替换次数是多少？
func main() {
	println(integerReplacement(8))
}

//递归的解法
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		return 1 + min(integerReplacement(n-1), integerReplacement(n+1))
	}
}

func min(num1, num2 int) int {
	if num1 >= num2 {
		return num2
	} else {
		return num1
	}
}

//TODO
//动态规划的解法

//TODO
//除了计算最少的次数，还要求输出变换的过程
//如：
//8->4->2->1

