package main

import (
    "fmt"
)
const N int = 1005

var (
	n , m int
	v [N]int
	w [N]int
	dp [N]int
)

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n , &m)
	for i := 1 ; i <= n ; i ++ {
		fmt.Scan(&v[i] , &w[i])
		fmt.Print(v[i] , " " , w[i])
	}
	for i := 1 ;i <= n ; i ++ {
		for j := m ; j >= v[i] ; j -- {
			dp[j] = max(dp[j] , dp[j - v[i]] + w[i])
		}
	}
	fmt.Print(dp[m])
}