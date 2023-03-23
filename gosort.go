package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{6, 3, 9, 8, 1, 2, 5, 7}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a)
}

