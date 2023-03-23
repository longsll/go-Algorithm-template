package main

import "fmt"

func main() {
    var a, b string
    fmt.Scanf("%s", &a) // a="123456"
    fmt.Scanf("%s", &b)

    var A, B []int
    for i := len(a) - 1; i >= 0; i-- { // A=[6,5,4,3,2,1]
        A = append(A, int(a[i]-'0')) // 存储数字到数组中
    }
    for i := len(b) - 1; i >= 0; i-- {
        B = append(B, int(b[i]-'0')) // 存储数字到数组中
    }

    c := add(A, B)
    for i := len(c) - 1; i >= 0; i-- {
        fmt.Printf("%d", c[i])
    }
}

// 模板
func add(A, B []int) []int {
    var c []int
    t := 0 // 进位
    for i := 0; i < len(A) || i < len(B); i++ {
        if i < len(A) {
            t += A[i]
        }
        if i < len(B) {
            t += B[i]
        }
        c = append(c, t%10)
        t /= 10
    }
    if t != 0 {
        c = append(c, 1)
    }
    return c
}