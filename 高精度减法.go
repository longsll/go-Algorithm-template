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

    if cmp(A, B) {

        c := sub(A, B)
        for i := len(c) - 1; i >= 0; i-- {
            fmt.Printf("%d", c[i])
        }
    } else {
        c := sub(B, A)
        fmt.Print("-")
        for i := len(c) - 1; i >= 0; i-- {
            fmt.Printf("%d", c[i])
        }
    }
}

// 模板
func sub(A, B []int) []int {
    var c []int
    t := 0                        // 进位
    for i := 0; i < len(A); i++ { // 从个位开始
        t = A[i] - t
        if i < len(B) { // 因为B比A小，看B是否存在
            t -= B[i]
        }
        c = append(c, (t+10)%10) // 如果t>=0还是t,如果t<0,t加上10
        if t < 0 {               // 如果t小于0
            t = 1 // 借位
        } else {
            t = 0 // 否则不借位
        }
    }
    // 去掉前导0
    for len(c) > 1 && c[len(c)-1] == 0 {
        c = c[:len(c)-1]
    }
    return c
}

// 判断 A 是否大于等于 B
func cmp(A, B []int) bool {
    if len(A) != len(B) {
        return len(A) > len(B)
    }
    for i := len(A) - 1; i >= 0; i-- {
        if A[i] != B[i] {
            return A[i] > B[i]
        }
    }
    return true // 正好相等的情况
}