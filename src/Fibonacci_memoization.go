package main

import "fmt"

func test_fib_mem() {
	fmt.Printf("fib 40=%d\n", fib_mem(40))
}
func fib_mem(n int) uint64 {
	var fibs = make([]uint64, n+1)
	f := func(n int) uint64 {
		if n <= 1 {
			fibs[n] = 1
			return 1
		} else {
			fibs[n] = fibs[n-1] + fibs[n-2]
		}
		return fibs[n]
	}
	for i := 1; i <= n; i++ {
		f(i)
	}
	return f(n)

}
