package main

import "fmt"

func main() {
    var a string
    var b int
    fmt.Scanf("%s", &a)
    fmt.Scanf("%d", &b)
    var A []int
    for i := len(a) - 1; i >= 0; i-- {
        A = append(A, int(a[i]-'0'))
    }

    r := 0
    c := div(A, b, &r)

    for i := len(c) - 1; i >= 0; i-- {
        fmt.Printf("%d", c[i])
    }
    fmt.Println()
    fmt.Println(r)
}

// A / b, 商
func div(A []int, b int, r *int) []int {
    var c []int
    *r = 0
    for i := len(A) - 1; i >= 0; i-- {
        *r = *r*10 + A[i]
        c = append(c, *r/b)
        *r %= b
    }
    for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
        c[i], c[j] = c[j], c[i]
    }
    for len(c) > 1 && c[len(c)-1] == 0 {
        c = c[:len(c)-1]
    }
    return c
}
