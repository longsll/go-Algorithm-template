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

    c := mul(A, b)
    for i := len(c) - 1; i >= 0; i-- {
        fmt.Printf("%d", c[i])
    }
}

func mul(A []int, b int) []int {
    var c []int
    t := 0 // 进位 从前往后循环
    for i := 0; i < len(A) || t != 0; i++ {
        if i < len(A) { // 如果 A 还没有遍历完
            t += A[i] * b
        }
        c = append(c, t%10)
        t /= 10 // 整除进位
    }
    for len(c) > 1 && c[len(c)-1] == 0 {
        c = c[:len(c)-1]
    }
    return c
}