package main

func main() {
	for i := 0; i < 32; i++ {
		println(fib(i))
	}

}
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}
