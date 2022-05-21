package main

import "fmt"

// fibonacci is a function that returns
func fibonacci() func() int {
	current, next, nextnext := 0, 1, 1
	return func() int {
		ret := current

		current = next
		next = nextnext
		nextnext = current + next
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//Golang recursive solution with memoization, 0ms
// Use a helper function so that we can pass in an initialized map.
func fib(N int) int {
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	return helper(N, cache)
}

// This helper function does all the heavy lifting. If our desired fib number
// isn't in the cache, then calculate it first and store it in the cache before
// returning the answer.
func helper(N int, Cache map[int]int) int {
	i, ok := Cache[N]
	if ok {
		return i
	}
	Cache[N] = helper(N-1, Cache) + helper(N-2, Cache)
	return Cache[N]
}
