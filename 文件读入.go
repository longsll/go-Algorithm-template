package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var scanner *bufio.Scanner

func main() {
    scanner = bufio.NewScanner(os.Stdin)
    res := readline()
    fmt.Print(res)
}

func readline() []int {
    scanner.Scan()
    l := strings.Split(scanner.Text(), " ")
    
    res := make([]int, len(l))
    for i, s := range l {
        x, _ := strconv.Atoi(s)
        res[i] = x
    }
    return res
}